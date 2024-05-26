package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getItems() ([]byte, error) {
	backendURL := os.Getenv("BACKEND_URL") + "/items"
	resp, err := http.Get(backendURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func handler(w http.ResponseWriter, r *http.Request) {
	items, err := getItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "<h1>Items</h1><pre>%s</pre>", items)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Frontend server running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
