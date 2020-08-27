package db

import (
	"blogger/model"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


func IsAdminExist(Auth *model.Auth) (exists bool, err error) {
	var id int64
	sqlstr := "select id from blog_auth where username=? and password=?"
	err = DB.Get(&id, sqlstr,Auth.Username,Auth.Password)
	if err == sql.ErrNoRows {
		exists = false
		return
	}
	if err != nil {
		return
	}
	exists = true
	return
}
