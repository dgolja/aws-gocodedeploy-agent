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

	"github.com/n1tr0g/aws-gocodedeploy-agent/utils"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop the AWS CodeDeploy agent",
	Long: `NAME
  stop - stop the AWS CodeDeploy agent`,
	Run: func(cmd *cobra.Command, args []string) {
		stop()
	},
}

func init() {
	RootCmd.AddCommand(stopCmd)
}

func stop() {
	process, err := utils.GetServiceProcess()
	if err != nil {
		fmt.Printf("No AWS CodeDeploy agent running - %s\n", err)
		os.Exit(1)
	}

	err = utils.StopRunningProcess(process)
	if err != nil {
		fmt.Printf("No AWS CodeDeploy agent running - %s\n", err)
		os.Exit(2)
	}

	pidfile, _ := utils.GetPidFile()

	if err := os.Remove(pidfile); err != nil {
		fmt.Printf("Unable to remove %s - %s\n", pidfile, err)
	}
}
