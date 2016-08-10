package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// CollectionController represents the controller for operating on the User resource
	CollectionController struct {
		session *mgo.Session
	}
)

// NewCollectionController provides a reference to a CollectionController with provided Mongo session
func NewCollectionController(s *mgo.Session) *CollectionController {
	return &CollectionController{s}
}

// GetCollections retrieves a list of Title resources
// NOTE: I made an assumption that Titles was lowercase.  I used this to see the list of Collection Names.
func (cc CollectionController) GetCollections(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	session := cc.session.Copy()
	defer session.Close()

	// get list
	collections, err := session.DB("dev-challenge").CollectionNames()
	if err != nil {
		w.WriteHeader(404)
		return
	}

	//marshal to json
	tj, _ := json.Marshal(collections)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", tj)
}

// GetMap retrieves a key value map of the collection
// I used this to get an idea of the layout of the document.
func (cc CollectionController) GetMap(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	session := cc.session.Copy()
	defer session.Close()

	// Grab Id
	//collectionName := p.ByName("collectionName")

	// get map
	var d bson.M
	err := session.DB("dev-challenge").C("Titles").Find(nil).One(&d)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	//marshal to json
	tj, _ := json.Marshal(d)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", tj)
}
