package main

import (
	"net/http"
	"fmt"
)

const (
	PORT = ":8080"
)

func main() {
	err:=http.ListenAndServe(PORT, http.FileServer(http.Dir("/home/jimersylee/projects/go/src/go-build-web-application/my-code/chapter1/staticServer")))
	if err!=nil{
		fmt.Println("start server failed")
	}
	fmt.Println("start server successfully")
}
