import axios, { AxiosInstance, AxiosRequestConfig, Method } from 'axios';
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

const sleep = (ms: number) => new Promise(resolve => setTimeout(resolve, ms));

export class JulesClient {
  private httpClient: AxiosInstance;
  private apiKey: string;
  private timeout: number;

  constructor(options: JulesClientOptions = {}) {
    const apiKey = options.apiKey || process.env.JULES_API_KEY;
    if (!apiKey) {
      throw new Error('API key must be provided or set as JULES_API_KEY environment variable');
    }

    this.apiKey = apiKey;
    this.timeout = options.timeout ?? 60000; // 60 seconds default

    const baseUrl = options.baseUrl || 'https://jules.googleapis.com/v1alpha';
    this.httpClient = axios.create({
      baseURL: baseUrl,
      headers: {
        'X-Goog-Api-Key': this.apiKey,
        'Content-Type': 'application/json',
      },
    });
  }

  private async _makeRequest<T>(
    method: Method,
    url: string,
    params: any = {},
    data: any = null,
    timeout?: number
  ): Promise<T> {
    const requestConfig: AxiosRequestConfig = {
      method,
      url,
      params,
      data,
      timeout: timeout ?? this.timeout,
    };

    if (requestConfig.timeout === -1) {
      // Infinite retry logic
      while (true) {
        try {
          // Use a default timeout for the individual request attempt
          const response = await this.httpClient.request({ ...requestConfig, timeout: 60000 });
          return response.data;
        } catch (error) {
          if (axios.isAxiosError(error) && error.response?.status === 404) {
            console.log(`Resource not found (404), retrying in 1 second...`);
          } else {
            console.log(`Request failed with ${error}, retrying in 1 second...`);
          }
          await sleep(1000);
        }
      }
    } else {
      // Normal request with specified timeout
      const response = await this.httpClient.request(requestConfig);
      return response.data;
    }
  }

  /**
   * List all available sources
   * @param nextPageToken - Token for pagination
   * @param timeout - Optional request timeout
   * @returns Promise<ListSourcesResponse>
   */
  async listSources(nextPageToken?: string, timeout?: number): Promise<ListSourcesResponse> {
    const params = nextPageToken ? { nextPageToken } : {};
    return this._makeRequest('GET', '/sources', params, null, timeout);
  }

  /**
   * Create a new session
   * @param request - Session creation parameters
   * @param timeout - Optional request timeout
   * @returns Promise<Session>
   */
  async createSession(request: CreateSessionRequest, timeout?: number): Promise<Session> {
    return this._makeRequest('POST', '/sessions', {}, request, timeout);
  }

  /**
   * List sessions
   * @param pageSize - Maximum number of sessions to return
   * @param nextPageToken - Token for pagination
   * @param timeout - Optional request timeout
   * @returns Promise<ListSessionsResponse>
   */
  async listSessions(pageSize?: number, nextPageToken?: string, timeout?: number): Promise<ListSessionsResponse> {
    const params: any = {};
    if (pageSize) params.pageSize = pageSize;
    if (nextPageToken) params.nextPageToken = nextPageToken;

    return this._makeRequest('GET', '/sessions', params, null, timeout);
  }

  /**
   * Approve the latest plan for a session
   * @param sessionId - The session ID
   * @param timeout - Optional request timeout
   * @returns Promise<void>
   */
  async approvePlan(sessionId: string, timeout?: number): Promise<void> {
    return this._makeRequest('POST', `/sessions/${sessionId}:approvePlan`, {}, null, timeout);
  }

  /**
   * List activities for a session
   * @param sessionId - The session ID
   * @param pageSize - Maximum number of activities to return
   * @param nextPageToken - Token for pagination
   * @param timeout - Optional request timeout
   * @returns Promise<ListActivitiesResponse>
   */
  async listActivities(
    sessionId: string,
    pageSize?: number,
    nextPageToken?: string,
    timeout?: number
  ): Promise<ListActivitiesResponse> {
    const params: any = {};
    if (pageSize) params.pageSize = pageSize;
    if (nextPageToken) params.nextPageToken = nextPageToken;

    return this._makeRequest('GET', `/sessions/${sessionId}/activities`, params, null, timeout);
  }

  /**
   * Send a message to the agent
   * @param sessionId - The session ID
   * @param request - Message parameters
   * @param timeout - Optional request timeout
   * @returns Promise<void>
   */
  async sendMessage(sessionId: string, request: SendMessageRequest, timeout?: number): Promise<void> {
    return this._makeRequest('POST', `/sessions/${sessionId}:sendMessage`, {}, request, timeout);
  }

  /**
   * Get a specific session
   * @param sessionId - The session ID
   * @param timeout - Optional request timeout
   * @returns Promise<Session>
   */
  async getSession(sessionId: string, timeout?: number): Promise<Session> {
    return this._makeRequest('GET', `/sessions/${sessionId}`, {}, null, timeout);
  }

  /**
   * Get a specific source
   * @param sourceId - The source ID
   * @param timeout - Optional request timeout
   * @returns Promise<Source>
   */
  async getSource(sourceId: string, timeout?: number): Promise<Source> {
    return this._makeRequest('GET', `/sources/${sourceId}`, {}, null, timeout);
  }
}

/**
 * Creates a new Jules API client.
 * @param options - Client configuration options.
 * @returns A new JulesClient instance.
 */
export function createClient(options?: JulesClientOptions): JulesClient {
  return new JulesClient(options);
}

export default JulesClient;
