# Gator 🐊

Gator is a CLI-based RSS feed aggregator designed for efficiency and ease of use. It allows users to fetch, store, and browse RSS feeds directly from the command line. Built with Go and PostgreSQL, Gator ensures fast performance and seamless integration for managing your favorite news sources.

## Prerequisites

Ensure you have the following installed:

- **PostgreSQL** v15 or later ([Download](https://www.postgresql.org/download/))
- **Go** v1.23 or later ([Download](https://go.dev/dl/))

You can check your installed versions using:

```sh
psql --version  # PostgreSQL
go version      # Go
```

## Installation

To install Gator, ensure you have **Go 1.23+** installed, then run:

```sh
go install github.com/yourusername/gator@latest
```

## Configuration

After installation, you need to set up a configuration file.  

Manually create a config file in your home directory:  

```sh
touch ~/.gatorconfig.json
```

Then add the following content:

```json
{
  "db_url": "postgres://example"
}
```

Replace "postgres://example" with your actual PostgreSQL connection string.

Gator will use this file to connect to your database.

## Gator Commands

Here are the available commands for interacting with Gator:

- **Register and Log In:**
  - `gator register <username>`  
    Registers a new user and logs them in.
  
  - `gator login <username>`  
    Logs the user into the application.
  
  - `gator users`  
    Lists all users and shows the currently logged-in user.

- **Feed Management:**
  - `gator addfeed <feed_name> <feed_url>`  
    Adds a new feed to your collection.

  - `gator feeds`  
    Lists all feeds stored in the database.

  - `gator follow <feed_url>`  
    Starts following a specific feed.

  - `gator following`  
    Lists all the feeds you're currently following.

  - `gator unfollow <feed_url>`  
    Unfollows a specific feed.

- **Post Aggregation:**
  - `gator agg [<aggregate_interval>]`  
    Aggregates posts at a specified interval. The interval format can be `1s`, `1m`, or `1m0s`.

  - `gator browse`  
    Lists all posts that have been aggregated by Gator.
