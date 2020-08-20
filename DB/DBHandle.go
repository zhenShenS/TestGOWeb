package DB

import (
	"TestGOWeb/recordlog"
)

var INSERT_DATA = `INSERT INTO student(sid,sname,age) VALUES(?,?,?);`

/*执行SQL语句insert或update(单数据库版)*/
func ExeSQLIU(query string) error {
	recordlog.Debug(query)
	// _, err := DB.Exec("INSERT INTO class (id,name) VALUES(?,?);", '1', "老张")
	_, err := DB.Exec(query)
	recordlog.Debug(err)
	recordlog.Debug("数据库操作中")
	return nil
}
