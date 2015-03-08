package controllers

import (
	"github.com/revel/revel"
	"twitterApp/app/models"
)

type Show struct {
	*revel.Controller
}

func (c Show) Index() revel.Result {
	// ユーザ情報取得
	user := getShowUser("kaisou_test")
	return c.Render(user)
}

// 表示用ユーザ情報取得
func getShowUser(name string) *models.ShowUser {
	return models.FindShowUser(name)
}
