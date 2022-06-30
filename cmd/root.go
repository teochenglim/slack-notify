/*
Copyright © 2022 Teo Cheng Lim teochenglim@gmail.com

*/
package cmd

import (
	"os"
	// "fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/teochenglim/slack-notify/pkg/slack-notify"
)

var (

	// used for register command
	channel     string
	color       string
	icon        string
	message     string
	title       string
	username    string
	verbose     bool
	webhook     string

	rootCmd = &cobra.Command{
		Use:   "slack-notify",
		Short: "A custom slack notify client",
		Long: `slack-notify is CLI client tools to send slack message by given simple inputs.
You need to get SLACK_WEBHOOK from slack website.
Web hook explain here https://api.slack.com/messaging/webhooks
The minimum inputs are SLACK_WEBHOOK and SLACK_MESSAGE which is the text you want to send.`,
		Run: func(cmd *cobra.Command, args []string) {

			s := slack_notify.New()

			/// No config file found. Workaround to write to Slack object
			s.Config.SLACK_CHANNEL  = os.Getenv("SLACK_CHANNEL")
			s.Config.SLACK_COLOR    = os.Getenv("SLACK_COLOR")
			s.Config.SLACK_ICON     = os.Getenv("SLACK_ICON")
			s.Config.SLACK_MESSAGE  = os.Getenv("SLACK_MESSAGE")
			s.Config.SLACK_TITLE    = os.Getenv("SLACK_TITLE")
			s.Config.SLACK_USERNAME = os.Getenv("SLACK_USERNAME")
			verbose_env, present := os.LookupEnv("SLACK_VERBOSE")
			if present && verbose_env == "true" { s.Config.SLACK_VERBOSE  = true }
			s.Config.SLACK_WEBHOOK  = os.Getenv("SLACK_WEBHOOK")

			if channel  != ""   { s.Config.SLACK_CHANNEL = channel }
			if color    != ""   { s.Config.SLACK_COLOR = color }
			if icon     != ""   { s.Config.SLACK_ICON = icon }
			if message  != ""   { s.Config.SLACK_MESSAGE = message }
			if title    != ""   { s.Config.SLACK_TITLE = title }
			if username != ""   { s.Config.SLACK_USERNAME = username }
			if verbose  == true { s.Config.SLACK_VERBOSE = verbose }
			if webhook  != ""   { s.Config.SLACK_WEBHOOK = webhook }

			s_err := s.SendMessage()

			if s_err != nil {
				panic(s_err)
			}
    },
	}
)


// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&channel , "channel",  "c", "#chatops", "Slack channel name")
	rootCmd.PersistentFlags().StringVarP(&color   , "color",    "o", "", "Slack message color #7CD197")
	rootCmd.PersistentFlags().StringVarP(&icon    , "icon",     "i", "", "Slack icon URL https://slack.com/img/icons/app-57.png")
	rootCmd.PersistentFlags().StringVarP(&message , "message",  "m", "", "Slack message text")
	rootCmd.PersistentFlags().StringVarP(&title   , "title",    "t", "", "Slack message title")
	rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "Slack message display name")
	rootCmd.PersistentFlags().StringVarP(&webhook , "webook",   "w", "", "Slack webhook url https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX")

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	// if viper config file is found. mix it in with the values
	viper.BindPFlag("SLACK_CHANNEL",  rootCmd.PersistentFlags().Lookup("channel"))
	viper.BindPFlag("SLACK_COLOR",    rootCmd.PersistentFlags().Lookup("color"))
	viper.BindPFlag("SLACK_ICON",     rootCmd.PersistentFlags().Lookup("icon"))
	viper.BindPFlag("SLACK_MESSAGE",  rootCmd.PersistentFlags().Lookup("message"))
	viper.BindPFlag("SLACK_TITLE",    rootCmd.PersistentFlags().Lookup("title"))
	viper.BindPFlag("SLACK_USERNAME", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("SLACK_WEBHOOK",  rootCmd.PersistentFlags().Lookup("webook"))
	viper.BindPFlag("SLACK_VERBOSE",  rootCmd.PersistentFlags().Lookup("verbose"))

}

func initConfig() {

}
