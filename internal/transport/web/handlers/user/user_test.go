package user

import (
	"github.com/MaxKudIT/messkudi/internal/domain/user"
	"github.com/MaxKudIT/messkudi/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"testing"
)

type ussmock struct {
}

func (uh *ussmock) CreateUser(userCr dto.UserDTO) (user.User, error) {
	return user.User{}, nil
}

func (uh *ussmock) GetUser(id uuid.UUID) (user.User, error) {
	return user.User{}, nil
}

func (uh *ussmock) DeleteUser(id uuid.UUID) error {
	return nil
}

func Test_userhandler_CreateUser(t *testing.T) {
	type fields struct {
		us userservice
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "success",
			fields: fields{
				us: &ussmock{},
			},
			args: args{
				c: &gin.Context{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uh := &userhandler{
				us: tt.fields.us,
			}
			uh.CreateUser(tt.args.c)
		})
	}
}

//func Test_userhandler_DeleteUser(t *testing.T) {
//	type fields struct {
//		us *userservice
//	}
//	type args struct {
//		c *gin.Context
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			uh := &userhandler{
//				us: tt.fields.us,
//			}
//			uh.DeleteUser(tt.args.c)
//		})
//	}
//}
//
//func Test_userhandler_GetUser(t *testing.T) {
//	type fields struct {
//		us *userservice
//	}
//	type args struct {
//		c *gin.Context
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			uh := &userhandler{
//				us: tt.fields.us,
//			}
//			uh.GetUser(tt.args.c)
//		})
//	}
//}
