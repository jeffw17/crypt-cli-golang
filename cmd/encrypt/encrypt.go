/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package encrypt

import (
	"personal/crypt/helper"
	"github.com/spf13/cobra"
)

// EncryptCmd represents the encrypt command
var EncryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt the given file",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		isSet := cmd.Flags().Lookup("path").Changed

		if !isSet {
			cmd.Help()
		} else {
			path, _ := cmd.Flags().GetString("path")
		
			helper.EncryptFile(path)
		}	
	},
}

func init() {
	EncryptCmd.Flags().StringP("path", "p", "", "Path to file")
}
