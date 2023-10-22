package main

import (
	"net/http"
	"time"
)

func main() {
	router := http.DefaultServeMux
	router.HandleFunc("/fast", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("fast response"))
	})
	router.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		// 对该endpoint单独设置超时
		rc := http.NewResponseController(w)
		rc.SetReadDeadline(time.Now().Add(6 * time.Second))
		rc.SetWriteDeadline(time.Now().Add(6 * time.Second))

		time.Sleep(5 * time.Second)
		w.Write([]byte("slow response"))
	})

	server := http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
