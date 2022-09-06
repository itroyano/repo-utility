# Repository Utility

Repo Utility generates commonly used CI integrations in a code repository. These include GitHub actions, dependency management, code coverage, and more.

## The problem Repo Utility Solves

When creating a new repository, CI integrations often need to be integrated manually. sometimes this can be done from a template, other times from scratch.
using templates partially resolves the manual effort, but it does not provide the flexibility of hand-picking which specific integrations are needed for this specific repo. 
Repo Utility breaks down different common integrations into separate generation commands, which can be run in any desired order.

## Who Is It For?

Anyone :)
the utility can be used in or outside GitHub, by selecting the approprite generation actions.

## Usage

```
  repo-utility [command]
```

Sub commands will show their own help and usage.

# How to Build and Test repo utility

A Makefile is included with common options:

```
  make build
  make test
  make fmt
```

to build a specific version binary use `REPO_UTILITY_VERSION=0.0.1 make build`

### Requirements

- Go version 1.18+
- Make 4+
- Linux or MacOS

### Using VSCode with launch.json

If you want to run repo utility with VSCode debugging features here is an example on how the launch.json file can look like:

```
{
    "version": "0.2.0",
    "configurations": [
      {
        "name": "Launch Package",
        "type": "go",
        "request": "launch",
        "mode": "auto",
        "program": "${fileDirname}",
        "args": ["version"],
      }
    ]
  }

```

If you want to know more about debugging on VSCode check this page https://code.visualstudio.com/docs/editor/debugging

### Contribution Guidelines

