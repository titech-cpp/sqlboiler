package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/titech-cpp/sqlboiler/boiler"
	"github.com/titech-cpp/sqlboiler/model"
)

var schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "スキーマの生成",
	Long: `スキーマの生成のみ行います。`,
	RunE: func(cmd *cobra.Command, args []string) error {
		yaml := new(model.Yaml)
		schema := boiler.NewSchema("", yaml)
		err := schema.BoilSchema()
		if err != nil {
			return fmt.Errorf("Create Schema Error: %w", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(schemaCmd)
}
