package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Hello World! Running on container ID %s", name)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Web server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
