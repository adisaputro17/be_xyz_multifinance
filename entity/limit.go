package entity

import "time"

type Limit struct {
	LimitID    string  `json:"limit_id"`
	KonsumenID string  `json:"konsumen_id"`
	Tenor      int64   `json:"tenor"`
	Limit      float64 `json:"limit"`
	CommonField
}

type UpdateLimit struct {
	LimitID   string    `json:"limit_id" validate:"required"`
	Limit     float64   `json:"limit" validate:"required"`
	UpdatedAt time.Time `json:"-"`
	UpdatedBy int64     `json:"-"`
}

type InsertLimitRequest struct {
	NIK   string  `json:"nik" validate:"required"`
	Tenor int64   `json:"tenor" validate:"required"`
	Limit float64 `json:"limit" validate:"required"`
}
