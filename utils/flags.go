package utils

import (
	"github.com/spf13/cobra"
)

func GetFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String("login", "", "Github Login")
	cmd.PersistentFlags().String("password", "", "Github Password")
	cmd.PersistentFlags().String("token", "", "Use OAUTH2 token")
	cmd.PersistentFlags().StringP("repo", "r", "", "The Repository To Analyze")
	cmd.PersistentFlags().StringP("output", "o", "", "The output path of the txt file")
	cmd.PersistentFlags().BoolP("help", "h", false, "Print information about each command")
}
