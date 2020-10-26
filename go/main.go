package main

import (
	"crypto"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"bufio"

	"github.com/rs/cors"
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

func handleRequests() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/go/sha256", sha256)
	mux.HandleFunc("/go/write", write)
	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":8888", handler))

}

func main() {
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "please use a POST method", http.StatusNotFound)
}

func readFile(line_number int) string{
	file, err := os.Open("test.txt")
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
 
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	var linesNumber int = len(txtlines)
 
	file.Close()

	if line_number > linesNumber {
		return "Error"
	} else {
		return txtlines[line_number - 1]
	}
}

func write(w http.ResponseWriter, r *http.Request) {
	fmt.Println("._. in Go file in write function")

	number, err := strconv.Atoi(r.URL.Query().Get("number"))

	var line string = readFile(number)
	fmt.Println("--line is :", line)

	_ = err

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

		_ = json.NewEncoder(w).Encode(sha256ResponseBody{hex.EncodeToString(hasher.Sum(nil))})
	default:
		http.Error(w, "please use a POST method", http.StatusMethodNotAllowed)
	}
}
