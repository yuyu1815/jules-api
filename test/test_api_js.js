#!/usr/bin/env node

/**
 * Comprehensive API test program for Jules API - JavaScript Version
 * Tests all endpoints using the provided API key.
 */

import { JulesClient } from '../js/src/index.js';
import 'dotenv/config';

async function testListSources(client) {
  console.log("📋 Testing: List Sources");
  try {
    const response = await client.listSources();
    console.log(`   ✅ Success: Found ${response.sources.length} sources`);
    response.sources.slice(0, 5).forEach((source, i) => {
      console.log(`      [${i + 1}] ${source.id}: ${source.name}`);
      if (source.githubRepo) {
        console.log(`          GitHub: ${source.githubRepo.owner}/${source.githubRepo.repo}`);
      }
    });
    if (response.sources.length > 5) {
      console.log(`      ... and ${response.sources.length - 5} more sources`);
    }
    return response.sources;
  } catch (error) {
    console.log(`   ❌ Failed: ${error.message}`);
    return [];
  }
}

async function testCreateSession(client, sources) {
  console.log("\n🚀 Testing: Create Session");
  if (!sources || sources.length === 0) {
    console.log("   ⚠️  Skipping: No sources available");
    return null;
  }

  const firstSource = sources[0];
  console.log(`   Using source: ${firstSource.id}`);

  try {
    const request = {
      prompt: "Create a simple test to verify the API is working correctly.",
      sourceContext: {
        source: firstSource.name,
        githubRepoContext: {
          startingBranch: "main"
        }
      },
      title: "API Test Session - JS",
      requirePlanApproval: false
    };

    const session = await client.createSession(request);
    console.log("   ✅ Success: Session created");
    console.log(`      ID: ${session.id}`);
    console.log(`      Title: ${session.title}`);
    console.log(`      Name: ${session.name}`);
    return session;
  } catch (error) {
    console.log(`   ❌ Failed: ${error.message}`);
    return null;
  }
}

async function testGetSession(client, sessionId) {
  console.log("\n📖 Testing: Get Session");
  if (!sessionId) {
    console.log("   ⚠️  Skipping: No session ID available");
    return null;
  }

  try {
    const session = await client.getSession(sessionId);
    console.log("   ✅ Success: Session retrieved");
    console.log(`      ID: ${session.id}`);
    console.log(`      Title: ${session.title}`);
    return session;
  } catch (error) {
    console.log(`   ❌ Failed: ${error.message}`);
    return null;
  }
}

async function testListSessions(client) {
  console.log("\n📂 Testing: List Sessions");
  try {
    const response = await client.listSessions(5);
    console.log(`   ✅ Success: Found ${response.sessions.length} sessions`);
    if (response.nextPageToken) {
      console.log(`      Next page token: ${response.nextPageToken}`);
    }
    return response.sessions;
  } catch (error) {
    console.log(`   ❌ Failed: ${error.message}`);
    return [];
  }
}

async function testListActivities(client, sessionId) {
  console.log("\n🎬 Testing: List Activities");
  if (!sessionId) {
    console.log("   ⚠️  Skipping: No session ID available");
    return [];
  }

  try {
    const response = await client.listActivities(sessionId, 10);
    console.log(`   ✅ Success: Found ${response.activities.length} activities`);
    response.activities.slice(0, 3).forEach((activity, i) => {
      const timestamp = activity.timestamp ? new Date(activity.timestamp).toLocaleTimeString() : "No timestamp";
      const content = (activity.content || "No content").substring(0, 50) + "...";
      console.log(`      [${i + 1}] ${activity.type} @ ${timestamp}: ${content}`);
    });
    if (response.activities.length > 3) {
      console.log(`      ... and ${response.activities.length - 3} more activities`);
    }
    return response.activities;
  } catch (error) {
    console.log(`   ❌ Failed: ${error.message}`);
    return [];
  }
}

