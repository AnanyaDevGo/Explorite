// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/interface/user.go

// Package mock is a generated GoMock package.
package mock

import (
	domain "authservice/pkg/domain"
	models "authservice/pkg/utils/models"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// AddProfile mocks base method.
func (m *MockUserRepository) AddProfile(id int, profile models.UserProfile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProfile", id, profile)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProfile indicates an expected call of AddProfile.
func (mr *MockUserRepositoryMockRecorder) AddProfile(id, profile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProfile", reflect.TypeOf((*MockUserRepository)(nil).AddProfile), id, profile)
}

// ChangePassword mocks base method.
func (m *MockUserRepository) ChangePassword(id int, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", id, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockUserRepositoryMockRecorder) ChangePassword(id, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockUserRepository)(nil).ChangePassword), id, password)
}

// CheckUserExistsByEmail mocks base method.
func (m *MockUserRepository) CheckUserExistsByEmail(email string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserExistsByEmail", email)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserExistsByEmail indicates an expected call of CheckUserExistsByEmail.
func (mr *MockUserRepositoryMockRecorder) CheckUserExistsByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserExistsByEmail", reflect.TypeOf((*MockUserRepository)(nil).CheckUserExistsByEmail), email)
}

// DeleteRecentOtpRequestsBefore5min mocks base method.
func (m *MockUserRepository) DeleteRecentOtpRequestsBefore5min() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRecentOtpRequestsBefore5min")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRecentOtpRequestsBefore5min indicates an expected call of DeleteRecentOtpRequestsBefore5min.
func (mr *MockUserRepositoryMockRecorder) DeleteRecentOtpRequestsBefore5min() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRecentOtpRequestsBefore5min", reflect.TypeOf((*MockUserRepository)(nil).DeleteRecentOtpRequestsBefore5min))
}

// EditProfile mocks base method.
func (m *MockUserRepository) EditProfile(id int, user models.EditProfile) (models.EditProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditProfile", id, user)
	ret0, _ := ret[0].(models.EditProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditProfile indicates an expected call of EditProfile.
func (mr *MockUserRepositoryMockRecorder) EditProfile(id, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditProfile", reflect.TypeOf((*MockUserRepository)(nil).EditProfile), id, user)
}

// FindUserByEmail mocks base method.
func (m *MockUserRepository) FindUserByEmail(user models.UserLogin) (models.UserSignup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", user)
	ret0, _ := ret[0].(models.UserSignup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockUserRepositoryMockRecorder) FindUserByEmail(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindUserByEmail), user)
}

// GetPassword mocks base method.
func (m *MockUserRepository) GetPassword(id int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPassword", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPassword indicates an expected call of GetPassword.
func (mr *MockUserRepositoryMockRecorder) GetPassword(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPassword", reflect.TypeOf((*MockUserRepository)(nil).GetPassword), id)
}

// GetProfile mocks base method.
func (m *MockUserRepository) GetProfile(id int) (models.UserProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfile", id)
	ret0, _ := ret[0].(models.UserProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfile indicates an expected call of GetProfile.
func (mr *MockUserRepositoryMockRecorder) GetProfile(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockUserRepository)(nil).GetProfile), id)
}

// GetUserName mocks base method.
func (m *MockUserRepository) GetUserName(email string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserName", email)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserName indicates an expected call of GetUserName.
func (mr *MockUserRepositoryMockRecorder) GetUserName(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserName", reflect.TypeOf((*MockUserRepository)(nil).GetUserName), email)
}

// IsValidEmail mocks base method.
func (m *MockUserRepository) IsValidEmail(email string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidEmail", email)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValidEmail indicates an expected call of IsValidEmail.
func (mr *MockUserRepositoryMockRecorder) IsValidEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidEmail", reflect.TypeOf((*MockUserRepository)(nil).IsValidEmail), email)
}

// IsValidWebsite mocks base method.
func (m *MockUserRepository) IsValidWebsite(website string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidWebsite", website)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValidWebsite indicates an expected call of IsValidWebsite.
func (mr *MockUserRepositoryMockRecorder) IsValidWebsite(website interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidWebsite", reflect.TypeOf((*MockUserRepository)(nil).IsValidWebsite), website)
}

// TemporarySavingUserOtp mocks base method.
func (m *MockUserRepository) TemporarySavingUserOtp(otp int, userEmail string, expiration time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TemporarySavingUserOtp", otp, userEmail, expiration)
	ret0, _ := ret[0].(error)
	return ret0
}

// TemporarySavingUserOtp indicates an expected call of TemporarySavingUserOtp.
func (mr *MockUserRepositoryMockRecorder) TemporarySavingUserOtp(otp, userEmail, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TemporarySavingUserOtp", reflect.TypeOf((*MockUserRepository)(nil).TemporarySavingUserOtp), otp, userEmail, expiration)
}

// UserSignUp mocks base method.
func (m *MockUserRepository) UserSignUp(userDetails models.UserSignup) (models.UserDetailResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserSignUp", userDetails)
	ret0, _ := ret[0].(models.UserDetailResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserSignUp indicates an expected call of UserSignUp.
func (mr *MockUserRepositoryMockRecorder) UserSignUp(userDetails interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserSignUp", reflect.TypeOf((*MockUserRepository)(nil).UserSignUp), userDetails)
}

// ValidateAlphabets mocks base method.
func (m *MockUserRepository) ValidateAlphabets(data string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateAlphabets", data)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateAlphabets indicates an expected call of ValidateAlphabets.
func (mr *MockUserRepositoryMockRecorder) ValidateAlphabets(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateAlphabets", reflect.TypeOf((*MockUserRepository)(nil).ValidateAlphabets), data)
}

// ValidatePhoneNumber mocks base method.
func (m *MockUserRepository) ValidatePhoneNumber(phone string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatePhoneNumber", phone)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ValidatePhoneNumber indicates an expected call of ValidatePhoneNumber.
func (mr *MockUserRepositoryMockRecorder) ValidatePhoneNumber(phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatePhoneNumber", reflect.TypeOf((*MockUserRepository)(nil).ValidatePhoneNumber), phone)
}

// VerifyOTP mocks base method.
func (m *MockUserRepository) VerifyOTP(email, otp string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyOTP", email, otp)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyOTP indicates an expected call of VerifyOTP.
func (mr *MockUserRepositoryMockRecorder) VerifyOTP(email, otp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyOTP", reflect.TypeOf((*MockUserRepository)(nil).VerifyOTP), email, otp)
}
