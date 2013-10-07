package controllers

import (
"fmt"
"gomgoweb2_src"
"html/template"
"net/http"
)
type Params struct {
Bands []gomgoweb2_src.MyDoc
Title string
}


func HomeIndex (r http.ResponseWriter, rq *http.Request) {
	docs := gomgoweb2_src.GetAll(gomgoweb2_src.BAND_COL)
	t, err := template.ParseFiles("src/gomgoweb2_src/views/home/index.html")
	if err != nil {
		panic(err)
	} else {
	    t.Execute(r, Params{Bands: docs, Title: "My CD Catalog"})
    }
}

func GenreList (r http.ResponseWriter, rq *http.Request) {
	genres := gomgoweb2_src.GetAll(gomgoweb2_src.GENRE_COL)
	t, err := template.ParseFiles("src/gomgoweb2_src/views/home/genrelist.html")
	if err != nil {
		panic(err)
	} else {
		t.Execute(r, genres)
	}
}

func ByGenre (r http.ResponseWriter, rq *http.Request) {
	values := rq.URL.Query()
	id := values["id"][0]
	genreId := gomgoweb2_src.ToObjectId(id)
	genreName := gomgoweb2_src.GetGenreName(genreId)
	title := fmt.Sprintf("%s Albums", genreName)
	bands := gomgoweb2_src.GetBandsByGenre(genreId)
		t, err := template.ParseFiles("src/gomgoweb2_src/views/home/bygenre.html")
	if err != nil {
		panic(err)
	} else {
	    t.Execute(r, Params{Bands: bands, Title: title})
    }
}
