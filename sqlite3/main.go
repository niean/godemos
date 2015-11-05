package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
	"time"
)

func main() {
	path, err := filepath.Abs(os.Args[0])
	checkErr(err)

	path = filepath.Dir(path) + "/data/foo.db"
	db, err := sql.Open("sqlite3", path)
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo (username, departname, created) VALUES (?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("niean", "sa-dev", "2015-03-12")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	//更新数据
	stmt, err = db.Prepare("UPDATE userinfo SET username=? WHERE uid=?")
	checkErr(err)

	res, err = stmt.Exec("niean.update", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created time.Time
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, username, department, created)
	}

	//删除数据
	stmt, err = db.Prepare("DELETE FROM userinfo")
	checkErr(err)

	res, err = stmt.Exec()
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
