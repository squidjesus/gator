package main

import (
	"errors"
)

type command struct {
	name string
	args []string
}

type commands struct {
	registry map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if _, ok := c.registry[cmd.name]; !ok {
		return errors.New("Command name not found")
	}
	err := c.registry[cmd.name](s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registry[name] = f
	return
}



