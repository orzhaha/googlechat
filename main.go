package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	astikit "github.com/asticode/go-astikit"
	astilectron "github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

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
					{Label: astikit.StrPtr("About"), Role: astilectron.MenuItemRoleAbout},
					{Type: astilectron.MenuItemTypeSeparator},
					{Checked: astikit.BoolPtr(true), Label: astikit.StrPtr("Bounce"), Type: astilectron.MenuItemTypeCheckbox},
					{Type: astilectron.MenuItemTypeSeparator},
					{Accelerator: astilectron.NewAccelerator("Command", "Q"), Label: astikit.StrPtr("Close"), Role: astilectron.MenuItemRoleClose},
					{Type: astilectron.MenuItemTypeSeparator},
				},
			},
			{
				Label: astikit.StrPtr("File"),
				SubMenu: []*astilectron.MenuItemOptions{
					{Type: astilectron.MenuItemTypeSeparator},
				},
			},
			{
				Label: astikit.StrPtr("Edit"),
				SubMenu: []*astilectron.MenuItemOptions{
					{Accelerator: astilectron.NewAccelerator("Command", "Z"), Label: astikit.StrPtr("Undo"), Role: astilectron.MenuItemRoleUndo},
					{Label: astikit.StrPtr("Redo"), Role: astilectron.MenuItemRoleRedo},
					{Type: astilectron.MenuItemTypeSeparator},
					{Accelerator: astilectron.NewAccelerator("Command", "X"), Label: astikit.StrPtr("Cut"), Role: astilectron.MenuItemRoleCut},
					{Accelerator: astilectron.NewAccelerator("Command", "C"), Label: astikit.StrPtr("Copy"), Role: astilectron.MenuItemRoleCopy},
					{Accelerator: astilectron.NewAccelerator("Command", "V"), Label: astikit.StrPtr("Paste"), Role: astilectron.MenuItemRolePaste},
					{Label: astikit.StrPtr("PasteAndMatchStyle"), Role: astilectron.MenuItemRolePasteAndMatchStyle},
					{Accelerator: astilectron.NewAccelerator("Command", "A"), Label: astikit.StrPtr("SelectAll"), Role: astilectron.MenuItemRoleSelectAll},
				},
			},
			{
				Label: astikit.StrPtr("View"),
				SubMenu: []*astilectron.MenuItemOptions{
					{Accelerator: astilectron.NewAccelerator("Command", "R"), Label: astikit.StrPtr("Reload"), Role: astilectron.MenuItemRoleReload},
					{Type: astilectron.MenuItemTypeSeparator},
					{Label: astikit.StrPtr("Minimize"), Role: astilectron.MenuItemRoleMinimize},
					{Label: astikit.StrPtr("Resetzoom"), Role: astilectron.MenuItemRoleResetZoom},
					{Label: astikit.StrPtr("ZoomOut"), Role: astilectron.MenuItemRoleZoomOut},
					{Label: astikit.StrPtr("ZoomIn"), Role: astilectron.MenuItemRoleZoomIn},
					{Type: astilectron.MenuItemTypeSeparator},
					{Accelerator: astilectron.NewAccelerator("F12"), Label: astikit.StrPtr("DevTool"), Role: astilectron.MenuItemRoleToggleDevTools},
				},
			},
			{
				Label: astikit.StrPtr("Window"),
				SubMenu: []*astilectron.MenuItemOptions{
					{Type: astilectron.MenuItemTypeSeparator},
				},
			},
			{
				Label: astikit.StrPtr("Help"),
				SubMenu: []*astilectron.MenuItemOptions{
					{Type: astilectron.MenuItemTypeSeparator},
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

		OnWait: func(a *astilectron.Astilectron, ws []*astilectron.Window, m *astilectron.Menu, t *astilectron.Tray, _ *astilectron.Menu) error {
			BounceMi, _ := m.Item(0, 2)
			BounceEnable := true

			BounceMi.On(astilectron.EventNameMenuItemEventClicked, func(e astilectron.Event) bool {
				if *e.MenuItemOptions.Checked {
					BounceEnable = true
				} else {
					BounceEnable = false
				}

				return false
			})

			// Get the dock
			var d = a.Dock()

			basePaht := a.Paths().BaseDirectory()

			// 畫面reload實在載入一次
			ws[0].On(astilectron.EventNameWindowEventReadyToShow, func(e astilectron.Event) (deleteListener bool) {
				ws[0].ExecuteJavaScript(executeJavaScript)
				return
			})

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
					if BounceEnable {
						d.Bounce(astilectron.DockBounceTypeInformational)
					}
					t.SetImage(basePaht + "/resources/icon201.png")
				case 2:
					d.SetBadge("2")
					if BounceEnable {
						d.Bounce(astilectron.DockBounceTypeInformational)
					}
					t.SetImage(basePaht + "/resources/icon202.png")
				case 3:
					d.SetBadge("3")
					if BounceEnable {
						d.Bounce(astilectron.DockBounceTypeInformational)
					}
					t.SetImage(basePaht + "/resources/icon203.png")
				case 4:
					d.SetBadge("4")
					if BounceEnable {
						d.Bounce(astilectron.DockBounceTypeInformational)
					}
					t.SetImage(basePaht + "/resources/icon204.png")
				case 5:
					d.SetBadge("5")
					if BounceEnable {
						d.Bounce(astilectron.DockBounceTypeInformational)
					}
					t.SetImage(basePaht + "/resources/icon205.png")
				case 6:
					d.SetBadge("6")
					if BounceEnable {
						d.Bounce(astilectron.DockBounceTypeInformational)
					}
					t.SetImage(basePaht + "/resources/icon206.png")
				case 7:
					d.SetBadge("7")
					if BounceEnable {
						d.Bounce(astilectron.DockBounceTypeInformational)
					}
					t.SetImage(basePaht + "/resources/icon207.png")
				case 8:
					d.SetBadge("8")
					if BounceEnable {
						d.Bounce(astilectron.DockBounceTypeInformational)
					}
					t.SetImage(basePaht + "/resources/icon208.png")
				case 9:
					d.SetBadge("9")
					if BounceEnable {
						d.Bounce(astilectron.DockBounceTypeInformational)
					}
					t.SetImage(basePaht + "/resources/icon209.png")
				default:
					d.SetBadge("...")
					if BounceEnable {
						d.Bounce(astilectron.DockBounceTypeInformational)
					}
					t.SetImage(basePaht + "/resources/icon20o.png")

				}

				return nil
			})

			ws[0].ExecuteJavaScript(executeJavaScript)

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
