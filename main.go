// main.go for migration
// package main

// import (
// 	"positiveMessage/models"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {
// 	models.Migrate()
// }

// main.go for main program
package main

import (
	"html/template"
	"net/http"
	"os"
	"positiveMessage/models"

	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
)

// db variable
var dbDriver = "mysql"
var dbDetail = "o48tblwd6bokwr8s:t8kapv0u7r4e937k@tcp(etdq12exrvdjisg6.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306)/t6kismi7cp6a11x1?charset=utf8mb4&parseTime=True"

// session variable
var sessionName = "login-session"
var cs *sessions.CookieStore = sessions.NewCookieStore([]byte("secret-key-1234"))

const redirectCode = 302

// login check
func checkLogin(w http.ResponseWriter, rq *http.Request) *models.User {
	ses, _ := cs.Get(rq, sessionName)
	if ses.Values["login"] == nil || !ses.Values["login"].(bool) {
		http.Redirect(w, rq, "/login", redirectCode)
	}
	un := ""
	if ses.Values["userName"] != nil {
		un = ses.Values["userName"].(string)
	}

	var user models.User
	db, _ := gorm.Open(dbDriver, dbDetail)
	defer db.Close()

	db.Where("user_name = ?", un).First(&user)

	return &user
}

// テンプレート名を用意
func page(fname string) *template.Template {
	tmps, _ := template.ParseFiles("templates/"+fname+".html",
		"templates/header.html", "templates/footer.html")
	return tmps
}

// main program.
func main() {
	// index handling
	http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		myPage(w, rq)
	})
	// explore handling
	http.HandleFunc("/explore", func(w http.ResponseWriter, rq *http.Request) {
		explore(w, rq)
	})
	// search handling
	http.HandleFunc("/search", func(w http.ResponseWriter, rq *http.Request) {
		search(w, rq)
	})
	// myPage handling
	http.HandleFunc("/myPage", func(w http.ResponseWriter, rq *http.Request) {
		myPage(w, rq)
	})
	// newPost handling
	http.HandleFunc("/newPost", func(w http.ResponseWriter, rq *http.Request) {
		newPost(w, rq)
	})
	// login handling
	http.HandleFunc("/login", func(w http.ResponseWriter, rq *http.Request) {
		login(w, rq)
	})
	// regisration handling
	http.HandleFunc("/regisration", func(w http.ResponseWriter, rq *http.Request) {
		regisration(w, rq)
	})
	// logout handling.
	http.HandleFunc("/logout", func(w http.ResponseWriter, rq *http.Request) {
		logout(w, rq)
	})
	// postDetail handling
	http.HandleFunc("/postDetail", func(w http.ResponseWriter, rq *http.Request) {
		postDetail(w, rq)
	})

	//imagesフォルダ内の画像を読み込むファイルサーバー
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images/"))))

	//JavaScriptフォルダ内のJSを読み込むファイルサーバー
	http.Handle("/JavaScript/", http.StripPrefix("/JavaScript/", http.FileServer(http.Dir("JavaScript/"))))

	port := os.Getenv("PORT") //環境変数を取得
	if port == "" {
		http.ListenAndServe("", nil) //localhost
	} else {
		http.ListenAndServe(":"+port, nil) //heroku
	}
}
