package main

import (
	"log"
	"net/http"
	"positiveMessage/models"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
)

// explore handler
func explore(w http.ResponseWriter, rq *http.Request) {
	user := checkLogin(w, rq)
	db, _ := gorm.Open(dbDriver, dbDetail)
	defer db.Close()
	//MySQLで使う条件文を用意
	selectPostAndUserDetail := "t6kismi7cp6a11x1.posts.*, t6kismi7cp6a11x1.users.user_name, t6kismi7cp6a11x1.users.icon_number"
	connectPostAndUser := "join t6kismi7cp6a11x1.users on  t6kismi7cp6a11x1.users.id = t6kismi7cp6a11x1.posts.authors_user_id"
	//人気の投稿とその投稿者名・アイコンを取得
	var popularPostOverviews []models.PostOverview
	db.Table("t6kismi7cp6a11x1.posts").Select(selectPostAndUserDetail).Joins(connectPostAndUser).Order("number_of_likes desc").Limit(8).Find(&popularPostOverviews)
	//最新の投稿とその投稿者名・アイコンを取得
	var latestPostOverviews []models.PostOverview
	db.Table("t6kismi7cp6a11x1.posts").Select(selectPostAndUserDetail).Joins(connectPostAndUser).Order("created_at desc").Limit(8).Find(&latestPostOverviews)
	item := struct {
		UserName         string
		IconNumber       int
		Title            string
		FavoritePostList []models.PostOverview
		LatestPostList   []models.PostOverview
	}{
		UserName:         user.UserName + "さん",
		IconNumber:       user.IconNumber,
		Title:            "人気の投稿&新作",
		FavoritePostList: popularPostOverviews,
		LatestPostList:   latestPostOverviews,
	}
	er := page("explore").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

// search handler
func search(w http.ResponseWriter, rq *http.Request) {
	user := checkLogin(w, rq)
	db, _ := gorm.Open(dbDriver, dbDetail)
	defer db.Close()
	item := struct {
		UserName    string
		IconNumber  int
		Title       string
		PostList    []models.PostOverview
		NoMatchFlag bool
	}{
		UserName:    user.UserName + "さん",
		IconNumber:  user.IconNumber,
		Title:       "検索",
		NoMatchFlag: false,
	}

	if rq.Method == "POST" {
		searchInput := strings.TrimSpace(rq.PostFormValue("searchInput"))
		//MySQLで使う条件文を用意
		selectPostAndUserDetail := "t6kismi7cp6a11x1.posts.*, t6kismi7cp6a11x1.users.user_name, t6kismi7cp6a11x1.users.icon_number"
		connectPostAndUser := "join t6kismi7cp6a11x1.users on  t6kismi7cp6a11x1.users.id = t6kismi7cp6a11x1.posts.authors_user_id"
		//検索結果と、その投稿者名・アイコンを取得
		var postOverviews []models.PostOverview
		db.Table("t6kismi7cp6a11x1.posts").Select(selectPostAndUserDetail).Where("t6kismi7cp6a11x1.posts.post_title LIKE ?", "%"+searchInput+"%").Joins(connectPostAndUser).Order("created_at desc").Find(&postOverviews)
		if len(postOverviews) == 0 {
			item.NoMatchFlag = true
		} else {
			item.PostList = postOverviews
		}
	}
	er := page("search").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

// myPage handler
func myPage(w http.ResponseWriter, rq *http.Request) {
	user := checkLogin(w, rq)
	db, _ := gorm.Open(dbDriver, dbDetail)
	defer db.Close()
	//MySQLで使う条件文を用意
	selectPostAndUserDetail := "t6kismi7cp6a11x1.posts.*, t6kismi7cp6a11x1.users.user_name, t6kismi7cp6a11x1.users.icon_number"
	connectSaveAndPost := "join t6kismi7cp6a11x1.posts on t6kismi7cp6a11x1.posts.id = t6kismi7cp6a11x1.saves.post_id"
	connectLikeAndPost := "join t6kismi7cp6a11x1.posts on t6kismi7cp6a11x1.posts.id = t6kismi7cp6a11x1.likes.post_id"
	connectPostAndUser := "join t6kismi7cp6a11x1.users on  t6kismi7cp6a11x1.users.id = t6kismi7cp6a11x1.posts.authors_user_id"
	//保存済みの投稿とその投稿者名・アイコンを取得
	var savedPostOverviews []models.PostOverview
	db.Table("t6kismi7cp6a11x1.saves").Select(selectPostAndUserDetail).Joins(connectSaveAndPost).Joins(connectPostAndUser).Where("t6kismi7cp6a11x1.saves.user_id=?", user.ID).Order("created_at desc").Find(&savedPostOverviews)
	//お気に入りの投稿とその投稿者名・アイコンを取得
	var favoritePostOverviews []models.PostOverview
	db.Table("t6kismi7cp6a11x1.likes").Select(selectPostAndUserDetail).Joins(connectLikeAndPost).Joins(connectPostAndUser).Where("t6kismi7cp6a11x1.likes.user_id=?", user.ID).Order("created_at desc").Find(&favoritePostOverviews)
	//ログインユーザーの投稿とその投稿者名・アイコンを取得
	var postOverviews []models.PostOverview
	db.Table("t6kismi7cp6a11x1.posts").Select(selectPostAndUserDetail).Joins(connectPostAndUser).Where("t6kismi7cp6a11x1.posts.authors_user_id=?", user.ID).Order("created_at desc").Find(&postOverviews)

	item := struct {
		UserName         string
		IconNumber       int
		Title            string
		SavedPostList    []models.PostOverview
		FavoritePostList []models.PostOverview
		PostList         []models.PostOverview
	}{
		UserName:         user.UserName + "さん",
		IconNumber:       user.IconNumber,
		Title:            "マイページ",
		SavedPostList:    savedPostOverviews,
		FavoritePostList: favoritePostOverviews,
		PostList:         postOverviews,
	}
	er := page("myPage").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

// newPost handler
func newPost(w http.ResponseWriter, rq *http.Request) {
	user := checkLogin(w, rq)
	db, _ := gorm.Open(dbDriver, dbDetail)
	defer db.Close()
	// add new post
	if rq.Method == "POST" {
		switch rq.PostFormValue("form") {
		case "addPost":
			postTitle := strings.TrimSpace(rq.PostFormValue("postTitle"))
			body := strings.TrimSpace(rq.PostFormValue("body"))
			post := models.Post{
				PostTitle:     postTitle,
				Body:          body,
				AuthorsUserId: int(user.ID),
				NumberOfLikes: 0,
			}
			db.Create(&post)
		}
	}

	item := struct {
		UserName   string
		IconNumber int
		Title      string
	}{
		UserName:   user.UserName + "さん",
		IconNumber: user.IconNumber,
		Title:      "新規投稿",
	}
	er := page("newPost").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

// login handler
func login(w http.ResponseWriter, rq *http.Request) {
	item := struct {
		UserName  string
		Title     string
		ErrorFlag bool
	}{
		UserName:  "",
		Title:     "ログイン",
		ErrorFlag: false,
	}
	if rq.Method == "GET" {
		er := page("login").Execute(w, item)
		if er != nil {
			log.Fatal(er)
		}
		return
	}
	if rq.Method == "POST" {
		db, _ := gorm.Open(dbDriver, dbDetail)
		defer db.Close()

		account := strings.TrimSpace(rq.PostFormValue("account"))

		// distinguish account (user name / e-mail address)
		accountIsEmailAddress := false
		if strings.Contains(account, "@") {
			accountIsEmailAddress = true
		}
		password := strings.TrimSpace(rq.PostFormValue("password"))

		// check account and password
		var numberOfMatch int
		var user models.User
		if accountIsEmailAddress {
			// change e-mail address to user name
			db.Where("email_address = ?", account).First(&user)
			account = user.UserName
		}
		db.Where("user_name = ? and password = ?", account, password).Find(&user).Count(&numberOfMatch)
		if numberOfMatch <= 0 {
			item.ErrorFlag = true //wrong input to login
			page("login").Execute(w, item)
			return
		}

		// logined
		ses, _ := cs.Get(rq, sessionName)
		ses.Values["login"] = true
		ses.Values["userName"] = account
		ses.Save(rq, w)
		http.Redirect(w, rq, "/myPage", redirectCode)
	}

	er := page("login").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

// logout handler
func logout(w http.ResponseWriter, rq *http.Request) {
	ses, _ := cs.Get(rq, sessionName)
	ses.Values["login"] = nil
	ses.Values["userName"] = nil
	ses.Save(rq, w)
	http.Redirect(w, rq, "/login", redirectCode)
}

// regisration handler
func regisration(w http.ResponseWriter, rq *http.Request) {
	//アイコン画像の番号を用意
	var iconNumbersTochoose [43]int
	for i := 0; i < 43; i++ {
		iconNumbersTochoose[i] = i + 1
	}

	item := struct {
		UserName            string
		Title               string
		ErrorFlag           bool
		IconNumbersTochoose [43]int
	}{
		UserName:            "",
		Title:               "新規アカウント登録",
		ErrorFlag:           false,
		IconNumbersTochoose: iconNumbersTochoose,
	}
	db, _ := gorm.Open(dbDriver, dbDetail)
	defer db.Close()

	// add new account
	if rq.Method == "POST" {
		iconNumber, _ := strconv.Atoi(rq.PostFormValue("iconNumber"))
		userName := strings.TrimSpace(rq.PostFormValue("userName"))
		emailAddress := strings.TrimSpace(rq.PostFormValue("emailAddress"))
		password := strings.TrimSpace(rq.PostFormValue("password"))
		user := models.User{
			UserName:     userName,
			EmailAddress: emailAddress,
			Password:     password,
			IconNumber:   iconNumber,
		}

		// ユーザー名とメールアドレスの重複をチェック
		var countOfUserName int
		var countOfEmailAddress int
		db.Where("user_name = ?", userName).Find(&user).Count(&countOfUserName)
		db.Where("email_address = ?", emailAddress).Find(&user).Count(&countOfEmailAddress)
		if countOfUserName > 0 || countOfEmailAddress > 0 {
			item.ErrorFlag = true
			page("regisration").Execute(w, item)
			return
		}
		db.Create(&user)

		// logined
		ses, _ := cs.Get(rq, sessionName)
		ses.Values["login"] = true
		ses.Values["userName"] = userName
		ses.Save(rq, w)
		http.Redirect(w, rq, "/myPage", redirectCode)
	}

	er := page("regisration").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

// postDetail handler
func postDetail(w http.ResponseWriter, rq *http.Request) {
	user := checkLogin(w, rq)
	db, _ := gorm.Open(dbDriver, dbDetail)
	defer db.Close()

	if rq.Method == "POST" {
		postId, _ := strconv.Atoi(rq.PostFormValue("postId"))
		userId, _ := strconv.Atoi(rq.PostFormValue("userId"))
		switch rq.PostFormValue("form") {
		//保存
		case "save":
			save := models.Save{
				PostId: postId,
				UserId: userId,
			}
			db.Create(&save)
		case "cancelSave":
			db.Where("post_id = ? and user_id = ?", postId, userId).Delete(&models.Save{})
		//お気に入り登録
		case "like":
			//likesテーブルを更新
			like := models.Like{
				PostId: postId,
				UserId: userId,
			}
			db.Create(&like)
			//postsテーブルのnumber_of_likesを更新
			db.Model(&models.Post{}).Where("id = ?", postId).UpdateColumn("number_of_likes", gorm.Expr("number_of_likes + ?", 1))
		case "cancelLike":
			//likesテーブルを更新
			db.Where("post_id = ? AND user_id = ?", postId, userId).Delete(&models.Like{})
			//postsテーブルのnumber_of_likesを更新
			db.Model(&models.Post{}).Where("id = ?", postId).UpdateColumn("number_of_likes", gorm.Expr("number_of_likes - ?", 1))
		}
	}

	//MySQLで使う条件文を用意
	selectPostAndUserDetail := "t6kismi7cp6a11x1.posts.*, t6kismi7cp6a11x1.users.user_name, t6kismi7cp6a11x1.users.icon_number"
	connectPostAndUser := "join t6kismi7cp6a11x1.users on  t6kismi7cp6a11x1.users.id = t6kismi7cp6a11x1.posts.authors_user_id"
	//トップに表示する投稿とその投稿者名・アイコンを取得
	postId := rq.FormValue("postId")
	var postOverview models.PostOverview
	db.Table("t6kismi7cp6a11x1.posts").Select(selectPostAndUserDetail).Joins(connectPostAndUser).Where("t6kismi7cp6a11x1.posts.id=?", postId).First(&postOverview)
	//下部に表示する、同投稿者の投稿とその投稿者名・アイコンを取得
	var relatedPostOverviews []models.PostOverview
	db.Table("t6kismi7cp6a11x1.posts").Select(selectPostAndUserDetail).Joins(connectPostAndUser).Where("t6kismi7cp6a11x1.posts.authors_user_id=?", postOverview.Post.AuthorsUserId).Order("created_at desc").Find(&relatedPostOverviews)

	item := struct {
		UserId          int
		UserName        string
		IconNumber      int
		Title           string
		PostOverview    models.PostOverview
		RelatedPostList []models.PostOverview
		SavedFlag       bool
		LikedFlag       bool
	}{
		UserId:          int(user.ID),
		UserName:        user.UserName + "さん",
		IconNumber:      user.IconNumber,
		Title:           "投稿の詳細",
		PostOverview:    postOverview,
		RelatedPostList: relatedPostOverviews,
		SavedFlag:       false,
		LikedFlag:       false,
	}

	// トップに表示する投稿が、アカウントに保存済みか確認
	var numberOfSave int
	var save models.Save
	db.Where("post_id = ? and user_id = ?", postId, user.ID).Find(&save).Count(&numberOfSave)
	if numberOfSave > 0 {
		item.SavedFlag = true //already saved
	}
	// 表示する投稿が、お気に入り登録されているか確認
	var numberOfLike int
	var like models.Like
	db.Where("post_id = ? and user_id = ?", postId, user.ID).Find(&like).Count(&numberOfLike)
	if numberOfLike > 0 {
		item.LikedFlag = true //already liked
	}

	er := page("postDetail").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}
