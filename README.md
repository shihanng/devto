# devto -- publish to [dev.to](https://dev.to) from your terminal

[![CI](https://github.com/shihanng/devto/workflows/main/badge.svg?branch=develop)](https://github.com/shihanng/devto/actions?query=workflow%3Amain)
[![Release](https://github.com/shihanng/devto/workflows/release/badge.svg)](https://github.com/shihanng/devto/actions?query=workflow%3Arelease)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/shihanng/devto)](https://github.com/shihanng/devto/releases)
[![Coverage Status](https://coveralls.io/repos/github/shihanng/devto/badge.svg?branch=develop)](https://coveralls.io/github/shihanng/devto?branch=develop)
[![Go Report Card](https://goreportcard.com/badge/github.com/shihanng/devto)](https://goreportcard.com/report/github.com/shihanng/devto)
[![GitHub](https://img.shields.io/github/license/shihanng/devto)](./LICENSE)

## What is this?

`devto` is a CLI tool that helps submit articles to DEV from the terminal. It makes use of the [APIs that DEV kindly provides in OpenAPI specification](https://docs.dev.to/api/). `devto` mainly does two things:

1. It collects all image links from the Markdown file into a `devto.yml` file with the `generate` subcommand. For example, if we have `./image-1.png` and `./image-2.png` in the Markdown file, we will get the following:

   ```yml
   images:
     ./image-1.png: ""
     ./image-2.png: ""
   ```

2. It submits the article to DEV with the `submit` subcommand. The `submit` subcommand creates a new article in DEV and updates the `devto.yml` with the resulting `article_id`. `devto` will use this `article_id` in the following execution to perform an update operation instead of creating a new entry for the same article.

The DEV API does not have a way of uploading images yet. If we submit a Markdown content with relative paths of image links, DEV will not be able to show those images. As a workaround of this problem, we need to provide a full path for the images either manually via the `devto.yml` file or using the `--prefix` flag.

The Markdown file must contains at least the title property of the Jekyll front matter, like in:

```
---
title: An example title
description: ...
tags: ...
cover_image: ...
---
```

You can find more information about the usage via the `--help` flag.

```sh
devto --help
```

## Installation

### [Homebrew (macOS)](https://github.com/shihanng/homebrew-devto)

```sh
brew install shihanng/devto/devto
```

### Debian, Ubuntu

```sh
curl -sLO https://github.com/shihanng/devto/releases/latest/download/devto_linux_amd64.deb
dpkg -i devto_linux_amd64.deb
```

### RedHat, CentOS

```sh
rpm -ivh https://github.com/shihanng/devto/releases/latest/download/devto_linux_amd64.rpm
```

### Binaries

The [release page](https://github.com/shihanng/devto/releases) contains binaries built for various platforms. Download the version matches your environment (e.g. `linux_amd64`) and place the binary in the executable `$PATH` e.g. `/usr/local/bin`:

```sh
curl -sL https://github.com/shihanng/devto/releases/latest/download/devto_linux_amd64.tar.gz | \
    tar xz -C /usr/local/bin/ devto
```

### For Gophers

With [Go](https://golang.org/doc/install) already installed in your system, use `go get`

```sh
go get github.com/shihanng/devto
```

or clone this repo and `make install`

```sh
git clone https://github.com/shihanng/devto.git
cd devto
make install
```

## Configuration

| Description                                                                                    | CLI Flag    | Environment Variable | `config.yml` |
| ---------------------------------------------------------------------------------------------- | ----------- | -------------------- | ------------ |
| [DEV API key](https://docs.dev.to/api/#section/Authentication) is needed to talk with DEV API. | `--api-key` | `DEVTO_API_KEY`      | `api-key`    |

### Sample config in YAML

```yaml
api-key: abcd1234
```

## Contributing

Want to add missing feature? Found bug :bug:? Pull requests and issues are welcome. For major changes, please open an issue first to discuss what you would like to change :heart:.

```sh
make lint
make test
```

should help with the idiomatic Go styles and unit-tests.

### How to generate [DEV's API](https://docs.dev.to/api/) client

```sh
make gen
```

See [`pkg/devto`](./pkg/devto) for client documentation.

## License

[MIT](./LICENSE)
