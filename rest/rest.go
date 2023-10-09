package rest

import (
	"github.com/adisaputro17/be_xyz_multifinance/usecase"
	"github.com/labstack/echo/v4"
)

type REST interface {
	// Konsumen
	InsertKonsumen(c echo.Context) error
	GetKonsumenByID(c echo.Context) error

	// Limit
	InsertLimit(c echo.Context) error
	GetLimitByKonsumenID(c echo.Context) error
	UpdateLimit(c echo.Context) error

	// Transaksi
	InsertTransaksi(c echo.Context) error
	GetTransaksiByKonsumenID(c echo.Context) error
}

type rest struct {
	Usecase usecase.UsecaseItf
}

func Init(usecase usecase.UsecaseItf) REST {
	return &rest{
		Usecase: usecase,
	}
}
