package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Hello, Kubernetes!")
		log.Println("connection from", req.RemoteAddr)

		names := make([]string, 0, len(req.Header))
		for k := range req.Header {
			names = append(names, k)
		}
				for _, name := range names {
			log.Printf("%s: %q\n", name, req.Header[name])
		}

		fmt.Fprintln(w, req.Proto)
		log.Println()
		log.Println()
	})

	http.ListenAndServe(":8080", nil)
}
