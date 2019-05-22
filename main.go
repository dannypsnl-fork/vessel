package main

import (
	"os"

	"github.com/dannypsnl/vessel/cmd/initcmd"
	"github.com/dannypsnl/vessel/cmd/run"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use: "vessel",
}

func init() {
	root.AddCommand(initcmd.Cmd)
	root.AddCommand(run.Cmd)

	logrus.SetOutput(os.Stdout)
}

func main() {
	err := root.Execute()
	if err != nil {
		logrus.Errorf("error: %s", err)
	}
}
