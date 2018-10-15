package cmd

import (
	"fmt"
	"log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

const description = "Sacred: Confluence Markdown Uploader"

var cfg Configuration
var cfgFile string
var outputDir string
var verbose bool
var dryRun bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use: "sacred",
	Short: description,
	Long: description,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.AddCommand(UploadCmd)
	RootCmd.AddCommand(versionCmd)

	viper.SetEnvPrefix("SACRED")
	viper.BindEnv("token")
	viper.BindEnv("domain")

	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default ./.sacred.yaml)")
	RootCmd.PersistentFlags().StringVarP(&outputDir, "output", "o", "", "output directory")
	RootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "verbose output")
	RootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "dry run")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName(".sacred")
	}

	// read in environment variables that match
	viper.AutomaticEnv()
}

func LoadConfig() {
	err := viper.ReadInConfig()
	if err == nil	&& verbose{
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Invalid config file: %s\n", cfgFile)
		log.Printf("Unable to decode into struct: %v\n", err)
	}

	MergeCredentials(&cfg)
}

func MergeCredentials(cfg *Configuration) {
	token := viper.Get("token")
	domain := viper.Get("domain")

	if token != nil {
		cfg.Auth.Token = token.(string)
	}

	if domain != nil {
		cfg.Auth.Domain = domain.(string)
	}
}
