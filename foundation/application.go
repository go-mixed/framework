package foundation

import (
	"gopkg.in/go-mixed/framework/container"
	"os"
	"strings"

	"gopkg.in/go-mixed/framework/config"
	"gopkg.in/go-mixed/framework/contracts"
	"gopkg.in/go-mixed/framework/facades"
	"gopkg.in/go-mixed/framework/support"
)

var app Application

func init() {
	setEnv()

	app = Application{
		Container: container.Initial(),
	}
	app.registerBaseServiceProviders()
	app.bootBaseServiceProviders()
}

type Application struct {
	container.Container
}

// Boot Register and bootstrap configured service providers.
func (app *Application) Boot() {
	app.registerConfiguredServiceProviders()
	app.bootConfiguredServiceProviders()

	app.bootArtisan()
	setRootPath()
}

// bootArtisan Boot artisan command.
func (app *Application) bootArtisan() {
	facades.Artisan.Run(os.Args, true)
}

// getBaseServiceProviders Get base service providers.
func (app *Application) getBaseServiceProviders() []contracts.IServiceProvider {
	return []contracts.IServiceProvider{
		&config.ServiceProvider{},
	}
}

// getConfiguredServiceProviders Get configured service providers.
func (app *Application) getConfiguredServiceProviders() []contracts.IServiceProvider {
	return facades.Config.Get("app.providers").([]contracts.IServiceProvider)
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
		serviceProvider.Register(app)
	}
}

// bootServiceProviders Bootstrap service providers.
func (app *Application) bootServiceProviders(serviceProviders []contracts.IServiceProvider) {
	for _, serviceProvider := range serviceProviders {
		serviceProvider.Boot(app)
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
	rootPath := getCurrentAbPath()

	// Hack air path
	airPath := "/storage/temp"
	if strings.HasSuffix(rootPath, airPath) {
		rootPath = strings.ReplaceAll(rootPath, airPath, "")
	}

	support.RootPath = rootPath
}
