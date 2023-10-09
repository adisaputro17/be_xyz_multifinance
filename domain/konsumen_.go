package domain

import (
	"context"
	"database/sql"
	"log"

	"github.com/adisaputro17/be_xyz_multifinance/entity"
)

func (d *domain) GetKonsumenByID(ctx context.Context, nik string) (entity.Konsumen, error) {
	return d.sqlReadKonsumenByID(ctx, nik)
}

func (d *domain) InsertKonsumen(ctx context.Context, v entity.Konsumen) (entity.Konsumen, error) {
	tx, err := d.DB.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return v, err
	}

	tx, err = d.sqlInsertKonsumen(tx, v)
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
