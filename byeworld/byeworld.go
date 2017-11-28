package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	curTime := time.Now().Format("02.01.2006 15:04:05")

	fmt.Fprintf(w, "Hello world, the time is %s", curTime)
}
