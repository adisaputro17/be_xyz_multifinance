package domain

import (
	"context"
	"database/sql"

	"github.com/adisaputro17/be_xyz_multifinance/entity"
)

type DomainItf interface {
	// Konsumen
	GetKonsumenByID(ctx context.Context, nik string) (entity.Konsumen, error)
	InsertKonsumen(ctx context.Context, v entity.Konsumen) (entity.Konsumen, error)

	// Limit
	GetLimitByKonsumenID(ctx context.Context, konsumenID string) ([]entity.Limit, error)
	InsertLimit(ctx context.Context, v entity.Limit) (entity.Limit, error)
	UpdateLimit(ctx context.Context, v entity.UpdateLimit) (entity.UpdateLimit, error)

	// Transaksi
	GetTransaksiByKonsumenID(ctx context.Context, konsumenID string) ([]entity.Transaksi, error)
	InsertTransaksi(ctx context.Context, v entity.Transaksi) (entity.Transaksi, error)
}

type domain struct {
	DB  *sql.DB
	Opt Options
}

type Options struct {
}

func InitDomain(
	db *sql.DB,
	opt Options,
) DomainItf {
	return &domain{
		DB:  db,
		Opt: opt,
	}
}
