package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type DualNum struct{
	A0 int `json:"first"`
	A1 int `json:"second"`
}

func main() {
	http.HandleFunc("/", reqHandler)
	http.HandleFunc("/go/sha256", sha256Response)
	http.HandleFunc("/go/write", lineSenderResponse)
	http.ListenAndServe(":8080", nil)
}

func reqHandler(w http.ResponseWriter, r * http.Request){
	if r.URL.Path != "/go/sha256" && r.URL.Path != "/go/write" {
		http.Error(w, "404 Error", http.StatusNotFound)
		return
	}
}

func lineSenderResponse(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
	//todo
}

func sha256Response(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case "POST":
		bod, err := ioutil.ReadAll(r.Body)
		var nums DualNum
		err = json.Unmarshal(bod, &nums)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println("sum result = ", nums.A0 + nums.A1)
		hasher := sha256.New()
		hasher.Write([]byte(strconv.Itoa(nums.A0 + nums.A1)))
		sha := hex.EncodeToString(hasher.Sum(nil))
		fmt.Fprintf(w, "Result: %s", sha)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}