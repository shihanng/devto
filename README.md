# devto -- publish to [dev.to](https://dev.to) from your terminal

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
