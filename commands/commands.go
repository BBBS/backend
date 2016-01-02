package commands

import (
	"github.com/spf13/cobra"

	"github.com/bbbs/backend/db"
	"github.com/bbbs/backend/db/get"
	"github.com/bbbs/backend/db/save"
	"github.com/bbbs/backend/http"
	"github.com/bbbs/backend/user"
)

// Commands
var Commands = []*cobra.Command{
	&cobra.Command{
		Use:   "get",
		Short: "Get number",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			get.Run()
		},
	},
	&cobra.Command{
		Use:   "save",
		Short: "Save a doc in the database",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			save.Run()
		},
	},
	&cobra.Command{
		Use:   "http",
		Short: "Start the http server",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			http.Run()
		},
	},
	&cobra.Command{
		Use:   "db",
		Short: "Start the db service",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			db.Run()
		},
	},
	&cobra.Command{
		Use:   "user",
		Short: "Start the user service",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			user.Run()
		},
	},
}

func init() {
	ConfigDefaults(Commands...)
}
