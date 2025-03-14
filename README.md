# Gator CLI

A command-line tool for managing feeds and posts, built with Go and PostgreSQL.

## Prerequisites

To run Gator, you need to have the following installed on your system:

- **PostgreSQL**: A powerful, open-source relational database. You can download it from the [official PostgreSQL website](https://www.postgresql.org/download/).
- **Go**: The Go programming language. Installation instructions are available on the [official Go website](https://go.dev/doc/install).


## Installation

### Install Gator CLI

To install the Gator CLI tool, use the following command in your terminal:

```bash
go install github.com/your-github-username/gator@latest
```

This will download and compile the binary for Gator, placing it in your Go bin directory.

### Setup Configuration File

Before running Gator, you need to configure it by setting up a YAML configuration file. Here’s an example of a simple configuration file:

```yaml
currentUserName: "user123"
database:
  user: "your-postgres-username"
  password: "your-postgres-password"
  dbname: "gator"
  host: "localhost"
  port: 5432
```

Place this configuration file at `~/.gator/config.yaml` or any location you prefer and make sure to specify its path when running Gator.

## Running Gator

After installation and configuration, you can run Gator using the following commands:

- **Register a new user**: `gator register <username>`
- **Log in as an existing user**: `gator login <username>`
- **Add a new feed and follow it**: `gator addfeed "feedname" "feedurl"`
- **List all feeds**: `gator feeds`
- **Follow an existing feed**: `gator follow feedurl`
- **List all feeds followed by the current user**: `gator following`
- **Unfollow a feed**: `gator unfollow feedurl`
- **Aggregate posts from RSS feeds at intervals**: `gator agg 10m`
- **Browse posts for the current user**: `gator browse`


## Deployment

Gator is designed to be a statically compiled binary. After running `go build` or `go install`, you can run your code without needing the Go toolchain. Use `go run .` for development purposes and the compiled binary for production.
