package command

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/pathcl/gf/version"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var (
	Search = &cobra.Command{
		Use:   "search",
		Short: "Search a given string in Github",
		Long:  fmt.Sprintf("version shows the version details for the %s application.", version.ApplicationName()),
		Run:   executeSearchCommand,
	}
)

type Page struct {
	url  string
	Body []byte
	size int
}

func getBody(url string, channel chan Page) {
	log.Println("Getting", url)
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	channel <- Page{url: url, size: len(body), Body: body}

}

// Returns github content containing code

func FetchRepos(client *github.Client, op []*github.CodeSearchResult) []string {
	var content []string
	for i := 0; i < len(op); i++ {
		for _, r := range op[i].CodeResults {
			url := r.GetHTMLURL()
			ns := strings.Replace(url, "github.com", "raw.githubusercontent.com", -1)
			ns = strings.Replace(ns, "/blob/", "/", -1)
			content = append(content, ns)
		}
	}
	return content
}

// we'll search something
func executeSearchCommand(cmd *cobra.Command, args []string) {

	token := os.Getenv("GITHUB_TOKEN")

	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	opts := &github.SearchOptions{TextMatch: true,
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}

	var resultsSearch []*github.CodeSearchResult

	// main search loop
	str2 := strings.Join(args, " ")
	gquery := fmt.Sprintf("extension:go %s", str2)
	log.Printf("Query %s", gquery)

	for {
		op, resp, err := client.Search.Code(ctx, gquery, opts)
		if _, ok := err.(*github.RateLimitError); ok {
			log.Println("Hit rate limit. Waiting 45 secs")
			time.Sleep(45 * time.Second)
		}

		// we will append every repo to get it
		resultsSearch = append(resultsSearch, op)

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	var urls = FetchRepos(client, resultsSearch)
	pages := make(chan Page)

	for _, url := range urls {
		go getBody(url, pages)
	}

	log.Printf("%d results", len(urls))

	// We'll throw everything to stdout

	for i := 0; i < len(urls); i++ {
		page := <-pages
		// spit everything to stdout
		fmt.Printf("%s", page.Body)

	}
}
