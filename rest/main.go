package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	P_Id    int32   `json:"id"`
	P_Name  string  `json:"name"`
	P_Price float32 `json:"price"`
}

var listProducts []Product

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}
func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi! You came to Home! :)")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error in Reading Request", err)
	}
	product := Product{}
	err = json.Unmarshal(reqBody, &product)
	if err != nil {
		fmt.Println("Error in Unmarshalling", err)
	}
	listProducts = append(listProducts, product)
	fmt.Println(listProducts)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}
func ReadAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(listProducts)
}
func ReadOne(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println("Error in Conversion", err)
	}
	product := Product{}
	for _, prd := range listProducts {
		if int32(id) == prd.P_Id {
			product.P_Id = prd.P_Id
			product.P_Name = prd.P_Name
			product.P_Price = prd.P_Price
		}
	}
	json.NewEncoder(w).Encode(product)
}
func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi! You came to Home! :)")
	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println("Error in Conversion", err)
	}
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Error in Reading Request", err)
	}
	product := Product{}
	err = json.Unmarshal(reqBody, &product)
	if err != nil {
		fmt.Println("Error in Unmarshalling", err)
	}
	for i, prd := range listProducts {
		if prd.P_Id == int32(id) {
			listProducts[i].P_Name = product.P_Name
			listProducts[i].P_Price = product.P_Price
			json.NewEncoder(w).Encode(listProducts[i])
			break
		}
	}

	fmt.Println(listProducts)
}
func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi! You came to Home! :)")
	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println("Error in Conversion", err)
	}

	for i, prd := range listProducts {
		if prd.P_Id == int32(id) {
			listProducts = append(listProducts[:i], listProducts[i+1:]...)
			json.NewEncoder(w).Encode(listProducts[i])
			break
		}
	}
	fmt.Println(listProducts)
	json.NewEncoder(w).Encode(listProducts)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/create", Create).Methods("POST")
	myRouter.HandleFunc("/read/{id}", ReadOne).Methods("GET")
	myRouter.HandleFunc("/readall", ReadAll).Methods("GET")
	myRouter.HandleFunc("/update/{id}", Update).Methods("PUT")
	myRouter.HandleFunc("/delete/{id}", Delete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
func main() {
	handleRequests()
}
