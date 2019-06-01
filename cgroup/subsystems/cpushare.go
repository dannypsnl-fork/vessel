package subsystems

import (
	"fmt"
	"io/ioutil"
	"path"
)

type CPUShareSubSystem struct {
	shareImplementation
}

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
