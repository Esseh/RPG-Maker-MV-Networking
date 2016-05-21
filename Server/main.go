package main

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
)



// ==============================
// Refer to API.go for specifics.
// ==============================
func init(){
	// Initialize Maps Container
	Maps = make(map[int64](map[string]Player))
	// Initialize Julien Schmidt Router
	router := httprouter.New()
	// Begin the Time Out Handler
	go TimeOut()
	// Logs the Player In
	router.POST("/playerLoggedIn/:uuid/:mapid/:x/:y",PlayerLogin)
	// Logs the Player Off 
	router.POST("/playerLoggedOff/:uuid/:mapid",PlayerLogout)
	// Adds to Action Queue
	router.POST("/aq/:uuid/:mapid/:actionid/:x/:y",AQ)
	// Gets the Map Information For Parsing By Player
	router.GET("/map/:mapid",Map)
	// Generates a UUID
	router.GET("/uuid/",UUID)
	// Make Appengine use the Julien Schmidt Router
	http.Handle("/", router)
}