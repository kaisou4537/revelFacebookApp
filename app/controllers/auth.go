package controllers

import (
	"encoding/json"
	"facebookApp/app/models"
	"fmt"
	"github.com/revel/revel"
	"golang.org/x/oauth2"
	// "io/ioutil"
)

type Auth struct {
	*revel.Controller
}

// scopeは何にアクセスするかを確認するもの ex)メールアドレスなど 下記はその一覧
// http://fb.dev-plus.jp/reference/coreconcepts/api/permissions/

var facebook = newConfig([]string{"email"})

func newConfig(scopes []string) *oauth2.Config {
	c := &oauth2.Config{
		ClientID:     "xxxxxxxxxxxxxxxxxxxx",
		ClientSecret: "xxxxxxxxxxxxxxxxxxxx",
		RedirectURL:  "http://localhost:9000/auth/callback",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.facebook.com/dialog/oauth",
			TokenURL: "https://graph.facebook.com/oauth/access_token",
		},
		Scopes: []string{"email"},
	}

	for _, scope := range scopes {
		c.Scopes = append(c.Scopes, scope)
	}

	return c
}

func (c Auth) Index() revel.Result {
	// facebookから情報を取得する
	url := facebook.AuthCodeURL("")
	return c.Redirect(url)
}

func (c Auth) Callback(code string) revel.Result {
	// codeを取得したらユーザ情報を取得

	// アクセストークンを取得
	tok, err := facebook.Exchange(oauth2.NoContext, code)
	// エラー処理は必須
	if err != nil {
		return c.Redirect(App.Index)
	}

	// ユーザ情報を取得
	client := facebook.Client(oauth2.NoContext, tok)
	result, err := client.Get("https://graph.facebook.com/me")
	if err != nil {
		// 失敗
		revel.ERROR.Println("アクセストークン取得失敗！！", err)
	}

	// 関数を抜ける際に必ずresponseをcloseするようにdeferでcloseを呼ぶ
	defer result.Body.Close()

	//jsonをパースする
	account := struct {
		MailAddress string `json:"email"`
	}{}
	_ = json.NewDecoder(result.Body).Decode(&account)

	fmt.Println(account.MailAddress)

	// 取得したhtml(json)を確認する方法
	// byteArray, _ := ioutil.ReadAll(result.Body)
	// fmt.Println(string(byteArray)) // htmlをstringで取得

	return c.Redirect(Show.Index)
}

func (c Auth) Show() revel.Result {
	// ユーザ情報取得
	user := getShowUser("kaisou_test")
	return c.Render(user)
}

// Twitterユーザ情報
func getUser() *models.User {
	return models.FindOrCreate("kaisou")
}

// 表示用ユーザ情報セット
func setUserData(name, imgURL string) {
	models.CreateShowUser(name, imgURL)
}
