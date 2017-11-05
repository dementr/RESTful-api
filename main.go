package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/gorilla/handlers"
	// "unicode"
	"strconv"
	"os"
	"fmt"
)

// https://4gophers.ru/articles/avtorizaciya-v-go-s-ispolzovaniem-jwt/#.WfofOddl8uU

func main() {

	router := mux.NewRouter()

	router.Handle("/", http.FileServer(http.Dir("./views/")))

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
										http.FileServer(http.Dir(".static/"))))

	router.Handle("/status", StatusHandler).Methods("GET")

	router.Handle("/products", ProductsHandler).Methods("GET")

	router.Handle("/product/{id}", ProductIdHandler).Methods("GET")

	router.Handle("/products/{slug}/feedback", AddFeedbackHandler).Methods("POST")

	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, router))
}

var NotImplemented = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Not Implemented"))
		})

type Product struct {
	Id int64
	Name string
	Slug string
	Description string
}

var	products = []Product {
		Product{Id: 1, Name: "Hover Shooters", Slug: "hover-shooters", Description : "Shoot your way to the top on 14 different hoverboards"},
		Product{Id: 2, Name: "Ocean Explorer", Slug: "ocean-explorer", Description : "Explore the depths of the sea in this one of a kind"},
		Product{Id: 3, Name: "Dinosaur Park", Slug : "dinosaur-park", Description : "Go back 65 million years in the past and ride a T-Rex"},
		Product{Id: 4, Name: "Cars VR", Slug : "cars-vr", Description: "Get behind the wheel of the fastest cars in the world."},
		Product{Id: 5, Name: "Robin Hood", Slug: "robin-hood", Description : "Pick up the bow and arrow and master the art of archery"},
		Product{Id: 6, Name: "Real World VR", Slug: "real-world-vr", Description : "Explore the seven wonders of the world in VR"}}

var StatusHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("API is up and running"))
	})

var ProductsHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request){
		payload, _ := json.Marshal(products)

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(payload))
	})

var ProductIdHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request){
		var product Product
		vars := mux.Vars(r)
		var idstr = vars["id"]

		var id int64

		id, err := strconv.ParseInt(idstr, 10, 0)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, p := range products {
			if p.Id == id {
				product = p
			}
		}

		w.Header().Set("Content-Type", "application/json")
		if product.Id != 0 {
			payload, _ := json.Marshal(product)
			w.Write([]byte(payload))
		} else {
			w.Write([]byte("Product Not Found"))
		}
	})

var AddFeedbackHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request){
		var product Product
		vars := mux.Vars(r)
		slug := vars["slug"]

		fmt.Println(slug)

		for _, p := range products {
			if p.Slug == slug {
				product = p
			}
		}

		w.Header().Set("Content-Type", "application/json")
		if product.Slug != "" {
			payload, _ := json.Marshal(product)
			w.Write([]byte(payload))
		} else {
			w.Write([]byte("Product Not Found"))
		}
	})