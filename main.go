package main

import (
	"goProject/config"
	"goProject/server/controllers"
	"goProject/server/repositories/gorm"
	"goProject/server/router"
	"goProject/server/services"
	"os"

	"github.com/joho/godotenv"
)

// @title goProject Haktiv8
// @version 1.0
// @description Hacktiv8 Scalable Web Service with Golang Final Project
// @termsOfService http://swagger.io/terms/
// @contact.name Swagger API Team
// @contact.url http://swagger.io
// @contact.email billy.tambunan93@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3001

func main() {
	db, err := config.ConnectMysqlGorm()
	if err != nil {
		panic(err)
	}

	userRepo := gorm.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	photoRepo := gorm.NewPhotoRepository(db)
	photoService := services.NewPhotoService(photoRepo)
	photoController := controllers.NewPhotoController(photoService, userService)

	commentRepo := gorm.NewCommentRepository(db)
	commentService := services.NewCommentService(commentRepo)
	commentController := controllers.NewCommentController(commentService, userService, photoService)

	socmedRepo := gorm.NewSocialMediaRepository(db)
	socmedService := services.NewSocialMediaService(socmedRepo)
	socmedController := controllers.NewSocmedController(socmedService, userService)

	app := router.NewRouter(userController, photoController, commentController, socmedController)

	err = godotenv.Load()

	if err != nil {
		panic(err)
	}

	app.SetupRouter(os.Getenv("PORT"))
}
