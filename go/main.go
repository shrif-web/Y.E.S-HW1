
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