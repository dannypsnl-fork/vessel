package subsystems

import (
	"fmt"
	"io/ioutil"
	"path"
)

type CpuSetSubSystem struct {
	shareImplementation
}

func (c *CpuSetSubSystem) Name() string {
	return "cpuset"
}

func (c *CpuSetSubSystem) Set(cgroupPath string, res *ResourceConfig) error {
	cgPath, err := GetCgroupPath(c.Name(), cgroupPath, true)
	if err != nil {
		return err
	}
	if res.MemoryLimit != "" {
		err := ioutil.WriteFile(path.Join(cgPath, "cpuset.cpus"), []byte(res.CPUSet), 0644)
		if err != nil {
			return fmt.Errorf("failed at set cgroup cpuset: %s", err)
		}
	}
	return nil
}
