package log

import (
	"context"
	"fmt"
	"gopkg.in/go-mixed/framework.v1/container"
	"os"
	"os/exec"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	configmocks "gopkg.in/go-mixed/framework.v1/contracts/config/mocks"
	"gopkg.in/go-mixed/framework.v1/support/file"
	"gopkg.in/go-mixed/framework.v1/support/time"
)

var singleLog = "storage/logs/laravel.log"
var dailyLog = fmt.Sprintf("storage/logs/laravel-%s.log", time.Now().Format("2006-01-02"))

type LogrusTestSuite struct {
	suite.Suite
}

func TestLogrusTestSuite(t *testing.T) {
	suite.Run(t, new(LogrusTestSuite))
}

func (s *LogrusTestSuite) SetupTest() {

}

func (s *LogrusTestSuite) TestLogrus() {
	var (
		mockConfig *configmocks.Config
		log        *Logger
	)

	beforeEach := func() {
		mockConfig = initMockConfig()
	}

	tests := []struct {
		name   string
		setup  func()
		assert func()
	}{
		{
			name: "WithContext",
			setup: func() {
				mockConfig.On("GetString", "logging.channels.daily.level").Return("debug").Once()
				mockConfig.On("GetString", "logging.channels.single.level").Return("debug").Once()

				log, _ = NewLogger(context.Background(), "stack")
			},
			assert: func() {
				writer := log.WithContext(context.Background())
				assert.Equal(s.T(), reflect.TypeOf(writer).String(), reflect.TypeOf(&Writer{}).String())
			},
		},
		{
			name: "Debug",
			setup: func() {
				mockDriverConfig(mockConfig)

				log, _ = NewLogger(context.Background(), "stack")
				log.Debug("Laravel")
			},
			assert: func() {
				assert.True(s.T(), file.Exists(dailyLog))
				assert.True(s.T(), file.Exists(singleLog))
				assert.True(s.T(), file.Contain(singleLog, "test.debug: Laravel"))
				assert.True(s.T(), file.Contain(dailyLog, "test.debug: Laravel"))
			},
		},
		{
			name: "No Debug",
			setup: func() {
				mockConfig.On("GetString", "logging.channels.daily.level").Return("info").Once()
				mockConfig.On("GetString", "logging.channels.single.level").Return("info").Once()

				log, _ = NewLogger(context.Background(), "stack")
				log.Debug("Laravel")
			},
			assert: func() {
				assert.False(s.T(), file.Exists(dailyLog))
				assert.False(s.T(), file.Exists(singleLog))
			},
		},
		{
			name: "Debugf",
			setup: func() {
				mockDriverConfig(mockConfig)

				log, _ = NewLogger(context.Background(), "stack")
				log.Debugf("Laravel: %s", "World")
			},
			assert: func() {
				assert.True(s.T(), file.Exists(dailyLog))
				assert.True(s.T(), file.Exists(singleLog))
				assert.True(s.T(), file.Contain(singleLog, "test.debug: Laravel: World"))
				assert.True(s.T(), file.Contain(dailyLog, "test.debug: Laravel: World"))
			},
		},
		{
			name: "Info",
			setup: func() {
				mockDriverConfig(mockConfig)

				log, _ = NewLogger(context.Background(), "stack")
				log.Info("Laravel")
			},
			assert: func() {
				assert.True(s.T(), file.Exists(dailyLog))
				assert.True(s.T(), file.Exists(singleLog))
				assert.True(s.T(), file.Contain(singleLog, "test.info: Laravel"))
				assert.True(s.T(), file.Contain(dailyLog, "test.info: Laravel"))
			},
		},
		{
			name: "Infof",
			setup: func() {
				mockDriverConfig(mockConfig)

				log, _ = NewLogger(context.Background(), "stack")
				log.Infof("Laravel: %s", "World")
			},
			assert: func() {
				assert.True(s.T(), file.Exists(dailyLog))
				assert.True(s.T(), file.Exists(singleLog))
				assert.True(s.T(), file.Contain(singleLog, "test.info: Laravel: World"))
				assert.True(s.T(), file.Contain(dailyLog, "test.info: Laravel: World"))
			},
		},
		{
			name: "Warning",
			setup: func() {
				mockDriverConfig(mockConfig)

				log, _ = NewLogger(context.Background(), "stack")
				log.Warning("Laravel")
			},
			assert: func() {
				assert.True(s.T(), file.Exists(dailyLog))
				assert.True(s.T(), file.Exists(singleLog))
				assert.True(s.T(), file.Contain(singleLog, "test.warning: Laravel"))
				assert.True(s.T(), file.Contain(dailyLog, "test.warning: Laravel"))
			},
		},
		{
			name: "Warningf",
			setup: func() {
				mockDriverConfig(mockConfig)

				log, _ = NewLogger(context.Background(), "stack")
				log.Warningf("Laravel: %s", "World")
			},
			assert: func() {
				assert.True(s.T(), file.Exists(dailyLog))
				assert.True(s.T(), file.Exists(singleLog))
				assert.True(s.T(), file.Contain(singleLog, "test.warning: Laravel: World"))
				assert.True(s.T(), file.Contain(dailyLog, "test.warning: Laravel: World"))
			},
		},
		{
			name: "Error",
			setup: func() {
				mockDriverConfig(mockConfig)

				log, _ = NewLogger(context.Background(), "stack")
				log.Error("Laravel")
			},
			assert: func() {
				assert.True(s.T(), file.Exists(dailyLog))
				assert.True(s.T(), file.Exists(singleLog))
				assert.True(s.T(), file.Contain(singleLog, "test.error: Laravel"))
				assert.True(s.T(), file.Contain(dailyLog, "test.error: Laravel"))
			},
		},
		{
			name: "Errorf",
			setup: func() {
				mockDriverConfig(mockConfig)

				log, _ = NewLogger(context.Background(), "stack")
				log.Errorf("Laravel: %s", "World")
			},
			assert: func() {
				assert.True(s.T(), file.Exists(dailyLog))
				assert.True(s.T(), file.Exists(singleLog))
				assert.True(s.T(), file.Contain(singleLog, "test.error: Laravel: World"))
				assert.True(s.T(), file.Contain(dailyLog, "test.error: Laravel: World"))
			},
		},
		{
			name: "Panic",
			setup: func() {
				mockDriverConfig(mockConfig)

				log, _ = NewLogger(context.Background(), "stack")
			},
			assert: func() {
				assert.Panics(s.T(), func() {
					log.Panic("Laravel")
				})
				assert.True(s.T(), file.Exists(dailyLog))
				assert.True(s.T(), file.Exists(singleLog))
				assert.True(s.T(), file.Contain(singleLog, "test.panic: Laravel"))
				assert.True(s.T(), file.Contain(dailyLog, "test.panic: Laravel"))
			},
		},
		{
			name: "Panicf",
			setup: func() {
				mockDriverConfig(mockConfig)

				log, _ = NewLogger(context.Background(), "stack")
			},
			assert: func() {
				assert.Panics(s.T(), func() {
					log.Panicf("Laravel: %s", "World")
				})
				assert.True(s.T(), file.Exists(dailyLog))
				assert.True(s.T(), file.Exists(singleLog))
				assert.True(s.T(), file.Contain(singleLog, "test.panic: Laravel: World"))
				assert.True(s.T(), file.Contain(dailyLog, "test.panic: Laravel: World"))
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			beforeEach()
			test.setup()
			test.assert()
			mockConfig.AssertExpectations(s.T())
			file.Remove("storage")
		})
	}
}

