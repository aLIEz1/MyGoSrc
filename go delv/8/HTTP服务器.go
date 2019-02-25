package main

import "net/http"

func handleconn(w http.ResponseWriter, req *http.Request)  {
	w.Write([]byte("Hello Go"))

}
func main()  {
	http.HandleFunc("/",handleconn)
	http.ListenAndServe(":8000",nil)//listenAndServer第一个参数是监听地址，第二个一般为空
}
