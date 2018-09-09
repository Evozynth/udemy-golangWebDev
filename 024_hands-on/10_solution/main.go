package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("./starting-files"))
	http.Handle("/", fs)
	http.Handle("/public/", http.StripPrefix("/public", fs))
	http.ListenAndServe(":8080", nil)
}
