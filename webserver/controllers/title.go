package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webserver/models"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// TitleController represents the controller for operating on the User resource
	TitleController struct {
		session *mgo.Session
	}
)

// NewTitleController provides a reference to a TitleController with provided Mongo session
func NewTitleController(s *mgo.Session) *TitleController {
	return &TitleController{s}
}

// GetTitles retrieves a list of Title resources
func (tc TitleController) GetTitles(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	session := tc.session.Copy()
	defer session.Close()

	// get list
	titles := []models.Title{}

	err := session.DB("dev-challenge").C("Titles").Find(nil).All(&titles)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	//marshal to json
	tj, _ := json.Marshal(titles)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", tj)
}

// GetTitle retrieves an individual Title resource
func (tc TitleController) GetTitle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab search field
	TitleName := p.ByName("searchterm")

	session := tc.session.Copy()
	defer session.Close()

	// Stub Title
	t := []models.Title{}

	// Fetch  by using a regular expression.
	// NOTE:  Not fond of concatenating the regex string.  Need to look into other options.
	if err := session.DB("dev-challenge").C("Titles").Find(bson.M{"TitleName": bson.RegEx{"(" + TitleName + "\\s*)", "i"}}).All(&t); err != nil {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	tj, _ := json.Marshal(t)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", tj)
}

/*
 * On a previous iteration, I had this working for pulling the record by TitleId.
 * I'm not certain how it changed.  I may not need this method, if I return the 
 * search results to the client, and have the entire document available.
 */
// GetTitleByID retrieves an individual Title resource
func (tc TitleController) GetTitleByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab search field
	TitleId := p.ByName("titleid")

	session := tc.session.Copy()
	defer session.Close()

	// Stub Title
	t := models.Title{}

	// Fetch by using Find.  There is an error.  Does it return an error on a null result?
	if err := session.DB("dev-challenge").C("Titles").Find(bson.M{"TitleId": TitleId}).One(&t); err != nil {
		w.WriteHeader(418)
		return
	}

	// Marshal provided interface into JSON structure
	tj, _ := json.Marshal(t)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", tj)
}
