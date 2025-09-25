# godx

`godx` is a Go package providing diagnostic utilities for Go applications. It offers functions to inspect application version, environment variables, and user information.

## Features

- **GitVersion**: Logs the short version information of the application.
- **EnvironmentVars**: Logs environment variables, masking sensitive information.
- **UserInfo**: Logs process ID, current user details, and group memberships.

## Installation

To use `godx` in your Go project, you can install it using `go get`:

```bash
go get github.com/rm-hull/godx
```

## Usage

Import the `godx` package into your Go application and call the desired diagnostic functions:

```go
package main

import (
	"github.com/rm-hull/godx"
)

func main() {
	godx.GitVersion()
	godx.EnvironmentVars()
	godx.UserInfo()
}
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
