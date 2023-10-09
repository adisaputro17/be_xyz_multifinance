package entity

import "time"

type Konsumen struct {
	NIK          string    `json:"nik"`
	FullName     string    `json:"full_name"`
	LegalName    string    `json:"legal_name"`
	TempatLahir  string    `json:"tempat_lahir"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	Gaji         float64   `json:"gaji"`
	FotoKTP      string    `json:"foto_ktp"`
	FotoSelfie   string    `json:"foto_selfie"`
	CommonField
}

type InsertKonsumenRequest struct {
	NIK          string  `json:"nik" validate:"required"`
	FullName     string  `json:"full_name" validate:"required"`
	LegalName    string  `json:"legal_name" validate:"required"`
	TempatLahir  string  `json:"tempat_lahir" validate:"required"`
	TanggalLahir string  `json:"tanggal_lahir" validate:"required"`
	Gaji         float64 `json:"gaji" validate:"required"`
	FotoKTP      string  `json:"foto_ktp" validate:"required"`
	FotoSelfie   string  `json:"foto_selfie" validate:"required"`
}
