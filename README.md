# What is gf?
Simple tool which looks for a string in Github and output Go code where the query is present.

Usage:

    $ gf search fmt.Println | vim -

Note: You need to have 'GITHUB_TOKEN' variable present in your environment.

## Building the application
Run `make build`.

## Running the application
Run `make run`.
To run a command or pass flags, pass them via `args`. For example: `make run args="deploy --help"`
