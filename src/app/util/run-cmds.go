package util

import (
	"bytes"
	"fmt"
	"os/exec"
)

func RunCmds(cmds []*exec.Cmd) error {
	for _, c := range cmds {
		var stderr bytes.Buffer
		c.Stderr = &stderr

		if err := c.Run(); err != nil {
			return fmt.Errorf("command %v failed: %v\nstderr:\n%s", c.Args, err, stderr.String())
		}
	}
	return nil
}
