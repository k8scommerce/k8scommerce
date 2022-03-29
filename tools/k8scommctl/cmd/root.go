/*
Copyright © 2022 K8sCommerce
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "k8scomm2hugo",
	Short: "Generates Hugo markdown files from K8sCommerce",
	Long:  `Generates Hugo markdown files from K8sCommerce using the K8sCommerce API for category & product generation.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
