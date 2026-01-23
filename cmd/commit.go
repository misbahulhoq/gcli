/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/misbahulhoq/gcm/utils"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !IsGitRepo() {
			fmt.Println("❌ Not a git repository. Did you run \"git init\" ?")
			return
		}
		diff, err := GetStagedChanges()

		if err != nil {
			fmt.Println(err)
			return
		}

		message := utils.GetMeaningfulCommitMessage(diff)
		Commit(message)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	commitCmd.Flags().BoolP("staged", "s", false, "Commit the staged changes")
	commitCmd.Flags().BoolP("all", "a", true, "Commit both staged and unstaged changes.")
}

func GetStagedChanges() (string, error) {
	cmd := exec.Command("git", "diff", "--staged")
	output, err := cmd.Output()

	if err != nil {
		return "", err
	}

	diff := strings.TrimSpace(string(output))

	if diff == "" {
		return "", fmt.Errorf("No staged changes found. Did you run \"git add\" ?")
	}

	return diff, nil
}

func GetAllChanges() {}

func Commit(message string) {
	// Print the commit message clearly
	fmt.Println("\n\nProposed commit message: ")
	fmt.Println("\n-----------------------------------------")
	fmt.Printf(" \n%s\n", message)
	fmt.Println("-----------------------------------------")
	// Ask for confirmation
	fmt.Print("Do you want to commit with this message? (Y/n): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(strings.ToLower(input))
	// We check for "y", "yes", or empty string (if you want Enter to mean Yes)
	if input == "y" || input == "yes" || input == "" {
		cmd := exec.Command("git", "commit", "-m", message)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error while committing ", err)
			return
		}
		fmt.Println("✅ Git commit successful")
	}

	fmt.Println("Commit Aborted ")

}
