package ksettings

import (
	"github.com/rivo/tview"
)

// SettingsUI - Allows modification of
func (s *User) SettingsUI(title string, editors []string) {
	app := tview.NewApplication()

	form := tview.NewForm().
		AddCheckbox("Update on starting katbox", s.AutoUpdate, nil).
		AddDropDown("Editor", editors, 0, nil).
		AddInputField("(optional) Custom editor Path", s.Editor, 30, nil, nil).
		AddInputField("Git clone path", s.GitPath, 30, nil, nil).
		AddCheckbox("Open URLs in Browser", s.OpenURL, nil).
		AddButton("Save Settings", func() { app.Stop() })

	form.SetBorder(true).SetTitle(title).SetTitleAlign(tview.AlignLeft)
	if err := app.SetRoot(form, true).Run(); err != nil {
		panic(err)
	}

	// Retrieve values and update settings

	_, s.Editor = form.GetFormItemByLabel("Editor").(*tview.DropDown).GetCurrentOption()
	// If a custom editor has been selected then set the value from the custom Editor field
	if s.Editor == "Custom" {
		s.CustomEditor = form.GetFormItemByLabel("Editor Path").(*tview.InputField).GetText()
	}

	// TODO - do a OS/Editor lookup and set the path accordingly

	s.OpenURL = form.GetFormItemByLabel("Open URLs in Browser").(*tview.Checkbox).IsChecked()
}
