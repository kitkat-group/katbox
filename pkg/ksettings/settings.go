package ksettings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"

	"github.com/mitchellh/go-homedir"
)

// User are the katbox settings that relate to a users configuration
type User struct {

	// Update Options
	AutoUpdate bool `json:"autoupdate"`

	// Which editor will be used to open file(s) or directories
	Editor       string `json:"editor"`
	CustomEditor string `json:"customEditor,omitEmpty"`

	// TODO - Might be configurable down the line
	GoPath string `json:"gopath"`

	// Which paths will be used to clone projects
	GitPath string `json:"gitpath"`
	// Should URLs be opened by the (default) browser
	OpenURL bool `json:"openUrl"`

	// ContentURL points to a different kitkat url for source material
	ContentURL string `json:"customURL"`
}

// GenerateDefaultSettings - This will create a set of defaults and allow the user to modify them
func GenerateDefaultSettings(path string) (*User, error) {

	// Start by checking the home directory
	homeDirectory, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	if path == "" {
		path = fmt.Sprintf("%s/%s/%s", homeDirectory, ksettingsDir, ksettingsUserConfig)
	}
	// Build new defaults
	usr := setDefaults(path)
	usr.SaveUserSettingsToFile(path)
	return usr, nil
}

// LoadUserSettings - this will attempt to populate the user settings from the default file
func LoadUserSettings() (*User, error) {
	// Start by checking the home directory
	homeDirectory, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s/%s/%s", homeDirectory, ksettingsDir, ksettingsUserConfig)

	// Use the load Settings function
	usr, err := LoadUserSettingsFromFile(path)
	if err != nil {
		log.Infof("%v", err)
	}
	return usr, nil
}

// LoadUserSettingsFromFile - this will attempt to populate the user settings from a local file
func LoadUserSettingsFromFile(filePath string) (*User, error) {
	// Check if the file exists, if so attempt to parse it
	if _, err := os.Stat(filePath); !os.IsExist(err) {
		// Attempt to read the file into a []byte buffer
		b, err := ioutil.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		// Create a struct to parse into
		var usr User
		// Unmarshall the rawbytes into a struct
		err = json.Unmarshal(b, &usr)
		if err != nil {
			return nil, err
		}
		// Return the user settings
		return &usr, nil
	}
	return nil, fmt.Errorf("File %s doesn't exist", filePath)
}

// SaveUserSettingsToFile - This will save User settings to a file
func (u *User) SaveUserSettingsToFile(settingsPath string) error {

	// Marshall the configuration
	b, err := json.MarshalIndent(u, "", "\t")
	if err != nil {
		return err
	}
	directoryPath := filepath.Dir(settingsPath)
	// Create the path if it doesn't exist
	os.MkdirAll(directoryPath, os.ModePerm)
	f, err := os.Create(settingsPath)
	if err != nil {
		return err
	}
	defer f.Close()
	byteCount, err := f.WriteString(string(b))
	if err != nil {
		return err
	}
	log.Infof("Written %d bytes to file [%s]", byteCount, settingsPath)
	f.Sync()

	return nil
}
