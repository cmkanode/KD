package main

import (
	"fmt"
	"net/http"
	"webserver/controllers"

	// Third Party packages
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func getSession() *mgo.Session {
	// Connect to mongo
	s, err := mgo.Dial("mongodb://readonly:turner@ds043348.mongolab.com:43348/dev-challenge")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}

func main() {
	r := httprouter.New()

	// Get a TitleController instance
	tc := controllers.NewTitleController(getSession())
	// Get a CollectionController instance
	cc := controllers.NewCollectionController(getSession())

	// Add a handler on /test
	r.GET("/alltitles", tc.GetTitles)
	r.GET("/title/search/:searchterm", tc.GetTitle)
	r.GET("/title/:titleid", tc.GetTitleByID)
	r.GET("/collections", cc.GetCollections)
	r.GET("/collections/map", cc.GetMap)

	r.GET("/ping", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// still alive message
		fmt.Fprint(w, "I'm alive!\n")
	})

	http.ListenAndServe(":8080", r)
}
