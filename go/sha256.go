package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type DualNum struct{
	A0 int `db:"a0" json:"a0,string"`
	A1 int `db:"a1" json:"a1,string"`
}

// request example :
//{
//    "a0": "2",
//    "a1": "3"
//}

func main() {
	http.HandleFunc("/", handle_func)
	http.ListenAndServe(":8080", nil)
}

func handle_func(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/go/sha256" {
		http.Error(w, "fuck you", http.StatusNotFound)
		return
	}
	enableCors(&w)
	switch r.Method {
	case "POST":
		var nums DualNum
		err := json.NewDecoder(r.Body).Decode(&nums)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println("sum result = ", nums.A0 + nums.A1)
		hasher := sha1.New()
		hasher.Write([]byte(strconv.Itoa(nums.A0 + nums.A1)))
		//sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil)) //todo which one ?
		sha := hex.EncodeToString(hasher.Sum(nil))
		fmt.Fprintf(w, "Result: %s", sha)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		//io.WriteString(w, "Sorry, only GET and POST methods are supported.")
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}