"""
Jules API Client implementation.
"""

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
        self.api_key = options.api_key
        self.base_url = options.base_url.rstrip('/')
        self.session = requests.Session()
        self.session.headers.update({
            'X-Goog-Api-Key': self.api_key,
            'Content-Type': 'application/json',
        })

    def _make_request(self, method: str, endpoint: str, params: Optional[dict] = None,
                     json_data: Optional[dict] = None) -> dict:
        """Make an HTTP request to the API."""
        url = f"{self.base_url}{endpoint}"
        response = self.session.request(method, url, params=params, json=json_data)
        response.raise_for_status()
        return response.json()

    def list_sources(self, next_page_token: Optional[str] = None) -> ListSourcesResponse:
        """
        List all available sources.

        Args:
            next_page_token: Token for pagination

        Returns:
            ListSourcesResponse: Available sources
        """
        params = {}
        if next_page_token:
            params['nextPageToken'] = next_page_token

        response = self._make_request('GET', '/sources', params=params)
        return ListSourcesResponse(**response)

    def create_session(self, request: CreateSessionRequest) -> Session:
        """
        Create a new session.

        Args:
            request: Session creation parameters

        Returns:
            Session: Created session
        """
        response = self._make_request('POST', '/sessions', json_data=request.dict())
        return Session(**response)

    def list_sessions(self, page_size: Optional[int] = None,
                     next_page_token: Optional[str] = None) -> ListSessionsResponse:
        """
        List sessions.

        Args:
            page_size: Maximum number of sessions to return
            next_page_token: Token for pagination

        Returns:
            ListSessionsResponse: List of sessions
        """
        params = {}
        if page_size:
            params['pageSize'] = page_size
        if next_page_token:
            params['nextPageToken'] = next_page_token

        response = self._make_request('GET', '/sessions', params=params)
        return ListSessionsResponse(**response)

    def approve_plan(self, session_id: str) -> None:
        """
        Approve the latest plan for a session.

        Args:
            session_id: The session ID
        """
        self._make_request('POST', f'/sessions/{session_id}:approvePlan')

    def list_activities(self, session_id: str, page_size: Optional[int] = None,
                       next_page_token: Optional[str] = None) -> ListActivitiesResponse:
        """
        List activities for a session.

        Args:
            session_id: The session ID
            page_size: Maximum number of activities to return
            next_page_token: Token for pagination

        Returns:
            ListActivitiesResponse: List of activities
        """
        params = {}
        if page_size:
            params['pageSize'] = page_size
        if next_page_token:
            params['nextPageToken'] = next_page_token

        response = self._make_request('GET', f'/sessions/{session_id}/activities', params=params)
        return ListActivitiesResponse(**response)

    def send_message(self, session_id: str, request: SendMessageRequest) -> None:
        """
        Send a message to the agent.

        Args:
            session_id: The session ID
            request: Message parameters
        """
        self._make_request('POST', f'/sessions/{session_id}:sendMessage',
                          json_data=request.dict())

    def get_session(self, session_id: str) -> Session:
        """
        Get details of a specific session.

        Args:
            session_id: The session ID

        Returns:
            Session: Session details
        """
        response = self._make_request('GET', f'/sessions/{session_id}')
        return Session(**response)

    def get_source(self, source_id: str) -> Source:
        """
        Get details of a specific source.

        Args:
            source_id: The source ID

        Returns:
            Source: Source details
        """
        response = self._make_request('GET', f'/sources/{source_id}')
        return Source(**response)


def create_client(api_key: str, base_url: Optional[str] = None) -> JulesClient:
    """
    Create a new Jules API client.

    Args:
        api_key: Your Jules API key
        base_url: API base URL (optional)

    Returns:
        JulesClient: Configured client instance
    """
    options = ClientOptions(api_key=api_key)
    if base_url:
        options.base_url = base_url
    return JulesClient(options)
