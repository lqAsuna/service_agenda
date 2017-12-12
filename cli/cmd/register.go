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
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"service_agenda/entity"

	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",
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

		b, err := json.Marshal(ur)
		if err != nil {
			fmt.Println("json err:", err)
		}
		body := bytes.NewBuffer([]byte(b))

		res, err := http.Post("http://localhost:8080/v1/users", "application/json;charset=utf-8", body)
		if err != nil {
			panic(err)
		}

		result, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			panic(err)
		}

		if res.StatusCode == http.StatusCreated {
			fmt.Println("created successfully")
			fmt.Println(string(result))
		} else {
			fmt.Println("created failed")
		}

	},
}

func init() {
	RootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("name", "n", "Anonymous", "提供用户名字")
	registerCmd.Flags().StringP("password", "p", "Anonymous", "提供用户密码")
	registerCmd.Flags().StringP("email", "e", "Anonymous", "提供用户邮箱地址")
	registerCmd.Flags().StringP("phone", "o", "Anonymous", "提供用户手机号码")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
