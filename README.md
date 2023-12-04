# Gools

Gools is a collection of GoLang scripts to manage, well anything really. The initial focus is on AWS Cloud Management.

## Getting Started

### Prerequisites

- Go 1.21.3 or higher
- AWS account with appropriate permissions

### Installation

Clone the repository and navigate to the project directory:

```bash
git clone https://github.com/krishanthisera/gools.git
cd gools
```

Build the project:

```bash
go build
```

## Usage

For instance to manage IAM users and perform access key cleanup:

```bash
gools aws iam --profile your-profile --region your-region
gools aws iam keyCleanup --profile your-profile --region your-region --dry-run --concurrency 1 --max-age 90
```
