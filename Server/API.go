package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/nu7hatch/gouuid"
	"net/http"
	"encoding/json"
	"strconv"
	"time"
)

/*  ==================================================
	Method: GET
	Handler: /uuid/
	Results: plain text
	Description: Generates a UUID
	=================================================*/
func UUID(res http.ResponseWriter, _ *http.Request, _ httprouter.Params){
	// Make UUID
	id,_ := uuid.NewV4()
	// Respond with it.
	fmt.Fprint(res,id.String())
}

/*  ==================================================
	Method: GET
	Handler: /map/:mapid
	Results: JSON
	Description: Collects the information about a map
	and the players there.
	=================================================*/
func Map(res http.ResponseWriter, req *http.Request, ps httprouter.Params){
	mapID, err0 := strconv.ParseInt(ps.ByName("mapid"),10,64)
	result:= Maps[mapID]
	obj,err1 := json.Marshal(result)
	HandleError(res,err0,err1)
	r := string(obj)
	if(r == "null"){
		Maps[mapID] = make(map[string]Player)
		Map(res,req,ps)
		return
	}
	fmt.Fprint(res,string(r))
}

/*  ==================================================
	Method: POST
	Handler: /playerLoggedIn/:uuid/:mapid/:x/:y
	Results: NONE
	Description: Logs the player on symbolically.
	The uuid is generated separately.
	=================================================*/
func PlayerLogin(res http.ResponseWriter, req *http.Request, ps httprouter.Params){
	// Get Player's UUID
	uuid        := ps.ByName("uuid")
	// Get the Player's MAP ID
	mapid, err0 := strconv.ParseInt(ps.ByName("mapid"),10,64)
	// Get the Player's X Position
	x, err1 := strconv.ParseInt(ps.ByName("x"),10,64)
	// Get the Player's Y Position
	y, err2 := strconv.ParseInt(ps.ByName("y"),10,64)
	HandleError(res,err0,err1,err2)
	// Run the entry by Map API to ensure an entry exists for the map.
	Map(res,req,ps)
	// Make an entry for the player.
	(Maps[mapid])[uuid] = 
		Player{
			uuid,		   					//The players unique id.
			MakeActionList(x,y,mapid),		//Initialize the player's list of actions.
			&[]time.Time{time.Now()}[0],	//Make player's initial time stamp
		}
}

/*  ==================================================
	Method: POST
	Handler: /aq/:uuid/:mapid/:actionid/:x/:y
	Results: NONE
	Description: Adds a symbolic entry to a player's actions
	for use by the game's parsing.
	This does present a security flaw but as UUIDs are 
	throwaway anyways not much can be done with it.
	(Changing how you appear to someone else's representation
	of the world.)
	=================================================*/
func AQ(res http.ResponseWriter, _ *http.Request, ps httprouter.Params){
	// Get Player's UUID
	uuid        := ps.ByName("uuid")
	// Get Player's Map ID
	mapid, err0 := strconv.ParseInt(ps.ByName("mapid"),10,64)
	// Get Player's Action ID
	actionid, err1 := strconv.ParseInt(ps.ByName("actionid"),10,64)
	// Get Player's X Position
	x, err2 := strconv.ParseInt(ps.ByName("x"),10,64)
	// Get Player's Y Position
	y, err3 := strconv.ParseInt(ps.ByName("y"),10,64)
	HandleError(res,err0,err1,err2,err3)
	
	// Copy player's current actions for increased thread safety.
	t := *(Maps[mapid][uuid].Actions)
	
	// Get next Enumerated Integer
	enum := t[len(t)-1].Enum+1
	// If the action queue has more than 10 members remove the first entry.
	if(len(t)) > 10{ t = t[1:] }
	// Append the new action.
	t = append(t,Action{enum,actionid,x,y,mapid})
	// Move the action into memory and update the timestamp.
	*Maps[mapid][uuid].Actions = t
	*Maps[mapid][uuid].timeStamp = time.Now()
}

/*  ==================================================
	Method: POST
	Handler: /playerLoggedOff/:uuid/:mapid
	Results: NONE
	Description: Removes the player's entry from the map.
	Not a big deal if it fails as TimeOut will handle it
	anyways.
	=================================================*/
func PlayerLogout(res http.ResponseWriter, _ *http.Request, ps httprouter.Params){
	// Get Player's UUID
	uuid  := ps.ByName("uuid")
	// Get Player's Map ID
	mapid, _ := strconv.ParseInt(ps.ByName("mapid"),10,64)
	// Remove Player From Game
	delete(Maps[mapid],uuid)
}

/*  ==================================================
	Method: NONE
	Handler: NONE
	Results: NONE
	Description: Removes Players who have been gone 
	for more than 5 minutes. This includes idling.
	=================================================*/
func TimeOut(){
	for(true){
		// Wait 3 Minutes Before Working. (This is a fairly expensive operation so its best not to run it too often.)
		time.Sleep(time.Minute * 3)
		// Make Collector for Timed Out Clients
		keys := make([]string,0)
		// Get the current time.
		now  := time.Now()
		// Iterate through Maps.
		for _,i := range Maps{
			// For each map iterate through players.
			for key,v := range i{
				// Get the individual player's time.
				then := *v.timeStamp
				// Get the difference in time in minutes.
				elapsedMinutes := int64(now.Sub(then)/time.Minute)
				// If they haven't communicated with the server in more than 5 minutes...
				if (elapsedMinutes > 5){
					// Add their UUID to the collected keys.
					keys = append(keys,key)
				} 
			}
		}
		// Iterate through each map.
		for _,i := range Maps{
			// Remove each player that has timed out.
			for _,v := range keys{
				delete(i,v)
			}	
		}	
	}
}