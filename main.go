package main

import (
	"Gurumu/config"
	_guruData "Gurumu/features/guru/data"
	_guruHandler "Gurumu/features/guru/handler"
	_guruService "Gurumu/features/guru/service"
	"Gurumu/features/siswa/data"
	"Gurumu/features/siswa/handler"
	"Gurumu/features/siswa/service"
	"Gurumu/migration"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	migration.Migrate(db)

	guruData := _guruData.New(db)
	guruSrv := _guruService.New(guruData)
	guruHdl := _guruHandler.New(guruSrv)

	studentData := data.New(db)
	studentSrv := service.New(studentData)
	studentHdl := handler.New(studentSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.POST("/siswa", studentHdl.Register())
	// e.POST("/login", userHdl.Login())
	// e.GET("/users", userHdl.Profile(), middleware.JWT([]byte(config.JWT_KEY)))
	// e.PUT("/users", userHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	// e.DELETE("/users", userHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))

	e.POST("/guru", guruHdl.Register())
	e.DELETE("/guru", guruHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))
	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
