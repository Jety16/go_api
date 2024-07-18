package main

import ("fmt"; "net/http")


func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, WOrLd!")
}

func main() {
	// Route setup
	http.HandleFunc("/", handler)
	fmt.Println("Starting Server at port 8080...")
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		fmt.Println("Failed to start server: ", err)
	}
}