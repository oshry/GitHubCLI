/*
Copyright © Oshry Levy

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// downloadsCmd represents the downloads command
var downloadsCmd = &cobra.Command{
	Use:   "downloads",
	Short: "Present the entire downloads for each asset",
	Long: `Github ApI: 
Present the entire downloads for each asset`,
	Run: func(cmd *cobra.Command, args []string) {
		//Get Flags
		RunCmd.Repo, _ = cmd.Flags().GetString("repo")
		RunCmd.Output, _ = cmd.Flags().GetString("output")
		RunCmd.Secure.Login, _ = cmd.Flags().GetString("login")
		RunCmd.Secure.Password, _ = cmd.Flags().GetString("password")
		RunCmd.Secure.Token, _ = cmd.Flags().GetString("token")
		if len(RunCmd.Repo) == 0 {
			log.Fatal("No repository provided")
		}
		RunCmd.Cmd = "releases"
		RunCmd.HttpObj.Url = fmt.Sprintf(RunCmd.HttpObj.Url+"%s/%s", RunCmd.Repo, RunCmd.Cmd)
		RunCmd.GithubApiPrepareReq()
		//Unmarshal request ro a ReleaseObj
		if err := json.Unmarshal(RunCmd.HttpExecuteReq(), &RunCmd.ReleaseObj); err != nil {
			fmt.Printf("Could not unmarshal reponseBytes. %v", err)
		}
		// Parse main module into the expected table structure
		RunCmd.ParseDownloadResponse()
	},
}

func init() {
	rootCmd.AddCommand(downloadsCmd)
}
