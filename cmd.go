package vcmd

import (
	"os"
	"strings"
	"errors"
)

const NoErrorCode = 0
const ArgError = 1


type Arg struct {
	K string
	V string
}

type cmd struct {
	s       string
	args    map[string]*Arg
	errCode int
	errMsg  string
}

var cmdInstance *cmd

// parse args
func parse(args []string) *cmd {
	s := parseString(args)
	cmdInstance = &cmd{s,make(map[string]*Arg), NoErrorCode, ""}
	length := len(args)
	if length < 1 {
		cmdInstance.errCode = ArgError
		cmdInstance.errMsg = "not found args"
	} else {
		for i := 0; i < length; i++ {
			arg := parseParam(os.Args[i])
			cmdInstance.args[arg.K] = arg

		}
	}

	return cmdInstance
}

// parse os.Args to string
func parseString(args []string) string {
	s := ""
	length := len(args)
	for i:=0; i<length ; i++ {
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
	if c.errCode > 0 {
		return nil, errors.New(c.errMsg)
	}
	return c.args, nil
}

// Get a arg by name
func (c *cmd) Get(name string) (*Arg, error) {
	if c.errCode > 0 {
		return nil, errors.New(c.errMsg)
	}
	if v, ok := c.args[name]; ok {
		return v, nil
	} else {
		return nil, errors.New("not found arg")
	}

}

func (c *cmd) ToString() string {
	return c.s
}

// get cmd instance
func Cmd(args []string) *cmd {
	if cmdInstance != nil {
		return cmdInstance
	}
	cmdInstance = parse(args)
	return cmdInstance
}
