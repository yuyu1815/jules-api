// Types and interfaces for Jules API

export interface Source {
  name: string;
  id: string;
  githubRepo?: {
    owner: string;
    repo: string;
  };
}

export interface GithubRepoContext {
  startingBranch?: string;
}

export interface SourceContext {
  source: string;
  githubRepoContext?: GithubRepoContext;
}

export interface Session {
  name: string;
  id: string;
  title: string;
  sourceContext?: SourceContext;
  prompt?: string;
}

export interface ListSourcesResponse {
  sources: Source[];
  nextPageToken?: string;
}

export interface CreateSessionRequest {
  prompt: string;
  sourceContext: SourceContext;
  title: string;
  requirePlanApproval?: boolean;
}

export interface ListSessionsResponse {
  sessions: Session[];
  nextPageToken?: string;
}

export interface Activity {
  name: string;
  id: string;
  type: string;
  content?: string;
  timestamp?: string;
}

export interface ListActivitiesResponse {
  activities: Activity[];
  nextPageToken?: string;
}

export interface SendMessageRequest {
  prompt: string;
}

export interface JulesClientOptions {
  apiKey: string;
  baseUrl?: string;
}
