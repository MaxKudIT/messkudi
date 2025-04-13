package user

import (
	"github.com/MaxKudIT/messkudi/internal/domain"
	"github.com/MaxKudIT/messkudi/internal/transport/web/dto"
	"github.com/google/uuid"
	"log/slog"
	"os"
	"testing"
)

type usmock struct {
}

func (us *usmock) GetUserById(id uuid.UUID) (domain.User, error) {
	return domain.User{}, nil
}

func (us *usmock) SaveUser(userp domain.User) error {
	return nil
}

func (us *usmock) DeleteUser(id uuid.UUID) error {
	return nil
}

func Test_userService_CreateUser(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	userdto := dto.UserDTO{Name: "Dima", LastName: "SCHERBAKOV", Password: "123456", PhoneNumber: "7900153134"}
	usmockp := &usmock{}
	type fields struct {
		us userStorage
		l  *slog.Logger
	}
	type args struct {
		userCr dto.UserDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				us: usmockp,
				l:  logger,
			},
			args:    args{userCr: userdto},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				us: tt.fields.us,
				l:  tt.fields.l,
			}
			_, err := u.CreateUser(tt.args.userCr)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func Test_userService_DeleteUser(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	usmockp := &usmock{}
	type fields struct {
		us userStorage
		l  *slog.Logger
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "success",
			fields: fields{
				us: usmockp,
				l:  logger,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				us: tt.fields.us,
				l:  tt.fields.l,
			}
			u.DeleteUser(tt.args.id)
		})
	}
}
func Test_userService_GetUserById(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	usmockp := &usmock{}
	type fields struct {
		us userStorage
		l  *slog.Logger
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "success",
			fields: fields{
				us: usmockp,
				l:  logger,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				us: tt.fields.us,
				l:  tt.fields.l,
			}
			u.GetUser(tt.args.id)
		})
	}
}
