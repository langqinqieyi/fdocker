package container

import (
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/sky-big/fdocker/container/config"
	"github.com/sky-big/fdocker/container/logs"
	fvolume "github.com/sky-big/fdocker/container/volume"

	"github.com/golang/glog"
)

func NewParentProcess(tty bool, containerName, volume, imageName, user string, envSlice []string) *exec.Cmd {
	initCmd, err := os.Readlink("/proc/self/exe")
	if err != nil {
		glog.Errorf("get init process error %v", err)
		return nil
	}

	cmd := exec.Command(initCmd, "init", "-c", containerName, "-u", user)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS |
			syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}

	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		stdLogFile := logs.NewLogFile(containerName, config.ContainerLogFile)
		if stdLogFile == nil {
			return nil
		}
		cmd.Stdout = stdLogFile
		stdErrFile := logs.NewLogFile(containerName, config.ContainerErrFile)
		if stdErrFile == nil {
			return nil
		}
		cmd.Stderr = stdErrFile
	}

	cmd.Env = append(os.Environ(), envSlice...)
	fvolume.NewWorkSpace(volume)
	cmd.Dir = filepath.Join(config.RootUrl, config.Runtime)
	return cmd
}
