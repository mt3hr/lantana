package lantana_server_cmd

import (
	"log"
	"os"
	"os/signal"

	lantana "github.com/mt3hr/lantana/src/app/lantana/lantana"
	"github.com/spf13/cobra"
)

func Execute() {
	if err := serverCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.MousetrapHelpText = "" // Windowsでマウスから起動しても怒られないようにする
	serverCmd.PersistentFlags().StringVarP(&lantana.ConfigFileName, "config_file", "c", "", "使用するコンフィグファイル")
}

var (
	serverCmd = &cobra.Command{
		Use:              "lantana_server",
		PersistentPreRun: lantana.PersistentPreRun,
		Run: func(_ *cobra.Command, _ []string) {
			err := lantana.LoadRepositories()
			if err != nil {
				log.Fatal(err)
			}
			defer lantana.LoadedRepositories.Close()
			interceptCh := make(chan os.Signal)
			signal.Notify(interceptCh, os.Interrupt)
			go func() {
				<-interceptCh
				lantana.LoadedRepositories.Close()
				os.Exit(0)
			}()
			lantana.LoadedRepositories, err = lantana.WrapT(lantana.LoadedRepositories)
			if err != nil {
				log.Fatal(err)
			}

			err = lantana.LaunchServer()
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)
