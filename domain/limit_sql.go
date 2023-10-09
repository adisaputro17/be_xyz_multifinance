package domain

import (
	"context"
	"database/sql"

	"github.com/adisaputro17/be_xyz_multifinance/entity"
	_ "github.com/go-sql-driver/mysql"
)

func (d *domain) sqlInsertLimit(tx *sql.Tx, v entity.Limit) (*sql.Tx, error) {
	_, err := tx.Exec(insertLimit,
		v.LimitID,
		v.KonsumenID,
		v.Tenor,
		v.Limit,
		v.CreatedAt,
		v.CreatedBy,
		v.UpdatedAt,
		v.UpdatedBy,
	)
	if err != nil {
		return tx, err
	}

	return tx, nil
}

func (d *domain) sqlReadLimitByKonsumenID(ctx context.Context, konsumenID string) ([]entity.Limit, error) {
	results := []entity.Limit{}
	rows, err := d.DB.QueryContext(ctx, readLimitByKonsumenID, konsumenID)
	if err != nil {
		return results, err
	}
	defer rows.Close()

	for rows.Next() {
		data := entity.Limit{}
		if err := rows.Scan(
			&data.LimitID,
			&data.KonsumenID,
			&data.Tenor,
			&data.Limit,
			&data.CreatedAt,
			&data.CreatedBy,
			&data.UpdatedAt,
			&data.UpdatedBy,
		); err != nil {
			return results, err
		}
		results = append(results, data)
	}

	return results, nil
}

func (d *domain) sqlUpdateLimitByLimitID(tx *sql.Tx, v entity.UpdateLimit) (*sql.Tx, error) {
	_, err := tx.Exec(updateLimitByLimitID,
		v.Limit,
		v.UpdatedAt,
		v.UpdatedBy,
		// where
		v.LimitID,
	)
	if err != nil {
		return tx, err
	}

	return tx, err
}
