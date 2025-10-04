"""
Pydantic models for Jules API.
"""

from datetime import datetime
from typing import Optional

from pydantic import BaseModel


class GithubRepo(BaseModel):
    """GitHub repository information."""
    owner: str
    repo: str


class GithubRepoContext(BaseModel):
    """Additional context for GitHub repositories."""
    starting_branch: Optional[str] = None


class SourceContext(BaseModel):
    """Source context for a session."""
    source: str
    github_repo_context: Optional[GithubRepoContext] = None


class Source(BaseModel):
    """Represents an input source (e.g., GitHub repository)."""
    name: str
    id: str
    github_repo: Optional[GithubRepo] = None


class Session(BaseModel):
    """Represents a continuous unit of work within a specific context."""
    name: str
    id: str
    title: str
    source_context: Optional[SourceContext] = None
    prompt: Optional[str] = None


class CreateSessionRequest(BaseModel):
    """Request to create a new session."""
    prompt: str
    source_context: SourceContext
    title: str
    require_plan_approval: Optional[bool] = False


class SendMessageRequest(BaseModel):
    """Request to send a message to the agent."""
    prompt: str


class ListSourcesResponse(BaseModel):
    """Response from listing sources."""
    sources: list[Source]
    next_page_token: Optional[str] = None


class ListSessionsResponse(BaseModel):
    """Response from listing sessions."""
    sessions: list[Session]
    next_page_token: Optional[str] = None


class Activity(BaseModel):
    """Represents a single unit of work within a Session."""
    name: str
    id: str
    type: str
    content: Optional[str] = None
    timestamp: Optional[datetime] = None


class ListActivitiesResponse(BaseModel):
    """Response from listing activities."""
    activities: list[Activity]
    next_page_token: Optional[str] = None


class ClientOptions(BaseModel):
    """Client configuration options."""
    api_key: str
    base_url: Optional[str] = "https://jules.googleapis.com/v1alpha"

    class Config:
        validate_assignment = True
