package main

import (
	"Gurumu/config"
	_guruData "Gurumu/features/guru/data"
	_guruHandler "Gurumu/features/guru/handler"
	_guruService "Gurumu/features/guru/service"

	_autentikasiData "Gurumu/features/autentikasi/data"
	_autentikasiHandler "Gurumu/features/autentikasi/handler"
	_autentikasiService "Gurumu/features/autentikasi/service"

	_jadwalData "Gurumu/features/jadwal/data"
	_jadwalHandler "Gurumu/features/jadwal/handler"
	_jadwalService "Gurumu/features/jadwal/service"

	_reservasiData "Gurumu/features/reservasi/data"
	_reservasiHandler "Gurumu/features/reservasi/handler"
	_reservasiService "Gurumu/features/reservasi/service"

	"Gurumu/features/siswa/data"
	"Gurumu/features/siswa/handler"
	"Gurumu/features/siswa/service"

	_ulasanData "Gurumu/features/ulasan/data"
	_ulasanHandler "Gurumu/features/ulasan/handler"
	_ulasanService "Gurumu/features/ulasan/service"
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

	middleware.ErrJWTMissing.Code = 401
	middleware.ErrJWTMissing.Message = "Unauthorized"

	autentikasiData := _autentikasiData.New(db)
	autentikasiSrv := _autentikasiService.New(autentikasiData)
	autentikasiHdl := _autentikasiHandler.New(autentikasiSrv)

	guruData := _guruData.New(db)
	guruSrv := _guruService.New(guruData)
	guruHdl := _guruHandler.New(guruSrv)

	jadwalData := _jadwalData.New(db)
	jadwalSrv := _jadwalService.New(jadwalData)
	jadwalHdl := _jadwalHandler.New(jadwalSrv)

	reservasiData := _reservasiData.New(db)
	reservasiSrv := _reservasiService.New(reservasiData)
	reservasiHdl := _reservasiHandler.New(reservasiSrv)

	studentData := data.New(db)
	studentSrv := service.New(studentData)
	studentHdl := handler.New(studentSrv)

	ulasanData := _ulasanData.New(db)
	ulasanSrv := _ulasanService.New(ulasanData)
	ulasanHdl := _ulasanHandler.New(ulasanSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.POST(("/login"), autentikasiHdl.Login())

	e.POST("/siswa", studentHdl.Register())
	e.GET("/siswa", studentHdl.Profile(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/siswa", studentHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/siswa", studentHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))

	e.POST("/guru", guruHdl.Register())
	e.DELETE("/guru", guruHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/guru", guruHdl.ProfileBeranda())
	e.GET("/guru/:guru_id", guruHdl.Profile())
	e.PUT("/guru", guruHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))

	e.POST("/jadwal", jadwalHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/jadwal", jadwalHdl.GetJadwal(), middleware.JWT([]byte(config.JWT_KEY)))

	e.POST("/ulasan/:guruid", ulasanHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/ulasan", ulasanHdl.GetAll())
	e.GET("/ulasan/:guruid", ulasanHdl.GetById())

	e.POST("/reservasi", reservasiHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/sesiku", reservasiHdl.Mysession(), middleware.JWT([]byte(config.JWT_KEY)))

	e.GET("/reservasi/callbackmid/:kode", reservasiHdl.CallbackMid())
	e.GET("/callback", reservasiHdl.Callback())
	e.POST("/paymentnotifications", reservasiHdl.NotificationTransactionStatus(), middleware.JWT([]byte(config.JWT_KEY)))

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
