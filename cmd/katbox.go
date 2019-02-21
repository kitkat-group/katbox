package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kitkat-group/katbox/pkg/kontent"

	"github.com/kitkat-group/katbox/pkg/ksettings"
	"github.com/kitkat-group/katbox/pkg/ui"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

var userSettingsPath *string

// Release - this struct contains the release information populated when building katbox
var Release struct {
	Version string
	Build   string
}

var katboxCmd = &cobra.Command{
	Use:   "katbox",
	Short: "The \"katbox\" is used to manage articles, artifacts and resources for the KitKat team",
	Run: func(cmd *cobra.Command, args []string) {
		log.SetLevel(log.Level(logLevel))

		// Load the User preferences
		settings, err := ksettings.LoadUserSettings()
		if err != nil {
			log.Fatalf("No settings for user generated, create with katbox user")
		}

		// Retreive the content
		articles, err := kontent.Retrieve(settings)
		if err != nil {
			log.Fatalf("%v", err)
		}

		// Retreive the content
		ui.MainUI(articles, settings)
		return
	},
}

var katboxCmdUser = &cobra.Command{
	Use:   "user",
	Short: "The user sub-command allows configuration of user settings",
	Run: func(cmd *cobra.Command, args []string) {
		log.SetLevel(log.Level(logLevel))
		cmd.Help()
		u, err := ksettings.LoadUserSettings()
		if err != nil {
			ksettings.GenerateDefaultSettings(*userSettingsPath)
			return
		}
		u.EditExistingSettings(*userSettingsPath, "Edit existing Settings")
		return
	},
}

var logLevel int

func init() {
	userSettingsPath = katboxCmdUser.Flags().String("settings", "", "Path to User Settings")
	katboxCmd.AddCommand(katboxCmdUser)
}

// Execute - starts the command parsing process
func Execute() {
	if os.Getenv("KATBOX_LOGLEVEL") != "" {
		i, err := strconv.ParseInt(os.Getenv("KATBOX_LOGLEVEL"), 10, 8)
		if err != nil {
			log.Fatalf("Error parsing environment variable [KATBOX_LOGLEVEL")
		}
		// We've only parsed to an 8bit integer, however i is still a int64 so needs casting
		logLevel = int(i)
	} else {
		// Default to logging anything Info and below
		logLevel = int(log.InfoLevel)
	}

	log.SetLevel(log.Level(logLevel))
	if err := katboxCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var katboxVersion = &cobra.Command{
	Use:   "version",
	Short: "Version and Release information about the plunder tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Plunder Release Information\n")
		fmt.Printf("Version:  %s\n", Release.Version)
		fmt.Printf("Build:    %s\n", Release.Build)
	},
}
