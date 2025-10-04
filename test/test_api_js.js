#!/usr/bin/env node

/**
 * Comprehensive API test program for Jules API - JavaScript Version
 * Tests all endpoints using the provided API key.
 */

import { JulesClient } from '../js/dist/index.js'; // Corrected import path
import dotenv from 'dotenv';

// Load .env file from the current directory
dotenv.config();

// Helper to wait for a specific duration
const sleep = (ms) => new Promise(resolve => setTimeout(resolve, ms));

// Helper to run a test and capture its result
async function runTest(testName, testFn, testResults) {
    try {
        const result = await testFn();
        testResults[testName] = result;
    } catch (error) {
        console.log(`   💥 Unexpected error in test '${testName}': ${error.message}`);
        testResults[testName] = false;
    }
}


async function testListSources(client) {
  console.log("📋 Testing: List Sources");
  try {
    const response = await client.listSources();
    console.log(`   ✅ Success: Found ${response.sources.length} sources`);
    response.sources.slice(0, 5).forEach((source, i) => {
      console.log(`      [${i + 1}] ${source.id}: ${source.name}`);
    });
    return response.sources;
  } catch (error) {
    console.log(`   ❌ Failed: ${error.message}`);
    return null;
  }
}

async function testCreateSession(client, sources) {
  console.log("\n🚀 Testing: Create Session");
  if (!sources || sources.length === 0) {
    console.log("   ⚠️  Skipping: No sources available");
    return null;
  }
  const firstSource = sources[0];
  try {
    const session = await client.createSession({
      prompt: "Create a simple test to verify the API is working correctly.",
      sourceContext: { source: firstSource.name },
      title: "API Test Session - JS",
    });
    console.log(`   ✅ Success: Session created with ID: ${session.id}`);
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
    console.log(`   ✅ Success: Session retrieved with ID: ${session.id}`);
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
    return response.sessions;
  } catch (error) {
    console.log(`   ❌ Failed: ${error.message}`);
    return null;
  }
}

async function testListActivities(client, sessionId) {
    console.log("\n🎬 Testing: List Activities");
    if (!sessionId) {
        console.log("   ⚠️  Skipping: No session ID available");
        return null;
    }

    let retries = 5; // Max 5 retries
    while (retries > 0) {
        try {
            const response = await client.listActivities(sessionId, 10);
            console.log(`   ✅ Success: Found ${response.activities.length} activities`);
            return response.activities;
        } catch (error) {
            if (error.message.includes("404")) {
                retries--;
                console.log(`   ... Received 404, retrying in 10 seconds (${retries} retries left)`);
                await sleep(10000);
            } else {
                console.log(`   ❌ Failed: ${error.message}`);
                return null;
            }
        }
    }
    console.log("   ❌ Failed: Could not get activities after multiple retries.");
    return null;
}


async function testSendMessage(client, sessionId) {
    console.log("\n💬 Testing: Send Message");
    if (!sessionId) {
        console.log("   ⚠️  Skipping: No session ID available");
        return false;
    }
    let retries = 5;
    while(retries > 0) {
        try {
            await client.sendMessage(sessionId, { prompt: "Test message." });
            console.log("   ✅ Success: Message sent.");
            return true;
        } catch (error) {
            if (error.message.includes("404")) {
                retries--;
                console.log(`   ... Received 404, retrying in 10 seconds (${retries} retries left)`);
                await sleep(10000);
            } else {
                 console.log(`   ❌ Failed: ${error.message}`);
                 return false;
            }
        }
    }
    console.log("   ❌ Failed: Could not send message after multiple retries.");
    return false;
}


// --- New Tests ---

async function testClientCreationFromEnv() {
    console.log("\n🔑 Testing: Client Creation from Env Var");
    try {
        const client = new JulesClient(); // No options
        await client.listSources();
        console.log(`   ✅ Success: Client created and attempted API call using JULES_API_KEY.`);
        return true;
    } catch (e) {
        if (e.message.includes("API key must be provided")) {
            console.log(`   ❌ Failed: ${e.message}`);
            return false;
        }
        console.log(`   ✅ Success: Client created, though API call failed as expected: ${e.message}`);
        return true;
    }
}

async function testTimeoutError() {
    console.log("\n⏱️  Testing: Request Timeout");
    try {
        const client = new JulesClient({ timeout: 1 }); // 1ms timeout
        await client.listSources();
        console.log("   ❌ Failed: API call did not time out as expected.");
        return false;
    } catch (e) {
        if (e.message.toLowerCase().includes('timeout')) {
            console.log("   ✅ Success: API call timed out as expected.");
            return true;
        } else {
            console.log(`   ❌ Failed: Received an error, but it was not a timeout error: ${e.message}`);
            return false;
        }
    }
}

async function testInvalidApiKey() {
    console.log("\n🚫 Testing: Invalid API Key");
    try {
        const client = new JulesClient({ apiKey: "invalid-key-for-testing" });
        await client.listSources();
        console.log("   ❌ Failed: API call succeeded with an invalid key.");
        return false;
    } catch (e) {
        if (e.message.includes("400") || e.message.includes("401") || e.message.includes("403")) {
            console.log(`   ✅ Success: API call failed with an invalid key as expected.`);
            return true;
        } else {
            console.log(`   ❌ Failed: API call failed, but not with the expected status code. Error: ${e.message}`);
            return false;
        }
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
        return 1;
    }

  console.log(`🔑 Using API Key from .env: ${apiKey.substring(0, 20)}...`);
  console.log();

  const client = new JulesClient({ apiKey });

  const testResults = {};
  let sources = [];
  let sessionId = null;

  // Standard API tests
  sources = await testListSources(client);
  testResults['List Sources'] = sources !== null && sources.length > 0;

  const session = await testCreateSession(client, sources);
  testResults['Create Session'] = session !== null;
  if(session) sessionId = session.id;

  testResults['Get Session'] = (await testGetSession(client, sessionId)) !== null;

  const sessionsList = await testListSessions(client);
  testResults['List Sessions'] = sessionsList !== null;

  const activities = await testListActivities(client, sessionId);
  testResults['List Activities'] = activities !== null;

  testResults['Send Message'] = await testSendMessage(client, sessionId);


  // New client/error handling tests
  console.log("\n" + "=".repeat(60));
  console.log("🚀 Running New Client/Error Handling Tests");
  console.log("=".repeat(60));
  testResults['Client Creation From Env'] = await testClientCreationFromEnv();
  testResults['Request Timeout'] = await testTimeoutError();
  testResults['Invalid API Key'] = await testInvalidApiKey();


  // Summary
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
    console.log(`  ${testName}: ${status}`);
  });

  console.log();
  if (failedTests === 0) {
    console.log("🎉 ALL TESTS PASSED!");
    return 0;
  } else {
    console.log(`⚠️  ${failedTests} test(s) failed.`);
    return 1;
  }
}

// Run the test
main().then(exitCode => {
  process.exit(exitCode);
}).catch(error => {
  console.error(`💥 Fatal error: ${error.message}`);
  process.exit(1);
});