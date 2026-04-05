# job-crawler

A Go REST API that concurrently crawls job postings from multiple sources and serves them filtered by location (Brazil or remote).

## Sources

| Source | Type | Filter |
|---|---|---|
| [Adzuna](https://www.adzuna.com.br) | Paid API | Configurable via env vars |
| [Greenhouse](https://www.greenhouse.io) | Public API | Per-company slug |
| [Lever](https://www.lever.co) | Public API | Per-company slug |

## Getting Started

### Prerequisites

- Go 1.23+
- An [Adzuna API key](https://developer.adzuna.com/)

### Environment variables

```env
ADZUNA_APP_ID=your_app_id
ADZUNA_API_KEY=your_api_key
ADZUNA_PAGES=3
ADZUNA_PER_PAGE=50
```

### Run locally

```bash
cp .env.example .env
# fill in your credentials
source .env && go run ./cmd/crawler
```

### Run with Docker

```bash
docker build -t job-crawler .
docker run -p 8080:8080 --env-file .env job-crawler
```

## API

The server starts immediately. Jobs are loaded asynchronously in the background — check the `status` field to know when results are ready.

### `GET /v1/jobs`

Returns a paginated list of job postings.

**Query parameters**

| Param | Default | Description |
|---|---|---|
| `page` | `1` | Page number |
| `page_size` | `50` | Results per page (max 100) |
| `title` | — | Filter by job title substring (case-insensitive) |
| `company` | — | Filter by company name substring (case-insensitive) |
| `source` | — | Filter by source: `adzuna`, `greenhouse`, `lever` |

**Example requests**

```
GET /v1/jobs
GET /v1/jobs?title=qa
GET /v1/jobs?title=developer&company=nubank
GET /v1/jobs?source=greenhouse&page=2&page_size=10
```

**Response**

```json
{
  "status": "done",
  "data": [
    {
      "id": "12345",
      "title": "QA Engineer",
      "company": "Nubank",
      "location": "São Paulo, Brazil",
      "description": "...",
      "url": "https://...",
      "source": "greenhouse"
    }
  ],
  "page": 1,
  "page_size": 50,
  "total": 42
}
```

> `status` is `"crawling"` while jobs are being fetched, and `"done"` once complete.

## Deploy on Render

1. Push this repository to GitHub
2. On Render → **New Web Service** → connect the repo
3. Render detects the `Dockerfile` automatically
4. Add the environment variables under **Environment**

## Project structure

```
cmd/crawler/        entry point
internal/
  api/
    handler/        HTTP request handlers
    render/         UTF-8 JSON response helper
    router.go       route registration
  classifier/       location & remote work detection
  config/           environment variable loading
  domain/           core data types
  service/          crawler orchestration
  sources/
    adzuna/         Adzuna API client
    greenhouse/     Greenhouse API client
    lever/          Lever API client
  store/            thread-safe in-memory job store
```
