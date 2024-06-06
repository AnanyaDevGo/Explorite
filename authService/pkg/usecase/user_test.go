 package usecase

// import (
// 	"authservice/pkg/domain"
// 	"authservice/pkg/usecase"
// 	"authservice/pkg/utils/models"
// 	"errors"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"golang.org/x/crypto/bcrypt"
// 	"gopkg.in/go-playground/assert.v1"
// 	"honnef.co/go/tools/config"
// )

// func Test_LoginHandler(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	userRepo := mockRepository.NewMockUserRepository(ctrl)
// 	helper := mockhelper.NewMockHelper(ctrl)
// 	cfg := config.Config{}

// 	userUseCase := usecase.NewUserUseCase(cfg, helper, userRepo)
// 	password := "password123"
// 	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

// 	testingData := map[string]struct {
// 		input   models.UserLogin
// 		stub    func(*mockRepository.MockUserRepository, *mockhelper.MockHelper, models.UserLogin)
// 		want    *domain.TokenUser
// 		wantErr error
// 	}{
// 		"successful login": {
// 			input: models.UserLogin{
// 				Email:    "testuser@gmail.com",
// 				Password: password,
// 			},
// 			stub: func(mur *mockRepository.MockUserRepository, mh *mockhelper.MockHelper, ul models.UserLogin) {
// 				mur.EXPECT().CheckUserExistsByEmail(ul.Email).Times(1).Return(true, nil)
// 				mur.EXPECT().FindUserByEmail(ul).Times(1).Return(models.UserSignup{
// 					ID:          1,
// 					Email:       "testuser@gmail.com",
// 					Password:    string(hashedPassword),
// 					Firstname:   "Test",
// 					Lastname:    "User",
// 					PhoneNumber: "1234567890",
// 					DateOfBirth: "1990-01-01",
// 					Gender:      "male",
// 					Bio:         "bio",
// 				}, nil)
// 				mh.EXPECT().GenerateTokenUser(gomock.Any()).Times(1).Return("token_string", nil)
// 			},
// 			want: &domain.TokenUser{
// 				User: models.UserDetailResponse{
// 					ID:          1,
// 					Firstname:   "Test",
// 					Lastname:    "User",
// 					Email:       "testuser@gmail.com",
// 					PhoneNumber: "1234567890",
// 					DateOfBirth: "1990-01-01",
// 					Gender:      "male",
// 					Bio:         "bio",
// 				},
// 				Token: "token_string",
// 			},
// 			wantErr: nil,
// 		},
// 		"user not found": {
// 			input: models.UserLogin{
// 				Email:    "nonexistent@gmail.com",
// 				Password: password,
// 			},
// 			stub: func(mur *mockRepository.MockUserRepository, mh *mockhelper.MockHelper, ul models.UserLogin) {
// 				mur.EXPECT().CheckUserExistsByEmail(ul.Email).Times(1).Return(false, nil)
// 			},
// 			want:    &domain.TokenUser{},
// 			wantErr: errors.New("email doesn't exist"),
// 		},
// 		"incorrect password": {
// 			input: models.UserLogin{
// 				Email:    "testuser@gmail.com",
// 				Password: "wrongpassword",
// 			},
// 			stub: func(mur *mockRepository.MockUserRepository, mh *mockhelper.MockHelper, ul models.UserLogin) {
// 				mur.EXPECT().CheckUserExistsByEmail(ul.Email).Times(1).Return(true, nil)
// 				mur.EXPECT().FindUserByEmail(ul).Times(1).Return(models.UserSignup{
// 					ID:          1,
// 					Email:       "testuser@gmail.com",
// 					Password:    string(hashedPassword),
// 					Firstname:   "Test",
// 					Lastname:    "User",
// 					PhoneNumber: "1234567890",
// 					DateOfBirth: "1990-01-01",
// 					Gender:      "male",
// 					Bio:         "bio",
// 				}, nil)
// 			},
// 			want:    &domain.TokenUser{},
// 			wantErr: errors.New("password not matching"),
// 		},
// 		"server error": {
// 			input: models.UserLogin{
// 				Email:    "testuser@gmail.com",
// 				Password: password,
// 			},
// 			stub: func(mur *mockRepository.MockUserRepository, mh *mockhelper.MockHelper, ul models.UserLogin) {
// 				mur.EXPECT().CheckUserExistsByEmail(ul.Email).Times(1).Return(true, errors.New("server error"))
// 			},
// 			want:    &domain.TokenUser{},
// 			wantErr: errors.New("error with server"),
// 		},
// 	}

// 	for name, test := range testingData {
// 		t.Run(name, func(t *testing.T) {
// 			test.stub(userRepo, helper, test.input)
// 			result, err := userUseCase.LoginHandler(test.input)
// 			assert.Equal(t, test.want, result)
// 			assert.Equal(t, test.wantErr, err)
// 		})
// 	}
// }
