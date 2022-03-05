/*
Copyright © Oshry Levy

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"githubCLI/models"
	"githubCLI/utils"
	"net/http"
)

var downloadsCmd = &cobra.Command{
	Use:   "downloads",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		repo, _ := cmd.Flags().GetString("repo")
		output, _ := cmd.Flags().GetString("output")
		login, _ := cmd.Flags().GetString("login")
		password, _ := cmd.Flags().GetString("password")
		token, _ := cmd.Flags().GetString("token")
		command := "releases"
		var req *http.Request
		var ReleaseObj []models.ReleaseObj
		r := models.RunCmd{
			command,
			repo,
			output,
			models.HttpObj{
				fmt.Sprintf("https://api.github.com/repos/%s/%s", repo, command),
				models.Secure{
					Login:    login,
					Password: password,
					Token:    token,
				},
			},
			req,
			ReleaseObj,
		}

		r.GithubApiPrepareReq()
		r.HttpExecuteReq()
		r.ParseResponse()
	},
}

func init() {
	rootCmd.AddCommand(downloadsCmd)
	utils.GetFlags(downloadsCmd)
}
