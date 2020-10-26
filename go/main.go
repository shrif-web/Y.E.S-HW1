package main

import (
	"crypto"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/rs/cors"
)

type sha256RequestBody struct {
	First  string `json:"first"`
	Second string `json:"second"`
}

type sha256ResponceBody struct {
	Result string `json:"result"`
}

func handleRequests() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home_page)
	mux.HandleFunc("/go/sha256", sha256)
	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":8888", handler))

}

func main() {
	handleRequests()
}

func home_page(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "please use a POST method", http.StatusNotFound)
}
func sha256(w http.ResponseWriter, r *http.Request) {
	fmt.Print()

	switch r.Method {
	case "POST":
		b, err := ioutil.ReadAll(r.Body)
		var operands sha256RequestBody
		err = json.Unmarshal(b, &operands)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		first, err := strconv.Atoi(operands.First)
		second, err := strconv.Atoi(operands.Second)
		result := []byte(strconv.Itoa(first + second))
		hasher := crypto.SHA256.New()
		hasher.Write(result)

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Accept", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(sha256ResponceBody{hex.EncodeToString(hasher.Sum(nil))})
	default:
		http.Error(w, "please use a POST method", http.StatusMethodNotAllowed)
	}
}
