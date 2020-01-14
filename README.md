# devto -- publish to [dev.to](https://dev.to) from your terminal

[![](https://github.com/shihanng/devto/workflows/main/badge.svg?branch=develop)](https://github.com/shihanng/devto/actions?query=workflow%3Amain)
[![Coverage Status](https://coveralls.io/repos/github/shihanng/devto/badge.svg?branch=develop)](https://coveralls.io/github/shihanng/devto?branch=develop)
[![Go Report Card](https://goreportcard.com/badge/github.com/shihanng/devto)](https://goreportcard.com/report/github.com/shihanng/devto)

## Configuration

| Description                                                    | CLI Flag    | Environment Variable | `config.yml` |
|----------------------------------------------------------------|-------------|----------------------|--------------|
| [DEV API key](https://docs.dev.to/api/#section/Authentication) | `--api-key` | `DEVTO_API_KEY`      | `api-key`    |

### Sample config in YAML

```yaml
api-key: abcd1234
```


## Generate [dev.to's API](https://docs.dev.to/api/) client

```
make gen
```

See [`pkg/devto`](./pkg/devto).
