package main

import (
    "fmt"
    "net/http"
)

func HandleError(res http.ResponseWriter,err ...error){
	for _, e := range err{
		if e != nil {
			fmt.Fprint(res,e,"\n")
		}
	}
}


func MakeActionList(x,y,mapid int64)(*[]Action){
	t:=[]Action{
		Action{
			0,	   //The first action.
			0,	   //Action ID 0 "Update Position"
			x, 	   //x position
			y,     //y position
			mapid, //player map location
		},
	}
	return &t
}