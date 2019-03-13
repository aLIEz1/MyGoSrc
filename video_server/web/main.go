/**
 * Created by super on 2019/3/12.
 *awesomeProject.
 */
package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandler() *httprouter.Router {
	router:=httprouter.New()
	router.GET("/",homeHandler)
	router.POST("/",homeHandler)
	router.GET("/userhome",userHomehandler)
	router.POST("/userhome",userHomehandler)
	router.POST("/api",apiHandler)
	router.ServeFiles("/statics/*filepath",http.Dir("./template"))
	return router
}
func main()  {
	r:=RegisterHandler()
	http.ListenAndServe("127.0.0.1:8080",r)
}