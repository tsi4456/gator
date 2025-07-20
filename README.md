# gator

A simple RSS feed aggregator (geddit), in Go. Polls selected feeds at a given frequency and caches posts in a Postgres database. Supports feed tracking for multiple users.

### Requirements

Requires Postgres and Go

### Installation

`go install github.com/tsi4456/gator`

#### config

gator looks for a config file, `.gatorconfig.json`, in your home directory. The file should contain `{"db_url": "postgres:USERNAME:PASSWORD@localhost:5432/gator?sslmost=disable"}`, where USERNAME and PASSWORD are the login details for your local Postgres database. Should you be running Postgres on a non-standard port, you will also need to modify this.

#### Running gator

gator supports the following commands, in the form `gator COMMAND <args>`:

`register <USERNAME>` adds USERNAME to the list of users and sets it as the current user
`login <USERNAME>` sets USERNAME as the current user, if registered
`users` lists the registered users

`agg <PERIOD>` begins feed aggregation; continually polls feeds on the given period (default: 5 minutes)

`addfeed <URL>` adds an RSS feed from the given URL to the feed database and to the current user's follow list
`feeds` prints a list of the current RSS feeds tracked by the program

`follow <URL>` if URL is already in the database, adds this feed to the current user's follow list
`unfollow <URL>` removes URL from the current user's follow list
`following` prints a list of the feeds currently followed by the current user
`browse <LIMIT>` prints a list of the most recent posts from the current user's feeds, up to LIMIT (default: 2)

`reset` clears the user and post databases; use with care!
