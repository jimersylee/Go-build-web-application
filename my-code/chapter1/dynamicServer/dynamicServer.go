package main

import (
	"fmt"
	"net/http"
	"time"
)

const(
	PORT =":8080"
)

func serveDynamic(w http.ResponseWriter,r *http.Request){
	response:="The time is "+time.Now().String()
	fmt.Fprintln(w,response)
}
func serveStatic(w http.ResponseWriter,r *http.Request){
	http.ServeFile(w,r,"static.html")
}


func main(){
	http.HandleFunc("/static",serveStatic)
	http.HandleFunc("/",serveDynamic)
	http.ListenAndServe(PORT,nil)
	fmt.Printf("success");

}
