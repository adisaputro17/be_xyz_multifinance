package usecase

import (
	"context"

	"github.com/adisaputro17/be_xyz_multifinance/domain"
	"github.com/adisaputro17/be_xyz_multifinance/entity"
	"github.com/go-playground/validator/v10"
)

type UsecaseItf interface {
	// Konsumen
	GetKonsumenByID(ctx context.Context, konsumenID string) (entity.Konsumen, error)
	InsertKonsumen(ctx context.Context, request entity.InsertKonsumenRequest) (entity.Konsumen, error)

	// Limit
	GetLimitByKonsumenID(ctx context.Context, konsumenID string) ([]entity.Limit, error)
	InsertLimit(ctx context.Context, request entity.InsertLimitRequest) (entity.Limit, error)
	UpdateLimit(ctx context.Context, request entity.UpdateLimit) (entity.UpdateLimit, error)

	// Transaksi
	GetTransaksiByKonsumenID(ctx context.Context, konsumenID string) ([]entity.Transaksi, error)
	InsertTransaksi(ctx context.Context, request entity.InsertTransaksiRequest) (entity.Transaksi, error)
}

type usecase struct {
	Domain   domain.DomainItf
	Validate *validator.Validate
}

func InitUsecase(domain domain.DomainItf, validate *validator.Validate) UsecaseItf {
	return &usecase{
		Domain:   domain,
		Validate: validate,
	}
}
