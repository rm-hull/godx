# Project: godx

## Project Overview

`godx` is a Go package that provides diagnostic utilities for Go applications. It offers functions to inspect and log application version information, environment variables (with automatic masking of sensitive values), and user/process details. The project is structured as a Go module and is intended to be used as a library in other Go applications.

## Building and Running

This is a library project, so there is no main application to run. However, you can build and test the package using standard Go tools.

### Key Commands:

*   **To run tests:**
    ```bash
    go test ./...
    ```
    The project is configured to use `gotestsum` for more detailed test output in CI:
    ```bash
    gotestsum -- -v ./...
    ```

*   **To run linter:**
    The project uses `golangci-lint`.
    ```bash
    golangci-lint run
    ```

*   **To build the package:**
    ```bash
    go build ./...
    ```

## Development Conventions

*   **Testing:** Tests are written using the standard `testing` package and the `testify/assert` library for assertions. All new functionality should be accompanied by corresponding unit tests.
*   **Linting:** The project uses `golangci-lint` to enforce code style and quality. All code should pass the linter before being committed.
*   **Continuous Integration:** The `.github/workflows/build.yml` file defines the CI pipeline, which includes steps for building, testing, and linting the code on every push and pull request to the `main` branch.
*   **Dependency Management:** Dependencies are managed using Go modules (`go.mod` and `go.sum`). The project uses `dependabot` to automatically keep dependencies up-to-date.
*   **Versioning:** The project uses an automated versioning and release process based on commit messages, managed by the `github-tag-action`.
