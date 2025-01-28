# Gator

A cli written in `go` for subscribing to and aggregating RSS feeds.

## Installation

You will need an instance of `Postgres` running.

To install, run `go install`, you can then run the `gator` command

## Configuration

You will need to create a `~/.gatorconfig.json` file with the following information:
```json
{
    "db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
}

```
## Running `gator`
### Create a new user
```bash
gator register <user>
```

### List all users
```bash
gator users
```
### Switch user
```bash
gator login <user> ```
### Add a new RSS feed
```bash
gator addfeed <rss_feed_name> <rss_feed_url>
```
### List all RSS feeds
```bash
gator feeds
```
### Follow a RSS feed for current user
```bash
gator follow <rss_feed_url>
```
### Unfollow a RSS feed for current user
```bash
gator unfollow <rss_feed_url>
```
### Aggregate RSS feeds
```bash
gator agg <time_between_aggregations>
```
### Browse RSS posts user is following
```bash
gator browse <*optional limit>
```
### Clear out the database
```bash
gator reset
```
## Developement

If you have `nix` installed, the easiest way to install all remaining depencencies is simply running the following in your root.
```bash
nix develop
```

Otherwise, you will also need:
- `psql` (for interacting with the database)
- `goose` (for database migrations)
- `sqlc` (for generating Go interfaces from SQL queries)


