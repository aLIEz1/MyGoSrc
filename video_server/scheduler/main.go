/**
 * Created by super on 2019/3/9.
 *awesomeProject.
 */
package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/scheduler/taskrunner"
)

func RegisterHandlers()*httprouter.Router  {
	router:=httprouter.New()
	router.GET("/video-delete-record/:vid-id",
		vidDelRecHandler)
	return router
}
func main()  {
	go taskrunner.Start()
	r:=RegisterHandlers()
	http.ListenAndServe("127.0.0.1:9001",r)
}

