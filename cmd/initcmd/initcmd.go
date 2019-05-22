package initcmd

import (
	"os"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "init",
	RunE: func(_ *cobra.Command, args []string) error {
		logrus.Infof("init container")
		cmd := args[0]
		logrus.Infof("command: %s", cmd)
		return RunContainerInitProcess(cmd)
	},
}

func RunContainerInitProcess(cmd string) error {
	logrus.Infof("command: %s", cmd)

	syscall.Mount("", "/", "", syscall.MS_PRIVATE|syscall.MS_REC, "")
	defaultMountFlag := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlag), "")
	argv := []string{cmd}
	if err := syscall.Exec(cmd, argv, os.Environ()); err != nil {
		logrus.Errorf("command exec failed: %s", err)
	}
	return nil
}
