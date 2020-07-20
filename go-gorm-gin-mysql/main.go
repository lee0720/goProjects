package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/leehom/go-gorm-gin-mysql/Config"
	"github.com/leehom/go-gorm-gin-mysql/Models"
	Routers "github.com/leehom/go-gorm-gin-mysql/Routes"
)

var err error
func main() {
	Config.DB, err = gorm.Open("mysql",
		Config.DbUrl(Config.BuildConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	r := Routers.SetupRouter()
	r.Run()
}
