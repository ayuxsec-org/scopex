package main

import (
	"errors"

	"github.com/ayuxsec-org/log"
	"github.com/ayuxsec-org/scopex/internal/utils"
	"github.com/ayuxsec-org/scopex/pkg/config"
	"github.com/ayuxsec-org/scopex/pkg/version"
	"github.com/spf13/cobra"
)

func (cmdi *Cmdi) RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "scopex",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			showVersion := utils.Must(cmd.Flags().GetBool("version"))
			if showVersion {
				log.WithTrace = false
				log.Infof("current scopex version: %s", version.String())
				return nil
			}

			createCfgPath := utils.Must(cmd.Flags().GetString("create-cfg"))
			if createCfgPath != "" {
				return config.Create(createCfgPath)
			}

			loadCfgPath := utils.Must(cmd.Flags().GetString("load-cfg"))
			if loadCfgPath != "" {
				cfg, err := config.Load(loadCfgPath)
				if err != nil {
					if errors.Is(err, config.ErrFileNotFound) {
						log.Warnf("failed to load config. '%s' doesn't exist", loadCfgPath)
						return nil
					}
					return err
				}
				cmdi.Config = cfg
				log.Infof("loaded config from: %s", loadCfgPath)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	rootCmd.PersistentFlags().String("load-cfg", defaultCfgPath, "path to the yaml config file to use")
	rootCmd.PersistentFlags().String("create-cfg", "", "path to write yaml config file to")
	rootCmd.PersistentFlags().BoolP("version", "v", false, "print version and exit")

	rootCmd.AddCommand(cmdi.NewHackerOneCmd())
	return rootCmd
}