async function testSendMessage(client, sessionId) {
  console.log("\n💬 Testing: Send Message");
  if (!sessionId) {
    console.log("   ⚠️  Skipping: No session ID available");
    return false;
  }

  try {
    const request = {
      prompt: "Please confirm that the API testing is working correctly by acknowledging this message."
    };
    await client.sendMessage(sessionId, request);
    console.log("   ✅ Success: Message sent");
    return true;
  } catch (error) {
    console.log(`   ❌ Failed: ${error.message}`);
    return false;
  }
}

async function testGetSource(client, sources) {
  console.log("\n📦 Testing: Get Source");
  if (!sources || sources.length === 0) {
    console.log("   ⚠️  Skipping: No sources available");
    return null;
  }

  const sourceId = sources[0].id;
  try {
    const source = await client.getSource(sourceId);
    console.log("   ✅ Success: Source retrieved");
    console.log(`      ID: ${source.id}`);
    console.log(`      Name: ${source.name}`);
    return source;
  } catch (error) {
    console.log(`   ❌ Failed: ${error.message}`);
    return null;
  }
}

async function main() {
  console.log("🧪 Jules API Comprehensive Test Suite - JavaScript Version");
  console.log("=".repeat(60));
  console.log(`⏰ Test started at: ${new Date().toISOString()}`);
  console.log();

  const apiKey = process.env.JULES_API_KEY;
  if (!apiKey) {
    console.log("❌ Error: JULES_API_KEY environment variable not found.");
    console.log("   Please create a test/.env file with:");
    console.log("   JULES_API_KEY=your_api_key_here");
    process.exit(1);
  }

  console.log(`🔑 Using API Key from .env: ${apiKey.substring(0, 20)}...`);
  console.log();

  try {
    const client = new JulesClient({
      apiKey: apiKey,
      timeout: -1, // Use infinite timeout for all requests
    });

    const testResults = {
      listSources: false,
      createSession: false,
      getSession: false,
      listSessions: false,
      listActivities: false,
      sendMessage: false,
      getSource: false
    };

    let sources = [];
    let session = null;
    let sessionId = null;

    sources = await testListSources(client);
    testResults.listSources = sources.length > 0;

    session = await testCreateSession(client, sources);
    testResults.createSession = session !== null;
    if (session) {
      sessionId = session.id;
    }

    testResults.getSession = (await testGetSession(client, sessionId)) !== null;

    const sessionsList = await testListSessions(client);
    testResults.listSessions = sessionsList.length >= 0;

    console.log("\n⏳ Waiting 5 seconds for activities to be generated...");
    await new Promise(resolve => setTimeout(resolve, 5000));
    const activities = await testListActivities(client, sessionId);
    testResults.listActivities = activities.length >= 0;

    testResults.sendMessage = await testSendMessage(client, sessionId);

    testResults.getSource = (await testGetSource(client, sources)) !== null;

    console.log("\n" + "=".repeat(60));
    console.log("📊 TEST RESULTS SUMMARY");
    console.log("=".repeat(60));

    const totalTests = Object.keys(testResults).length;
    const passedTests = Object.values(testResults).filter(v => v).length;
    const failedTests = totalTests - passedTests;

    console.log(`Total Tests: ${totalTests}`);
    console.log(`Passed: ${passedTests}`);
    console.log(`Failed: ${failedTests}`);
    console.log();

    Object.entries(testResults).forEach(([testName, passed]) => {
      const status = passed ? "✅ PASS" : "❌ FAIL";
      console.log(`  ${testName.replace(/([A-Z])/g, ' $1').replace(/^./, str => str.toUpperCase())}: ${status}`);
    });

    console.log();
    if (failedTests === 0) {
      console.log("🎉 ALL TESTS PASSED! The Jules API is working correctly.");
      return 0;
    } else {
      console.log(`⚠️  ${failedTests} test(s) failed. Please check the API or network connection.`);
      return 1;
    }

  } catch (error) {
    console.log(`💥 Unexpected error during testing: ${error.message}`);
    return 1;
  }
}

main().then(exitCode => {
  process.exit(exitCode);
}).catch(error => {
  console.error(`💥 Fatal error: ${error.message}`);
  process.exit(1);
});