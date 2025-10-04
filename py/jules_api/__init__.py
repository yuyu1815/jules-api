"""
Jules API Python Client

Official Python client library for the Jules API.
"""

from .client import JulesClient, create_client
from .models import (
    Source,
    GithubRepo,
    GithubRepoContext,
    SourceContext,
    Session,
    CreateSessionRequest,
    SendMessageRequest,
    ListSourcesResponse,
    ListSessionsResponse,
    Activity,
    ListActivitiesResponse,
)

__version__ = "0.1.0a"
__all__ = [
    "JulesClient",
    "create_client",
    "Source",
    "GithubRepo",
    "GithubRepoContext",
    "SourceContext",
    "Session",
    "CreateSessionRequest",
    "SendMessageRequest",
    "ListSourcesResponse",
    "ListSessionsResponse",
    "Activity",
    "ListActivitiesResponse",
]
