package cmd

import (
	"fmt"
	"html"
	"os"
	"strings"

	"github.com/maryayi/news-cli/internal/news"
	"github.com/spf13/cobra"
)

var count int

var rootCmd = &cobra.Command{
	Use:     "news [keyword]",
	Short:   "Fetch top news headlines for a keyword",
	Long:    `news fetches the latest news headlines for a given keyword using NewsAPI.org.`,
	Version: "0.2.1",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := os.Getenv("NEWS_API_KEY")
		if apiKey == "" {
			return fmt.Errorf("NEWS_API_KEY environment variable is not set.\nGet a free API key at https://newsapi.org/register")
		}

		var keyword string
		if len(args) > 0 {
			keyword = args[0]
		}

		articles, err := news.FetchHeadlines(apiKey, keyword, count)
		if err != nil {
			return err
		}

		if len(articles) == 0 {
			if keyword == "" {
				fmt.Println("No headlines found.")
			} else {
				fmt.Printf("No headlines found for %q.\n", keyword)
			}
			return nil
		}

		if keyword == "" {
			fmt.Printf("Top %d headlines around the world:\n\n", len(articles))
		} else {
			fmt.Printf("Top %d headlines for \"%s\":\n\n", len(articles), keyword)
		}
		for i, a := range articles {
			source := a.Source.Name
			if source == "" {
				source = "Unknown"
			}
			title := strings.TrimSpace(a.Title)
			url := html.UnescapeString(a.URL)
			fmt.Printf("%d. %s\n   \033[34m%s\033[0m  \033[2m%s\033[0m\n\n", i+1, title, source, url)
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&count, "count", "n", 5, "number of headlines to fetch")
}
