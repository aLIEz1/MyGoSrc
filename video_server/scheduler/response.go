/**
 * Created by super on 2019/3/9.
 *awesomeProject.
 */
package main

import (
	"io"
	"net/http"
)

func sendResponse(w http.ResponseWriter,
	sc int, resp string) {
		w.WriteHeader(sc)
		io.WriteString(w,resp)

}
