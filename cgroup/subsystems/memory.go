package subsystems

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type MemorySubSystem struct{}

func (m *MemorySubSystem) Name() string { return "memory" }

func (m *MemorySubSystem) Set(cgroupPath string, res *ResourceConfig) error {
	cgPath, err := GetCgroupPath(m.Name(), cgroupPath, true)
	if err != nil {
		return err
	}
	if res.MemoryLimit != "" {
		// memory limit is not none
		err := ioutil.WriteFile(path.Join(cgPath, "memory.limit_in_bytes"), []byte(res.MemoryLimit), 0644)
		if err != nil {
			return fmt.Errorf("failed at set cgroup memory: %s", err)
		}
	}
	return nil
}

func (m *MemorySubSystem) Apply(cgroupPath string, pid int) error {
	cgPath, err := GetCgroupPath(m.Name(), cgroupPath, false)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path.Join(cgPath, "tasks"), []byte(strconv.Itoa(pid)), 0644)
	if err != nil {
		return fmt.Errorf("failed at set cgroup to proc: %d, error: %s", pid, err)
	}
	return nil
}

func (m *MemorySubSystem) Remove(cgroupPath string) error {
	cgPath, err := GetCgroupPath(m.Name(), cgroupPath, false)
	if err != nil {
		return err
	}
	return os.RemoveAll(cgPath)
}
