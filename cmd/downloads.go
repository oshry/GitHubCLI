/*
Copyright © Oshry Levy

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"githubCLI/models"
	"githubCLI/utils"
	"log"
	"net/http"
)

// downloadsCmd represents the downloads command
var downloadsCmd = &cobra.Command{
	Use:   "downloads",
	Short: "Present the entire downloads for each asset",
	Long: `Github ApI: 
Present the entire downloads for each asset`,
	Run: func(cmd *cobra.Command, args []string) {
		var req *http.Request
		var ReleaseObj []models.ReleaseObj
		var Contributions []models.Contributions
		var Repository models.Repository

		repo, _ := cmd.Flags().GetString("repo")
		output, _ := cmd.Flags().GetString("output")
		login, _ := cmd.Flags().GetString("login")
		password, _ := cmd.Flags().GetString("password")
		token, _ := cmd.Flags().GetString("token")
		command := "releases"

		if len(repo) == 0 {
			log.Fatal("No repository provided")
		}

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
			Repository,
			0,
			0,
			Contributions,
			0,
			"",
		}

		r.GithubApiPrepareReq()
		if err := json.Unmarshal(r.HttpExecuteReq(), &r.ReleaseObj); err != nil {
			fmt.Printf("Could not unmarshal reponseBytes. %v", err)
		}
		r.ParseDownloadResponse()
	},
}

func init() {
	rootCmd.AddCommand(downloadsCmd)
	utils.GetFlags(downloadsCmd)
}
