package service

import (
	"blogger/dao/db"
	"blogger/model"
	"blogweb_gin/utils"
	"fmt"
)
func GetLogin(username, password string) (err error) {
	//exist, err := db.IsArticleExist(username)
	//if err != nil {
	//	fmt.Printf("query database failed, err:%v\n", err)
	//	return
	//}
	//if exist == false {
	//	err = fmt.Errorf("article id:%d not found", articleId)
	//	return
	//}
	var exists bool
	var c model.Auth
	c.Username = username
	c.Password = utils.MD5(password)
	exists,err = db.IsAdminExist(&c)
	if err != nil {
		fmt.Printf("query database failed, err:%v\n", err)
		return
	}
	if exists == false {
		err = fmt.Errorf("Auth id:%d not found", username)
		return
	}
	return
}
