package crypt

import (
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/facades/crypt"
	"gopkg.in/go-mixed/framework.v1/testing/mock"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ApplicationTestSuite struct {
	suite.Suite
}

func TestApplicationTestSuite(t *testing.T) {
	mockConfig := mock.Config()
	mockConfig.On("GetString", "app.key").Return("11111111111111111111111111111111").Once()

	container.Instance("crypt", NewCrypt())
	suite.Run(t, new(ApplicationTestSuite))
	mockConfig.AssertExpectations(t)
}

func (s *ApplicationTestSuite) SetupTest() {

}

func (s *ApplicationTestSuite) TestEncryptString() {
	encryptString, err := crypt.EncryptString("Goravel")
	s.NoError(err)
	s.NotEmpty(encryptString)
}

func (s *ApplicationTestSuite) TestDecryptString() {
	payload, err := crypt.EncryptString("Goravel")
	s.NoError(err)
	s.NotEmpty(payload)

	value, err := crypt.DecryptString(payload)
	s.NoError(err)
	s.Equal("Goravel", value)

	_, err = crypt.DecryptString("Goravel")
	s.Error(err)

	_, err = crypt.DecryptString("R29yYXZlbA==")
	s.Error(err)

	_, err = crypt.DecryptString("eyJpIjoiMTIzNDUiLCJ2YWx1ZSI6IjEyMzQ1In0=")
	s.Error(err)

	_, err = crypt.DecryptString("eyJpdiI6IjEyMzQ1IiwidiI6IjEyMzQ1In0=")
	s.Error(err)

	_, err = crypt.DecryptString("eyJpdiI6IjEyMzQ1IiwidmFsdWUiOiIxMjM0NSJ9")
	s.Error(err)
}
