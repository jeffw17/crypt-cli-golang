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
	Short: "Decrypt the given file",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		
		helper.DecryptFile(path)
	},
}

func init() {
	DecryptCmd.Flags().StringP("path", "p", "", "Path to file")
}
