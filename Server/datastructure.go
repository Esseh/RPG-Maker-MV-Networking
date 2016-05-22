package main
import "time"
// ==============================================================================
// An action contains a lot of information.
// The player's x and y location when performing the action(used for correction),
// the Action_Id so that the type of action can be resolve,
// an Enumerated value which will keep track of in what sequence an action has occurred,
// and lastly the current map id. 
// ==============================================================================
type Action struct{
	Enum,Action_Id,X,Y,Map_Id int64
}
// ==============================================================================
// The player structure.
// Contains the player's UUID and a list of their actions.
// ==============================================================================
type Player struct{
	UUID string
	Actions *([]Action)
	timeStamp *time.Time
}
// ==============================================================================
// Int is the Map ID 
// This can be used to find a list of players in the same map.
// ==============================================================================
var Maps map[int64](map[string]Player) 
// ==============================================================================
// This is a way for third party scripters to store and retrieve data.
// ==============================================================================
var MetaBlob map[string](map[string](map[string](map[string](map[string](string)))))
// ==============================================================================
// Switches and Variable Container
// Used for Consistency with Client
// ==============================================================================
var GameSwitches map[string]bool
var GameVariable map[string]int64