package main

import (
	"fmt"
	"log"

	astikit "github.com/asticode/go-astikit"
	astilectron "github.com/asticode/go-astilectron"
)

var (
	AppName            string
	BuiltAt            string
	VersionAstilectron string
	VersionElectron    string

	w *astilectron.Window
)

func main() {
	// Set logger
	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	// Create astilectron
	a, err := astilectron.New(l, astilectron.Options{
		AppName:            "Chat",
		AppIconDarwinPath:  "resources/icon.icns",
		AppIconDefaultPath: "resources/icon.png",
		SingleInstance:     true,
		VersionAstilectron: VersionAstilectron,
		VersionElectron:    VersionElectron,
	})

	if err != nil {
		l.Fatal(fmt.Errorf("main: creating astilectron failed: %w", err))
	}
	defer a.Close()

	// Handle signals
	a.HandleSignals()

	// Start
	if err = a.Start(); err != nil {
		l.Fatal(fmt.Errorf("main: starting astilectron failed: %w", err))
	}

	// Init a new app menu
	var m = a.NewMenu([]*astilectron.MenuItemOptions{
		{
			Label: astikit.StrPtr("Chat"),
			SubMenu: []*astilectron.MenuItemOptions{
				{Label: astikit.StrPtr("Minimize"), Role: astilectron.MenuItemRoleMinimize},
				{Label: astikit.StrPtr("Close"), Role: astilectron.MenuItemRoleClose},
				{Label: astikit.StrPtr("About"), Role: astilectron.MenuItemRoleAbout},
			},
		},
	})

	m.Create()

	// New window
	var w *astilectron.Window
	if w, err = a.NewWindow("https://mail.google.com/chat", &astilectron.WindowOptions{
		WebPreferences: &astilectron.WebPreferences{},
		Center:         astikit.BoolPtr(true),
		MinHeight:      astikit.IntPtr(700),
		MinWidth:       astikit.IntPtr(500),
		Height:         astikit.IntPtr(700),
		Width:          astikit.IntPtr(700),
	}); err != nil {
		l.Fatal(fmt.Errorf("main: new window failed: %w", err))
	}

	// Create windows
	if err = w.Create(); err != nil {
		l.Fatal(fmt.Errorf("main: creating window failed: %w", err))
	}

	var t = a.NewTray(&astilectron.TrayOptions{
		Image:   astikit.StrPtr("resources/icon.png"),
		Tooltip: astikit.StrPtr("tooltip"),
	})

	// Create tray
	t.Create()

	t.On(astilectron.EventNameTrayEventClicked, func(e astilectron.Event) (deleteListener bool) {
		// err = w.Focus()
		// if err != nil {
		// 	fmt.Println("err###############", err)
		// }
		// err = w.Show()
		// if err != nil {
		// 	fmt.Println("err###############", err)
		// }
		// fmt.Println("Tray item clicked")
		return
	})

	w.OnMessage(func(m *astilectron.EventMessage) interface{} {
		// Unmarshal
		var count int
		m.Unmarshal(&count)

		switch count {
		case 0:
			t.SetImage("resources/icon200.png")
		case 1:
			t.SetImage("resources/icon201.png")
		case 2:
			t.SetImage("resources/icon202.png")
		case 3:
			t.SetImage("resources/icon203.png")
		case 4:
			t.SetImage("resources/icon204.png")
		case 5:
			t.SetImage("resources/icon205.png")
		case 6:
			t.SetImage("resources/icon206.png")
		case 7:
			t.SetImage("resources/icon207.png")
		case 8:
			t.SetImage("resources/icon208.png")
		case 9:
			t.SetImage("resources/icon209.png")
		default:
			t.SetImage("resources/icon20o.png")

		}

		return nil
	})

	// 開發者工具
	// w.OpenDevTools()

	// 注入監聽未讀訊息的Js
	w.ExecuteJavaScript(`
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

	a.Wait()
}
