/**
 * Created by super on 2019/3/7.
 *awesomeProject.
 */
package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	return router
}
func main() {
	r := RegisterHandlers()
	http.ListenAndServe("127.0.0.1:8000", r)
}
