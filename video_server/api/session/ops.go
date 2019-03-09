/**
 * Created by super on 2019/3/8.
 *awesomeProject.
 */
package session

import (
	"sync"
	"time"
	"video_server/api/dbops"
	"video_server/api/defs"
	"video_server/api/utils"
)

var sessionMap *sync.Map
func init()  {
	sessionMap=&sync.Map{}
}
func nowInMilli() int64 {
	return time.Now().UnixNano()/1000000
}
func LoadSessionFromDB()  {
	r,err:=dbops.RetrieveAllSessions()
	if err!=nil {
		return
	}

	r.Range(func(k, v interface{}) bool {
		ss:=v.(defs.SimpleSession)
		sessionMap.Store(k,ss)
		return true
	})
}
func GenerateMewSessionId(un string) string {
	id,_:=utils.NewUUID()
	ct:=nowInMilli()
	ttl:=ct+30*60*1000

	ss:=&defs.SimpleSession{un,ttl}
	sessionMap.Store(id,ss)
	dbops.InsertSession(id, ttl, un)

	return id
}
func IsSessionExpired(sid string)(string,bool)  {
	ss,ok:=sessionMap.Load(sid)
	if ok {
		ct:=nowInMilli()
		if ss.(*defs.SimpleSession).TTL<ct {
			DeleteExpiredSession(sid)
			return "",true
		}

		return ss.(*defs.SimpleSession).Username,false
	}
	return "",true
}
func DeleteExpiredSession(sid string)  {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}