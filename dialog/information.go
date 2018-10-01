package dialog

import (
	"github.com/fyne-io/fyne"
	"github.com/fyne-io/fyne/theme"
	"github.com/fyne-io/fyne/widget"
)

// ShowInformation shows a dialog over the specified application for user
// information. The title is used for the dialog window and message is the content.
func ShowInformation(title, message string, parent fyne.App) {
	dialog := newDialog(title, message, theme.InfoIcon(), nil, parent)
	dialog.setButtons(newButtonList(&widget.Button{Text: "OK",
		OnTapped: func() {
			dialog.response <- false
		},
	}))

	go dialog.wait()
	dialog.win.Show()
}
