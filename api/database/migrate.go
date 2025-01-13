package database

import (
	"ariga.io/atlas-go-sdk/atlasexec"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

var models []interface{}

func RegisterModel(m ...interface{}) {
	models = append(models, m...)
}

func Migrate() error {
	if len(os.Args) > 1 && os.Args[1] == "--load" {
		return LoadMStatements()
	}

	err := runMigration(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}

	return nil
}

func LoadMStatements() error {
	stmts, err := LoadStatements()
	if err != nil {
		return err
	}

	io.WriteString(os.Stdout, stmts)
	return nil
}

func runMigration(ctx context.Context, dbURL string) error {
	// Define the execution context, supplying a migration directory
	// and potentially an `atlas.hcl` configuration file using `atlasexec.WithHCL`.
	workdir, err := atlasexec.NewWorkingDir(
		atlasexec.WithMigrations(
			os.DirFS("./migrations"),
		),
	)
	if err != nil {
		log.Printf("creating workdir: %v", err)
		return err
	}

	// atlasexec works on a temporary directory, so we need to close it
	defer workdir.Close()

	// Initialize the client.
	client, err := atlasexec.NewClient(workdir.Path(), "atlas")
	if err != nil {
		log.Printf("creating atlase client: %v", err)
		return err
	}

	// Run `atlas migrate apply` on a SQLite database under /tmp.
	res, err := client.MigrateApply(ctx, &atlasexec.MigrateApplyParams{
		URL:        dbURL,
		AllowDirty: true,
	})
	if err != nil {
		log.Printf("migrating: %v", err)
		return err
	}

	targetVersion := res.Target
	if res.Target == "" {
		targetVersion = "Already at latest version"
	}

	log.Printf("Migration Status: %v", res.Error == "")
	if res.Error != "" {
		log.Printf("\t--Error: %v", res.Error)
		return errors.New(res.Error)
	}
	log.Printf("\t-- Current Version: %s", res.Current)
	log.Printf("\t-- Next Version: %s", targetVersion)
	log.Printf("\t-- Executed Files %d", len(res.Applied))
	log.Printf("\t-- Pending Files %d", len(res.Pending))

	if res.Error != "" {
		return errors.New("Migration failed")
	}

	return nil
}

func AutoMigrate() error {
	db := GetDB()

	// Example migration
	err := db.AutoMigrate(models...)
	if err != nil {
		return fmt.Errorf("migration failed: %v", err)
	}

	return nil
}
