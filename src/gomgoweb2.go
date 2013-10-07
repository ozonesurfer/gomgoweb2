package main 

import (
"fmt"
"net/http"
"gomgoweb2_src/controllers"
)

func main() {
 	fmt.Println("Program starting")
 	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources")))) 
	http.HandleFunc("/", controllers.HomeIndex)
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	} else {
		fmt.Println("Server running")
	}	
}

