package subsystems

import (
	"fmt"
	"io/ioutil"
	"path"
)

type MemorySubSystem struct {
	shareImplementation
}

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
