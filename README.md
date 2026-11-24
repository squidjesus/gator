# gator

RSS Feed Aggre"gator" :)


## Requirements

- Postgres v15 or later
- Goose
- Go


## Installation

Install with go

```bash
go install github.com/squidjesus/gator
```
Note: You will also need to clone the repo and run the database migrations from 'sql/schema' with psql or goose (see database config below)


## Database Setup

- Start the Postgres server in the background

Mac:
```bash
brew services start postgresql@15
```

Linux:
```bash
sudo service postgresql start
```
- Run psql and create a new database, 'gator'.

```bash
psql
CREATE DATABASE gator;
```

- Set the password (Linux/WSL only)
```bash
ALTER USER postgres PASSWORD 'postgres';
```

- Run the migration files from the cloned repo inside 'sql/schema', update your connection string as needed
```bash
goose postgres 'postgres://postgres:postgress@localhost:5432/gator' up
```


## Config File

- Config file: Create a file called .gatorconfig.xml in your home directory.
```bash
{
    "db_url": "postgres://postgres:postgress@localhost:5432/gator?sslmode=disable",
}
```


## Basic Usage

- Register a user with `gator register <username>`
- Add feeds with `gator addfeed <name> <url>`
- Run `gator agg <time interval>` to start collecting feed posts
- Wait, or open another terminal to start browsing with `gator browse <#>`


## All Commands

`gator register <username>` - register a new user
`gator login <username>` - login as user
`gator users` - list registered users
`gator addfeed <name> <url>` - adds feed under current user and follows it
`gator feeds` - lists all saved feeds
`gator agg <time interval>` - scrapes feed with oldest updated timestamp every <time interval>, example: gator agg 5s
`gator follow <url>` - follows feed at given url
`gator following` - lists followed feeds for current user
`gator unfollow <url>` - unfollows feed at given url
`gator browse <#>` - prints # most recent posts of followed feeds to terminal, defaults to 2
`gator reset` - clears all db tables