func (s *LogrusTestSuite) TestTestWriter() {
	log := WrapLogger(NewTestWriter())
	assert.Equal(s.T(), log.WithContext(context.Background()), &TestWriter{})
	assert.NotPanics(s.T(), func() {
		log.Debug("Laravel")
		log.Debugf("Laravel")
		log.Info("Laravel")
		log.Infof("Laravel")
		log.Warning("Laravel")
		log.Warningf("Laravel")
		log.Error("Laravel")
		log.Errorf("Laravel")
		log.Fatal("Laravel")
		log.Fatalf("Laravel")
		log.Panic("Laravel")
		log.Panicf("Laravel")
	})
}

func TestLogrus_Fatal(t *testing.T) {
	mockConfig := initMockConfig()
	mockDriverConfig(mockConfig)
	log, _ := NewLogger(context.Background(), "stack")

	if os.Getenv("FATAL") == "1" {
		log.Fatal("Laravel")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestLogrus_Fatal")
	cmd.Env = append(os.Environ(), "FATAL=1")
	err := cmd.Run()

	assert.EqualError(t, err, "exit status 1")
	assert.True(t, file.Exists(dailyLog))
	assert.True(t, file.Exists(singleLog))
	assert.True(t, file.Contain(singleLog, "test.fatal: Laravel"))
	assert.True(t, file.Contain(dailyLog, "test.fatal: Laravel"))
	file.Remove("storage")
}

func TestLogrus_Fatalf(t *testing.T) {
	mockConfig := initMockConfig()
	mockDriverConfig(mockConfig)
	log, _ := NewLogger(context.Background(), "stack")

	if os.Getenv("FATAL") == "1" {
		log.Fatalf("Laravel")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestLogrus_Fatal")
	cmd.Env = append(os.Environ(), "FATAL=1")
	err := cmd.Run()

	assert.EqualError(t, err, "exit status 1")
	assert.True(t, file.Exists(dailyLog))
	assert.True(t, file.Exists(singleLog))
	assert.True(t, file.Contain(singleLog, "test.fatal: Laravel"))
	assert.True(t, file.Contain(dailyLog, "test.fatal: Laravel"))
	file.Remove("storage")
}

func initMockConfig() *configmocks.Config {
	mockConfig := &configmocks.Config{}
	container.Instance("config", mockConfig)

	mockConfig.On("GetString", "logging.default").Return("stack").Once()
	mockConfig.On("GetString", "logging.channels.stack.driver").Return("stack").Once()
	mockConfig.On("Get", "logging.channels.stack.channels").Return([]string{"single", "daily"}).Once()
	mockConfig.On("GetString", "logging.channels.daily.driver").Return("daily").Once()
	mockConfig.On("GetString", "logging.channels.daily.path").Return(singleLog).Once()
	mockConfig.On("GetInt", "logging.channels.daily.days").Return(7).Once()
	mockConfig.On("GetBool", "logging.channels.daily.print").Return(false).Once()
	mockConfig.On("GetString", "logging.channels.single.driver").Return("single").Once()
	mockConfig.On("GetString", "logging.channels.single.path").Return(singleLog).Once()
	mockConfig.On("GetBool", "logging.channels.single.print").Return(false).Once()

	return mockConfig
}

func mockDriverConfig(mockConfig *configmocks.Config) {
	mockConfig.On("GetString", "logging.channels.daily.level").Return("debug").Once()
	mockConfig.On("GetString", "logging.channels.single.level").Return("debug").Once()
	mockConfig.On("GetString", "app.timezone").Return("UTC")
	mockConfig.On("GetString", "app.env").Return("test")
}
