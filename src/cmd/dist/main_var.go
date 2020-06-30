package main

// commands records the available commands.
var commands = map[string]func(){
	"banner":    cmdbanner,
	"bootstrap": cmdbootstrap,
	"clean":     cmdclean,
	"env":       cmdenv,
	"install":   cmdinstall,
	"list":      cmdlist,
	"test":      cmdtest,
	"version":   cmdversion,
}
