package subsystems

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
}
