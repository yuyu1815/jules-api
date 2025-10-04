"""
Example usage of the Jules API Python client.
"""

import os
import time
from jules_api import create_client, CreateSessionRequest, SourceContext, GithubRepoContext, SendMessageRequest


def main():
    """Main example function."""
    # Initialize the client with your API key from environment variable
    api_key = os.getenv("JULES_API_KEY")
    if not api_key:
        print("❌ Please set the JULES_API_KEY environment variable")
        print("   Example: export JULES_API_KEY=your_api_key_here")
        return

    client = create_client(api_key)

    print("🔍 Listing available sources...")
    sources_response = client.list_sources()

    print("Available sources:")
    for source in sources_response.sources:
        print(f"  - {source.id}: {source.name}")
        if source.github_repo:
            print(f"    GitHub: {source.github_repo.owner}/{source.github_repo.repo}")

    if not sources_response.sources:
        print("\n❌ No sources found. Please connect a GitHub repository in the Jules web app first.")
        return

    # Create a new session
    first_source = sources_response.sources[0]
    print(f"\n🚀 Creating a new session using source: {first_source.id}")

    session_request = CreateSessionRequest(
        prompt="Create a simple web app that displays 'Hello from Jules!'",
        source_context=SourceContext(
            source=first_source.name,
            github_repo_context=GithubRepoContext(starting_branch="main")
        ),
        title="Hello World App Session"
    )

    try:
        session = client.create_session(session_request)
        print("✅ Created session:")
        print(f"   ID: {session.id}")
        print(f"   Title: {session.title}")
        print(f"   Prompt: {session.prompt}")
    except Exception as e:
        print(f"❌ Failed to create session: {e}")
        return

    # Wait a moment for the agent to start working
    print("\n⏳ Waiting a moment for the agent to start working...")
    time.sleep(3)

    # List activities
    print("\n📋 Listing activities...")
    try:
        activities_response = client.list_activities(session.id, page_size=10)
        print(f"Found {len(activities_response.activities)} activities:")
        for activity in activities_response.activities:
            content = activity.content or "No content"
            if len(content) > 100:
                content = content[:100] + "..."
            print(f"  - {activity.type}: {content}")
    except Exception as e:
        print(f"⚠️  Could not list activities: {e}")

    # Send a follow-up message
    print("\n💬 Sending a follow-up message...")
    try:
        message_request = SendMessageRequest(
            prompt="Please add some styling to make it look more attractive."
        )
        client.send_message(session.id, message_request)
        print("✅ Message sent to the agent!")
        print("   The agent will respond in future activities.")
    except Exception as e:
        print(f"❌ Failed to send message: {e}")

    print("\n🎉 Example completed!")


if __name__ == "__main__":
    main()
