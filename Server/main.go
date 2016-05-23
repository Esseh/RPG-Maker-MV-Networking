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
	// Make an Entry in the MetaBlob
	router.POST("/MetaBlob/Post/:owner/:purpose/:client/:primarykey/:secondarykey/:input",MetaBlobEntry)
	// Delete an Entry in the MetaBlob
	router.POST("/MetaBlob/Delete/:owner/:purpose/:client/:primarykey/:secondarykey",MetaBlobDeleteEntry)
	// Delete All Entries Associated with a Primary Key
	router.POST("/MetaBlob/Delete/:owner/:purpose/:client/:primarykey",MetaBlobDeletePrimary)
	// Delete All Entries Associated with a Client
	router.POST("/MetaBlob/Delete/:owner/:purpose/:client",MetaBlobDeleteClient)
	// Delete All Entries Associated with a Purpose
	router.POST("/MetaBlob/Delete/:owner/:purpose",MetaBlobDeletePurpose)
	// Delete All Entries Associated with an Owner
	router.POST("/MetaBlob/Delete/:owner",MetaBlobDeleteOwner)
	// Get All Entries Associated with an Owner
	router.GET("/MetaBlob/Get/:owner",BlobOwners)
	// Get All Entries Associated with a Purpose
	router.GET("/MetaBlob/Get/:owner/:purpose",BlobPurpose)
	// Get All Entries Associated with a Client
	router.GET("/MetaBlob/Get/:owner/:purpose/:client",BlobPrimaryKey)
	// Get All Entries Associated with a Primary Key
	router.GET("/MetaBlob/Get/:owner/:purpose/:client/:primarykey",BlobClient)
	// 
	router.GET("/MetaBlob/Get/:owner/:purpose/:client/:primarykey/:secondarykey",BlobSecondaryKey)
	//==============
	// Retrieves the Client Server Switch States
	router.GET("/switches/",GetSwitches)
	// Retrieves the Client server Variables States
	router.GET("/variables/",GetVariables)
	// Assigns to a switch state
	router.GET("/switches/:id/:value",SetSwitch)
	// Assigns to a variable state
	router.POST("/variables/:id/:value",SetVariable)
	http.Handle("/", router)
}