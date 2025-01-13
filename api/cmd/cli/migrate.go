package cli

import (
	"fmt"
	"github.com/azeezlala/assessment/api/database"
	"github.com/spf13/cobra"
	"log"
)

var load bool
var addCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		if load {
			if err := database.LoadMStatements(); err != nil {
				log.Printf("error loading database statements: %v", err)
			}
			return
		}
		if err := database.Migrate(); err != nil {
			fmt.Printf("Error during migration: %v\n", err)
		} else {
			fmt.Println("Database migrations completed successfully.")
		}
	},
}

func init() {
	addCmd.Flags().BoolVar(&load, "load", false, "Load statements instead of running migrations")
	rootCmd.AddCommand(addCmd)
}
