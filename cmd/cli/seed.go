package cli

import (
	"fmt"
	"github.com/azeezlala/assessment/database/seeder"
	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed the database with initial data",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Seeding the database...")
		if err := seeder.Seed(); err != nil {
			fmt.Printf("Error during seeding: %v\n", err)
		} else {
			fmt.Println("Database seeding completed successfully.")
		}
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
