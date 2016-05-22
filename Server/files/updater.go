package main
import (
    "net/http"
    "io/ioutil"
	"time"
	"fmt"
)


var name string


func init(){
	// Just change this to your server location.
	name = "localhost:8080"
}


func main(){
	for(true){
		// Adjust the number here to set how many minutes apart a backup should be done.
		time.Sleep(time.Minute*30)
		// Get Responses
		response1, err := http.Get(`http://`+name+"/switches")
		if(err != nil){
			fmt.Println(err)
			continue 
		}
		response2, errr := http.Get(`http://`+name+"/variables")
		if(errr != nil){ 
			fmt.Println(errr)
			continue 
		}
		// Get Contents
		contents1, err1    := ioutil.ReadAll(response1.Body)
		if(err1 != nil){ 
			fmt.Println(err1)
			continue 
		}
		contents2, err2    := ioutil.ReadAll(response2.Body)
		if(err2 != nil){ 
			fmt.Println(err2)
			continue 
		}
		// Update Files
		err3 := ioutil.WriteFile("Switches.ini",contents1,0777)
		if(err3 != nil){ 
			fmt.Println(err3)
			continue 
		}
		err4 := ioutil.WriteFile("Variables.ini",contents2,0777)
		if(err4 != nil){ 
			fmt.Println(err4)
			continue 
		}
	}
}