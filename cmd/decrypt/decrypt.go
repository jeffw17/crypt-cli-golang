/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package decrypt

import (
	"personal/crypt/helper"
	"github.com/spf13/cobra"
)

// DecryptCmd represents the decrypt command
var DecryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt a file with a secret key",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		isSet := cmd.Flags().Lookup("path").Changed

		if !isSet {
			cmd.Help()
		} else {
			path, _ := cmd.Flags().GetString("path")
			key, _ := cmd.Flags().GetString("key")
		
			helper.DecryptFile(path, key)
		}	
	},
}

func init() {
	DecryptCmd.Flags().StringP("path", "p", "", "Path to file")
	DecryptCmd.Flags().StringP("key", "k", "", "Secret key")
}
