package mock

import (
	"gopkg.in/go-mixed/framework.v1/container"
	accessmocks "gopkg.in/go-mixed/framework.v1/contracts/auth/access/mocks"
	authmocks "gopkg.in/go-mixed/framework.v1/contracts/auth/mocks"
	cachemocks "gopkg.in/go-mixed/framework.v1/contracts/cache/mocks"
	configmocks "gopkg.in/go-mixed/framework.v1/contracts/config/mocks"
	consolemocks "gopkg.in/go-mixed/framework.v1/contracts/console/mocks"
	ormmocks "gopkg.in/go-mixed/framework.v1/contracts/database/orm/mocks"
	eventmocks "gopkg.in/go-mixed/framework.v1/contracts/event/mocks"
	filesystemmocks "gopkg.in/go-mixed/framework.v1/contracts/filesystem/mocks"
	grpcmocks "gopkg.in/go-mixed/framework.v1/contracts/grpc/mocks"
	mailmocks "gopkg.in/go-mixed/framework.v1/contracts/mail/mocks"
	queuemocks "gopkg.in/go-mixed/framework.v1/contracts/queue/mocks"
	validatemocks "gopkg.in/go-mixed/framework.v1/contracts/validation/mocks"
	"gopkg.in/go-mixed/framework.v1/log"
)

func Cache() *cachemocks.Store {
	mockCache := &cachemocks.Store{}
	container.Instance("cache.store", mockCache)

	return mockCache
}

func Config() *configmocks.Config {
	mockConfig := &configmocks.Config{}
	container.Instance("config", mockConfig)

	return mockConfig
}

func Artisan() *consolemocks.Artisan {
	mockArtisan := &consolemocks.Artisan{}
	container.Instance("artisan", mockArtisan)

	return mockArtisan
}

func Orm() (*ormmocks.Orm, *ormmocks.DB, *ormmocks.Transaction, *ormmocks.Association) {
	mockOrm := &ormmocks.Orm{}
	container.Instance("orm", mockOrm)

	return mockOrm, &ormmocks.DB{}, &ormmocks.Transaction{}, &ormmocks.Association{}
}

func Event() (*eventmocks.Instance, *eventmocks.Task) {
	mockEvent := &eventmocks.Instance{}
	container.Instance("event", mockEvent)

	return mockEvent, &eventmocks.Task{}
}

func Log() {
	container.Instance("log", log.WrapLogger(log.NewTestWriter()))
}

func Mail() *mailmocks.Mail {
	mockMail := &mailmocks.Mail{}
	container.Instance("mail", mockMail)
	//facades.Mail = mockMail

	return mockMail
}

func Queue() *queuemocks.Queue {
	mockQueue := &queuemocks.Queue{}

	return mockQueue
}

func Storage() (*filesystemmocks.Storage, *filesystemmocks.Driver, *filesystemmocks.File) {
	mockStorage := &filesystemmocks.Storage{}
	mockDriver := &filesystemmocks.Driver{}
	mockFile := &filesystemmocks.File{}

	container.Instance("filesystem", mockStorage)

	return mockStorage, mockDriver, mockFile
}

func Validation() (*validatemocks.Validation, *validatemocks.Validator, *validatemocks.Errors) {
	mockValidation := &validatemocks.Validation{}
	mockValidator := &validatemocks.Validator{}
	mockErrors := &validatemocks.Errors{}
	container.Instance("validation", mockValidation)

	return mockValidation, mockValidator, mockErrors
}

func Auth() *authmocks.Auth {
	mockAuth := &authmocks.Auth{}
	container.Instance("auth", mockAuth)

	return mockAuth
}

func Gate() *accessmocks.Gate {
	mockGate := &accessmocks.Gate{}
	container.Instance("gate", mockGate)

	return mockGate
}

func Grpc() *grpcmocks.Grpc {
	mockGrpc := &grpcmocks.Grpc{}
	container.Instance("grpc", mockGrpc)

	return mockGrpc
}
