package repository

import (
	"authservice/pkg/utils/models"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestFindUserByEmail(t *testing.T) {
	tests := []struct {
		name    string
		arg     models.UserLogin
		stub    func(mock sqlmock.Sqlmock)
		want    models.UserSignup
		wantErr bool
	}{
		{
			name: "successful, user found",
			arg:  models.UserLogin{Email: "testuser@gmail.com"},
			stub: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT \* FROM users WHERE email=\$1`).
					WithArgs("testuser@gmail.com").
					WillReturnRows(sqlmock.NewRows([]string{"id", "firstname", "lastname", "email", "password", "phone_number", "date_of_birth", "gender", "bio"}).
						AddRow(1, "Test", "User", "testuser@gmail.com", "password123", "1234567890", "1990-01-01", "male", "bio"))
			},
			want: models.UserSignup{
				ID:          1,
				Firstname:   "Test",
				Lastname:    "User",
				Email:       "testuser@gmail.com",
				Password:    "password123",
				PhoneNumber: "1234567890",
				DateOfBirth: "1990-01-01",
				Gender:      "male",
				Bio:         "bio",
			},
			wantErr: false,
		},
		{
			name: "database error",
			arg:  models.UserLogin{Email: "dberror@gmail.com"},
			stub: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT \* FROM users WHERE email=\$1`).
					WithArgs("dberror@gmail.com").
					WillReturnError(errors.New("error checking user details"))
			},
			want:    models.UserSignup{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSql, _ := sqlmock.New()

			DB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			userRepository := NewUserRepository(DB)
			tt.stub(mockSql)

			got, err := userRepository.FindUserByEmail(tt.arg)
			if tt.wantErr {
				assert.Error(t, err)
				if tt.name == "database error" {
					assert.EqualError(t, err, "error checking user details")
				} else {
					assert.EqualError(t, err, "error checking user details")
				}
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)

			if err := mockSql.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
