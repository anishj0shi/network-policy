package main

import "net/http"

func main() {
	http.HandleFunc("/hello", respond)

	http.ListenAndServe(":8080", nil)
}

func respond(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello World!!!"))
}
