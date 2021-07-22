package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "learn_blockchain",
	Short: "Blockchain tool basic",
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "print the version of tool",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("v1.0.0")
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "print",
		Short: "print all block in blockchain",
		Run: func(cmd *cobra.Command, args []string) {
			bc, closeFunc := NewBlockChain()
			defer closeFunc()
			bc.Print()
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "add",
		Short: "add blocks with data",
		Run: func(cmd *cobra.Command, args []string) {
			bc, closeFunc := NewBlockChain()
			defer closeFunc()
			for _, val := range args {
				bc.AddBlock([]byte(val))
			}
		},
	})
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
