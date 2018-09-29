package main

import (
	"fmt"
	"net/http"
)

const (
	PORT = ":8080"
)

func main() {
	err:=http.ListenAndServe(PORT, http.FileServer(http.Dir("/home/jimersylee")))
	if err!=nil{
		fmt.Println("start server failed")
	}else{
		fmt.Println("start server successfully")
	}

}
