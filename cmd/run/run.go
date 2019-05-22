package run

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	tty bool
)

func init() {
	Cmd.PersistentFlags().BoolVarP(&tty, "tty", "t", false, "enable tty")
}

var Cmd = &cobra.Command{
	Use: "run",
	RunE: func(_ *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("missing container command")
		}
		cmd := args[0]
		Run(tty, cmd)
		return nil
	},
}

func NewParentProcess(tty bool, cmd string) *exec.Cmd {
	args := []string{"init", cmd}
	command := exec.Command("/proc/self/exe", args...)
	command.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}
	if tty {
		//command.Stdin = os.Stdin
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
	}
	return command
}

func Run(tty bool, cmd string) {
	parent := NewParentProcess(tty, cmd)
	if err := parent.Start(); err != nil {
		logrus.Error(err)
		os.Exit(-1)
	}
	if err := parent.Wait(); err != nil {
		logrus.Error(err)
	}
	os.Exit(-1)
}
