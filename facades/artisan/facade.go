package artisan

import (
	"gopkg.in/go-mixed/framework.v1/container"
	"gopkg.in/go-mixed/framework.v1/contracts/console"
)

func getArtisan() console.IArtisan {
	return container.MustMakeAs("artisan", console.IArtisan(nil))
}

func Register(commands []console.Command) {
	getArtisan().Register(commands)
}

// Call Run an Artisan console command by name.
func Call(command string) {
	getArtisan().Call(command)
}

// CallAndExit Run an Artisan console command by name and exit.
func CallAndExit(command string) {
	getArtisan().CallAndExit(command)
}

// Run a command. args include: ["./main", "artisan", "command"]
func Run(args []string, exitIfArtisan bool) {
	getArtisan().Run(args, exitIfArtisan)
}
