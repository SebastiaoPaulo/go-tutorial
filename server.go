package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
)

func main() {

	stripe.Key = "sk_test_51PvwFn09vYM5hXNFkquFiGphB3gsiP4fXJviUvPYWMnNhaVRpv69HcUhmfqZkQ1mIakRxPuoKTedi4YkflnLT8Zh00rKAxQo6o"
	http.HandleFunc("/hello-world", handleHelloWorld)
	http.HandleFunc("/payment", handlePayment)

	log.Println("Server is running at localhost:3333...")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func handleHelloWorld(writer http.ResponseWriter, request *http.Request) {

	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	response := []byte("Hello World!!!")

	_, err := writer.Write(response)

	if err != nil {
		log.Fatal(err)
	}
}

func handlePayment(writer http.ResponseWriter, request *http.Request) {

	if request.Method != "POST" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// That's the struct that will receive the Body from the request
	var reqBody struct {
		ProductId string `json:"productId"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}

	// Decoding the request Body to the struct
	err := json.NewDecoder(request.Body).Decode(&reqBody)

	// returning Bad Request Error case the decoding process didn't success
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Accessing data from the struct, which has the request's body values
	fmt.Println("Product ID: ", reqBody.ProductId)
	fmt.Println("First Name: ", reqBody.FirstName)
	fmt.Println("last Name: ", reqBody.LastName)

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(calculateOrderAmount(reqBody.ProductId)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	paymentIntent, err := paymentintent.New(params)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	fmt.Println(paymentIntent.ClientSecret)

	// Creating the response as a json object...
	// Define the struct of the json response
	var response struct {
		ClientSecret string `json:"clientSecret"`
	}

	// Assigning the value to the response's struct property
	response.ClientSecret = paymentIntent.ClientSecret

	// Declaring the buffer that will receive the data to respond
	var buf bytes.Buffer
	// Enconding the Json struct/data to the buffer
	err = json.NewEncoder(&buf).Encode(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	// Setting the response's content type
	writer.Header().Set("Content-Type", "application/json")

	// Copying the buf's data to Writer, which gives the response to the request
	_, err = io.Copy(writer, &buf)
	if err != nil {
		fmt.Println(err)
	}
}

func calculateOrderAmount(productId string) int64 {
	switch productId {
	case "product111":
		return 26000
	case "product112":
		return 30000
	case "product113":
		return 15000
	}
	return 0
}
