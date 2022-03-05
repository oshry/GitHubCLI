package models

import (
	"encoding/csv"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type AssetsObj struct {
	Url      string `json:"url"`
	Id       int    `json:"id"`
	NodeId   string `json:"node_id"`
	Name     string `json:"name"`
	Label    string `json:"label"`
	Uploader struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"uploader"`
	ContentType        string    `json:"content_type"`
	State              string    `json:"state"`
	Size               int       `json:"size"`
	DownloadCount      int       `json:"download_count"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	BrowserDownloadUrl string    `json:"browser_download_url"`
}

type ReleaseObj struct {
	Url       string `json:"url"`
	AssetsUrl string `json:"assets_url"`
	UploadUrl string `json:"upload_url"`
	HtmlUrl   string `json:"html_url"`
	Id        int    `json:"id"`
	Author    struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
	NodeId          string      `json:"node_id"`
	TagName         string      `json:"tag_name"`
	TargetCommitish string      `json:"target_commitish"`
	Name            string      `json:"name"`
	Draft           bool        `json:"draft"`
	Prerelease      bool        `json:"prerelease"`
	CreatedAt       time.Time   `json:"created_at"`
	PublishedAt     time.Time   `json:"published_at"`
	Assets          []AssetsObj `json:"assets"`
	TarballUrl      string      `json:"tarball_url"`
	ZipballUrl      string      `json:"zipball_url"`
	Body            string      `json:"body"`
	Reactions       struct {
		Url        string `json:"url"`
		TotalCount int    `json:"total_count"`
		Field3     int    `json:"+1"`
		Field4     int    `json:"-1"`
		Laugh      int    `json:"laugh"`
		Hooray     int    `json:"hooray"`
		Confused   int    `json:"confused"`
		Heart      int    `json:"heart"`
		Rocket     int    `json:"rocket"`
		Eyes       int    `json:"eyes"`
	} `json:"reactions"`
	MentionsCount int `json:"mentions_count"`
}

type Repository struct {
	Id       int    `json:"id"`
	NodeId   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
	Owner    struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"owner"`
	HtmlUrl          string      `json:"html_url"`
	Description      string      `json:"description"`
	Fork             bool        `json:"fork"`
	Url              string      `json:"url"`
	ForksUrl         string      `json:"forks_url"`
	KeysUrl          string      `json:"keys_url"`
	CollaboratorsUrl string      `json:"collaborators_url"`
	TeamsUrl         string      `json:"teams_url"`
	HooksUrl         string      `json:"hooks_url"`
	IssueEventsUrl   string      `json:"issue_events_url"`
	EventsUrl        string      `json:"events_url"`
	AssigneesUrl     string      `json:"assignees_url"`
	BranchesUrl      string      `json:"branches_url"`
	TagsUrl          string      `json:"tags_url"`
	BlobsUrl         string      `json:"blobs_url"`
	GitTagsUrl       string      `json:"git_tags_url"`
	GitRefsUrl       string      `json:"git_refs_url"`
	TreesUrl         string      `json:"trees_url"`
	StatusesUrl      string      `json:"statuses_url"`
	LanguagesUrl     string      `json:"languages_url"`
	StargazersUrl    string      `json:"stargazers_url"`
	ContributorsUrl  string      `json:"contributors_url"`
	SubscribersUrl   string      `json:"subscribers_url"`
	SubscriptionUrl  string      `json:"subscription_url"`
	CommitsUrl       string      `json:"commits_url"`
	GitCommitsUrl    string      `json:"git_commits_url"`
	CommentsUrl      string      `json:"comments_url"`
	IssueCommentUrl  string      `json:"issue_comment_url"`
	ContentsUrl      string      `json:"contents_url"`
	CompareUrl       string      `json:"compare_url"`
	MergesUrl        string      `json:"merges_url"`
	ArchiveUrl       string      `json:"archive_url"`
	DownloadsUrl     string      `json:"downloads_url"`
	IssuesUrl        string      `json:"issues_url"`
	PullsUrl         string      `json:"pulls_url"`
	MilestonesUrl    string      `json:"milestones_url"`
	NotificationsUrl string      `json:"notifications_url"`
	LabelsUrl        string      `json:"labels_url"`
	ReleasesUrl      string      `json:"releases_url"`
	DeploymentsUrl   string      `json:"deployments_url"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	PushedAt         time.Time   `json:"pushed_at"`
	GitUrl           string      `json:"git_url"`
	SshUrl           string      `json:"ssh_url"`
	CloneUrl         string      `json:"clone_url"`
	SvnUrl           string      `json:"svn_url"`
	Homepage         string      `json:"homepage"`
	Size             int         `json:"size"`
	StargazersCount  int         `json:"stargazers_count"`
	WatchersCount    int         `json:"watchers_count"`
	Language         string      `json:"language"`
	HasIssues        bool        `json:"has_issues"`
	HasProjects      bool        `json:"has_projects"`
	HasDownloads     bool        `json:"has_downloads"`
	HasWiki          bool        `json:"has_wiki"`
	HasPages         bool        `json:"has_pages"`
	ForksCount       int         `json:"forks_count"`
	MirrorUrl        interface{} `json:"mirror_url"`
	Archived         bool        `json:"archived"`
	Disabled         bool        `json:"disabled"`
	OpenIssuesCount  int         `json:"open_issues_count"`
	License          struct {
		Key    string      `json:"key"`
		Name   string      `json:"name"`
		SpdxId string      `json:"spdx_id"`
		Url    interface{} `json:"url"`
		NodeId string      `json:"node_id"`
	} `json:"license"`
	AllowForking   bool        `json:"allow_forking"`
	IsTemplate     bool        `json:"is_template"`
	Topics         []string    `json:"topics"`
	Visibility     string      `json:"visibility"`
	Forks          int         `json:"forks"`
	OpenIssues     int         `json:"open_issues"`
	Watchers       int         `json:"watchers"`
	DefaultBranch  string      `json:"default_branch"`
	TempCloneToken interface{} `json:"temp_clone_token"`
	Organization   struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"organization"`
	NetworkCount     int `json:"network_count"`
	SubscribersCount int `json:"subscribers_count"`
}

type Contributions struct {
	Login             string `json:"login"`
	Id                int    `json:"id"`
	NodeId            string `json:"node_id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
	Contributions     int    `json:"contributions"`
}

type Secure struct {
	Login    string
	Password string
	Token    string
}

type HttpObj struct {
	Url string
	Secure
}

type RunCmd struct {
	Cmd    string
	Repo   string
	Output string
	HttpObj
	HttpReq         *http.Request
	ReleaseObj      []ReleaseObj
	Repository      Repository
	Stars           int
	Forks           int
	Contributors    []Contributions
	ContributorsSum int
	Language        string
}

func (r *RunCmd) GithubApiPrepareReq() {
	newReq, err := http.NewRequest(
		http.MethodGet,
		r.Url,
		nil,
	)
	if err != nil {
		log.Printf("Could not request repository from gitub. %v", err)
	}

	if len(r.Secure.Login)+len(r.Secure.Password) > 0 {
		newReq.SetBasicAuth(r.Secure.Login, r.Secure.Password)
	} else if len(r.Secure.Token) > 0 {
		newReq.Header.Add("Authorization", fmt.Sprintf("token %s", r.Secure.Token))
	}

	// headers
	newReq.Header.Add("Accept", "application/vnd.github.v3+json")
	r.HttpReq = newReq
}

func (r *RunCmd) HttpExecuteReq() []byte {
	response, err := http.DefaultClient.Do(r.HttpReq)
	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}

	return responseBytes
}

