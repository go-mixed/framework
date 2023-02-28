package foundation

import (
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/facades/artisan"
	configfacade "gopkg.in/go-mixed/framework.v1/facades/config"
	logfacade "gopkg.in/go-mixed/framework.v1/facades/log"
	"gopkg.in/go-mixed/framework.v1/log"
	"os"
	"strings"

	"gopkg.in/go-mixed/framework.v1/config"
	"gopkg.in/go-mixed/framework.v1/contracts"
	"gopkg.in/go-mixed/framework.v1/support"
)

type Application struct {
}

func NewApplication() *Application {
	setEnv()

	container.Initialize()
	app := &Application{}
	app.registerBaseServiceProviders()
	app.bootBaseServiceProviders()

	return app
}

// Boot Register and bootstrap configured service providers.
func (app *Application) Boot() {
	logfacade.Debugf("Application start.")

	app.registerConfiguredServiceProviders()
	app.bootConfiguredServiceProviders()

	app.bootArtisan()
	setRootPath()
}

// bootArtisan Boot artisan command.
func (app *Application) bootArtisan() {
	artisan.Run(os.Args, true)
}

// getBaseServiceProviders Get base service providers.
func (app *Application) getBaseServiceProviders() []contracts.IServiceProvider {
	return []contracts.IServiceProvider{
		&config.ServiceProvider{},
		&log.ServiceProvider{},
	}
}

// getConfiguredServiceProviders Get configured service providers.
func (app *Application) getConfiguredServiceProviders() []contracts.IServiceProvider {
	return configfacade.Get("app.providers").([]contracts.IServiceProvider)
}

// registerBaseServiceProviders Register base service providers.
func (app *Application) registerBaseServiceProviders() {
	app.registerServiceProviders(app.getBaseServiceProviders())
}

// bootBaseServiceProviders Bootstrap base service providers.
func (app *Application) bootBaseServiceProviders() {
	app.bootServiceProviders(app.getBaseServiceProviders())
}

// registerConfiguredServiceProviders Register configured service providers.
func (app *Application) registerConfiguredServiceProviders() {
	app.registerServiceProviders(app.getConfiguredServiceProviders())
}

// bootConfiguredServiceProviders Bootstrap configured service providers.
func (app *Application) bootConfiguredServiceProviders() {
	app.bootServiceProviders(app.getConfiguredServiceProviders())
}

// registerServiceProviders Register service providers.
func (app *Application) registerServiceProviders(serviceProviders []contracts.IServiceProvider) {
	for _, serviceProvider := range serviceProviders {
		serviceProvider.Register()
	}
}

// bootServiceProviders Bootstrap service providers.
func (app *Application) bootServiceProviders(serviceProviders []contracts.IServiceProvider) {
	for _, serviceProvider := range serviceProviders {
		serviceProvider.Boot()
	}
}

func setEnv() {
	args := os.Args
	if len(args) >= 2 {
		if args[1] == "artisan" {
			support.Env = support.EnvArtisan
		}
	}
}

func setRootPath() {
	rootPath := getCurrentAbsPath()

	// Hack air path
	airPath := "/storage/temp"
	if strings.HasSuffix(rootPath, airPath) {
		rootPath = strings.ReplaceAll(rootPath, airPath, "")
	}

	support.RootPath = rootPath
}
