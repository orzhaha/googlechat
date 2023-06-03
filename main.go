package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

// Constants
const htmlAbout = `Welcome on <b>Astilectron</b> demo!<br>
This is using the bootstrap and the bundler.`

// Vars injected via ldflags by bundler
var (
	AppName            string
	BuiltAt            string
	VersionAstilectron string
	VersionElectron    string
)

// Application Vars
var (
	fs    = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	debug = fs.Bool("d", false, "enables the debug mode")
	w     *astilectron.Window
)

func main() {
	// Create logger
	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	// Parse flags
	fs.Parse(os.Args[1:])

	// Run bootstrap
	l.Printf("Running app built at %s\n", BuiltAt)

	err := bootstrap.Run(bootstrap.Options{
		Asset:    Asset,
		AssetDir: AssetDir,
		AstilectronOptions: astilectron.Options{
			AppName:            "Chat",
			AppIconDarwinPath:  "resources/icon.icns",
			AppIconDefaultPath: "resources/icon.png",
			SingleInstance:     true,
			VersionAstilectron: VersionAstilectron,
			VersionElectron:    VersionElectron,
		},
		Debug:  *debug,
		Logger: l,
		MenuOptions: []*astilectron.MenuItemOptions{
			{
				Label: astikit.StrPtr("Chat"),
				SubMenu: []*astilectron.MenuItemOptions{
					{Label: astikit.StrPtr("Minimize"), Role: astilectron.MenuItemRoleMinimize},
					{Accelerator: astilectron.NewAccelerator("Command", "Q"), Label: astikit.StrPtr("Close"), Role: astilectron.MenuItemRoleClose},
					{Accelerator: astilectron.NewAccelerator("Command", "R"), Label: astikit.StrPtr("Reload"), Role: astilectron.MenuItemRoleReload},
					{Accelerator: astilectron.NewAccelerator("F12"), Label: astikit.StrPtr("DevTool"), Role: astilectron.MenuItemRoleToggleDevTools},
					{Label: astikit.StrPtr("About"), Role: astilectron.MenuItemRoleAbout},
				},
			},
		},

		RestoreAssets: RestoreAssets,
		TrayOptions: &astilectron.TrayOptions{
			Image:   astikit.StrPtr("resources/icon200.png"),
			Tooltip: astikit.StrPtr("tooltip"),
		},
		Windows: []*bootstrap.Window{
			{
				Homepage: "https://mail.google.com/chat",
				Options: &astilectron.WindowOptions{
					Center:    astikit.BoolPtr(true),
					MinHeight: astikit.IntPtr(700),
					MinWidth:  astikit.IntPtr(500),
					Height:    astikit.IntPtr(700),
					Width:     astikit.IntPtr(700),
				},
			},
		},

		OnWait: func(a *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, t *astilectron.Tray, _ *astilectron.Menu) error {
			// Get the dock
			var d = a.Dock()

			basePaht := a.Paths().BaseDirectory()

			ws[0].OnMessage(func(m *astilectron.EventMessage) interface{} {
				// Unmarshal
				var count int
				m.Unmarshal(&count)

				switch count {
				case 0:
					d.SetBadge("0")
					t.SetImage(basePaht + "/resources/icon200.png")
				case 1:
					d.SetBadge("1")
					d.Bounce(astilectron.DockBounceTypeInformational)
					t.SetImage(basePaht + "/resources/icon201.png")
				case 2:
					d.SetBadge("2")
					d.Bounce(astilectron.DockBounceTypeInformational)
					t.SetImage(basePaht + "/resources/icon202.png")
				case 3:
					d.SetBadge("3")
					d.Bounce(astilectron.DockBounceTypeInformational)
					t.SetImage(basePaht + "/resources/icon203.png")
				case 4:
					d.SetBadge("4")
					d.Bounce(astilectron.DockBounceTypeInformational)
					t.SetImage(basePaht + "/resources/icon204.png")
				case 5:
					d.SetBadge("5")
					d.Bounce(astilectron.DockBounceTypeInformational)
					t.SetImage(basePaht + "/resources/icon205.png")
				case 6:
					d.SetBadge("6")
					d.Bounce(astilectron.DockBounceTypeInformational)
					t.SetImage(basePaht + "/resources/icon206.png")
				case 7:
					d.SetBadge("7")
					d.Bounce(astilectron.DockBounceTypeInformational)
					t.SetImage(basePaht + "/resources/icon207.png")
				case 8:
					d.SetBadge("8")
					d.Bounce(astilectron.DockBounceTypeInformational)
					t.SetImage(basePaht + "/resources/icon208.png")
				case 9:
					d.SetBadge("9")
					d.Bounce(astilectron.DockBounceTypeInformational)
					t.SetImage(basePaht + "/resources/icon209.png")
				default:
					d.SetBadge("...")
					t.SetImage(basePaht + "/resources/icon20o.png")

				}

				return nil
			})

			ws[0].ExecuteJavaScript(`
          // This will send a message to GO
          astilectron.sendMessage("hello", function(message) {
              console.log("received " + message)
          });

          const aktDoms = document.getElementsByClassName('akt');
          let count = 0;

          function computeCount() {
            count = 0;

            for (let i = 0; i < aktDoms.length; i++) {
              const span = aktDoms[i].getElementsByTagName('span')[0];

              count += Number(span.textContent);
            }

            const link = document.createElement('A');

              astilectron.sendMessage(count, function() {

              });
          }

          for (let i = 0; i < aktDoms.length; i++) {
            aktDoms[i].addEventListener('DOMSubtreeModified', () => {
              computeCount();
            });
          }

          computeCount();
      `)

			t.On(astilectron.EventNameTrayEventClicked, func(e astilectron.Event) (deleteListener bool) {

				return
			})

			return nil
		},
	})

	if err != nil {
		l.Fatal(fmt.Errorf("running bootstrap failed: %w", err))
	}
}
