package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)


const(
	//数据库配置
	DBHost="127.0.0.1"
	DBPort=":3306"
	DBUser="root"
	DBPass=""
	DBDbase="go"


	//web服务器配置
	PORT=":8080"
)


var database *sql.DB


type Page struct{
	Title string
	Content string
	Date string
}


/**
数据库连接测试
 */
func main(){
	dbConn:=fmt.Sprintf("%s:%s@/%s",DBUser,DBPass,DBDbase)
	db,err:=sql.Open("mysql",dbConn)
	if err!=nil{
		log.Println("Couldn't connect")
		log.Println(err.Error())
	}
	log.Println("Connect successfully")
	database=db

	//设置路由
	routes:=mux.NewRouter()
	routes.HandleFunc("/page/{id:[0-9a-zA\\-]+",ServePage)
	routes.HandleFunc("/page/{guid:[0-9a-zA\\-]+}",ServePage)
	http.Handle("/",routes)
	http.ListenAndServe(PORT,nil)
}

func ServePage(w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r)
	pageID:=vars["id"]
	pageGUID:=vars["guid"]
	thisPage:=Page{}
	fmt.Println("pageID:"+pageID,"guid:"+pageGUID)
	err:=database.QueryRow("select page_title,page_content,page_date from pages where page_guid=?",pageGUID).Scan(&thisPage.Title,&thisPage.Content,&thisPage.Date)
	if err!=nil{
		http.Error(w,http.StatusText(404),http.StatusNotFound)
		log.Println("Couldn't get page: +pageID")
		log.Println(err.Error())
		return
	}

	html:=`<html><head><title>` + thisPage.Title +
`</title></head><body><h1>` + thisPage.Title + `</h1><div>` +
thisPage.Content + `</div></body></html>`
	fmt.Fprintln(w,html)
}

