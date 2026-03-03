# secdiag — Product Roadmap

Terminal security diagnostics for backend engineers. Phases are ordered to grow both security coverage and engineering maturity.

---

## Phase 0 — Foundation

**Goal:** Solid architecture before adding features.

**Deliverables:**

- **Project structure:** `cmd/secdiag`, `internal/app`, `internal/core`, `internal/analyzers`, `internal/tui`, `internal/infra`
- **Core contracts:** `Analyzer` interface (`Name`, `Analyze(ctx, target) (Result, error)`), `Result` with `Target`, `Score`, `Findings`
- **Analyzer registry:** Register analyzers by type; no switch-case in app logic
- **Basic TUI:** Bubble Tea — tool selection, URL input, loading state, result view

**Concepts:** Event-driven flow, state modeling, interface-driven design, clear boundaries.

---

## Phase 1 — Core Security

**Goal:** Real backend security checks.

**v0.1 – TLS Analyzer**

- TLS version detection, weak TLS (1.0/1.1) flagging
- Certificate expiration and self-signed detection
- Cipher suite listing

**Concepts:** `crypto/tls`, manual handshake, x509, timeouts, error classification.

**v0.2 – HTTP Security Headers**

- Check HSTS, CSP, X-Frame-Options, X-Content-Type-Options, Referrer-Policy
- Severity model and weighted risk score

**Concepts:** HTTP client tuning, response inspection, security/risk modeling.

---

## Phase 2 — Networking & Infra

**Goal:** Infra-level visibility.

**v0.3 – Port Scanner**

- Common ports (e.g. 80, 443, 22, 3306, 5432), worker pool, rate limit, timeouts

**Concepts:** TCP dialing, concurrency safety, backpressure.

**v0.4 – DNS Analyzer**

- A, AAAA, CNAME, MX; multiple IPs; basic CDN heuristics

**Concepts:** `net.Resolver`, infra and attack-surface thinking.

**v0.5 – Open Redirect & Basic Vuln Checks**

- Open redirects, server header leaks, framework exposure

**Concepts:** Redirect policy, header analysis, heuristics.

---

## Phase 3 — Engineering Maturity

**Goal:** Polish and automation.

**v0.6 – Headless CLI**

- `secdiag scan <url> --tool tls`, JSON output, CI-friendly exit codes

**Concepts:** CLI parsing, UI vs core separation, automation.

**v0.7 – JSON & HTML Reports**

- Export results as JSON and HTML, severity filtering

**Concepts:** Serialization, report generation, DTOs.

**v0.8 – Config File**

- `secdiag.yaml`: timeouts, port ranges, severity thresholds

**Concepts:** Config layering, defaults, validation.

---

## Phase 4 — Production-Grade Open Source

**Goal:** Extensibility and release quality.

**v1.0 – Plugin System**

- Dynamic analyzer registration behind a stable interface

**v1.1 – Performance**

- Benchmark analyzers, tune concurrency, memory profiling (`pprof`)

**v1.2 – CI/CD & Releases**

- GitHub Actions, multi-platform builds, version tags, release notes

**Concepts:** DevOps, release engineering.

---

## End State
