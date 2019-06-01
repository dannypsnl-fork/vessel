package subsystems

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type CpuSetSubSystem struct {
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

func (c *CpuSetSubSystem) Apply(cgroupPath string, pid int) error {
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

func (c *CpuSetSubSystem) Remove(cgroupPath string) error {
	cgPath, err := GetCgroupPath(c.Name(), cgroupPath, false)
	if err != nil {
		return err
	}
	return os.RemoveAll(cgPath)
}
