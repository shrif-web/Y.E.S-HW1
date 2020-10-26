package main

import (
	"bufio"
	"crypto"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rs/cors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type sha256RequestBody struct {
	First  string `json:"first"`
	Second string `json:"second"`
}

type sha256ResponseBody struct {
	Result string `json:"result"`
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	handleRequests("", 8888)
}

func handleRequests(domain string, port int) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/go/sha256", sha256)
	mux.HandleFunc("/go/write", write)

	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", domain, port), handler))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		log.Printf("new session initiated with client:[%s]", r.RemoteAddr)
		http.ServeFile(w, r, "./../front")
	} else {
		log.Printf("client [%s] tried to reach unsupported path , refusing ", r.RemoteAddr)
		http.Error(w, "the page you are looking for is not yet backed!", http.StatusNotFound)
	}
}

func readFile(lineNumber int) (string, error) {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var textlines []string
	for scanner.Scan() {
		textlines = append(textlines, scanner.Text())
	}
	linesNumber := len(textlines)

	if lineNumber > linesNumber || lineNumber < 1 {
		return "", errors.New("line out of bound exception")
	} else {
		return textlines[lineNumber-1], nil
	}
}

func write(w http.ResponseWriter, r *http.Request) {

	number, _ := strconv.Atoi(r.URL.Query().Get("number"))

	line,err := readFile(number)
	if err != nil {
		http.Error(w, "please enter a valid number", http.StatusMethodNotAllowed)
		return
	}
	_, _ = fmt.Fprintf(w, line)
}

func sha256(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		b, err := ioutil.ReadAll(r.Body)
		var operands sha256RequestBody
		err = json.Unmarshal(b, &operands)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		first, err := strconv.Atoi(operands.First)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		second, err := strconv.Atoi(operands.Second)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		result := []byte(strconv.Itoa(first + second))
		hasher := crypto.SHA256.New()
		hasher.Write(result)

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Accept", "application/json")
		w.WriteHeader(http.StatusCreated)

		_ = json.NewEncoder(w).Encode(sha256ResponseBody{hex.EncodeToString(hasher.Sum(nil))})
	default:
		http.Error(w, "please use a POST method", http.StatusMethodNotAllowed)
	}
}
