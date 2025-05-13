# Gitfluence CLI

Gitfluence CLI is a command-line tool to interact with Gitfluence API and execute generated commands.

## Installation

To install the tool, run:

```bash
go install github.com/pzaeemfar/gitfluence-cli@latest
````

## Usage

Run the command with your desired prompt:

```bash
gitfluence-cli <your-prompt>
```

Example:

```bash
gitfluence-cli reset whole repo
```

The tool will fetch a command from the Gitfluence API and ask if you want to execute it.

## Note
Using a shell alias can make it easier to use the command.
