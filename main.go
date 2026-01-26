/*
Copyright © 2026 Md Mezbah Uddin extraordinarymisbah@gmail.com
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/misbahulhoq/gcm/cmd"
)

func main() {
	_ = godotenv.Load()
	key := GetAPIKey()
	os.Setenv("GEMINI_API_KEY", key)
	cmd.Execute()
}

// GetAPIKey handles the "Login" logic
func GetAPIKey() string {
	// A. Check Environment Variable first (Standard for servers/CI/Devs)
	// If godotenv.Load() found a file, this will already be set.
	if key := os.Getenv("GEMINI_API_KEY"); key != "" {
		return key
	}

	// B. Check Config File (Standard for installed users)
	// We look for a file named .gcm_config in the user's home directory
	home, err := os.UserHomeDir()
	if err == nil {
		configPath := filepath.Join(home, ".gcm_config")
		if content, err := os.ReadFile(configPath); err == nil {
			// Found it! Return the saved key
			return strings.TrimSpace(string(content))
		}
	}

	// C. If neither exists, ASK the user (First Run Experience)
	fmt.Println("\n🔑 Gemini API Key not found.")
	fmt.Println("   To use this tool, you need a free API key from Google AI Studio.")
	fmt.Println("   Get it here: https://aistudio.google.com/app/apikey")
	fmt.Print("\n👉 Paste your API Key here: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	key := strings.TrimSpace(input)

	if key == "" {
		fmt.Println("❌ API Key is required. Exiting.")
		os.Exit(1)
	}

	// D. Save it for next time
	if home != "" {
		configPath := filepath.Join(home, ".gcm_config")
		// 0600 means "only the owner can read/write this file" (secure)
		err := os.WriteFile(configPath, []byte(key), 0600)
		if err == nil {
			fmt.Printf("✅ Key saved securely to %s\n\n", configPath)
		} else {
			fmt.Println("⚠️  Could not save key to file, but proceeding for this session.")
		}
	}

	return key
}
