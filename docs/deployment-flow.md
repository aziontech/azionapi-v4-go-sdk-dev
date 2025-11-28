# Deployment Flow

```mermaid
flowchart TD
    A[openapi-v4 Repo: Push to dev branch] --> B[Detect Modified YAML Files]
    B --> C[Copy YAMLs to go-sdk-dev Repo]
    C --> D[Create PR with Updated YAMLs]
    D --> E[PR Opened in go-sdk-dev]
    E --> F[Check Commit Message]
    F -->|"Updated YAML files"| G[Find Modified YAMLs]
    F -->|No Match| H[Skip Workflow]
    G --> I[Patch YAML Files]
    I --> J[Run OpenAPI Generator]
    J --> K[Generate Go SDK Code]
    K --> L[Commit Generated SDKs]
    L --> M[Remove YAML Files]
    M --> N[Push to PR Branch]
    N --> O[Notify Slack]
```

## Flow Overview

### Stage 1: openapi-v4 Repository
- **Trigger**: Push to `dev` branch
- **Action**: Detects modified YAML files and copies them to `azionapi-v4-go-sdk-dev`
- **Output**: Creates PR with updated YAML files

### Stage 2: go-sdk-dev Repository
- **Trigger**: PR opened/synchronized
- **Validation**: Checks commit message matches "Updated YAML files"
- **Generation**: 
  - Patches YAML files
  - Runs OpenAPI Generator for each modified YAML
  - Generates Go SDK code
- **Cleanup**: Removes YAML files after generation
- **Notification**: Sends Slack notification with PR link
