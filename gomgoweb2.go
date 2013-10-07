package main

import (
"fmt"
"net/http"
"gomgoweb2_src/controllers"
)

func main() {
 	fmt.Println("Program starting at http://localhost:8000")

 	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
 	http.HandleFunc("/", controllers.HomeIndex)
 	http.HandleFunc("/home/genrelist", controllers.GenreList)
 	http.HandleFunc("/home/bygenre/", controllers.ByGenre)
 	http.HandleFunc("/band/add", controllers.BandAdd)
 	http.HandleFunc("/band/verify", controllers.BandVerify)
 	http.HandleFunc("/album/index/", controllers.AlbumIndex)
 	http.HandleFunc("/album/add/", controllers.AlbumAdd)
 	http.HandleFunc("/album/verify/", controllers.AlbumVerify)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	} else {
		fmt.Println("Server running")
	}
}
