//copied from https://github.com/hashicorp/packer/blob/81522dced0b25084a824e79efda02483b12dc7cd/builder/amazon/chroot/command.go

package builder

import (
	"os/exec"
)

// CommandWrapper is a type that given a command, will possibly modify that
// command in-flight. This might return an error.
type CommandWrapper func(string) (string, error)

// ShellCommand takes a command string and returns an *exec.Cmd to execute
// it within the context of a shell (/bin/sh).
func ShellCommand(command string) *exec.Cmd {
	return exec.Command("/bin/sh", "-c", command)
}
