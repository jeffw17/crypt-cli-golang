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
	Short: "Encrypt a text file with a secret key.",
	Long: "Encrypt a text file with a secret key. \n" + 
	      "Must provide path to text file (absolute/relative) and a 32 byte secret key string (keep this safe as its required for decryption).",
	Run: func(cmd *cobra.Command, args []string) {
		isSet := cmd.Flags().Lookup("path").Changed

		if !isSet {
			cmd.Help()
		} else {
			path, _ := cmd.Flags().GetString("path")
			key, _ := cmd.Flags().GetString("key")
		
			helper.EncryptFile(path, key)
		}	
	},
}

func init() {
	EncryptCmd.Flags().StringP("path", "p", "", "path to text file")
	EncryptCmd.Flags().StringP("key", "k", "", "secret key (size 32)")
}
