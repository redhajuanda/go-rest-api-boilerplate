package main

import (
	"fmt"
	"net/http"
	"os"

	"rest-api-boilerplate/api"

	"rest-api-boilerplate/common"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type App struct {
	Router *gin.Engine
	DB     *gorm.DB
}

// Initialize application
func (app *App) Initialize() {
	app.SetLog()
	app.SetDB()
	app.SetRouter()
}

// Setting up logs
func (app *App) SetLog() {
	logrus.SetOutput(&lumberjack.Logger{
		Filename:   "log/log.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})
}

// Setting up Database
func (app *App) SetDB() {
	common.NewDB()
	app.DB = common.GetDB()
}

// Setting up router
func (app *App) SetRouter() {
	app.Router = gin.Default()
	app.RouteHandlers()
}

func (app *App) RouteHandlers() {
	api.ApplyRoutes(app.Router)
}

func (app *App) Run() {
	port := fmt.Sprintf(":%v", os.Getenv("PORT"))
	fmt.Println("Live on PORT", port)
	logrus.Fatal(http.ListenAndServe(port, app.Router))
}
