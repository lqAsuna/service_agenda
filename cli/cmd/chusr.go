// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"service_agenda/cli/req"
	"service_agenda/entity"

	"github.com/spf13/cobra"
)

// chusrCmd represents the chusr command
var chusrCmd = &cobra.Command{
	Use:   "chusr",
	Short: "only change password, email and phone, user's name can't be changed",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("phone")
		ur := entity.GetUser(name, password, email, phone)

		req.UserPatch(ur)
	},
}

func init() {
	RootCmd.AddCommand(chusrCmd)
	chusrCmd.Flags().StringP("name", "n", "", "提供用户名字")
	chusrCmd.Flags().StringP("password", "p", "", "提供用户密码")
	chusrCmd.Flags().StringP("email", "e", "", "提供用户邮箱地址")
	chusrCmd.Flags().StringP("phone", "o", "", "提供用户手机号码")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chusrCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chusrCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