func (r *RunCmd) ParseDownloadResponse() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"RELEASE NAME", "DISTRIBUTION", "DOWNLOAD COUNT"})
	total := 0
	for _, f := range r.ReleaseObj {
		if len(f.Assets) > 0 {
			for _, d := range f.Assets {
				t.AppendRows([]table.Row{
					{f.Name, d.Name, d.DownloadCount},
				})
				total += d.DownloadCount
			}
			t.AppendSeparator()
		} else {
			t.AppendRows([]table.Row{
				{f.Name, "no asset for this repository", 0},
			})
			fmt.Println("no asset for this repository")
		}
	}
	t.AppendFooter(table.Row{"", "Total", total})
	if r.Output != "" {
		var err error
		csvFile, errCreateFile := os.Create(r.Output)
		if errCreateFile != nil {
			log.Fatal("failed creating file")
		}
		csvwriter := csv.NewWriter(csvFile)
		rCSV := []string{t.RenderCSV()}

		csvwriter.Write(rCSV)
		if err = csvwriter.Write(rCSV); err != nil {
			//return err
		}

		csvwriter.Flush()
		if err = csvFile.Close(); err != nil {
			//return err
		}

	} else {
		t.Render()
	}

}

func (r *RunCmd) ParseStatsResponse() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"STAT", "VALUE"})
	t.AppendRows([]table.Row{
		{"Stars", r.Repository.StargazersCount},
	})
	t.AppendRows([]table.Row{
		{"Forks", r.Repository.Forks},
	})
	t.AppendRows([]table.Row{
		{"Contributors", len(r.Contributors)},
	})
	t.AppendRows([]table.Row{
		{"Language", r.Repository.Language},
	})

	if r.Output != "" {
		var err error
		csvFile, errCreateFile := os.Create(r.Output)
		if errCreateFile != nil {
			log.Fatal("failed creating file")
		}
		csvwriter := csv.NewWriter(csvFile)
		rCSV := []string{t.Render()}

		if err = csvwriter.Write(rCSV); err != nil {
			//return err
		}

		csvwriter.Flush()
		if err = csvFile.Close(); err != nil {
			//return err
		}

	} else {
		t.Render()
	}

}
