/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"personal/crypt/cmd/encrypt"
	"personal/crypt/cmd/decrypt"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "crypt",
	Short: "A CLI tool for encrypting and decrypting files.",
	Long: "A CLI tool for encrypting and decrypting files. \n \n" + 
			"This tool utilizes secret key encryption for encrypting/decrypting text files.\n" +
			"AES is used as the block cipher which is further enhanced by using a block cipher mode, " + 
			"GCM, to allow encrpt/decrypting arbitrary file sizes.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubcommands() {
	rootCmd.AddCommand(encrypt.EncryptCmd)
	rootCmd.AddCommand(decrypt.DecryptCmd)
}

func init() {
	addSubcommands()
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
}


