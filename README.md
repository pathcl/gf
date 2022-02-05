# Why?

I like reading code.

Let's say we want to find all Go code in 'kubernetes' github org for function 'fmt.Println':

    $ gf search extension:go org:kubernetes fmt.Println | vim -

Note: You need to have 'GITHUB_TOKEN' variable present in your environment.

# Install

    $ go install github.com/pathcl/gf@latest
    go: downloading github.com/pathcl/gf v0.0.2
    go: downloading github.com/spf13/cobra v1.3.0
    go: downloading github.com/google/go-github v17.0.0+incompatible
    go: downloading golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8
    go: downloading golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d
    go: downloading github.com/spf13/pflag v1.0.5
    go: downloading github.com/google/go-querystring v1.1.0

    $ gf --help
    Built with <3
    Usage:
      gf [command]

    Available Commands:
      completion  Generate the autocompletion script for the specified shell
      help        Help about any command
      search      Search a given string in Github
      version     Display version information

    Flags:
      -h, --help      help for gf
      -v, --version   version for gf

    Use "gf [command] --help" for more information about a command.


## Building the application
Run `make build`.

## Running the application
Run `make run`.
To run a command or pass flags, pass them via `args`. For example: `make run args="deploy --help"`
