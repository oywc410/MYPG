package oTsl

import "net"

type Command struct {
	commands map[string]func(net.Conn) error
}

func NewCommand() *Command {
	return &Command{make(map[string]func(net.Conn) error)}
}

func (command *Command) AddCommand(key string, fun func(net.Conn) error) bool {

	_, ok := command.commands[key]
	if ok {
		return false
	}

	command.commands[key] = fun
	return true
}

func (command *Command) DelCommand(key string) {
	delete(command.commands, key)
}

func (command *Command) GetCommand(key string) (func(net.Conn) error, bool) {
	fun, ok := command.commands[key]
	if ok {
		return fun, true
	}
	return nil, false
}

func (command *Command) GetAllCommands() map[string]func(net.Conn) error {
	return command.commands
}
