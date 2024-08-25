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
	Short: "Decrypt a file with a secret key.",
	Long: "Decrypt a file with a secret key. \n" + 
		  "Must provide path to text file (absolute/relative) and the 32 byte secret key string used during encryption.",
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
	DecryptCmd.Flags().StringP("path", "p", "", "path to text file")
	DecryptCmd.Flags().StringP("key", "k", "", "secret key (size 32)")
}
