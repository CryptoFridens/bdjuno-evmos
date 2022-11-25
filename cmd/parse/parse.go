package parse

import (
	parse "github.com/Fueav/juno/cmd/parse/types"
	"github.com/spf13/cobra"

	parseblocks "github.com/Fueav/juno/cmd/parse/blocks"

	parsegenesis "github.com/Fueav/juno/cmd/parse/genesis"

	parsetransaction "github.com/Fueav/juno/cmd/parse/transactions"
	parseauth "github.com/forbole/bdjuno/v3/cmd/parse/auth"
	parsefeegrant "github.com/forbole/bdjuno/v3/cmd/parse/feegrant"
	parsegov "github.com/forbole/bdjuno/v3/cmd/parse/gov"
	parsestaking "github.com/forbole/bdjuno/v3/cmd/parse/staking"
)

// NewParseCmd returns the Cobra command allowing to parse some chain data without having to re-sync the whole database
func NewParseCmd(parseCfg *parse.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "parse",
		Short:             "Parse some data without the need to re-syncing the whole database from scratch",
		PersistentPreRunE: runPersistentPreRuns(parse.ReadConfigPreRunE(parseCfg)),
	}

	cmd.AddCommand(
		parseauth.NewAuthCmd(parseCfg),
		parseblocks.NewBlocksCmd(parseCfg),
		parsefeegrant.NewFeegrantCmd(parseCfg),
		parsegenesis.NewGenesisCmd(parseCfg),
		parsegov.NewGovCmd(parseCfg),
		parsestaking.NewStakingCmd(parseCfg),
		parsetransaction.NewTransactionsCmd(parseCfg),
	)

	return cmd
}

func runPersistentPreRuns(preRun func(_ *cobra.Command, _ []string) error) func(_ *cobra.Command, _ []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if root := cmd.Root(); root != nil {
			if root.PersistentPreRunE != nil {
				err := root.PersistentPreRunE(root, args)
				if err != nil {
					return err
				}
			}
		}

		return preRun(cmd, args)
	}
}
