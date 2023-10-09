package rest

import (
	"net/http"

	"github.com/adisaputro17/be_xyz_multifinance/entity"
	"github.com/labstack/echo/v4"
)

func (rest *rest) InsertKonsumen(c echo.Context) error {
	req := new(entity.InsertKonsumenRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	konsumen, err := rest.Usecase.InsertKonsumen(c.Request().Context(), *req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.HTTPErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, entity.HTTPResp{
		Data: konsumen,
	})
}

func (rest *rest) GetKonsumenByID(c echo.Context) error {
	konsumenID := c.Param("konsumen_id")

	konsumen, err := rest.Usecase.GetKonsumenByID(c.Request().Context(), konsumenID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.HTTPErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.HTTPResp{
		Data: konsumen,
	})
}

func (rest *rest) InsertLimit(c echo.Context) error {
	req := new(entity.InsertLimitRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	limit, err := rest.Usecase.InsertLimit(c.Request().Context(), *req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.HTTPErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, entity.HTTPResp{
		Data: limit,
	})
}

func (rest *rest) GetLimitByKonsumenID(c echo.Context) error {
	konsumenID := c.Param("konsumen_id")

	limit, err := rest.Usecase.GetLimitByKonsumenID(c.Request().Context(), konsumenID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.HTTPErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.HTTPResp{
		Data: limit,
	})
}

func (rest *rest) UpdateLimit(c echo.Context) error {
	req := new(entity.UpdateLimit)
	if err := c.Bind(req); err != nil {
		return err
	}

	limit, err := rest.Usecase.UpdateLimit(c.Request().Context(), *req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.HTTPErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.HTTPResp{
		Data: limit,
	})
}

func (rest *rest) InsertTransaksi(c echo.Context) error {
	req := new(entity.InsertTransaksiRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	transaksi, err := rest.Usecase.InsertTransaksi(c.Request().Context(), *req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.HTTPErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, entity.HTTPResp{
		Data: transaksi,
	})
}

func (rest *rest) GetTransaksiByKonsumenID(c echo.Context) error {
	konsumenID := c.Param("konsumen_id")

	transaksi, err := rest.Usecase.GetTransaksiByKonsumenID(c.Request().Context(), konsumenID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.HTTPErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.HTTPResp{
		Data: transaksi,
	})
}
