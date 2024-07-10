// DBのマイグレーション用(普段は使用しない)
package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Migrate program.
func Migrate() {
	db, er := gorm.Open("mysql", "root:jfewoiHfw4ji2!ffo2@tcp(localhost:3306)/db0529")
	if er != nil {
		fmt.Println(er)
		return
	}
	defer db.Close()

	db.AutoMigrate(&User{}, &Post{}, &Save{}, &Like{})

}
