// Package cmd ...
/*
/* Copyright © 2022 Abhijit Wakchaure<abhijitwakchaure.2014@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// VERSION represents the program version
var VERSION string

const versionTemplate = ` Slack Notifier Go v%s
--------------------------
Developer: Abhijit Wakchaure
Github: https://github.com/abhijitWakchaure/slack-notifier-go
`

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(versionTemplate, VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
