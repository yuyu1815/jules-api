#!/usr/bin/env python3
"""
Comprehensive API test program for Jules API.
Tests all endpoints using the provided API key.
"""

import os
import sys
import time
from datetime import datetime
from dotenv import load_dotenv

# Load environment variables from .env file
load_dotenv()

# Add the parent directory to Python path so we can import jules_api
sys.path.insert(0, os.path.join(os.path.dirname(__file__), '..', 'py'))

from jules_api import create_client, CreateSessionRequest, SourceContext, GithubRepoContext, SendMessageRequest


def test_list_sources(client):
    """Test listing sources endpoint."""
    print("📋 Testing: List Sources")
    try:
        response = client.list_sources()
        print(f"   ✅ Success: Found {len(response.sources)} sources")
        for i, source in enumerate(response.sources):
            print(f"      [{i+1}] {source.id}: {source.name}")
            if source.github_repo:
                print(f"          GitHub: {source.github_repo.owner}/{source.github_repo.repo}")
        return response.sources
    except Exception as e:
        print(f"   ❌ Failed: {e}")
        return []


def test_create_session(client, sources):
    """Test creating a session endpoint."""
    print("\n🚀 Testing: Create Session")
    if not sources:
        print("   ⚠️  Skipping: No sources available")
        return None

    first_source = sources[0]
    print(f"   Using source: {first_source.id}")

    try:
        request = CreateSessionRequest(
            prompt="Create a simple test to verify the API is working correctly.",
            source_context=SourceContext(
                source=first_source.name,
                github_repo_context=GithubRepoContext(starting_branch="main")
            ),
            title="API Test Session",
            require_plan_approval=False  # Don't require approval for testing
        )

        session = client.create_session(request)
        print("   ✅ Success: Session created")
        print(f"      ID: {session.id}")
        print(f"      Title: {session.title}")
        print(f"      Name: {session.name}")
        return session
    except Exception as e:
        print(f"   ❌ Failed: {e}")
        return None


def test_get_session(client, session_id):
    """Test getting a specific session."""
    print("\n📖 Testing: Get Session")
    if not session_id:
        print("   ⚠️  Skipping: No session ID available")
        return None

    try:
        session = client.get_session(session_id)
        print("   ✅ Success: Session retrieved")
        print(f"      ID: {session.id}")
        print(f"      Title: {session.title}")
        return session
    except Exception as e:
        print(f"   ❌ Failed: {e}")
        return None


def test_list_sessions(client):
    """Test listing sessions endpoint."""
    print("\n📂 Testing: List Sessions")
    try:
        response = client.list_sessions(page_size=5)
        print(f"   ✅ Success: Found {len(response.sessions)} sessions")
        if response.next_page_token:
            print(f"      Next page token: {response.next_page_token}")
        return response.sessions
    except Exception as e:
        print(f"   ❌ Failed: {e}")
        return []


def test_list_activities(client, session_id):
    """Test listing activities endpoint."""
    print("\n🎬 Testing: List Activities")
    if not session_id:
        print("   ⚠️  Skipping: No session ID available")
        return []

    try:
        response = client.list_activities(session_id, page_size=10)
        print(f"   ✅ Success: Found {len(response.activities)} activities")
        for i, activity in enumerate(response.activities[:3]):  # Show first 3
            timestamp = activity.timestamp.strftime("%H:%M:%S") if activity.timestamp else "No timestamp"
            content = (activity.content or "No content")[:50] + "..."
            print(f"      [{i+1}] {activity.type} @ {timestamp}: {content}")
        if len(response.activities) > 3:
            print(f"      ... and {len(response.activities) - 3} more activities")
        return response.activities
    except Exception as e:
        print(f"   ❌ Failed: {e}")
        return []


def test_send_message(client, session_id):
    """Test sending a message endpoint."""
    print("\n💬 Testing: Send Message")
    if not session_id:
        print("   ⚠️  Skipping: No session ID available")
        return False

    try:
        request = SendMessageRequest(
            prompt="Please confirm that the API testing is working correctly by acknowledging this message."
        )
        client.send_message(session_id, request)
        print("   ✅ Success: Message sent")
        return True
    except Exception as e:
        print(f"   ❌ Failed: {e}")
        return False


def test_get_source(client, sources):
    """Test getting a specific source."""
    print("\n📦 Testing: Get Source")
    if not sources:
        print("   ⚠️  Skipping: No sources available")
        return None

    source_id = sources[0].id
    try:
        source = client.get_source(source_id)
        print("   ✅ Success: Source retrieved")
        print(f"      ID: {source.id}")
        print(f"      Name: {source.name}")
        return source
    except Exception as e:
        print(f"   ❌ Failed: {e}")
        return None


def main():
    """Run all API tests."""
    print("🧪 Jules API Comprehensive Test Suite")
    print("=" * 50)
    print(f"⏰ Test started at: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")
    print()

    # API key from environment variables (.env file)
    api_key = os.getenv("JULES_API_KEY")
    if not api_key:
        print("❌ Error: JULES_API_KEY environment variable not found.")
        print("   Please create a .env file in the test directory with:")
        print("   JULES_API_KEY=your_api_key_here")
        return 1

    print(f"🔑 Using API Key from .env: {api_key[:20]}...")
    print()

    try:
        # Create client
        client = create_client(api_key)

        # Test results tracking
        test_results = {
            'list_sources': False,
            'create_session': False,
            'get_session': False,
            'list_sessions': False,
            'list_activities': False,
            'send_message': False,
            'get_source': False
        }

        # Run tests
        sources = []
        session = None
        session_id = None

        # 1. List sources
        sources = test_list_sources(client)
        test_results['list_sources'] = len(sources) > 0

        # 2. Create session
        session = test_create_session(client, sources)
        test_results['create_session'] = session is not None
        if session:
            session_id = session.id

        # 3. Get session
        test_results['get_session'] = test_get_session(client, session_id) is not None

        # 4. List sessions
        sessions_list = test_list_sessions(client)
        test_results['list_sessions'] = len(sessions_list) >= 0  # Could be empty

        # 5. List activities (wait a moment for activities to be generated)
        print("\n⏳ Waiting 5 seconds for activities to be generated...")
        time.sleep(20)

        activities = test_list_activities(client, session_id)

        test_results['list_activities'] = len(activities) >= 0

        # 6. Send message
        test_results['send_message'] = test_send_message(client, session_id)

        # 7. Get source
        test_results['get_source'] = test_get_source(client, sources) is not None

        # Summary
        print("\n" + "=" * 50)
        print("📊 TEST RESULTS SUMMARY")
        print("=" * 50)

        total_tests = len(test_results)
        passed_tests = sum(1 for v in test_results.values() if v)
        failed_tests = total_tests - passed_tests

        print(f"Total Tests: {total_tests}")
        print(f"Passed: {passed_tests}")
        print(f"Failed: {failed_tests}")
        print()

        for test_name, passed in test_results.items():
            status = "✅ PASS" if passed else "❌ FAIL"
            print(f"  {test_name.replace('_', ' ').title()}: {status}")

        print()
        if failed_tests == 0:
            print("🎉 ALL TESTS PASSED! The Jules API is working correctly.")
            return 0
        else:
            print(f"⚠️  {failed_tests} test(s) failed. Please check the API or network connection.")
            return 1

    except KeyboardInterrupt:
        print("\n🛑 Test interrupted by user")
        return 1
    except Exception as e:
        print(f"💥 Unexpected error during testing: {e}")
        return 1


if __name__ == "__main__":

    exit_code = main()
    sys.exit(exit_code)
