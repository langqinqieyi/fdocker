package subsystems

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type MemorySubSystem struct {
}

func (s *MemorySubSystem) Set(cgroupPath string, res *ResourceConfig) error {
	if subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, true); err == nil {
		if res.MemoryLimit != "" {
			if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "memory.limit_in_bytes"), []byte(res.MemoryLimit), 0644); err != nil {
				return fmt.Errorf("set cgroup memory limit_in_bytes fail %v", err)
			}
			//set swappiness to 0
			if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "memory.swappiness"), []byte("0"), 0644); err != nil {
				return fmt.Errorf("set cgroup memory swappiness error: %v", err)
			}
		}
		return nil
	} else {
		return err
	}

}

func (s *MemorySubSystem) Remove(cgroupPath string) error {
	if subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, false); err == nil {
		return os.RemoveAll(subsysCgroupPath)
	} else {
		return err
	}
}

func (s *MemorySubSystem) Apply(cgroupPath string, pid int) error {
	if subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, false); err == nil {
		if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "cgroup.procs"), []byte(strconv.Itoa(pid)), 0644); err != nil {
			return fmt.Errorf("set cgroup proc fail %v", err)
		}
		return nil
	} else {
		return fmt.Errorf("get cgroup %s error: %v", cgroupPath, err)
	}
}

func (s *MemorySubSystem) Name() string {
	return "memory"
}
