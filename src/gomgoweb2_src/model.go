package gomgoweb2_src

import (
"fmt"
"labix.org/v2/mgo"
"labix.org/v2/mgo/bson"
"strings"
)

func ToObjectId(in string) bson.ObjectId {
	var result bson.ObjectId

	id1 := strings.Replace(in, "ObjectIdHex(", "", -1)
	id1 = strings.Replace(id1, ")", "", -1)
	id1 = strings.Replace(id1, "\"", "", -1)
	id1 = strings.TrimSpace(id1)
	result = bson.ObjectIdHex(id1)
	return result
}

func GetDB() (*mgo.Database, *mgo.Session) {
	session, err := mgo.Dial(DATABASE_DSN)
	if err != nil {
		panic(err.Error())
	}
	database := session.DB(DATABASE)
	return database, session
}

type MyDoc struct {
	Id    bson.ObjectId `bson:"_id"`
	Value bson.M        `bson:"Values"`
}

type Band struct {
	Name       string        `bson:"Name"`
	LocationId bson.ObjectId `bson:"LocationId"`
	Albums     []Album       `bson:"Albums"`
}

type Album struct {
	Name    string        `bson:"AlbumName"`
	Year    int           `bson:"Year"`
	GenreId bson.ObjectId `bson:"GenreId"`
}

type Genre struct {
	Name string `bson:"Name"`
}

type Location struct {
	City    string `bson:"City"`
	State   string `bson:"State"`
	Country string `bson:"Country"`
}

func GenerateId() bson.ObjectId {
	id := bson.NewObjectId()
	return id
}

func GetAll(docType string) []MyDoc {
	database, session := GetDB()
	defer session.Close()
	collection := database.C(docType)
	var results []MyDoc
	collection.Find(nil).All(&results)
	return results
}

func AddDoc(doc MyDoc, docType string) error {
	database, session := GetDB()
	defer session.Close()
	collection := database.C(docType)
	err := collection.Insert(doc)
	return err
}

func GetDoc(id bson.ObjectId, docType string) MyDoc {
	database, session := GetDB()
	defer session.Close()
	collection := database.C(docType)
	var doc MyDoc
	collection.Find(bson.M{"_id": id}).One(&doc)
	return doc
}

func (this *Album) GetGenreName() string {
	database, session := GetDB()
	defer session.Close()
	collection := database.C(GENRE_COL)
	var doc MyDoc
	collection.Find(bson.M{"_id": this.GenreId}).One(&doc)
	result := doc.Value["Name"].(string)
	return result

}

func GetGenreName(id bson.ObjectId) string {
	database, session := GetDB()
	defer session.Close()
	collection := database.C(GENRE_COL)
	var doc MyDoc
	collection.Find(bson.M{"_id": id}).One(&doc)
	result := doc.Value["Name"].(string)
	return result
}

func GetBandsByGenre(genreId bson.ObjectId) []MyDoc {
	fmt.Println("GetBandsByGenre received:", genreId)
	database, session := GetDB()
	defer session.Close()
	collection := database.C(BAND_COL)
	col2 := database.C(GENRE_COL)
	var doc MyDoc
	var docs []MyDoc
	//	params := bson.M{"$elemMatch": []Album{{GenreId: genreId}}}

	err := collection.Find(bson.M{"Values.Albums": bson.M{"$elemMatch": bson.M{"GenreId": genreId}}}).All(&docs)
	col2.Find(bson.M{"_id": genreId}).One(&doc)
	fmt.Println("Genre name =", doc.Value["Name"].(string))
	if err != nil {
		fmt.Println("Find failed:", err)
	} else {
		if docs == nil {
			fmt.Println("Returned empty set")
		} else {
			fmt.Println("Found", len(docs))
		}
	}
	return docs
}
func (this *MyDoc) LocToString() string {
	var cityStr, stateStr string
	m := this.Value
	var location Location
	for key, value := range m {
		switch key {
		case "City":
			location.City = value.(string)
			break
		case "State":
			location.State = value.(string)
			break
		case "Country":
			location.Country = value.(string)
			break
		}
	}

	if location.City != "" {
		cityStr = location.City
	} else {
		cityStr = "(city)"
	}
	if location.State != "" {
		stateStr = location.State
	} else {
		stateStr = "(state/province)"
	}

	result := fmt.Sprintf("%s, %s %s", cityStr, stateStr, location.Country)
	return result
}

func (this *MyDoc) GetLocation() string {
	database, session := GetDB()
	defer session.Close()
	collection := database.C(LOCATION_COL)
	id := this.Value["LocationId"].(bson.ObjectId)
	var doc MyDoc
	collection.Find(bson.M{"_id": id}).One(&doc)
	locString := doc.LocToString()
	return locString
}

func (this *MyDoc) AddAlbum(album Album) error {
	database, session := GetDB()
	defer session.Close()
	collection := database.C(BAND_COL)
	/*band := Band{Name: this.Value["Name"].(string),
		LocationId: this.Value["LocationId"].(bson.ObjectId)}
	band.Albums = []Album{}
	if this.Value["Albums"] != nil {
		for _, a := range this.Value["Albums"].([]interface{}) {
			x := a.(bson.M)

			q := Album{Name: x["AlbumName"].(string), Year: x["Year"].(int),
				GenreId: x["GenreId"].(bson.ObjectId)}
			band.Albums = append(band.Albums, q)
			fmt.Println("Found album", q.Name)
		}
	} else {
		band.Albums = []Album{}
	}

	band.Albums = append(band.Albums, album)
	doc := MyDoc{Id: this.Id, Value: bson.M{"Name": band.Name,
		"LocationId": band.LocationId, "Albums": band.Albums}}
	err := collection.Update(bson.M{"_id": doc.Id}, doc)
*/
err := collection.Update(bson.M{"_id": this.Id}, bson.M{"$push": bson.M{"Values.Albums": album}})
	return err
}

func (this MyDoc) GetAlbums() []Album {
	var a []Album
	q := this.Value["Albums"].([]interface{})

	for _, c := range q {
		z := c.(bson.M)
		y := Album{Name: z["AlbumName"].(string),
			Year:    z["Year"].(int),
			GenreId: z["GenreId"].(bson.ObjectId)}
		a = append(a, y)
	}
	return a
}


