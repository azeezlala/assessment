package cli

import (
	"fmt"
	"github.com/azeezlala/assessment/database"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running database migrations...")
		if err := database.Migrate(); err != nil {
			fmt.Printf("Error during migration: %v\n", err)
		} else {
			fmt.Println("Database migrations completed successfully.")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
