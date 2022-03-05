package models

import (
	"encoding/csv"
	"encoding/json"
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
	HttpReq    *http.Request
	ReleaseObj []ReleaseObj
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

func (r *RunCmd) HttpExecuteReq() {
	response, err := http.DefaultClient.Do(r.HttpReq)
	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}

	if err := json.Unmarshal(responseBytes, &r.ReleaseObj); err != nil {
		fmt.Printf("Could not unmarshal reponseBytes. %v", err)
	}
}

func (r *RunCmd) ParseResponse() {
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
