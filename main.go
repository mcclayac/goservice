package main

import "net/http"

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n\nThese Nutz \n\n goService\n\nDeveloper = mcclayac\n\n mcclay-goapp"))
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":9090", nil)
}
