package user

import (
	"database/sql"
	"fmt"
	"github.com/MaxKudIT/messkudi/internal/domain"
	"github.com/MaxKudIT/messkudi/internal/domain/auth"
	"github.com/MaxKudIT/messkudi/internal/storage"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"os"
	"testing"
	"time"
)

func Test_userStorage_DeleteUser(t *testing.T) {
	if godotenv.Load("../../../.env") != nil {
		return
	}
	uuids, _ := uuid.Parse("8a61ba71-e15a-457d-821c-b60853958f77")
	dbstruct := storage.NewDatabase(fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_NAME"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT")))
	db, _ := dbstruct.ConnectionDB()
	type fields struct {
		db *sql.DB
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "test",
			fields:  fields{db: db},
			args:    args{id: uuids},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &userStorage{
				db: tt.fields.db,
			}
			if err := us.DeleteUser(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userStorage_GetUserById(t *testing.T) {
	if godotenv.Load("../../../.env") != nil {
		return
	}
	dbstruct := storage.NewDatabase(fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_NAME"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT")))
	db, _ := dbstruct.ConnectionDB()
	uuids, _ := uuid.Parse("977e8c74-791a-40b3-bdab-0495a787eb2a")
	type fields struct {
		db *sql.DB
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "test",
			fields:  fields{db: db},
			args:    args{id: uuids},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &userStorage{
				db: tt.fields.db,
			}
			_, err := us.GetUserById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func Test_userStorage_SaveUser(t *testing.T) {
	if godotenv.Load("../../../.env") != nil {
		return
	}
	var userp domain.User = domain.User{
		Name:        "Maxos",
		Password:    "3211",
		LastName:    "Kud",
		PhoneNumber: "79001013704",
		CreatedAt:   time.Now().Format("2006-01-02"),
		UpdatedAt:   time.Now().Format("2006-01-02"),
		Token:       auth.Token{},
	}
	dbstruct := storage.NewDatabase(fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_NAME"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT")))
	db, _ := dbstruct.ConnectionDB()
	type fields struct {
		db *sql.DB
	}

	type args struct {
		userf domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "test",
			fields:  fields{db: db},
			args:    args{userf: userp},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &userStorage{
				db: tt.fields.db,
			}
			if err := us.SaveUser(tt.args.userf); (err != nil) != tt.wantErr {
				t.Errorf("SaveUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
