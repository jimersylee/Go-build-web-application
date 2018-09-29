package main

import(
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

const(
	PORT=":8080"
)

/**
mux库是路由库,可以自由定义路由
 */
func main(){
	router:=mux.NewRouter()
	router.HandleFunc("/pages/{id:[0-9]+}",pageHandler)
	http.Handle("/",router)
	http.ListenAndServe(PORT,nil)

}


func pageHandler(w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r)
	pageID:=vars["id"]
	fileName:="files/"+pageID+".html"

	_, err := os.Stat(fileName)
	if err != nil {
		fileName = "files/404.html"
	}

	http.ServeFile(w,r,fileName)
}