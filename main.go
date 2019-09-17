package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type product struct {
	ID        int
	Name      string
	Category  string
	CreatedAt string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Busble Api with Golang Sample")
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	myProduct := []product{}

	myProduct = append(myProduct, product{
		ID:        1,
		Name:      "IMac Pro 2019",
		Category:  "Desktop",
		CreatedAt: formatted,
	})

	myProduct = append(myProduct, product{
		ID:        2,
		Name:      "MacBook Pro 2019",
		Category:  "Labtop",
		CreatedAt: formatted,
	})

	myProduct = append(myProduct, product{
		ID:        3,
		Name:      "IPhone 11",
		Category:  "Mobile",
		CreatedAt: formatted,
	})

	myProduct = append(myProduct, product{
		ID:        4,
		Name:      "IPhone 11 Pro Max",
		Category:  "Mobile",
		CreatedAt: formatted,
	})
	json.NewEncoder(w).Encode(myProduct)
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/products", getProduct)
	http.ListenAndServe(getPort(), nil)
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("No Port In Heroku" + port)
	}
	return ":" + port
}

func main() {
	handleRequest()
}
