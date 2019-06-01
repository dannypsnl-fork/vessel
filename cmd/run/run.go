package run

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/dannypsnl/vessel/cgroup"
	"github.com/dannypsnl/vessel/cgroup/subsystems"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	tty bool
	// cgroups related flags
	memoryLimit, cpuShare, cpuSet string
)

func init() {
	Cmd.PersistentFlags().BoolVarP(&tty, "tty", "t", false, "enable tty")
	Cmd.PersistentFlags().StringVar(&memoryLimit, "memory", "", "limitation of memory")
	Cmd.PersistentFlags().StringVar(&cpuShare, "cpu", "", "limitation of cpu share")
	Cmd.PersistentFlags().StringVar(&cpuSet, "cpuset", "", "limitation of cpu set")
}

var Cmd = &cobra.Command{
	Use: "run",
	RunE: func(_ *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("missing container command")
		}
		resource := &subsystems.ResourceConfig{
			MemoryLimit: memoryLimit,
			CPUShare:    cpuShare,
			CPUSet:      cpuSet,
		}
		return Run(tty, args, resource)
	},
}

func Run(tty bool, cmd []string, res *subsystems.ResourceConfig) error {
	parent, writePipe, err := NewParentProcess(tty)
	if err != nil {
		return fmt.Errorf("failed at new parent process: %s", err)
	}
	if err := parent.Start(); err != nil {
		return fmt.Errorf("failed at start parent process: %s", err)
	}
	cgroupManager := cgroup.NewManager("vessel-cgroup")
	defer cgroupManager.Destroy()
	cgroupManager.Set(res)
	cgroupManager.Apply(parent.Process.Pid)
	sendInitCommand(cmd, writePipe)
	if err := parent.Wait(); err != nil {
		return fmt.Errorf("process has some error: %s", err)
	}
	return nil
}

func sendInitCommand(cmd []string, writePipe *os.File) {
	command := strings.Join(cmd, " ")
	logrus.Infof("command all is `%s`", command)
	_, err := writePipe.WriteString(command)
	if err != nil {
		logrus.Warn(err)
	}
	err = writePipe.Close()
	if err != nil {
		logrus.Warn(err)
	}
}

func NewParentProcess(tty bool) (*exec.Cmd, *os.File, error) {
	readPipe, writePipe, err := NewPipe()
	if err != nil {
		return nil, nil, err
	}
	command := exec.Command("/proc/self/exe", "init")
	command.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}
	if tty {
		command.Stdin = os.Stdin
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
	}
	command.ExtraFiles = []*os.File{readPipe}
	return command, writePipe, nil
}

func NewPipe() (*os.File, *os.File, error) {
	read, write, err := os.Pipe()
	if err != nil {
		return nil, nil, err
	}
	return read, write, nil
}
