package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	jules "github.com/yuyu1815/jules-api/go"
)

func testListSources(client *jules.Client) ([]jules.Source, bool) {
	fmt.Println("📋 Testing: List Sources")
	sources, err := client.ListSources("", nil)
	if err != nil {
		fmt.Printf("   ❌ Failed: %v\n", err)
		return nil, false
	}

	fmt.Printf("   ✅ Success: Found %d sources\n", len(sources.Sources))
	for i, source := range sources.Sources[:5] {
		fmt.Printf("      [%d] %s: %s\n", i+1, source.ID, source.Name)
		if source.GithubRepo != nil {
			fmt.Printf("          GitHub: %s/%s\n", source.GithubRepo.Owner, source.GithubRepo.Repo)
		}
	}
	if len(sources.Sources) > 5 {
		fmt.Printf("      ... and %d more sources\n", len(sources.Sources)-5)
	}
	return sources.Sources, len(sources.Sources) > 0
}

func testCreateSession(client *jules.Client, sources []jules.Source) (*jules.Session, bool) {
	fmt.Println("\n🚀 Testing: Create Session")
	if len(sources) == 0 {
		fmt.Println("   ⚠️  Skipping: No sources available")
		return nil, false
	}

	firstSource := sources[0]
	fmt.Printf("   Using source: %s\n", firstSource.ID)

	request := &jules.CreateSessionRequest{
		Prompt:  "Create a simple test to verify the API is working correctly.",
		SourceContext: jules.SourceContext{
			Source: firstSource.Name,
			GithubRepoContext: &jules.GithubRepoContext{
				StartingBranch: "main",
			},
		},
		Title: "API Test Session - Go",
		RequirePlanApproval: false,
	}

	session, err := client.CreateSession(request, nil)
	if err != nil {
		fmt.Printf("   ❌ Failed: %v\n", err)
		return nil, false
	}

	fmt.Println("   ✅ Success: Session created")
	fmt.Printf("      ID: %s\n", session.ID)
	fmt.Printf("      Title: %s\n", session.Title)
	fmt.Printf("      Name: %s\n", session.Name)
	return session, true
}

func testGetSession(client *jules.Client, sessionID string) (*jules.Session, bool) {
	fmt.Println("\n📖 Testing: Get Session")
	if sessionID == "" {
		fmt.Println("   ⚠️  Skipping: No session ID available")
		return nil, false
	}

	session, err := client.GetSession(sessionID, nil)
	if err != nil {
		fmt.Printf("   ❌ Failed: %v\n", err)
		return nil, false
	}

	fmt.Println("   ✅ Success: Session retrieved")
	fmt.Printf("      ID: %s\n", session.ID)
	fmt.Printf("      Title: %s\n", session.Title)
	return session, true
}

func testListSessions(client *jules.Client) ([]jules.Session, bool) {
	fmt.Println("\n📂 Testing: List Sessions")
	sessions, err := client.ListSessions(5, "", nil)
	if err != nil {
		fmt.Printf("   ❌ Failed: %v\n", err)
		return nil, false
	}

	fmt.Printf("   ✅ Success: Found %d sessions\n", len(sessions.Sessions))
	if sessions.NextPageToken != "" {
		fmt.Printf("      Next page token: %s\n", sessions.NextPageToken)
	}
	return sessions.Sessions, true
}

func testListActivities(client *jules.Client, sessionID string) ([]jules.Activity, bool) {
	fmt.Println("\n🎬 Testing: List Activities")
	if sessionID == "" {
		fmt.Println("   ⚠️  Skipping: No session ID available")
		return nil, false
	}

	activities, err := client.ListActivities(sessionID, 10, "", nil)
	if err != nil {
		fmt.Printf("   ❌ Failed: %v\n", err)
		return nil, false
	}

	fmt.Printf("   ✅ Success: Found %d activities\n", len(activities.Activities))
	for i, activity := range activities.Activities[:3] {
		timestamp := "No timestamp"
		if !activity.Timestamp.IsZero() {
			timestamp = activity.Timestamp.Format("15:04:05")
		}
		content := "No content"
		if activity.Content != "" {
			if len(activity.Content) > 50 {
				content = activity.Content[:50] + "..."
			} else {
				content = activity.Content
			}
		}
		fmt.Printf("      [%d] %s @ %s: %s\n", i+1, activity.Type, timestamp, content)
	}
	if len(activities.Activities) > 3 {
		fmt.Printf("      ... and %d more activities\n", len(activities.Activities)-3)
	}
	return activities.Activities, len(activities.Activities) >= 0
}

