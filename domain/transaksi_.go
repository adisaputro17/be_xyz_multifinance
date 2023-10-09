package domain

import (
	"context"
	"database/sql"
	"log"

	"github.com/adisaputro17/be_xyz_multifinance/entity"
)

func (d *domain) GetTransaksiByKonsumenID(ctx context.Context, konsumenID string) ([]entity.Transaksi, error) {
	return d.sqlReadTransaksiByKonsumenID(ctx, konsumenID)
}

func (d *domain) InsertTransaksi(ctx context.Context, v entity.Transaksi) (entity.Transaksi, error) {
	tx, err := d.DB.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return v, err
	}

	tx, err = d.sqlInsertTransaksi(tx, v)
	if err != nil {
		return v, err
	}

	defer func() {
		err := recover()
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		} else {
			tx.Commit()
		}
	}()

	return v, nil
}
