/**
 * Created by super on 2019/3/9.
 *awesomeProject.
 */
package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/scheduler/dbops"
)

func vidDelRecHandler(w http.ResponseWriter,
	r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	if len(vid) == 0 {
		sendResponse(w, 400, "Video id should not be empty")
		return
	}
	err := dbops.AddVideoDeletionRecord(vid)
	if err != nil {
		sendResponse(w, 500, "Internal server error")
		return
	}
	sendResponse(w, 200, "")
	return

}
