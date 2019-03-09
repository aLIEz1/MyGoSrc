/**
 * Created by super on 2019/3/8.
 *awesomeProject.
 */
package main

import (
	"encoding/json"
	"io"
	"net/http"
	"video_server/api/defs"
)

func sendErrorResponse(w http.ResponseWriter,errResp defs.ErroResponse) {
	w.WriteHeader(errResp.HttpSC)
	resStr,_:=json.Marshal(&errResp.Error)
	io.WriteString(w,string(resStr))

}

func sendNormalResponse(w http.ResponseWriter,resp string,sc int) {
	w.WriteHeader(sc)
	io.WriteString(w,resp)

}
