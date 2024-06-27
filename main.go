package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST Request Successful\n")
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	emailAddress := r.FormValue("emailAddress")
	phoneNumber := r.FormValue("phoneNumber")
	fmt.Fprintf(w, "First Name: %v\n", firstName)
	fmt.Fprintf(w, "Last Name: %v\n", lastName)
	fmt.Fprintf(w, "Email Address: %v\n", emailAddress)
	fmt.Fprintf(w, "Phone Number: %v\n", phoneNumber)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
	}

	fmt.Fprintf(w, "Hello!")
}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server http://127.0.0.1:8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
