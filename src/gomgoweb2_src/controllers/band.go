package controllers

import (
//"fmt"
"html/template"
"net/http"
"gomgoweb2_src"
"labix.org/v2/mgo/bson"
)

func BandAdd (r http.ResponseWriter, rq *http.Request) {
	locations := gomgoweb2_src.GetAll(gomgoweb2_src.LOCATION_COL)
	t, err := template.ParseFiles("src/gomgoweb2_src/views/band/add.html")
	if err != nil {
		panic(err)
	} else {
		t.Execute(r, locations)
	}	
}

func BandVerify (r http.ResponseWriter, rq *http.Request) {
	name := rq.FormValue("name")
	locType := rq.FormValue("loctype")
	var locationId bson.ObjectId
	errorString := "no errors"
	switch locType {
		case "existing":
			if rq.FormValue("location_id") == "" {
				errorString = "No location was selected"
			} else {
				locationId = gomgoweb2_src.ToObjectId(rq.FormValue("location_id"))
			}
			break
		case "new":
			if rq.FormValue("country") != "" {
				locationId = gomgoweb2_src.GenerateId()
				location := gomgoweb2_src.Location{
					City: rq.FormValue("city"),
					State: rq.FormValue("state"),
					Country: rq.FormValue("country"),
					}
				doc := gomgoweb2_src.MyDoc{ Id: locationId, Value: bson.M{"City": location.City, "State": location.State, "Country": location.Country}}	
				err := gomgoweb2_src.AddDoc(doc, gomgoweb2_src.LOCATION_COL)
				if err != nil {
					errorString = "error on location add: " + err.Error()
				}
			} else {
				errorString = "Country is required"
			}
			break
	}
	if errorString == "no errors" {
	id := gomgoweb2_src.GenerateId()
	band := gomgoweb2_src.Band{ Name: name, LocationId: locationId, Albums: []gomgoweb2_src.Album{} }
	doc := gomgoweb2_src.MyDoc{ Id: id, Value: bson.M{"Name": band.Name, "LocationId": band.LocationId, "Albums": band.Albums} }
	err := gomgoweb2_src.AddDoc(doc, gomgoweb2_src.BAND_COL)
	if err != nil {
		errorString = "error on band add: " + err.Error()
	} 	
	t, err := template.ParseFiles("src/gomgoweb2_src/views/band/verify.html")
	if err != nil {
		panic(err)
	} else {
		t.Execute(r, errorString)
	}	
}
}
			
											
