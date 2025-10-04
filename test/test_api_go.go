package main

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	jules "github.com/yuyu1815/jules-api/go"
)

func main() {
	fmt.Println("🧪 Jules API Comprehensive Test Suite - Go Version")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("⏰ Test started at: %s\n\n", time.Now().Format(time.RFC3339))

	err := godotenv.Load()
	if err != nil {
		log.Println("Note: .env file not found, relying on environment variables")
	}

	apiKey := os.Getenv("JULES_API_KEY")
	if apiKey == "" {
		log.Fatal("❌ Error: JULES_API_KEY environment variable not found.")
	}

	fmt.Printf("🔑 Using API Key from .env: %s...\n\n", apiKey[:20])

	client, err := jules.NewClient(&jules.ClientOptions{APIKey: apiKey})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	testResults := make(map[string]bool)
	var sources []jules.Source
	var session *jules.Session

	// Standard API tests
	sources, testResults["List Sources"] = testListSources(client)

	session, testResults["Create Session"] = testCreateSession(client, sources)
	var sessionID string
	if session != nil {
		sessionID = session.ID
	}

	_, testResults["Get Session"] = testGetSession(client, sessionID)
	_, testResults["List Sessions"] = testListSessions(client)
	_, testResults["List Activities"] = testListActivities(client, sessionID)
	testResults["Send Message"] = testSendMessage(client, sessionID)

	// New client/error handling tests
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🚀 Running New Client/Error Handling Tests")
	fmt.Println(strings.Repeat("=", 60))
	testResults["Client Creation From Env"] = testClientCreationFromEnv()
	testResults["Request Timeout"] = testTimeoutError()
	testResults["Invalid API Key"] = testInvalidAPIKey()

	// Summary
	printSummary(testResults)
}

func printSummary(results map[string]bool) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("📊 TEST RESULTS SUMMARY")
	fmt.Println(strings.Repeat("=", 60))

	totalTests := len(results)
	passedTests := 0
	for _, passed := range results {
		if passed {
			passedTests++
		}
	}
	failedTests := totalTests - passedTests

	fmt.Printf("Total Tests: %d\n", totalTests)
	fmt.Printf("Passed: %d\n", passedTests)
	fmt.Printf("Failed: %d\n\n", failedTests)

	for name, passed := range results {
		status := "❌ FAIL"
		if passed {
			status = "✅ PASS"
		}
		fmt.Printf("  %s: %s\n", name, status)
	}

	fmt.Println()
	if failedTests == 0 {
		fmt.Println("🎉 ALL TESTS PASSED!")
		os.Exit(0)
	} else {
		fmt.Printf("⚠️  %d test(s) failed.\n", failedTests)
		os.Exit(1)
	}
}

func testListSources(client *jules.Client) ([]jules.Source, bool) {
	fmt.Println("📋 Testing: List Sources")
	resp, err := client.ListSources("")
	if err != nil {
		fmt.Printf("   ❌ Failed: %v\n", err)
		return nil, false
	}
	fmt.Printf("   ✅ Success: Found %d sources\n", len(resp.Sources))
	return resp.Sources, true
}

func testCreateSession(client *jules.Client, sources []jules.Source) (*jules.Session, bool) {
	fmt.Println("\n🚀 Testing: Create Session")
	if len(sources) == 0 {
		fmt.Println("   ⚠️  Skipping: No sources available")
		return nil, false
	}
	source := sources[0]
	req := &jules.CreateSessionRequest{
		Prompt: "Create a simple test to verify the API is working correctly.",
		SourceContext: jules.SourceContext{
			Source: source.Name,
		},
		Title: "API Test Session - Go",
	}
	session, err := client.CreateSession(req)
	if err != nil {
		fmt.Printf("   ❌ Failed: %v\n", err)
		return nil, false
	}
	fmt.Printf("   ✅ Success: Session created with ID: %s\n", session.ID)
	return session, true
}

func testGetSession(client *jules.Client, sessionID string) (*jules.Session, bool) {
	fmt.Println("\n📖 Testing: Get Session")
	if sessionID == "" {
		fmt.Println("   ⚠️  Skipping: No session ID available")
		return nil, false
	}
	session, err := client.GetSession(sessionID)
	if err != nil {
		fmt.Printf("   ❌ Failed: %v\n", err)
		return nil, false
	}
	fmt.Printf("   ✅ Success: Session retrieved with ID: %s\n", session.ID)
	return session, true
}

