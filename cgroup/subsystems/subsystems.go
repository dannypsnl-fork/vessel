package subsystems

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type SubSystem interface {
	// Name return name of SubSystem
	Name() string
	// set cgroups's resource limit in this SubSystem
	Set(path string, res *ResourceConfig) error
	// add process into cgroups
	Apply(path string, pid int) error
	// remove cgroups
	Remove(path string) error
}

type ResourceConfig struct {
	MemoryLimit, CPUShare, CPUSet string
}

var SubSystemInstances = []SubSystem{
	&MemorySubSystem{},
	&CPUShareSubSystem{},
	&CpuSetSubSystem{},
}

type shareImplementation struct {
	SubSystem
}

func (c *shareImplementation) Apply(cgroupPath string, pid int) error {
	cgPath, err := GetCgroupPath(c.Name(), cgroupPath, false)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path.Join(cgPath, "tasks"), []byte(strconv.Itoa(pid)), 0644)
	if err != nil {
		return fmt.Errorf("failed at set cgroup to proc: %d, error: %s", pid, err)
	}
	return nil
}

func (c *shareImplementation) Remove(cgroupPath string) error {
	cgPath, err := GetCgroupPath(c.Name(), cgroupPath, false)
	if err != nil {
		return err
	}
	return os.RemoveAll(cgPath)
}
