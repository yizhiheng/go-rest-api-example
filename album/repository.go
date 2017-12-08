package album

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Repository ...
type Repository struct{}

// SERVER the DB server
const SERVER = "localhost:27017"

// DBNAME the name of the DB instance
const DBNAME = "musicstore"

// CollectionName the name of the document
const CollectionName = "albums"

// GetAlbums returns the list of Albums
func (r Repository) GetAlbums() Albums {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(CollectionName)
	results := Albums{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}
	return results
}

// AddAlbum inserts an Album in the DB
func (r Repository) AddAlbum(album Album) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	album.ID = bson.NewObjectId()
	session.DB(DBNAME).C(CollectionName).Insert(album)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// UpdateAlbum updates an Album in the DB (not used for now)
func (r Repository) UpdateAlbum(album Album) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	album.ID = bson.NewObjectId()
	session.DB(DBNAME).C(CollectionName).UpdateId(album.ID, album)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// DeleteAlbum deletes an Album (not used for now)
func (r Repository) DeleteAlbum(id string) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		return "NOT FOUND"
	}
	// Grab id
	oid := bson.ObjectIdHex(id)
	// Remove user
	if err = session.DB(DBNAME).C(CollectionName).RemoveId(oid); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}
	// Write status
	return "OK"
}
