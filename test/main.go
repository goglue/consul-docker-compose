package main

import (
	"net/http"
	"time"
	"os"
	"syscall"
	"os/signal"
	"io/ioutil"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	sigs := make(chan os.Signal, 1)
	go func() {
		for _ = range ticker.C {
			res, err := http.Get("http://app1.service.consul:8080/")
			if nil == err {
				b, err := ioutil.ReadAll(res.Body)
				if nil != err {
					println("Error:", err.Error())
				} else {
					println("Success:", string(b))
				}
			} else {
				println("error", err.Error())
			}
		}
	}()
	time.Sleep(1600 * time.Millisecond)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	r := <-sigs
	println("Received signal", r.String())
}
