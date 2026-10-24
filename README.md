# go-microservice-template

My starting point for new Go services and tools. It is deliberately small: a
working HTTP entrypoint, a dev container, live reload, linting, and Renovate —
the pieces I would otherwise copy by hand into every new repository. Anything
else gets added per project, once the project actually needs it.

## What's included

| Path               | Purpose                                                             |
| ------------------ | ------------------------------------------------------------------- |
| `cmd/server/`      | Minimal HTTP server on `:8080` using `net/http` and `log/slog`      |
| `internal/`        | Private application code, not importable by other modules           |
| `pkg/`             | Code intended to be imported by other modules                       |
| `.devcontainer/`   | Dev container with Go, gopls, Delve, Air, and golangci-lint         |
| `.air.toml`        | Live reload configuration for [Air](https://github.com/air-verse/air) |
| `.golangci.yml`    | Linter configuration                                                |
| `.vscode/`         | Debug configurations and tasks (`go mod tidy`, start/stop Air)      |
| `renovate.json`    | Automated dependency and Go version updates                         |

## Requirements

- Go 1.27.3 or newer
- Optional: Docker and VS Code with the Dev Containers extension
- Optional: Air and golangci-lint when working outside the container

## Starting a new project

1. Create a repository from this template (GitHub **Use this template**) or clone it.
2. Point the module path at the new repository:

   ```bash
   go mod edit -module github.com/<owner>/<repo>
   ```

3. Replace this README and adjust `LICENSE` if needed.
4. Drop the `internal/` or `pkg/` placeholders that the project does not use.

## Development

Run the server:

```bash
go run ./cmd/server
```

```bash
curl localhost:8080
```

Run with live reload (rebuilds on changes to `.go` and template files):

```bash
air
```

Test and lint:

```bash
go test ./...
```

```bash
golangci-lint run ./...
```

### Dev container

Open the folder in VS Code and choose **Reopen in Container**. The image ships
the Go toolchain plus gopls, Delve, Air, and golangci-lint, and runs
`go mod tidy` after creation.

### Debugging

`.vscode/launch.json` provides a launch configuration for the current package
and a remote attach configuration for a headless Delve server on port `2345`.

## Conventions

- Standard library first; a dependency has to earn its place.
- Application code lives in `internal/` by default and moves to `pkg/` only when it is meant to be shared.
- Structured logging with `log/slog`.
- Table-driven tests.

## License

MIT — see [LICENSE](LICENSE).
