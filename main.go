package main

import "net/http"

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n\nThese Nutz \n\n goService"))
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":9090", nil)
}
