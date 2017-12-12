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
	"net/http"
	"service_agenda/entity"

	"github.com/spf13/cobra"
)

// dmCmd represents the dm command
var dmCmd = &cobra.Command{
	Use:   "dm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		name, _ := cmd.Flags().GetString("name")

		mt := entity.GetMeeting(name, []string{}, "2017-01-22 12:00", "2017-02-10 10:00", title)
		b, err := json.Marshal(mt)
		if err != nil {
			fmt.Println("json err:", err)
		}
		body := bytes.NewBuffer([]byte(b))

		client := &http.Client{}

		req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/v1/meetings/"+title, body)

		if err != nil {
			panic(err)
		}

		req.Header.Set("Content-Type", "application/json;charset=utf-8")

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		if http.StatusNoContent == resp.StatusCode {
			fmt.Println("Deleted successfully")
		} else {
			fmt.Println("Deleted failed")
		}
	},
}

func init() {
	RootCmd.AddCommand(dmCmd)

	dmCmd.Flags().StringP("title", "t", "Anonymous", "提供会议主题名字")
	dmCmd.Flags().StringP("name", "n", "", "提供用户名字")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
