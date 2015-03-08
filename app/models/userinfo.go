package models

import "github.com/mrjones/oauth"

// Twitter用ユーザ情報
type User struct {
	Username     string
	RequestToken *oauth.RequestToken
	AccessToken  *oauth.AccessToken
}

var db = make(map[string]*User)

func FindOrCreate(username string) *User {
	if user, ok := db[username]; ok {
		return user
	}
	user := &User{Username: username}
	db[username] = user
	return user
}
