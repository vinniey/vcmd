package vcmd

import (
	"testing"
)

// test cmd
func Test_Cmd_1(t *testing.T) {
	//os.Args
	//"cmd_name address=0.0.0.0 port=8080"
	args := make([]string, 3)
	args[0] = "cmd_name"
	args[1] = "address=0.0.0.0"
	args[2] = "port=8080"
	cmd := New(args)
	err := cmd.getError()
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(cmd.ToString())
	}

}