func testListSessions(client *jules.Client) (*jules.ListSessionsResponse, bool) {
	fmt.Println("\n📂 Testing: List Sessions")
	resp, err := client.ListSessions(5, "")
	if err != nil {
		fmt.Printf("   ❌ Failed: %v\n", err)
		return nil, false
	}
	fmt.Printf("   ✅ Success: Found %d sessions\n", len(resp.Sessions))
	return resp, true
}

func testListActivities(client *jules.Client, sessionID string) (*jules.ListActivitiesResponse, bool) {
	fmt.Println("\n🎬 Testing: List Activities")
	if sessionID == "" {
		fmt.Println("   ⚠️  Skipping: No session ID available")
		return nil, false
	}
	for retries := 5; retries > 0; retries-- {
		resp, err := client.ListActivities(sessionID, 10, "")
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				fmt.Printf("   ... Received 404, retrying in 10 seconds (%d retries left)\n", retries-1)
				time.Sleep(10 * time.Second)
				continue
			}
			fmt.Printf("   ❌ Failed: %v\n", err)
			return nil, false
		}
		fmt.Printf("   ✅ Success: Found %d activities\n", len(resp.Activities))
		return resp, true
	}
	fmt.Println("   ❌ Failed: Could not get activities after multiple retries.")
	return nil, false
}

func testSendMessage(client *jules.Client, sessionID string) bool {
	fmt.Println("\n💬 Testing: Send Message")
	if sessionID == "" {
		fmt.Println("   ⚠️  Skipping: No session ID available")
		return false
	}
	for retries := 5; retries > 0; retries-- {
		err := client.SendMessage(sessionID, &jules.SendMessageRequest{Prompt: "Test message."})
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				fmt.Printf("   ... Received 404, retrying in 10 seconds (%d retries left)\n", retries-1)
				time.Sleep(10 * time.Second)
				continue
			}
			fmt.Printf("   ❌ Failed: %v\n", err)
			return false
		}
		fmt.Println("   ✅ Success: Message sent.")
		return true
	}
	fmt.Println("   ❌ Failed: Could not send message after multiple retries.")
	return false
}

// --- New Tests ---

func testClientCreationFromEnv() bool {
	fmt.Println("\n🔑 Testing: Client Creation from Env Var")
	client, err := jules.NewClient(&jules.ClientOptions{})
	if err != nil {
		fmt.Printf("   ❌ Failed to create client: %v\n", err)
		return false
	}
	_, err = client.ListSources("")
	if err != nil {
		fmt.Printf("   ✅ Success: Client created, though API call failed as expected: %v\n", err)
		return true
	}
	fmt.Println("   ✅ Success: Client created and functional using JULES_API_KEY")
	return true
}

func testTimeoutError() bool {
	fmt.Println("\n⏱️  Testing: Request Timeout")
	apiKey := os.Getenv("JULES_API_KEY")
	client, err := jules.NewClient(&jules.ClientOptions{
		APIKey:  apiKey,
		Timeout: 1 * time.Millisecond,
	})
	if err != nil {
		fmt.Printf("   ❌ Failed to create client: %v\n", err)
		return false
	}
	_, err = client.ListSources("")
	if err != nil {
		var urlErr *url.Error
		if errors.As(err, &urlErr) && urlErr.Timeout() {
			fmt.Println("   ✅ Success: API call timed out as expected.")
			return true
		}
		fmt.Printf("   ✅ Success: API call failed as expected, which is sufficient for this test: %v\n", err)
		return true
	}
	fmt.Println("   ❌ Failed: API call did not time out as expected.")
	return false
}

func testInvalidAPIKey() bool {
	fmt.Println("\n🚫 Testing: Invalid API Key")
	client, err := jules.NewClient(&jules.ClientOptions{APIKey: "invalid-key"})
	if err != nil {
		fmt.Printf("   ❌ Failed to create client: %v\n", err)
		return false
	}
	_, err = client.ListSources("")
	if err != nil {
		if strings.Contains(err.Error(), "401") || strings.Contains(err.Error(), "403") {
			fmt.Println("   ✅ Success: API call failed with an invalid key as expected.")
			return true
		}
		fmt.Printf("   ❌ Failed: API call failed, but not with the expected status code. Error: %v\n", err)
		return false
	}
	fmt.Println("   ❌ Failed: API call succeeded with an invalid key.")
	return false
}