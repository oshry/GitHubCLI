/*
Copyright © 2022 Oshry Levy

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Present the stats of the repo (stars, forks, language, contributors)",
	Long: `Github ApI: 
Present the stats of the repo (stars, forks, language, contributors)`,
	Run: func(cmd *cobra.Command, args []string) {
		RunCmd.Repo, _ = cmd.Flags().GetString("repo")
		RunCmd.Output, _ = cmd.Flags().GetString("output")
		RunCmd.Secure.Login, _ = cmd.Flags().GetString("login")
		RunCmd.Secure.Password, _ = cmd.Flags().GetString("password")
		RunCmd.Secure.Token, _ = cmd.Flags().GetString("token")
		if len(RunCmd.Repo) == 0 {
			log.Fatal("No repository provided")
		}
		RunCmd.HttpObj.Url = fmt.Sprintf(RunCmd.HttpObj.Url+"%s", RunCmd.Repo)
		RunCmd.GithubApiPrepareReq()
		// Execute repository request
		if err := json.Unmarshal(RunCmd.HttpExecuteReq(), &RunCmd.Repository); err != nil {
			fmt.Printf("Could not unmarshal reponseBytes. %v", err)
		}
		// Adjust main module state to work with contributors api
		RunCmd.Cmd = "contributors"
		RunCmd.HttpObj.Url = fmt.Sprintf(ApiURL+"/%s/%s", RunCmd.Repo, RunCmd.Cmd)
		RunCmd.GithubApiPrepareReq()
		// Execute contributors request
		if err := json.Unmarshal(RunCmd.HttpExecuteReq(), &RunCmd.Contributors); err != nil {
			fmt.Printf("Could not unmarshal reponseBytes. %v", err)
		}
		// Parse main module into the expected table structure
		RunCmd.ParseStatsResponse()
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
