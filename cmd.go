package vcmd

import (
	"strings"
	"errors"
)

type Arg struct {
	K string
	V string
}

type cmd struct {
	s    string
	args map[string]*Arg
	err  error
}

var cmdInstance *cmd

// parse args
func parse(args []string) *cmd {
	s := parseString(args)
	cmdInstance = &cmd{s, make(map[string]*Arg), nil}
	length := len(args)
	if length < 1 {
		cmdInstance.err = errors.New("not found args")
	} else {
		for i := 0; i < length; i++ {
			arg := parseParam(args[i])
			cmdInstance.args[arg.K] = arg

		}
	}

	return cmdInstance
}

// parse os.Args to string
func parseString(args []string) string {
	s := ""
	length := len(args)
	for i := 0; i < length; i++ {
		s += args[i] + " "
	}

	return s
}

// parse string to Arg
func parseParam(s string) *Arg {
	param := strings.Split(s, "=")
	arg := &Arg{}
	length := len(param)
	if length == 2 {
		arg.K = param[0]
		arg.V = param[1]
	}

	if length == 1 {
		arg.K = param[0]
		arg.V = ""
	}

	return arg
}

// Get all args
func (c *cmd) GetAll() (map[string]*Arg, error) {
	if c.err != nil {
		return nil, c.err
	}
	return c.args, nil
}

// Get a arg by name
func (c *cmd) Get(name string) (*Arg, error) {
	if c.err != nil {
		return nil, c.err
	}
	if v, ok := c.args[name]; ok {
		return v, nil
	} else {
		return nil, errors.New("not found arg")
	}

}
// cmd to string
func (c *cmd) ToString() string {
	return c.s
}

// get cmd error
func (c *cmd) getError() error {
	return c.err
}


// get cmd instance
func New(args []string) *cmd {
	if cmdInstance != nil {
		return cmdInstance
	}
	cmdInstance = parse(args)
	return cmdInstance
}
