package main

import (
	"bytes"
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

func addItem(name, value string) error {
	backendURL := os.Getenv("BACKEND_URL") + "/items"
	item := fmt.Sprintf(`{"name":"%s","value":"%s"}`, name, value)
	resp, err := http.Post(backendURL, "application/json", bytes.NewBuffer([]byte(item)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	return err
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		value := r.FormValue("value")
		err := addItem(name, value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	items, err := getItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `
        <html>
        <body>
            <h1>Items</h1>
            <pre>%s</pre>
            <h2>Add New Item</h2>
            <form method="post">
                <label for="name">Name:</label><br>
                <input type="text" id="name" name="name"><br>
                <label for="value">Value:</label><br>
                <input type="text" id="value" name="value"><br><br>
                <input type="submit" value="Add Item">
            </form>
        </body>
        </html>
    `, items)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Frontend server running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
