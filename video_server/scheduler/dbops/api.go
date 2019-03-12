/**
 * Created by super on 2019/3/9.
 *awesomeProject.
 */
package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func AddVideoDeletionRecord(vid string) error {
	stmtIns,err:=dbConn.Prepare(
		"INSERT INTO video_del_rec (video_id) VALUES(?) " )
	if err!=nil {
		return err
	}
	_,err= stmtIns.Exec(vid)
	if err!=nil {
		log.Printf("AddVideoDeletionDrecprd error:%v",
			err)
	}
	defer stmtIns.Close()
	return nil
}