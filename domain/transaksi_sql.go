package domain

import (
	"context"
	"database/sql"

	"github.com/adisaputro17/be_xyz_multifinance/entity"
	_ "github.com/go-sql-driver/mysql"
)

func (d *domain) sqlInsertTransaksi(tx *sql.Tx, v entity.Transaksi) (*sql.Tx, error) {
	_, err := tx.Exec(insertTransaksi,
		v.NomorKontrak,
		v.KonsumenID,
		v.NamaAsset,
		v.OTR,
		v.AdminFee,
		v.JumlahCicilan,
		v.JumlahBunga,
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

func (d *domain) sqlReadTransaksiByKonsumenID(ctx context.Context, konsumenID string) ([]entity.Transaksi, error) {
	results := []entity.Transaksi{}
	rows, err := d.DB.QueryContext(ctx, readTransaksiByKonsumenID, konsumenID)
	if err != nil {
		return results, err
	}
	defer rows.Close()

	for rows.Next() {
		data := entity.Transaksi{}
		if err := rows.Scan(
			&data.NomorKontrak,
			&data.KonsumenID,
			&data.NamaAsset,
			&data.OTR,
			&data.AdminFee,
			&data.JumlahCicilan,
			&data.JumlahBunga,
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
