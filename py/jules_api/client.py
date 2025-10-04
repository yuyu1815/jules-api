"""
Jules API Client implementation.
"""

import os
import time
import requests
from typing import Optional

from .models import (
    ClientOptions,
    Source,
    Session,
    CreateSessionRequest,
    SendMessageRequest,
    ListSourcesResponse,
    ListSessionsResponse,
    ListActivitiesResponse,
    Activity,
)


class JulesClient:
    """Official Python client for the Jules API."""

    def __init__(self, options: ClientOptions):
        """
        Initialize the Jules API client.

        Args:
            options: Client configuration options
        """
        self.api_key = options.api_key or os.environ.get("JULES_API_KEY")
        if not self.api_key:
            raise ValueError("API key must be provided or set as JULES_API_KEY environment variable")
        self.base_url = options.base_url.rstrip('/')
        self.timeout = options.timeout
        self.session = requests.Session()
        self.session.headers.update({
            'X-Goog-Api-Key': self.api_key,
            'Content-Type': 'application/json',
        })

    def _make_request(self, method: str, endpoint: str, params: Optional[dict] = None,
                     json_data: Optional[dict] = None, timeout: Optional[int] = None) -> dict:
        """Make an HTTP request to the API."""
        url = f"{self.base_url}{endpoint}"

        request_timeout = timeout if timeout is not None else self.timeout

        if request_timeout == -1:
            while True:
                try:
                    response = self.session.request(method, url, params=params, json=json_data, timeout=60) # Default timeout for single request
                    response.raise_for_status()
                    return response.json()
                except requests.exceptions.RequestException as e:
                    if isinstance(e, requests.exceptions.HTTPError) and e.response.status_code == 404:
                        print(f"Resource not found (404), retrying in 1 second...")
                        time.sleep(1)
                    else:
                        print(f"Request failed with {e}, retrying in 1 second...")
                        time.sleep(1)
        else:
            response = self.session.request(method, url, params=params, json=json_data, timeout=request_timeout)
            response.raise_for_status()
            return response.json()

    def list_sources(self, next_page_token: Optional[str] = None, timeout: Optional[int] = None) -> ListSourcesResponse:
        """
        List all available sources.

        Args:
            next_page_token: Token for pagination
            timeout: Request timeout in seconds

        Returns:
            ListSourcesResponse: Available sources
        """
        params = {}
        if next_page_token:
            params['nextPageToken'] = next_page_token

        response = self._make_request('GET', '/sources', params=params, timeout=timeout)
        return ListSourcesResponse(**response)

    def create_session(self, request: CreateSessionRequest, timeout: Optional[int] = None) -> Session:
        """
        Create a new session.

        Args:
            request: Session creation parameters
            timeout: Request timeout in seconds

        Returns:
            Session: Created session
        """
        response = self._make_request('POST', '/sessions', json_data=request.dict(), timeout=timeout)
        return Session(**response)

    def list_sessions(self, page_size: Optional[int] = None,
                     next_page_token: Optional[str] = None, timeout: Optional[int] = None) -> ListSessionsResponse:
        """
        List sessions.

        Args:
            page_size: Maximum number of sessions to return
            next_page_token: Token for pagination
            timeout: Request timeout in seconds

        Returns:
            ListSessionsResponse: List of sessions
        """
        params = {}
        if page_size:
            params['pageSize'] = page_size
        if next_page_token:
            params['nextPageToken'] = next_page_token

        response = self._make_request('GET', '/sessions', params=params, timeout=timeout)
        return ListSessionsResponse(**response)

    def approve_plan(self, session_id: str, timeout: Optional[int] = None) -> None:
        """
        Approve the latest plan for a session.

        Args:
            session_id: The session ID
            timeout: Request timeout in seconds
        """
        self._make_request('POST', f'/sessions/{session_id}:approvePlan', timeout=timeout)

    def list_activities(self, session_id: str, page_size: Optional[int] = None,
                       next_page_token: Optional[str] = None, timeout: Optional[int] = None) -> ListActivitiesResponse:
        """
        List activities for a session.

        Args:
            session_id: The session ID
            page_size: Maximum number of activities to return
            next_page_token: Token for pagination
            timeout: Request timeout in seconds

        Returns:
            ListActivitiesResponse: List of activities
        """
        params = {}
        if page_size:
            params['pageSize'] = page_size
        if next_page_token:
            params['nextPageToken'] = next_page_token

        response = self._make_request('GET', f'/sessions/{session_id}/activities', params=params, timeout=timeout)
        return ListActivitiesResponse(**response)

    def send_message(self, session_id: str, request: SendMessageRequest, timeout: Optional[int] = None) -> None:
        """
        Send a message to the agent.

        Args:
            session_id: The session ID
            request: Message parameters
            timeout: Request timeout in seconds
        """
        self._make_request('POST', f'/sessions/{session_id}:sendMessage',
                          json_data=request.dict(), timeout=timeout)

    def get_session(self, session_id: str, timeout: Optional[int] = None) -> Session:
        """
        Get details of a specific session.

        Args:
            session_id: The session ID
            timeout: Request timeout in seconds

        Returns:
            Session: Session details
        """
        response = self._make_request('GET', f'/sessions/{session_id}', timeout=timeout)
        return Session(**response)

    def get_source(self, source_id: str, timeout: Optional[int] = None) -> Source:
        """
        Get details of a specific source.

        Args:
            source_id: The source ID
            timeout: Request timeout in seconds

        Returns:
            Source: Source details
        """
        response = self._make_request('GET', f'/sources/{source_id}', timeout=timeout)
        return Source(**response)


def create_client(api_key: Optional[str] = None, base_url: Optional[str] = None, timeout: Optional[int] = None) -> JulesClient:
    """
    Create a new Jules API client.

    Args:
        api_key: Your Jules API key. If not provided, it will be read from JULES_API_KEY env var.
        base_url: API base URL (optional)
        timeout: Request timeout in seconds (optional)

    Returns:
        JulesClient: Configured client instance
    """
    options = ClientOptions()
    if api_key:
        options.api_key = api_key
    if base_url:
        options.base_url = base_url
    if timeout is not None:
        options.timeout = timeout
    return JulesClient(options)
