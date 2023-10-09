package main

import (
	"fmt"
	"log"
	"os"

	"github.com/adisaputro17/be_xyz_multifinance/app"
	"github.com/adisaputro17/be_xyz_multifinance/domain"
	"github.com/adisaputro17/be_xyz_multifinance/entity"
	"github.com/adisaputro17/be_xyz_multifinance/rest"
	"github.com/adisaputro17/be_xyz_multifinance/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	appConfig := entity.AppConfig{
		AppPort: os.Getenv("APP_PORT"),
		DBConfig: entity.DatabaseConfig{
			DBUser:     os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBHost:     os.Getenv("DB_HOST"),
			DBName:     os.Getenv("DB_NAME"),
		},
	}

	validate := validator.New()
	db, err := app.NewDB(appConfig.DBConfig)
	if err != nil {
		log.Fatal(err)
	}

	opt := domain.Options{}

	domain := domain.InitDomain(db, opt)
	usecase := usecase.InitUsecase(domain, validate)
	rest := rest.Init(usecase)

	e := echo.New()
	e.POST("/konsumen", rest.InsertKonsumen)
	e.GET("/konsumen/:konsumen_id", rest.GetKonsumenByID)

	e.POST("/limit", rest.InsertLimit)
	e.GET("/limit/:konsumen_id", rest.GetLimitByKonsumenID)
	e.PUT("/limit", rest.UpdateLimit)

	e.POST("/transaksi", rest.InsertTransaksi)
	e.GET("/transaksi/:konsumen_id", rest.GetTransaksiByKonsumenID)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", appConfig.AppPort)))
}
