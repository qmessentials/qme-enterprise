/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Database migrations",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
var databaseName string

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.PersistentFlags().StringVar(&databaseName, "database", "", "The database to connect to")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getCurrentDbVersion(conn *pgx.Conn) (int, error) {
	sql := "select max(version_number) from schema_version;"
	var versionNumber int
	if err := conn.QueryRow(context.Background(), sql).Scan(&versionNumber); err != nil {
		return -1, err
	}
	return versionNumber, nil
}

func applyDbVersion(conn *pgx.Conn, versionNumber int, sql string) error {
	cv, err := getCurrentDbVersion(conn)
	if err != nil {
		return err
	}
	if cv <= versionNumber {
		return fmt.Errorf("cannot apply version %v; database is currently at version %v", versionNumber, cv)
	}
	_, err = conn.Exec(context.Background(), sql)
	if err != nil {
		return err
	}
	vsql := fmt.Sprintf("insert into schema_version (version_number) values (%v);", versionNumber)
	tag, err := conn.Exec(context.Background(), vsql)
	if err != nil {
		return err
	}
	if tag.RowsAffected() != 1 {
		return fmt.Errorf("schema_version insert affected %v rows, expected 1", tag.RowsAffected())
	}
	return nil
}
