package main

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
)



// ==============================
// Refer to API.go for specifics.
// ==============================
func init(){
	// Initialize Containers
	InitializeMaps()
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
	/*=============
	 MetaBlob stuff
	==============*/
	router.POST("/MetaBlob/Post/:owner/:purpose/:client/:primarykey/:secondarykey/:input",MetaBlobEntry)
	router.POST("/MetaBlob/Delete/:owner/:purpose/:client/:primarykey/:secondarykey",MetaBlobDeleteEntry)
	router.POST("/MetaBlob/Delete/:owner/:purpose/:client/:primarykey",MetaBlobDeletePrimary)
	router.POST("/MetaBlob/Delete/:owner/:purpose/:client",MetaBlobDeleteClient)
	router.POST("/MetaBlob/Delete/:owner/:purpose",MetaBlobDeletePurpose)
	router.POST("/MetaBlob/Delete/:owner",MetaBlobDeleteOwner)
	router.GET("/MetaBlob/Get/:owner",BlobOwners)
	router.GET("/MetaBlob/Get/:owner/:purpose",BlobPurpose)
	router.GET("/MetaBlob/Get/:owner/:purpose/:client",BlobPrimaryKey)
	router.GET("/MetaBlob/Get/:owner/:purpose/:client/:primarykey",BlobClient)
	router.GET("/MetaBlob/Get/:owner/:purpose/:client/:primarykey/:secondarykey",BlobSecondaryKey)
	//==============
	// Retrieves the Client Server Switch States
	router.GET("/switches/",GetSwitches)
	// Retrieves the Client server Variables States
	router.GET("/variables/",GetVariables)
	// Assigns to a switch state
	router.POST("/switches/:id/:value",SetSwitch)
	// Assigns to a variable state
	router.POST("/variables/:id/:value",SetVariable)
	http.Handle("/", router)
}