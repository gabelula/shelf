package cmdrelationship

import (
	"github.com/coralproject/shelf/cmd/xenia/web"
	"github.com/spf13/cobra"
)

var getLong = `Retrieves relationships record from the system with the optional supplied predicate.

Example:
	relationship get

	relationship get -p predicate
`

// get contains the state for this command.
var get struct {
	predicate string
}

// addGet handles the retrival of relationship records, displayed in json formatted response.
func addGet() {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Retrieves all relationship records, or those matching an optional predicate.",
		Long:  getLong,
		RunE:  runGet,
	}

	cmd.Flags().StringVarP(&get.predicate, "predicate", "p", "", "Relationship predicate.")

	relationshipCmd.AddCommand(cmd)
}

// runGet issues the command talking to the web service.
func runGet(cmd *cobra.Command, args []string) error {
	verb := "GET"
	url := "/v1/relationship"

	if get.predicate != "" {
		url += "/" + get.predicate
	}

	resp, err := web.Request(cmd, verb, url, nil)
	if err != nil {
		return err
	}

	cmd.Printf("\n%s\n\n", resp)
	return nil
}
