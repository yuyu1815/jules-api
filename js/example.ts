import { JulesClient } from './src';

async function example() {
  // Initialize the client with your API key
  const client = new JulesClient({
    apiKey: process.env.JULES_API_KEY || 'YOUR_API_KEY_HERE'
  });

  try {
    console.log('🔍 Listing available sources...');
    const sourcesResponse = await client.listSources();
    console.log('Available sources:');
    sourcesResponse.sources.forEach(source => {
      console.log(`- ${source.id}: ${source.name}`);
      if (source.githubRepo) {
        console.log(`  GitHub: ${source.githubRepo.owner}/${source.githubRepo.repo}`);
      }
    });

    if (sourcesResponse.sources.length === 0) {
      console.log('No sources found. Please connect a GitHub repository in the Jules web app first.');
      return;
    }

    console.log('\n🚀 Creating a new session...');
    const firstSource = sourcesResponse.sources[0];
    const session = await client.createSession({
      prompt: 'Create a simple web app that displays "Hello from Jules!"',
      sourceContext: {
        source: firstSource.name,
        githubRepoContext: {
          startingBranch: 'main'
        }
      },
      title: 'Hello World App Session'
    });
    console.log('✅ Created session:', session.id);
    console.log('📝 Title:', session.title);
    console.log('🎯 Prompt:', session.prompt);

    console.log('\n⏳ Waiting a moment for the agent to start working...');
    await new Promise(resolve => setTimeout(resolve, 3000));

    console.log('\n📋 Listing activities...');
    const activitiesResponse = await client.listActivities(session.id, 10);
    console.log(`Found ${activitiesResponse.activities.length} activities:`);
    activitiesResponse.activities.forEach(activity => {
      console.log(`- ${activity.type}: ${activity.content?.substring(0, 100) || 'No content'}...`);
    });

    console.log('\n💬 Sending a follow-up message...');
    await client.sendMessage(session.id, {
      prompt: 'Please add some styling to make it look more attractive.'
    });
    console.log('✅ Message sent. The agent will respond in future activities.');

  } catch (error) {
    console.error('❌ Error:', error.message);
    if (error.response) {
      console.error('Response status:', error.response.status);
      console.error('Response data:', error.response.data);
    }
  }
}

// Run the example if this file is executed directly
if (require.main === module) {
  example().catch(console.error);
}

export { example };
