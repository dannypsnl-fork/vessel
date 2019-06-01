package cgroup

import (
	"github.com/dannypsnl/vessel/cgroup/subsystems"

	"github.com/sirupsen/logrus"
)

type Manager struct {
	Path     string
	Resource *subsystems.ResourceConfig
}

func NewManager(path string) *Manager {
	return &Manager{
		Path: path,
	}
}

func (c *Manager) Apply(pid int) {
	for _, subSys := range subsystems.SubSystemInstances {
		err := subSys.Apply(c.Path, pid)
		if err != nil {
			logrus.Error(err)
		}
	}
}

func (c *Manager) Set(res *subsystems.ResourceConfig) {
	for _, subSys := range subsystems.SubSystemInstances {
		err := subSys.Set(c.Path, res)
		if err != nil {
			logrus.Error(err)
		}
	}
}

func (c *Manager) Destroy() {
	for _, subSys := range subsystems.SubSystemInstances {
		err := subSys.Remove(c.Path)
		if err != nil {
			logrus.Error(err)
		}
	}
}
