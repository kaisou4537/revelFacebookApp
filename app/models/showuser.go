package models

// 表示用ユーザ情報
type ShowUser struct {
	Username string
	ImgURL   string
}

var dbShowUser = make(map[string]*ShowUser)

func CreateShowUser(name, imgURL string) {
	user := &ShowUser{Username: name, ImgURL: imgURL}
	dbShowUser[name] = user
}

func FindShowUser(name string) *ShowUser {
	if user, ok := dbShowUser[name]; ok {
		return user
	} else {
		return nil
	}
}
