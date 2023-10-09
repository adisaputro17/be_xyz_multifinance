package domain

import (
	"context"
	"database/sql"

	"github.com/adisaputro17/be_xyz_multifinance/entity"
	_ "github.com/go-sql-driver/mysql"
)

func (d *domain) sqlInsertKonsumen(tx *sql.Tx, v entity.Konsumen) (*sql.Tx, error) {
	_, err := tx.Exec(insertKonsumen,
		v.NIK,
		v.FullName,
		v.LegalName,
		v.TempatLahir,
		v.TanggalLahir,
		v.Gaji,
		v.FotoKTP,
		v.FotoSelfie,
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

func (d *domain) sqlReadKonsumenByID(ctx context.Context, nik string) (entity.Konsumen, error) {
	result := entity.Konsumen{}
	row := d.DB.QueryRowContext(ctx, readKonsumenByID, nik)

	err := row.Scan(
		&result.NIK,
		&result.FullName,
		&result.LegalName,
		&result.TempatLahir,
		&result.TanggalLahir,
		&result.Gaji,
		&result.FotoKTP,
		&result.FotoSelfie,
		&result.CreatedAt,
		&result.CreatedBy,
		&result.UpdatedAt,
		&result.UpdatedBy,
	)

	if err != nil {
		return result, err
	}

	return result, nil
}
