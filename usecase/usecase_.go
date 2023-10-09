package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/adisaputro17/be_xyz_multifinance/entity"
)

func (u *usecase) GetKonsumenByID(ctx context.Context, konsumenID string) (entity.Konsumen, error) {
	return u.Domain.GetKonsumenByID(ctx, konsumenID)
}

func (u *usecase) InsertKonsumen(ctx context.Context, request entity.InsertKonsumenRequest) (entity.Konsumen, error) {
	err := u.Validate.Struct(request)
	if err != nil {
		return entity.Konsumen{}, err
	}

	tanggalLahir, err := time.Parse(entity.LayoutISO, request.TanggalLahir)
	if err != nil {
		return entity.Konsumen{}, err
	}

	timeNowUTC := time.Now().UTC()
	commonField := entity.CommonField{
		CreatedAt: timeNowUTC,
		CreatedBy: 0,
		UpdatedAt: timeNowUTC,
		UpdatedBy: 0,
	}

	konsumen, err := u.Domain.InsertKonsumen(ctx, entity.Konsumen{
		NIK:          request.NIK,
		FullName:     request.FullName,
		LegalName:    request.LegalName,
		TempatLahir:  request.TempatLahir,
		TanggalLahir: tanggalLahir,
		Gaji:         request.Gaji,
		FotoKTP:      request.FotoKTP,
		FotoSelfie:   request.FotoSelfie,
		CommonField:  commonField,
	})
	if err != nil {
		return entity.Konsumen{}, err
	}

	return konsumen, nil
}

func (u *usecase) GetLimitByKonsumenID(ctx context.Context, konsumenID string) ([]entity.Limit, error) {
	return u.Domain.GetLimitByKonsumenID(ctx, konsumenID)
}

func (u *usecase) InsertLimit(ctx context.Context, request entity.InsertLimitRequest) (entity.Limit, error) {
	err := u.Validate.Struct(request)
	if err != nil {
		return entity.Limit{}, err
	}

	_, err = u.Domain.GetKonsumenByID(ctx, request.NIK)
	if err != nil {
		return entity.Limit{}, err
	}

	timeNowUTC := time.Now().UTC()
	limitID := fmt.Sprintf("%s-%d", request.NIK, request.Tenor)
	commonField := entity.CommonField{
		CreatedAt: timeNowUTC,
		CreatedBy: 0,
		UpdatedAt: timeNowUTC,
		UpdatedBy: 0,
	}

	limit, err := u.Domain.InsertLimit(ctx, entity.Limit{
		LimitID:     limitID,
		KonsumenID:  request.NIK,
		Tenor:       request.Tenor,
		Limit:       request.Limit,
		CommonField: commonField,
	})
	if err != nil {
		return entity.Limit{}, err
	}

	return limit, nil
}

func (u *usecase) UpdateLimit(ctx context.Context, request entity.UpdateLimit) (entity.UpdateLimit, error) {
	err := u.Validate.Struct(request)
	if err != nil {
		return entity.UpdateLimit{}, err
	}

	timeNowUTC := time.Now().UTC()
	updatedLimit, err := u.Domain.UpdateLimit(ctx, entity.UpdateLimit{
		LimitID:   request.LimitID,
		Limit:     request.Limit,
		UpdatedAt: timeNowUTC,
		UpdatedBy: 0,
	})
	if err != nil {
		return entity.UpdateLimit{}, err
	}

	return updatedLimit, nil
}

func (u *usecase) GetTransaksiByKonsumenID(ctx context.Context, konsumenID string) ([]entity.Transaksi, error) {
	return u.Domain.GetTransaksiByKonsumenID(ctx, konsumenID)
}

func (u *usecase) InsertTransaksi(ctx context.Context, request entity.InsertTransaksiRequest) (entity.Transaksi, error) {
	err := u.Validate.Struct(request)
	if err != nil {
		return entity.Transaksi{}, err
	}

	_, err = u.Domain.GetKonsumenByID(ctx, request.NIK)
	if err != nil {
		return entity.Transaksi{}, err
	}

	timeNowUTC := time.Now().UTC()
	commonField := entity.CommonField{
		CreatedAt: timeNowUTC,
		CreatedBy: 0,
		UpdatedAt: timeNowUTC,
		UpdatedBy: 0,
	}

	transaksi, err := u.Domain.InsertTransaksi(ctx, entity.Transaksi{
		NomorKontrak:  request.NomorKontrak,
		KonsumenID:    request.NIK,
		NamaAsset:     request.NamaAsset,
		OTR:           request.OTR,
		AdminFee:      request.AdminFee,
		JumlahCicilan: request.JumlahCicilan,
		JumlahBunga:   request.JumlahBunga,
		CommonField:   commonField,
	})
	if err != nil {
		return entity.Transaksi{}, err
	}

	return transaksi, nil
}
