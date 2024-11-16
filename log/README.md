# go-cm/log

- ZeroLog integration
- Support request logging with Resty

## Installation

```bash
go get -u github.com/phe-lab/go-cm/log
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
