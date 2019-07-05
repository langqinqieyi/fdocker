package fdocker

import (
	log "common/clog"
	"fdocker/container"

	"github.com/urfave/cli"
)

// init command
var InitCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container. Do not call it outside",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "c",
			Usage: "container name",
		},
		cli.StringFlag{
			Name:  "u",
			Usage: "user command owner",
		},
	},
	Action: func(context *cli.Context) error {
		log.Blog.Infof("init come on")
		containerName := context.String("c")
		user := context.String("u")
		err := Init(containerName, user)
		return err
	},
}

func Init(containerName, user string) error {
	return container.RunContainerInitProcess(containerName, user)
}
