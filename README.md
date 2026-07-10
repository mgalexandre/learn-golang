# Learn Go by Building 🐹

A job-market-driven path for learning Go, built for backend roles across
**Portugal · Netherlands · Luxembourg · Belgium · UK** (listings surveyed July 2026).

It uses an evidence-based method: **backwards design**, **grill-before-code**,
**retrieval practice**, and **verify-by-running**. You learn a concept, then apply it
in a project that grows toward a production-shaped Go service.

## Two tracks

**Track A · Learnings** (concept modules) — the knowledge, demand-ordered:

| # | Module | Focus |
|---|--------|-------|
| L01 | Foundations | package/main, errors, structs & tags |
| L02 | Concurrency | goroutines, channels, `context`, worker pools |
| L03 | APIs & Distributed Systems | net/http, REST, gRPC, GraphQL, resilience |
| L04 | Containers & Cloud | Docker, Kubernetes, CI/CD, AWS |
| L05 | Data & Messaging | Postgres, Redis, Kafka/RabbitMQ |
| L06 | Testing & Observability | table tests, mocks, slog, metrics, tracing |
| L07 | Security & Auth | JWT/OAuth2, TLS, hashing, input validation |
| L08 | Idiomatic Go | generics, error wrapping, interfaces, layout |
| L09 | System Design | scaling, caching, load balancing, CAP |

**Track B · Projects** (build modules) — one evolving codebase:

| # | Project | Applies |
|---|---------|---------|
| P01 | Weather CLI | L01 |
| P02 | Concurrent Multi-City Fetcher | L02 |
| P03 | Weather Microservice (REST+gRPC, Postgres, Redis, auth) | L03·L05·L07 |
| P04 | Ship It: Containers → Cloud | L04·L06 |

The interactive modules live in [`modules/`](./modules/) — open `modules/index.html`
in a browser, or view the published versions on claude.ai.

## Getting started (Project P01)

```sh
go run . London
```

Then open **Project P01 · Weather CLI** and fill in the TODOs in `main.go`.

## Study plan

See the **Study plan** section at the bottom of `modules/index.html` for a suggested
6-week schedule that interleaves learning and building.
