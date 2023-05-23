package lantana_server_cmd

import (
	"log"
	"os"
	"os/signal"

	"github.com/mt3hr/lantana/src/app/lantana/lantana"
	"github.com/mt3hr/rykv/kyou"
	"github.com/spf13/cobra"
)

func Execute() {
	if err := serverCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.MousetrapHelpText = "" // Windowsでマウスから起動しても怒られないようにする
	serverCmd.PersistentFlags().StringVarP(&lantanaServer.ConfigFileName, "config_file", "c", "", "使用するコンフィグファイル")
}

var (
	lantanaServer = &lantana.LantanaServer{}
	serverCmd     = &cobra.Command{
		Use: "lantana_server",
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			err := lantanaServer.LoadConfig()
			if err != nil {
				log.Fatal(err)
			}
			err = lantanaServer.LoadTagStruct()
			if err != nil {
				log.Fatal(err)
			}
			lantanaServer.Config.ApplicationConfig.HiddenTags = append(lantanaServer.Config.ApplicationConfig.HiddenTags, kyou.DeletedTagName)
		},
		Run: func(_ *cobra.Command, _ []string) {
			err := lantanaServer.LoadRepositories()
			if err != nil {
				log.Fatal(err)
			}
			defer lantanaServer.Repositories.Close()
			interceptCh := make(chan os.Signal)
			signal.Notify(interceptCh, os.Interrupt)
			go func() {
				<-interceptCh
				lantanaServer.Repositories.Close()
				os.Exit(0)
			}()

			err = lantanaServer.WrapT()
			if err != nil {
				log.Fatal(err)
			}

			err = lantanaServer.LaunchServer()
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)
