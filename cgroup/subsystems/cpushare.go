package subsystems

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type CPUShareSubSystem struct{}

func (c *CPUShareSubSystem) Name() string { return "cpu" }

func (c *CPUShareSubSystem) Set(cgroupPath string, res *ResourceConfig) error {
	cgPath, err := GetCgroupPath(c.Name(), cgroupPath, true)
	if err != nil {
		return err
	}
	if res.CPUShare != "" {
		err := ioutil.WriteFile(path.Join(cgPath, "cpu.shares"), []byte(res.CPUShare), 0644)
		if err != nil {
			return fmt.Errorf("failed at set cgroup cpu share: %s", err)
		}
	}
	return nil
}

func (c *CPUShareSubSystem) Apply(cgroupPath string, pid int) error {
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

func (c *CPUShareSubSystem) Remove(cgroupPath string) error {
	cgPath, err := GetCgroupPath(c.Name(), cgroupPath, false)
	if err != nil {
		return err
	}
	return os.RemoveAll(cgPath)
}
