package entity

import "time"

const (
	LayoutISO = "2006-01-02"
)

type AppConfig struct {
	AppPort  string
	DBConfig DatabaseConfig
}

type DatabaseConfig struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBName     string
}

type CommonField struct {
	CreatedAt time.Time `json:"-"`
	CreatedBy int64     `json:"-"`
	UpdatedAt time.Time `json:"-"`
	UpdatedBy int64     `json:"-"`
}

type HTTPResp struct {
	Data interface{} `json:"data,omitempty"`
}

type HTTPErrorResponse struct {
	Message string `json:"message"`
}
