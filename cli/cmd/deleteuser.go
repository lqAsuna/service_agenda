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
var deleteuserCmd = &cobra.Command{
	Use:   "deleteuser",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		password, _ := cmd.Flags().GetString("password")

		client := &http.Client{}
		u := entity.GetUser(name, password, "", "")

		b, err := json.Marshal(u)
		if err != nil {
			fmt.Println("json err:", err)
		}
		body := bytes.NewBuffer([]byte(b))

		req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/v1/users/"+name, body)

		if err != nil {
			// handle error
		}

		req.Header.Set("Content-Type", "application/json;charset=utf-8")

		resp, err := client.Do(req)
		if err != nil {
			// handle error
		}
		defer resp.Body.Close()

		if http.StatusNoContent == resp.StatusCode {
			fmt.Println("deleted successfully")
		}
	},
}

func init() {
	RootCmd.AddCommand(deleteuserCmd)

	deleteuserCmd.Flags().StringP("name", "n", "Anonymous", "提供用户名字")
	deleteuserCmd.Flags().StringP("password", "p", "Anonymous", "提供用户密码")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
