// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package main

import (
	"log"
	"net/http_my"
	"time"
	"io/ioutil"
)

func main() {
	srv := &http.Server{
		Addr:        "localhost:8089",
		Handler:     myhandler{},
		ReadTimeout: 2 * time.Second,
		IdleTimeout: 10 *time.Second,
		ReadHeaderTimeout: 2 *time.Second,
	}
	go func() {
		// wait server to start
		time.Sleep(time.Second)
		log.Print("[Client]: Sending request.")
		resp, err := http.Get("http://" + srv.Addr)
		log.Print("Client]: Checking resp && err")
		if err != nil {
			log.Fatalf("[Client]: %v\n", err)
			return
		}
		resp.Body.Close()
	}()
	log.Fatal(srv.ListenAndServe())
}

type myhandler struct{}

func (h myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print("Serving.")
	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("[Server]: %v\n", err)
		return
	}
	if cn, ok := w.(http.CloseNotifier); ok {
		<-cn.CloseNotify()
		log.Panic("Unexpected closed client")
	}
}
