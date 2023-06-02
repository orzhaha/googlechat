package main

import (
	"fmt"
	"log"
	// "time"

	astikit "github.com/asticode/go-astikit"
	astilectron "github.com/asticode/go-astilectron"

	"github.com/getlantern/systray"
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

	// // Init a new app menu
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

	w.OnMessage(func(m *astilectron.EventMessage) interface{} {
		// Unmarshal
		var count int
		m.Unmarshal(&count)

		systray.SetTitle(fmt.Sprintf("未讀訊息數:%d", count))

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

	go func() {
		a.Wait()
		systray.Quit()
	}()

	systray.Run(onReady, onExit)

	a.Wait()

}

func onReady() {
	// mIcon := systray.AddMenuItem(" ", "")
	// mQuit := systray.AddMenuItem("Quit", "Quit the application")
	// 監聽菜單項目的點擊事件，當點擊時向 clickCh 通道發送信號
	go func() {
		// for {
		// 	<-mIcon.ClickedCh
		// 	fmt.Println("###")
		// 	<-mQuit.ClickedCh
		// 	fmt.Println("###@")
		// }
	}()
	systray.SetTitle(fmt.Sprintf("未讀訊息數:%d", 0))
	// systray.SetTooltip("Pretty awesome超级棒")
}

func onExit() {
	// clean up here
}
