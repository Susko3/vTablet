package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/getlantern/systray"
)

func initTray() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(AppIconData_Disconnected)
	systray.SetTitle("vTablet")
	systray.SetTooltip("vTablet")

	mAbout := systray.AddMenuItem("vTablet v2.0.1", "About")

	go func() {
		for {
			<-mAbout.ClickedCh
			openBrowser("https://github.com/Teages/vTablet")
		}
	}()

	// systray.AddSeparator()

	// sStartwithos := systray.AddMenuItemCheckbox("Start with OS", "Start with OS", false)
	// go func ()  {
	// 	for {
	// 		<-sStartwithos.ClickedCh
	// 		// ...
	// 	}
	// }()

	systray.AddSeparator()

	mAdb := systray.AddMenuItem("Restart ADB", "Restart ADB services")
	go func() {
		for {
			<-mAdb.ClickedCh
			restartAdbServices()
		}
	}()

	mQuit := systray.AddMenuItem("Quit", "Quit the app")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}

func onExit() {
	// clean up here
}

func updateConnectState(clientCount int) {
	if clientCount > 0 {
		onConnected()
	} else {
		onNoConnected()
	}
}

func onConnected() {
	systray.SetIcon(AppIconData_Connected)
}

func onNoConnected() {
	systray.SetIcon(AppIconData_Disconnected)
}
func openBrowser(url string) {
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
	if err != nil {
		log.Fatal(err)
	}

}