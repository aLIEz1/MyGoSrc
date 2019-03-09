/**
 * Created by super on 2019/3/8.
 *awesomeProject.
 */
package defs

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}
type SingedUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
} 
type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}
type Comment struct {
	Id      string
	VideoId string
	Author  string
	Content string
}

type SimpleSession struct {
	Username string
	TTL      int64
}
