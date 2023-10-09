package domain

import (
	"context"
	"database/sql"
	"log"

	"github.com/adisaputro17/be_xyz_multifinance/entity"
)

func (d *domain) GetLimitByKonsumenID(ctx context.Context, konsumenID string) ([]entity.Limit, error) {
	return d.sqlReadLimitByKonsumenID(ctx, konsumenID)
}

func (d *domain) InsertLimit(ctx context.Context, v entity.Limit) (entity.Limit, error) {
	tx, err := d.DB.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return v, err
	}

	tx, err = d.sqlInsertLimit(tx, v)
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

func (d *domain) UpdateLimit(ctx context.Context, v entity.UpdateLimit) (entity.UpdateLimit, error) {
	tx, err := d.DB.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return v, err
	}

	tx, err = d.sqlUpdateLimitByLimitID(tx, v)
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
