/*
Copyright © 2022 Oshry Levy

*/
package cmd

import (
	"fmt"
	"githubCLI/models"
	"githubCLI/utils"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var req *http.Request
var Release []models.Release
var Contributions []models.Contributions
var Repository models.Repository
var ApiURL string
var RunCmd models.RunCmd

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "githubCLI",
	Short: "This tool is used to get some stats from from Github for a specific repo",
	Long: `This tool is used to get some stats from from Github for a specific repo
this tool present the result as a table and write the output for a given file or just print it`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	ApiURL = "https://api.github.com/repos"
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.githubCLI.yaml)")
	// Application main module
	RunCmd = models.RunCmd{
		"",
		"",
		"",
		models.HttpObj{
			fmt.Sprintf(ApiURL+"/%s", ""),
			models.Secure{
				Login:    "",
				Password: "",
				Token:    "",
			},
		},
		req,
		Release,
		Repository,
		0,
		0,
		Contributions,
		0,
		"",
	}
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	utils.GetFlags(rootCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
