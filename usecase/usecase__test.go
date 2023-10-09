package usecase

import (
	"context"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/adisaputro17/be_xyz_multifinance/app"
	"github.com/adisaputro17/be_xyz_multifinance/domain"
	"github.com/adisaputro17/be_xyz_multifinance/entity"
	"github.com/go-playground/validator/v10"
)

func Test_usecase_InsertKonsumen(t *testing.T) {
	validate := validator.New()
	db, err := app.NewDB(entity.DatabaseConfig{
		DBUser:     "root",
		DBPassword: "",
		DBHost:     "localhost",
		DBName:     "be_xyz_multifinance",
	})
	if err != nil {
		log.Fatal(err)
	}

	opt := domain.Options{}

	type fields struct {
		Domain   domain.DomainItf
		Validate *validator.Validate
	}
	type args struct {
		ctx     context.Context
		request entity.InsertKonsumenRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Konsumen
		wantErr bool
	}{
		{
			name: "Test Insert Konsumen",
			fields: fields{
				Domain:   domain.InitDomain(db, opt),
				Validate: validate,
			},
			args: args{
				ctx: context.Background(),
				request: entity.InsertKonsumenRequest{
					NIK:          "nik",
					FullName:     "full_name",
					LegalName:    "legal_name",
					TempatLahir:  "nganjuk",
					TanggalLahir: "2023-10-09",
					Gaji:         10000,
					FotoKTP:      "foto_ktp",
					FotoSelfie:   "foto_selfie",
				},
			},
			want: entity.Konsumen{
				NIK:          "nik",
				FullName:     "full_name",
				LegalName:    "legal_name",
				TempatLahir:  "nganjuk",
				TanggalLahir: time.Time{},
				Gaji:         10000,
				FotoKTP:      "foto_ktp",
				FotoSelfie:   "foto_selfie",
				CommonField: entity.CommonField{
					CreatedAt: time.Time{},
					CreatedBy: 0,
					UpdatedAt: time.Time{},
					UpdatedBy: 0,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				Domain:   tt.fields.Domain,
				Validate: tt.fields.Validate,
			}
			got, err := u.InsertKonsumen(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.InsertKonsumen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecase.InsertKonsumen() = %v, want %v", got, tt.want)
			}
		})
	}
}
