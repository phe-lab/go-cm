# go-cm/log

- [ZeroLog](https://github.com/rs/zerolog) integration
- Support logging [Resty](https://github.com/go-resty/resty) HTTP requests

## Installation

```bash
go get github.com/phe-lab/go-cm/log@v0.2.0
```

## Usage

### Quickstart

```go
package main

import "github.com/phe-lab/go-cm/log"

func main() {
	logger := log.NewLogger()
	logger.Debug().Str("host", "0.0.0.0").Int("port", 3000).Send()
	logger.Info().Msg("Application starting")
}
```

```bash
# output:
{"level":"debug","host":"0.0.0.0","port":3000,"time":"2024-11-16T13:38:53+07:00"}
{"level":"info","time":"2024-11-16T13:38:53+07:00","message":"Application starting"}
```

### Global setup

```go
package main

import "github.com/phe-lab/go-cm/log"

func main() {
	log.WithCaller(true)
	log.SetGlobalFormat("json")
	log.SetGlobalLevel("info")

    logger := log.NewLogger()
	logger.Debug().Str("host", "0.0.0.0").Int("port", 3000).Send()
	logger.Info().Msg("Application starting")
}
```
