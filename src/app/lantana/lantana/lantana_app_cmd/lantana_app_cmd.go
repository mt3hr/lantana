package lantana_app_cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	lantana "github.com/mt3hr/lantana/src/app/lantana/lantana"
	"github.com/mt3hr/rykv/kyou"
	"github.com/spf13/cobra"
)

func Execute() {
	if err := appCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.MousetrapHelpText = "" // Windowsでマウスから起動しても怒られないようにする
	appCmd.PersistentFlags().StringVarP(&lantanaServer.ConfigFileName, "config_file", "c", "", "使用するコンフィグファイル")
	appCmd.PersistentFlags().StringVarP(&lantana.TagStructFile, "tag_struct_file", "t", "", "使用するタグ構造ファイル")
}

var (
	lantanaServer = lantana.NewLantanaServer()
	appCmd        = &cobra.Command{
		Use: "lantana",
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			err := lantanaServer.LoadConfig()
			if err != nil {
				log.Fatal(err)
			}
			err = lantanaServer.LoadTagStructFromFile()
			if err != nil {
				log.Fatal(err)
			}
			lantanaServer.Config.ApplicationConfig.HiddenTags = append(lantanaServer.Config.ApplicationConfig.HiddenTags, kyou.DeletedTagName)
		},
		Run: func(_ *cobra.Command, _ []string) {
			func() {
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

				go func() {
					err := lantanaServer.LaunchServer()
					if err != nil {
						log.Fatal(err)
					}
				}()

				address := ""
				if lantanaServer.Config.ServerConfig.TLS.Enable {
					address += "https://localhost"
				} else {
					address += "http://localhost"
				}
				address += lantanaServer.Config.ServerConfig.Address

				// Initialize astilectron
				a, err := astilectron.New(nil, astilectron.Options{
					AppName:            "lantana",
					VersionAstilectron: "0.51.0",
					VersionElectron:    "22.0.0",
					AppIconDefaultPath: "C:/Users/yamat/Git/lantana/public/favicon.png",
					AppIconDarwinPath:  "C:/Users/yamat/Git/lantana/public/favicon.ico",
				})
				if err != nil {
					fmt.Println("Electronが動かない環境であるかもしれません。その場合lantanaは動きませんので変わりにlantana serverを起動し、ブラウザからのアクセスを試みてください。")
					log.Fatal(err)
				}
				defer a.Close()

				// Start astilectron
				a.Start()

				contextIsolation := false
				// Create a new window
				w, err := a.NewWindow(address, &astilectron.WindowOptions{
					Height: astikit.IntPtr(1200),
					Width:  astikit.IntPtr(1500),
					WebPreferences: &astilectron.WebPreferences{
						AllowRunningInsecureContent: &contextIsolation,
					},
				})
				if err != nil {
					err = fmt.Errorf("error at new window: %w", err)
					log.Fatal(err)
				}

				openInDefaultBrowserMessagePrefix := "open_in_default_browser:"
				w.OnMessage(func(m *astilectron.EventMessage) interface{} {
					msg := ""
					m.Unmarshal(&msg)

					if strings.HasPrefix(msg, openInDefaultBrowserMessagePrefix) {
						url := strings.TrimSpace(strings.TrimPrefix(msg, openInDefaultBrowserMessagePrefix))
						openbrowser(url)
						return nil
					}
					return nil
				})
				w.Create()
				w.ExecuteJavaScript(`// aタグがクリックされた時にelectronで開かず、デフォルトのブラウザで開く
document.addEventListener('click', (e) => {
  for (let i = 0; i < e.path.length; i++) {
    let element = e.path[i]
	if (element.tagName === 'A') {
      e.preventDefault()
	  let aTag = element
	  let href = aTag.href
      astilectron.sendMessage('` + openInDefaultBrowserMessagePrefix + ` ' + href)
	}
  }
})
`)

				// Blocking pattern
				a.Wait()
			}()
			os.Exit(0)
		},
	}
)

func openbrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return err
}
