package cmdregex

import (
	"github.com/coralproject/shelf/cmd/xenia/web"
	"github.com/spf13/cobra"
)

var listLong = `Retrieves a list of all available Regex names.

Example:
	regex list
`

// addList handles the retrival Regex records names.
func addList() {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Retrieves a list of all available Regex names.",
		Long:  listLong,
		RunE:  runList,
	}
	regexCmd.AddCommand(cmd)
}

// runList issues the command talking to the web service.
func runList(cmd *cobra.Command, args []string) error {
	verb := "GET"
	url := "/v1/regex"

	resp, err := web.Request(cmd, verb, url, nil)
	if err != nil {
		return err
	}

	cmd.Printf("\n%s\n\n", resp)
	return nil
}
