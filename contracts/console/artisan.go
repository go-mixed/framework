package console

//go:generate mockery --name=Artisan
type IArtisan interface {
	//Register commands.
	Register(commands []ICommand)

	//Call Run an Artisan console command by name.
	Call(command string)

	//CallAndExit Run an Artisan console command by name and exit.
	CallAndExit(command string)

	//Run a command. args include: ["./main", "artisan", "command"]
	Run(args []string, exitIfArtisan bool)
}
