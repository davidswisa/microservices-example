package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func gui() func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(wrt http.ResponseWriter, req *http.Request) {
		content, err := ioutil.ReadFile("index.html")
		if err != nil {
			log.Fatal(err)
		}
		wrt.Write([]byte(content))
	})
}

func imagesHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./images/restaurant-circular.svg")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./images/favicon.ico")
}

func main() {

	// Add handle func for producer.
	http.HandleFunc("/", gui())
	http.HandleFunc("/images/restaurant-circular.svg", imagesHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)

	// Run the web server.
	fmt.Println("starting producer-api...")
	log.Fatal(http.ListenAndServe(":8084", nil))
}
