package gui

import (
	"fmt"
	"havoc-launcher/launcher"
	"havoc-launcher/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func StartGUI() {
	a := app.New()
	w := a.NewWindow("Havoc Launcher")
	w.Resize(fyne.NewSize(500, 400))
	w.SetFixedSize(true)

	config := util.GetConfig()

	//Play Page
	fnPathEntry := widget.NewEntry()
	fnPathEntry.SetPlaceHolder("Fortnite Path")

	selectButton := widget.NewButtonWithIcon("", theme.FileIcon(), func() {
		dialog.ShowFolderOpen(func(selectedFolder fyne.ListableURI, err error) {
			if err == nil {
				// Update the Entry widget with the selected folder path
				fnPathEntry.SetText(selectedFolder.Path())
			} else {
				fmt.Println("Error opening folder:", err)
			}
		}, w)
	})

	pathForm := container.New(layout.NewFormLayout(), selectButton, fnPathEntry)

	//startServerCheck := widget.NewCheck("Start As Server", func(b bool) {
	//})

	//Launch Button
	launchButton := widget.NewButtonWithIcon("Launch", theme.MediaPlayIcon(), func() {
		//Check to make sure fn is in given path
		if launcher.VerifyFnPath(fnPathEntry.Text) {
			launcher.StartFortnite(launcher.GetLauncherType())
		} else {
			showWarningDialog(w, "path does not contain fortnite")
		}
	})

	playContent := container.NewVBox(
		widget.NewSeparator(),
		pathForm,
		layout.NewSpacer(),
		widget.NewSeparator(),
		//startServerCheck,
		launchButton,
	)

	//Options Page
	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Username")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Password")

	randPassCheck := widget.NewCheck("Use Random Password", func(b bool) {
		if b {
			passwordEntry.Disable()
		} else {
			passwordEntry.Enable()
		}
	})

	saveButton := widget.NewButtonWithIcon("Save", theme.ConfirmIcon(), func() {
		util.UpdateConfigCredentials(config, usernameEntry.Text, passwordEntry.Text, randPassCheck.Checked)
	})

	consoleCheck := widget.NewCheck("Inject Console Unlocker (Not Working)", func(b bool) {
	})

	memleakCheck := widget.NewCheck("Inject Memory Leak Fix", func(b bool) {
	})

	optionsContent := container.NewVBox(
		usernameEntry,
		passwordEntry,
		randPassCheck,
		saveButton,
		layout.NewSpacer(),
		widget.NewSeparator(),
		consoleCheck,
		memleakCheck,
	)

	//Set the ui
	usernameEntry.SetText(launcher.Username)
	passwordEntry.SetText(launcher.Password)
	randPassCheck.SetChecked(config.Credentials.RandomPass)

	consoleCheck.SetChecked(config.Options.Console)
	memleakCheck.SetChecked(config.Options.MemoryFix)

	//About Page
	text1 := widget.NewLabel("Havoc Launcher V1.0 - Created by Walter Man#0857")
	text2 := widget.NewLabel("Cobalt - https://github.com/Milxnor/Cobalt")

	aboutContent := container.NewVBox(
		text1,
		text2,
	)

	tabs := container.NewAppTabs(
		container.NewTabItem("Play", playContent),
		container.NewTabItem("Options", optionsContent),
		container.NewTabItem("About", aboutContent),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	w.SetContent(tabs)
	w.ShowAndRun()
}

func showWarningDialog(win fyne.Window, message string) {
	dialog.ShowError(fmt.Errorf(message), win)
}

/*func showRunningDialog(window fyne.Window) {
	confirmDialog := dialog.NewConfirm("Fortnite is already running", "Are you sure you want to start another process?", func(response bool) {
		if response {
			fmt.Println("User confirmed.")
			// Perform the desired action when the user confirms
		} else {
			fmt.Println("User canceled.")
			// Perform the desired action when the user cancels
		}
	}, window)

	confirmDialog.Show()
}*/
