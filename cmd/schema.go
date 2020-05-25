package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/titech-cpp/sqlboiler/boiler"
)

type schemaOptions struct {
	yamlPath   string
	schemaPath string
}

var (
	schemaOpt schemaOptions
)

var schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "スキーマの生成",
	Long:  `スキーマの生成のみ行います。`,
	RunE: func(cmd *cobra.Command, args []string) error {
		yaml, err := boiler.NewYaml(schemaOpt.yamlPath)
		if err != nil {
			return fmt.Errorf("Yaml Error: %w", err)
		}
		schema := boiler.NewSchema(schemaOpt.schemaPath, yaml)
		if err != nil {
			return fmt.Errorf("Schema Error: %w", err)
		}
		err = schema.BoilSchema()
		if err != nil {
			return fmt.Errorf("Create Schema Error: %w", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(schemaCmd)
	schemaCmd.Flags().StringVarP(&schemaOpt.yamlPath, "yaml", "y", "sqlboiler.yaml", "string option")
	schemaCmd.Flags().StringVarP(&schemaOpt.schemaPath, "schema", "s", "docs", "string option")
}
