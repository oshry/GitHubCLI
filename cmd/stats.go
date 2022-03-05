/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"githubCLI/models"
	"githubCLI/utils"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Present the stats of the repo (stars, forks, language, contributors)",
	Long: `Github ApI: 
Present the stats of the repo (stars, forks, language, contributors)`,
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
		command := ""

		if len(repo) == 0 {
			log.Fatal("No repository provided")
		}

		r := models.RunCmd{
			command,
			repo,
			output,
			models.HttpObj{
				fmt.Sprintf("https://api.github.com/repos/%s", repo),
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
		if err := json.Unmarshal(r.HttpExecuteReq(), &r.Repository); err != nil {
			fmt.Printf("Could not unmarshal reponseBytes. %v", err)
		}
		//r.HttpExecuteReq()
		r.Cmd = "contributors"
		r.HttpObj.Url = fmt.Sprintf("https://api.github.com/repos/%s/%s", repo, r.Cmd)
		r.GithubApiPrepareReq()
		if err := json.Unmarshal(r.HttpExecuteReq(), &r.Contributors); err != nil {
			fmt.Printf("Could not unmarshal reponseBytes. %v", err)
		}

		r.ParseStatsResponse()
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
	utils.GetFlags(statsCmd)
}
