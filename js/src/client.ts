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

  constructor(options: JulesClientOptions = {}) {
    let { apiKey, baseUrl = 'https://jules.googleapis.com/v1alpha', timeout = 60000 } = options;

    if (!apiKey) {
      apiKey = process.env.JULES_API_KEY;
    }

    if (!apiKey) {
      throw new Error('API key must be provided either in options or as a JULES_API_KEY environment variable.');
    }

    this.httpClient = axios.create({
      baseURL: baseUrl,
      timeout: timeout,
      headers: {
        'X-Goog-Api-Key': apiKey,
        'Content-Type': 'application/json',
      },
    });
  }

  private async _request<T>(requestFn: () => Promise<{ data: T }>): Promise<T> {
    try {
      const { data } = await requestFn();
      return data;
    } catch (error) {
      if (axios.isAxiosError(error)) {
        // Prefer the API's error message, but fall back to the default Axios message
        const message = (error.response?.data as any)?.error?.message || error.message;
        const status = error.response?.status ? ` (status: ${error.response.status})` : '';
        throw new Error(`Jules API request failed${status}: ${message}`);
      }
      // Re-throw non-Axios errors
      throw error;
    }
  }

  /**
   * List all available sources
   * @param nextPageToken - Token for pagination
   * @returns Promise<ListSourcesResponse>
   */
  async listSources(nextPageToken?: string): Promise<ListSourcesResponse> {
    const params = nextPageToken ? { nextPageToken } : {};
    return this._request<ListSourcesResponse>(() => this.httpClient.get('/sources', { params }));
  }

  /**
   * Create a new session
   * @param request - Session creation parameters
   * @returns Promise<Session>
   */
  async createSession(request: CreateSessionRequest): Promise<Session> {
    return this._request<Session>(() => this.httpClient.post('/sessions', request));
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

    return this._request<ListSessionsResponse>(() => this.httpClient.get('/sessions', { params }));
  }

  /**
   * Approve the latest plan for a session
   * @param sessionId - The session ID
   * @returns Promise<void>
   */
  async approvePlan(sessionId: string): Promise<void> {
    await this._request<void>(() => this.httpClient.post(`/sessions/${sessionId}:approvePlan`));
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

    return this._request<ListActivitiesResponse>(() => this.httpClient.get(`/sessions/${sessionId}/activities`, { params }));
  }

  /**
   * Send a message to the agent
   * @param sessionId - The session ID
   * @param request - Message parameters
   * @returns Promise<void>
   */
  async sendMessage(sessionId: string, request: SendMessageRequest): Promise<void> {
    await this._request<void>(() => this.httpClient.post(`/sessions/${sessionId}:sendMessage`, request));
  }

  /**
   * Get a specific session
   * @param sessionId - The session ID
   * @returns Promise<Session>
   */
  async getSession(sessionId: string): Promise<Session> {
    return this._request<Session>(() => this.httpClient.get(`/sessions/${sessionId}`));
  }

  /**
   * Get a specific source
   * @param sourceId - The source ID
   * @returns Promise<Source>
   */
  async getSource(sourceId: string): Promise<Source> {
    return this._request<Source>(() => this.httpClient.get(`/sources/${sourceId}`));
  }
}

export default JulesClient;