func testSendMessage(client *jules.Client, sessionID string) bool {
	fmt.Println("\n💬 Testing: Send Message")
	if sessionID == "" {
		fmt.Println("   ⚠️  Skipping: No session ID available")
		return false
	}

	request := &jules.SendMessageRequest{
		Prompt: "Please confirm that the API testing is working correctly by acknowledging this message.",
	}

	err := client.SendMessage(sessionID, request, nil)
	if err != nil {
		fmt.Printf("   ❌ Failed: %v\n", err)
		return false
	}

	fmt.Println("   ✅ Success: Message sent")
	return true
}

func testGetSource(client *jules.Client, sources []jules.Source) (*jules.Source, bool) {
	fmt.Println("\n📦 Testing: Get Source")
	if len(sources) == 0 {
		fmt.Println("   ⚠️  Skipping: No sources available")
		return nil, false
	}

	sourceID := sources[0].ID
	source, err := client.GetSource(sourceID, nil)
	if err != nil {
		fmt.Printf("   ❌ Failed: %v\n", err)
		return nil, false
	}

	fmt.Println("   ✅ Success: Source retrieved")
	fmt.Printf("      ID: %s\n", source.ID)
	fmt.Printf("      Name: %s\n", source.Name)
	return source, true
}

func main() {
	fmt.Println("🧪 Jules API Comprehensive Test Suite - Go Version")
	fmt.Println("=" + strings.Repeat("=", 50))
	fmt.Printf("⏰ Test started at: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println()

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("⚠️  Warning: Could not load .env file. Continuing without it.")
	}

	apiKey := os.Getenv("JULES_API_KEY")
	if apiKey == "" {
		fmt.Println("❌ Error: JULES_API_KEY environment variable not found.")
		fmt.Println("   Please create a test/.env file with:")
		fmt.Println("   JULES_API_KEY=your_api_key_here")
		os.Exit(1)
	}

	fmt.Printf("🔑 Using API Key from .env: %s...\n", apiKey[:20])
	fmt.Println()

	// Create client with infinite timeout
	options := jules.NewClientOptions()
	options.APIKey = &apiKey
	timeout := -1 * time.Second
	options.Timeout = &timeout
	client, err := jules.NewClient(options)
	if err != nil {
		fmt.Printf("❌ Error creating client: %v\n", err)
		os.Exit(1)
	}

	testResults := map[string]bool{
		"listSources":   false,
		"createSession": false,
		"getSession":    false,
		"listSessions":  false,
		"listActivities": false,
		"sendMessage":   false,
		"getSource":     false,
	}

	var sources []jules.Source
	var session *jules.Session
	var sessionID string

	sources, testResults["listSources"] = testListSources(client)
	session, testResults["createSession"] = testCreateSession(client, sources)
	if session != nil {
		sessionID = session.ID
	}

	_, testResults["getSession"] = testGetSession(client, sessionID)
	_, testResults["listSessions"] = testListSessions(client)

	fmt.Println("\n⏳ Waiting 5 seconds for activities to be generated...")
	time.Sleep(5)
	_, testResults["listActivities"] = testListActivities(client, sessionID)

	testResults["sendMessage"] = testSendMessage(client, sessionID)
	_, testResults["getSource"] = testGetSource(client, sources)

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("📊 TEST RESULTS SUMMARY")
	fmt.Println(strings.Repeat("=", 60))

	totalTests := len(testResults)
	passedTests := 0
	for _, passed := range testResults {
		if passed {
			passedTests++
		}
	}
	failedTests := totalTests - passedTests

	fmt.Printf("Total Tests: %d\n", totalTests)
	fmt.Printf("Passed: %d\n", passedTests)
	fmt.Printf("Failed: %d\n", failedTests)
	fmt.Println()

	for testName, passed := range testResults {
		status := "✅ PASS"
		if !passed {
			status = "❌ FAIL"
		}
		displayName := strings.ToUpper(testName[:1]) + strings.ReplaceAll(strings.ReplaceAll(testName[1:], "([A-Z])", " $1"), "([A-Z])", " $1")
		fmt.Printf("  %s: %s\n", displayName, status)
	}

	fmt.Println()
	if failedTests == 0 {
		fmt.Println("🎉 ALL TESTS PASSED! The Jules API is working correctly.")
	} else {
		fmt.Printf("⚠️  %d test(s) failed. Please check the API or network connection.\n", failedTests)
	}
}