package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(echoCmd)
}

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "Echo anything to the screen",
	Long:  `echo is for echoing anything back.`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")

		line, err := reader.ReadSlice('\n')
		if err != nil {
			fmt.Println("An error occurred while reading input. Please try again.")
			return
		}
		fmt.Println("Text entered: ", strings.Split(string(line), " ")[2])
	},
}
