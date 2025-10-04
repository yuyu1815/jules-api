import axios, { AxiosInstance } from 'axios';
import {
  JulesClientOptions,
  Source,
  ListSourcesResponse,
  Session,
  CreateSessionRequest,
  ListSessionsResponse,
  ListActivitiesResponse,
  SendMessageRequest,
  Activity
} from './types';

export class JulesClient {
  private httpClient: AxiosInstance;
  private apiKey: string;

  constructor(options: JulesClientOptions) {
    const { apiKey, baseUrl = 'https://jules.googleapis.com/v1alpha' } = options;

    this.apiKey = apiKey;
    this.httpClient = axios.create({
      baseURL: baseUrl,
      headers: {
        'X-Goog-Api-Key': apiKey,
        'Content-Type': 'application/json',
      },
    });
  }

  /**
   * List all available sources
   * @param nextPageToken - Token for pagination
   * @returns Promise<ListSourcesResponse>
   */
  async listSources(nextPageToken?: string): Promise<ListSourcesResponse> {
    const params = nextPageToken ? { nextPageToken } : {};
    const response = await this.httpClient.get('/sources', { params });
    return response.data;
  }

  /**
   * Create a new session
   * @param request - Session creation parameters
   * @returns Promise<Session>
   */
  async createSession(request: CreateSessionRequest): Promise<Session> {
    const response = await this.httpClient.post('/sessions', request);
    return response.data;
  }

  /**
   * List sessions
   * @param pageSize - Maximum number of sessions to return
   * @param nextPageToken - Token for pagination
   * @returns Promise<ListSessionsResponse>
   */
  async listSessions(pageSize?: number, nextPageToken?: string): Promise<ListSessionsResponse> {
    const params: any = {};
    if (pageSize) params.pageSize = pageSize;
    if (nextPageToken) params.nextPageToken = nextPageToken;

    const response = await this.httpClient.get('/sessions', { params });
    return response.data;
  }

  /**
   * Approve the latest plan for a session
   * @param sessionId - The session ID
   * @returns Promise<void>
   */
  async approvePlan(sessionId: string): Promise<void> {
    await this.httpClient.post(`/sessions/${sessionId}:approvePlan`);
  }

  /**
   * List activities for a session
   * @param sessionId - The session ID
   * @param pageSize - Maximum number of activities to return
   * @param nextPageToken - Token for pagination
   * @returns Promise<ListActivitiesResponse>
   */
  async listActivities(
    sessionId: string,
    pageSize?: number,
    nextPageToken?: string
  ): Promise<ListActivitiesResponse> {
    const params: any = {};
    if (pageSize) params.pageSize = pageSize;
    if (nextPageToken) params.nextPageToken = nextPageToken;

    const response = await this.httpClient.get(`/sessions/${sessionId}/activities`, { params });
    return response.data;
  }

  /**
   * Send a message to the agent
   * @param sessionId - The session ID
   * @param request - Message parameters
   * @returns Promise<void>
   */
  async sendMessage(sessionId: string, request: SendMessageRequest): Promise<void> {
    await this.httpClient.post(`/sessions/${sessionId}:sendMessage`, request);
  }

  /**
   * Get a specific session
   * @param sessionId - The session ID
   * @returns Promise<Session>
   */
  async getSession(sessionId: string): Promise<Session> {
    const response = await this.httpClient.get(`/sessions/${sessionId}`);
    return response.data;
  }

  /**
   * Get a specific source
   * @param sourceId - The source ID
   * @returns Promise<Source>
   */
  async getSource(sourceId: string): Promise<Source> {
    const response = await this.httpClient.get(`/sources/${sourceId}`);
    return response.data;
  }
}

export default JulesClient;
