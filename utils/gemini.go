package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"google.golang.org/genai"
)

func GetMeaningfulCommitMessage(changes string) string {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)

	s.Suffix = " Generating commit message..."
	s.Color("cyan")
	s.Start()

	if err != nil {
		fmt.Println(err)
		return ""
	}

	result, err := client.Models.GenerateContent(ctx, "gemini-2.5-flash", genai.Text("You are a git commit message generator. Output ONLY the raw commit message using Conventional Commits format. DO NOT use markdown code blocks (```). DO NOT add conversational text like 'Here is the commit message'. DO NOT wrap the output in quotes."+changes), nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	s.Stop()

	return result.Text()

}
