package ai

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var (
	client *genai.Client // Pointer to hold the Gemini client instance
	once   sync.Once     // Once object for ensuring single client creation
)

func NewClient() (*genai.Client, error) {
	once.Do(func() {
		ctx := context.Background()

		apiKey := os.Getenv("GEMINI_API_KEY")

		if apiKey == "" {
			log.Fatal("GEMINI_API_KEY not set")
		}

		var err error
		client, err = genai.NewClient(ctx, option.WithAPIKey(apiKey))
		if err != nil {
			log.Fatal(err)
		}
	})
	return client, nil
}

func Generate(ctx context.Context, template string, prompt string) (string, error) {
	client, err := NewClient()
	if err != nil {
		return "", err
	}

	model := client.GenerativeModel("gemini-pro")

	userPrompt := fmt.Sprintf(template, prompt)

	resp, err := model.GenerateContent(ctx, genai.Text(userPrompt))

	if err != nil {
		return "", err
	}

	part := getResponse(resp)

	if part == nil {
		return "", errors.New("please provide a valid prompt")
	}

	return fmt.Sprint(part), nil
}

func getResponse(resp *genai.GenerateContentResponse) genai.Part {
	var foundPart genai.Part

	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				if strings.Contains(fmt.Sprint(part), "NAVI_AI_ERROR") {
					return nil
				}
				foundPart = part
			}
		}
	}

	if foundPart == nil {
		return nil
	}

	return foundPart

}
