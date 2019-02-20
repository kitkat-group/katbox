package ksettings

// The Default settings directory should be {home_dir}/.katbox/
const ksettingsDir = ".katbox"

// The Default settings directory should be {home_dir}/.katbox/usrcfg.cfg
const ksettingsUserConfig = "usrcfg.json"

// KsettingsEditors -Editors (arrays are typically mutable, and can't be a const)
var KsettingsEditors = []string{"vi", "ed", "VSCode", "custom"}

// Paths MacOS
// - Visual Studio Code
const ksettingsMacVSCodePath = "/Applications/Visual Studio Code.app/Contents/Resources/app/bin/code"

// - Visual Studio Code Download Link (TODO - revisit)
const ksettingsMacVSCodeURL = "https://go.microsoft.com/fwlink/?LinkID=620882"

// TODO
// lookupEditor - will lookup the OS/Editor selection and set the default path
func lookupEditor(editor string) string {
	// OS lookup
	// -> Editor lookup
	// --> find default editor path
	return "TODO"
}

// setDefaults - will set a default configuration
func setDefaults(settingsFilePath string) *User {

	// Populate the default configuration struct
	u := User{
		AutoUpdate: false,
		GitPath:    "~/Documents",
		OpenURL:    true,
		Editor:     "vim",
		ContentURL: "https://raw.githubusercontent.com/kitkat-group/kontent/master/kitkat.json",
	}

	u.SettingsUI("Set new default settings", KsettingsEditors)

	// Save the updated settings

	return &u
}
