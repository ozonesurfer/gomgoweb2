package controllers

import (
"net/http"
"gomgoweb2_src"
"html/template"
"strconv"
"labix.org/v2/mgo/bson"
"strings"
)

func AlbumIndex(r http.ResponseWriter, rq *http.Request) {
	value := rq.URL.Query()
	rawId := value["id"][0]
	bandId := gomgoweb2_src.ToObjectId(rawId)
	bandDoc := gomgoweb2_src.GetDoc(bandId, gomgoweb2_src.BAND_COL)
	t, err := template.ParseFiles("src/gomgoweb2_src/views/album/index.html")
	if err != nil {
		panic(err)
	} else {
		t.Execute(r, struct{Id string; Title string; Band gomgoweb2_src.MyDoc}{Id: rawId, Title: bandDoc.Value["Name"].(string), Band: bandDoc}) 
	}	
}

func AlbumAdd(r http.ResponseWriter, rq *http.Request) {
	value := rq.URL.Query()
	id := value["id"][0]
	genres := gomgoweb2_src.GetAll(gomgoweb2_src.GENRE_COL)
	t, err := template.ParseFiles("src/gomgoweb2_src/views/album/add.html")
	if err != nil {
		panic(err)
	} else {
		t.Execute(r, struct{Id string; Genres []gomgoweb2_src.MyDoc}{Id: id, Genres: genres})
	}		
}

func AlbumVerify(r http.ResponseWriter, rq *http.Request) {
	// implement
	rawId := gomgoweb2_src.ToObjectId(rq.URL.Query()["id"][0])
	id := rawId
	name := rq.FormValue("name")
	yearString := rq.FormValue("year")
	year, _ := strconv.Atoi(yearString)
	genreType := rq.FormValue("genretype")
	var genreId bson.ObjectId
	errorString := "no errors"
	switch genreType {
		case "existing":
			if rq.FormValue("genre_id") == "" {
				errorString = "No genre was selected"
			} else {
				genreId = gomgoweb2_src.ToObjectId(rq.FormValue("genre_id"))
			}
			break
		case "new":
			genre_name := strings.ToUpper(rq.FormValue("genre_name")) 
			if rq.FormValue("genre_name") == "" {
				errorString = "Genre name is required to create a genre"
			} else {
				database, session := gomgoweb2_src.GetDB()
				col := database.C(gomgoweb2_src.GENRE_COL)
				query := col.Find(bson.M{"Values.Name": genre_name})
				if q, _ := query.Count(); q == 0 {
					session.Close()
					genreId = gomgoweb2_src.GenerateId()
					genre := gomgoweb2_src.Genre{ Name: genre_name }
					doc := gomgoweb2_src.MyDoc{ Id: genreId, Value: bson.M{ "Name": genre.Name }}
					err := gomgoweb2_src.AddDoc(doc, gomgoweb2_src.GENRE_COL)
					if err != nil {
						errorString = "Genre add error: " + err.Error()
					} 
				} else {
					var genres []gomgoweb2_src.MyDoc
					query.All(&genres)
					genreId = genres[0].Id
					session.Close()
				}
			}
			break
		}
		
		if errorString == "no errors" {
			bandDoc := gomgoweb2_src.GetDoc(rawId, gomgoweb2_src.BAND_COL)
			album := gomgoweb2_src.Album{ Name: name, Year: year, GenreId: genreId }
			err := bandDoc.AddAlbum(album)
			if err != nil {
				errorString = "Album add error: " + err.Error()
			}
		}
		t, err := template.ParseFiles("src/gomgoweb2_src/views/album/verify.html")
	if err != nil {
		panic(err)
	} else {
		t.Execute(r, struct{Id string; Message string}{Id: id.String(), Message: errorString})
	}						 				
}