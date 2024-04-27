package main

import "net/http"

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Anpadoma!"))
	}
	http.ListenAndServe("localhost:8080", http.HandlerFunc(handler))
}
