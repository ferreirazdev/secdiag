# secdiag

**Terminal-based security diagnostics for backend engineers.**

secdiag is an interactive TUI (Terminal UI) built in Go. Run quick security checks against your servers and APIs—TLS configuration, security headers, and common misconfigurations—without leaving the terminal.

---

## About

Backend services often ship with weak TLS, missing security headers, or exposed ports. secdiag helps you catch these issues early with a focused, keyboard-driven interface. No browser, no dashboards—just point it at a host and get a clear security snapshot.

**Best for:** API maintainers, SREs, and developers who want fast, repeatable security checks during development and deployment.

---

## Features

- **Interactive TUI** — Navigate and run checks from the keyboard
- **TLS analysis** — Versions, cipher suites, certificate validity, self-signed detection
- **HTTP security headers** — HSTS, CSP, X-Frame-Options, X-Content-Type-Options, Referrer-Policy
- **Severity levels** — Findings tagged as LOW / MEDIUM / HIGH
- **Security score** — Simple 0–100 score per target
- **Graceful exit** — Ctrl+C and timeouts handled cleanly

---

## Installation

**From source (requires Go 1.21+):**

```bash
git clone https://github.com/ferreirazdev/secdiag.git
cd secdiag
go build -o secdiag ./cmd/secdiag
```

Or run without installing:

```bash
make run
# or
go run ./cmd/secdiag
```

---

## Quick start

```bash
make run
```

Launch the TUI, then follow the on-screen prompts. Use **q** or **Ctrl+C** to quit.

---

## Project structure

```
secdiag/
├── cmd/secdiag/          # Entrypoint
├── internal/
│   ├── app/              # Application flow & orchestration
│   ├── core/             # Analyzer interface & result types
│   ├── analyzers/        # TLS, headers, ports, etc.
│   ├── tui/              # Terminal UI (Bubble Tea)
│   └── infra/            # HTTP client, DNS, shared infra
└── Makefile
```

---

## Roadmap

| Version | Focus |
|--------|--------|
| **v0.1** | TUI, TLS analyzer, headers analyzer, severity levels, security score |
| **v0.2** | Port scanner, DNS analyzer, JSON export, headless CLI (`secdiag scan <url>`) |
| **v0.3** | Reverse proxy checks, open redirects, rate-limit detection, configurable rules |
| **v1.0** | Plugin system, config file, HTML reports, CI mode, cross-platform releases |

---

## Contributing

Contributions are welcome. Open an issue to discuss ideas or send a pull request. For larger changes, start with an issue so we can align on design.

---

## License

[MIT](LICENSE) — use it in your own projects and portfolios.
