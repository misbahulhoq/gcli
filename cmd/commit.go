/*
Copyright © 2026 Md Mezbah Uddin extraordinarymisbah@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/misbahulhoq/gcm/utils"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit all the changes by default. Use --staged to commit only the staged changes.",
	Long:  `Run this command to commit all the changes by default. Use --staged to commit only the staged changes. If you want to commit only the staged changes, run "git add" first.`,

	Run: func(cmd *cobra.Command, args []string) {
		if !IsGitRepo() {
			fmt.Println("❌ Not a git repository. Did you run \"git init\" ?")
			return
		}
		if IsWorkDirClean() {
			fmt.Println("Work directory is clean. No changes to commit.")
			return
		}

		var commitStagedChanges bool
		var diff string
		var err error
		commitStagedChanges, _ = cmd.Flags().GetBool("staged")

		if commitStagedChanges {
			diff, err = GetStagedChanges()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			diff, err = GetAllChanges()
			if err != nil {
				fmt.Println(err)
			}
		}

		message := utils.GetMeaningfulCommitMessage(diff)
		Commit(message)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	commitCmd.Flags().BoolP("staged", "s", false, "Commit the staged changes")
	commitCmd.Flags().BoolP("all", "a", true, "Commit both staged and unstaged changes.")
}

func IsWorkDirClean() bool {
	cmd := exec.Command("git", "status", "--porcelain")
	output, _ := cmd.Output()

	diff := strings.TrimSpace(string(output))
	if diff == "" {
		return true
	}
	return false
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

func GetAllChanges() (string, error) {
	var didStage = CheckAndStage()

	// If did CheckAndStage() function returns nil that means no changes found or the user has run  "git add .". In that case we operate it as if the user has run "git add ."
	if didStage == nil {
		cmd := exec.Command("git", "diff", "HEAD")
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

	if !*didStage {
		return GetStagedChanges()
	}
	return "", nil
}

func Commit(message string) {
	// Print the commit message clearly
	fmt.Println("\n\nProposed commit message: ")
	fmt.Println("\n-----------------------------------------")
	fmt.Printf(" \n%s\n", message)
	fmt.Println("-----------------------------------------")
	// Ask for confirmation
	if !utils.Confirm("Do you want to commit with this message? (Y/n): ") {
		color.Red("Commit Aborted ")
		return
	}

	cmd := exec.Command("git", "commit", "-a", "-m", message)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error while committing ", err)
		return
	}
	color.Green("✅ Git commit successful")
}

// CheckAndStage checks for unstaged work and prompts the user to add it
func CheckAndStage() *bool {
	// 1. Get the status in a machine-readable format
	cmd := exec.Command("git", "status", "--porcelain")
	outputBytes, _ := cmd.Output()
	output := string(outputBytes)

	// If output is empty, the tree is clean (no changes at all)
	if strings.TrimSpace(output) == "" {
		return nil
	}

	// 2. Check if there are any "unstaged" changes
	// In porcelain:
	// "?? file.txt" -> Untracked (New)
	// " M file.txt" -> Modified (Unstaged)
	// " D file.txt" -> Deleted (Unstaged)
	// "M  file.txt" -> Modified (Staged) - We ignore this one here

	var unstagedFiles []string
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if len(line) < 4 {
			continue
		}

		status := line[0:2]
		filePath := line[3:]
		green := color.New(color.FgGreen).SprintFunc()
		yellow := color.New(color.FgYellow).SprintfFunc()
		red := color.New(color.FgRed).SprintfFunc()

		// Check the 2-character status code
		// If it's "Untracked" (??) or the second char is NOT a space, it's unstaged
		if strings.HasPrefix(status, "??") {
			unstagedFiles = append(unstagedFiles, green("🆕  %s", filePath))
		} else if string(status[1]) == "M" {
			unstagedFiles = append(unstagedFiles, yellow("📝  %s", filePath))
		} else if string(status[1]) == "D" {
			unstagedFiles = append(unstagedFiles, red("🗑️  %s", filePath))
		}
	}

	// 3. If we found work that isn't staged, ask the user
	if len(unstagedFiles) > 0 {
		fmt.Println("\n⚠️  The following files are not staged yet: ")
		fmt.Println("-----------------------------------------")
		for _, f := range unstagedFiles {
			fmt.Println(f)
		}
		fmt.Println("-----------------------------------------")

		// Use your standard confirmation function
		if !utils.Confirm(" Do you want to run 'git add .' to include them? (Y/n): ") {
			fmt.Println("Committing only the staged changes.")
			// If the user declined to stage, return false so that we commit only the staged changes.
			var stage = false
			return &stage
		}

		color.Yellow("📦 Running git add .")
		err := exec.Command("git", "add", ".").Run()
		if err != nil {
			fmt.Println("❌ Error staging files:", err)
			os.Exit(1)
		}
		color.Green("✅ Files staged.")

	}
	return nil
}
