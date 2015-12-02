package cmdquery

import (
	"github.com/coralproject/shelf/pkg/db"
	"github.com/coralproject/shelf/pkg/srv/query"

	"github.com/spf13/cobra"
)

var listLong = `Retrieves a list of all available query names.

Example:
	query list
`

// addList handles the retrival query records names.
func addList() {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Retrieves a list of all available query names.",
		Long:  listLong,
		Run:   runList,
	}
	queryCmd.AddCommand(cmd)
}

// runList is the code that implements the lists command.
func runList(cmd *cobra.Command, args []string) {
	cmd.Println("Getting List")

	db := db.NewMGO()
	defer db.CloseMGO()

	names, err := query.GetSetNames("commands", db)
	if err != nil {
		cmd.Println("Getting Query : ", err)
		return
	}

	cmd.Println("")

	for _, name := range names {
		cmd.Println(name)
	}

	cmd.Println("")
}
