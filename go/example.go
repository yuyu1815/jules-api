package main

import (
	"fmt"
	"log"
	"os"
	"time"

	jules "github.com/jules-ai/jules-api-go"
)

func main() {
	// Initialize the client with your API key
	apiKey := os.Getenv("JULES_API_KEY")
	if apiKey == "" {
		log.Fatal("Please set the JULES_API_KEY environment variable")
	}

	client := jules.NewClient(jules.NewClientOptions(apiKey))

	// List available sources
	fmt.Println("🔍 Listing available sources...")
	sources, err := client.ListSources("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Available sources:")
	for _, source := range sources.Sources {
		fmt.Printf("- %s: %s\n", source.ID, source.Name)
		if source.GithubRepo != nil {
			fmt.Printf("  GitHub: %s/%s\n", source.GithubRepo.Owner, source.GithubRepo.Repo)
		}
	}

	if len(sources.Sources) == 0 {
		fmt.Println("No sources found. Please connect a GitHub repository in the Jules web app first.")
		return
	}

	// Create a new session
	fmt.Println("\n🚀 Creating a new session...")
	firstSource := sources.Sources[0]
	session, err := client.CreateSession(&jules.CreateSessionRequest{
		Prompt: "Create a simple web app that displays 'Hello from Jules!'",
		SourceContext: jules.SourceContext{
			Source: firstSource.Name,
			GithubRepoContext: &jules.GithubRepoContext{
				StartingBranch: "main",
			},
		},
		Title: "Hello World App Session",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("✅ Created session: %s\n", session.ID)
	fmt.Printf("📝 Title: %s\n", session.Title)
	fmt.Printf("🎯 Prompt: %s\n", session.Prompt)

	// Wait a moment for the agent to start working
	fmt.Println("\n⏳ Waiting a moment for the agent to start working...")
	time.Sleep(3 * time.Second)

	// List activities
	fmt.Println("\n📋 Listing activities...")
	activities, err := client.ListActivities(session.ID, 10, "")
	if err != nil {
		log.Printf("Error listing activities: %v", err)
	} else {
		fmt.Printf("Found %d activities:\n", len(activities.Activities))
		for _, activity := range activities.Activities {
			content := activity.Content
			if len(content) > 100 {
				content = content[:100] + "..."
			}
			fmt.Printf("- %s: %s\n", activity.Type, content)
		}
	}

	// Send a follow-up message
	fmt.Println("\n💬 Sending a follow-up message...")
	err = client.SendMessage(session.ID, &jules.SendMessageRequest{
		Prompt: "Please add some styling to make it look more attractive.",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ Message sent. The agent will respond in future activities.")
}
