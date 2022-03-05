#**GitHub CLI**
GitHub CLI get information from the GitHub repository, create a report and save it locally.

##**Install**
###**run**
go install githubCLI
githubCLI downloads --repo=target/flottbot
##**Usage**
###**Stats**

Usage:
githubCLI downloads [flags]

Flags:
-h, --help string       Print information about each command
--login string      Github Login
-o, --output string     The output path of the txt file
--password string   Github Password
-r, --repo string       The Repository To Analyze
--token string      Use OAUTH2 token

###**Downloads**

Usage:
githubCLI downloads [flags]

Flags:
-h, --help              Print information about each command
--login string      Github Login
-o, --output string     The output path of the txt file
--password string   Github Password
-r, --repo string       The Repository To Analyze
--token string      Use OAUTH2 token


**cobra**
mkdir githubCLI && cd githubCLI
go mod init githubCLI
cobra init
git init
git add .
git
cobra add stats
cobra add downloads
