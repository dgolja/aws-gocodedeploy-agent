// Copyright © 2016 NAME HERE Dejan Golja
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/n1tr0g/aws-gocodedeploy-agent/utils"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start the AWS CodeDeploy agent",
	Long: `NAME
  start - start the AWS CodeDeploy agent`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		start()
	},
}

func init() {
	RootCmd.AddCommand(startCmd)
}

func start() {
	signalStop := make(chan os.Signal, 1)

	pidfile, err := utils.GetPidFile()
	if err != nil {
		fmt.Printf("Ignoring ... pid file definition missing - %s", err)
	}

	process, err := utils.GetServiceProcess()
	if err == nil {
		err = utils.IsProcessRunning(process)
		if err == nil {
			fmt.Printf("The AWS CodeDeploy agent is already running as PID %d", process.Pid)
			os.Exit(1)
		}
	}

	myPID := os.Getpid()

	if err := utils.SavePidIDToFile(pidfile, myPID); err != nil {
		fmt.Printf("Unable to save pid %d to %s - %s", myPID, pidfile, err)
	}

	go func() {
		for {
			time.Sleep(10 * time.Second)
			fmt.Println("Working ...")
		}
	}()

	signal.Notify(signalStop, os.Interrupt, syscall.SIGTERM)
	<-signalStop
	fmt.Printf("Quiting...%s \n", pidfile)
}
