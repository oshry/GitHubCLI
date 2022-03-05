/*
Copyright © 2022 Oshry Levy
Main Object + Parsers
*/
package models

import (
	"encoding/csv"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"log"
	"net/http"
	"os"
)

type RunCmd struct {
	Cmd    string
	Repo   string
	Output string
	HttpObj
	HttpReq         *http.Request
	ReleaseObj      []Release
	Repository      Repository
	Stars           int
	Forks           int
	Contributors    []Contributions
	ContributorsSum int
	Language        string
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
