package main

import (
	"errors"
	"sync"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	command_list map[string]func(*state, command) error
	mux          sync.RWMutex
}

func (c *commands) run(s *state, cmd command) error {
	c.mux.RLock()
	defer c.mux.RUnlock()

	com, exists := c.command_list[cmd.Name]
	if !exists {
		return errors.New("Invalid command")
	}
	return com(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.command_list[name] = f
}
