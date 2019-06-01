package subsystems

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/sirupsen/logrus"
)

func GetCgroupPath(subsystem, cgroupPath string, autoCreate bool) (string, error) {
	cgroupRoot := FindCgroupMountPoint(subsystem)
	cgPath := path.Join(cgroupRoot, cgroupPath)
	_, err := os.Stat(cgPath)
	if autoCreate {
		if os.IsNotExist(err) {
			err := os.Mkdir(cgPath, 0755)
			if err != nil {
				return "", fmt.Errorf("failed at create cgroup: %s", err)
			}
			return cgPath, nil
		}
		return cgPath, nil
	}
	if err != nil {
		return "", fmt.Errorf("failed at get cgroup path: %s", err)
	}
	return cgPath, nil
}

func FindCgroupMountPoint(subsystem string) string {
	f, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		logrus.Errorf("can't open /proc/self/mountinfo: %s", err)
		return ""
	}
	defer f.Close()

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		text := scan.Text()
		fields := strings.Split(text, " ")
		for _, opt := range strings.Split(fields[len(fields)-1], ",") {
			if opt == subsystem {
				return fields[4]
			}
		}
	}
	err = scan.Err()
	if err != nil {
		logrus.Errorf("read /proc/self/mountinfo got error: %s", err)
	}
	return ""
}
