package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/titech-cpp/sqlboiler/boiler"
)

type codeOptions struct {
	yamlPath string
	codePath string
}

var (
	codeOpt codeOptions
)

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:   "code",
	Short: "コードの生成",
	Long:  `コードの生成のみ行います。`,
	RunE: func(cmd *cobra.Command, args []string) error {
		yaml, err := boiler.NewYaml(codeOpt.yamlPath)
		if err != nil {
			return fmt.Errorf("Yaml Error: %w", err)
		}
		code, err := boiler.NewCode(codeOpt.codePath, yaml.Yaml)
		if err != nil {
			return fmt.Errorf("Code Constructor Error: %w", err)
		}
		err = code.BoilCode()
		if err != nil {
			return fmt.Errorf("Create Code Error: %w", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(codeCmd)
	codeCmd.Flags().StringVarP(&codeOpt.yamlPath, "yaml", "y", "sqlboiler.yaml", "string option")
	codeCmd.Flags().StringVarP(&codeOpt.codePath, "code", "c", "models", "string option")
}
