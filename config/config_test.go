package config

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"gopkg.in/go-mixed/framework.v1/support/file"
)

type ApplicationTestSuite struct {
	suite.Suite
	config *Config
}

func TestApplicationTestSuite(t *testing.T) {
	file.Create(".env", "APP_KEY=12345678901234567890123456789012")

	suite.Run(t, &ApplicationTestSuite{
		config: NewModule(".env"),
	})

	file.Remove(".env")
}

func (s *ApplicationTestSuite) SetupTest() {

}

func (s *ApplicationTestSuite) TestEnv() {
	s.Equal("goravel", s.config.Env("APP_NAME", "goravel").(string))
	s.Equal("127.0.0.1", s.config.Env("DB_HOST", "127.0.0.1").(string))
}

func (s *ApplicationTestSuite) TestAdd() {
	s.config.Add("app", map[string]any{
		"env": "local",
	})

	s.Equal("local", s.config.GetString("app.env"))
}

func (s *ApplicationTestSuite) TestGet() {
	s.Equal("goravel", s.config.Get("APP_NAME", "goravel").(string))
}

func (s *ApplicationTestSuite) TestGetString() {
	s.config.Add("database", map[string]any{
		"default": s.config.Env("DB_CONNECTION", "mysql"),
		"connections": map[string]any{
			"mysql": map[string]any{
				"host": s.config.Env("DB_HOST", "127.0.0.1"),
			},
		},
	})

	s.Equal("goravel", s.config.GetString("APP_NAME", "goravel"))
	s.Equal("127.0.0.1", s.config.GetString("database.connections.mysql.host"))
	s.Equal("mysql", s.config.GetString("database.default"))
}

func (s *ApplicationTestSuite) TestGetInt() {
	s.Equal(s.config.GetInt("DB_PORT", 3306), 3306)
}

func (s *ApplicationTestSuite) TestGetBool() {
	s.Equal(true, s.config.GetBool("APP_DEBUG", true))
}
