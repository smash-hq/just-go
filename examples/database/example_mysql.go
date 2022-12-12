package database

import (
	"database/sql"
	"fmt"
	log "github.com/corgi-kx/logcustom"
	_ "github.com/go-sql-driver/mysql"
	. "h7/tools"
	"h7/tools/result"
	"time"
)

const (
	USERNAME = "root"
	PASSWORD = "1qa2ws#ED"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "golang"
)

type userInfo struct {
	Uid        int    `json:"uid"`
	Username   string `json:"username"`
	Department string `json:"department"`
	Created    string `json:"created"`
}

var connection = fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
var db, err = sql.Open("mysql", connection)

func init() {
	CheckErr(err)
}

func InsertData() {
	stmt, err := db.Prepare("INSERT INTO user_info SET username = ?,department = ?,created=?;")
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		CheckErr(err)
	}(stmt)
	CheckErr(err)
	exec, err := stmt.Exec("h7", "师傅你是做什么的？", time.Now())
	CheckErr(err)
	id, err := exec.LastInsertId()
	CheckErr(err)
	log.Info("本次插入的自增id是：", id)
}

func SelectData() result.MultiResult {
	rows, err := db.Query("SELECT * FROM user_info")
	CheckErr(err)
	defer func() {
		err := rows.Close()
		CheckErr(err)
	}()
	user := userInfo{}
	var multi result.MultiResult
	multi.BuildSuccess()

	for rows.Next() {
		err := rows.Scan(&user.Uid, &user.Username, &user.Department, &user.Created)
		CheckErr(err)
		multi.Data = append(multi.Data, user)
	}

	return multi
}
