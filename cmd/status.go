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

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Report running status of the AWS CodeDeploy agent",
	Long: `NAME
  status - Report running status of the AWS CodeDeploy agent`,
	Run: func(cmd *cobra.Command, args []string) {
		status()
	},
}

func init() {
	RootCmd.AddCommand(statusCmd)
}

func status() {
	process, err := utils.GetServiceProcess()
	if err != nil {
		fmt.Printf("No AWS CodeDeploy agent running - %s\n", err)
		os.Exit(1)
	}

	err = utils.IsProcessRunning(process)
	if err != nil {
		fmt.Printf("No AWS CodeDeploy agent running - %s\n", err)
		os.Exit(2)
	}
	fmt.Printf("The AWS CodeDeploy agent is running as PID %d\n", process.Pid)
}
