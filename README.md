# Navi: AI-Powered Shell Guide

"Navi" means "guide" in Hindi, and that's exactly what this tool aims to be. Navi is a command-line tool that utilizes the Gemini API to leverage artificial intelligence in generating shell commands based on your prompts. It simplifies your workflow by understanding your intent and providing the necessary commands to achieve your tasks.

## Usage

```shell
Navi - Your AI-powered Shell Guide

Usage:
  navi [flags]
  navi [command]

Examples:
navi "List all files in the current directory."

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  explain     Explain - Understand your shell commands
  help        Help about any command

Flags:
  -h, --help   help for navi

Use "navi [command] --help" for more information about a command.
```

## Commands

### 1. navi
The `navi` command takes a string as an argument and generates a shell command based on that string.

Example:

```shell
$ navi "list all files`

$ >> ls -a
```

### 2. navi explain
The `navi explain` command also takes a string (shell command) as an argument, and explains the functionality of a given shell command.

Example:

```shell
$ navi explain "ls -a`

$ Explanation >>

the command explanation
```

## Installation

You can install Navi using the `go install` command:

```shell
go install github.com/nitintf/navi
```

After installation, don't forget to export your Gemini API key:
```shell
export GEMINI_API_KEY="..."
```
You can obtain your Gemini API key from the [Gemini API Management page](https://aistudio.google.com/app/apikey).

## Caution

While Navi uses AI to generate shell commands, it's important to understand that AI isn't perfect. Always review the generated commands before executing them, especially if you're working in a production environment or dealing with sensitive data. Navi is a tool designed to assist you, but it doesn't replace good judgment and understanding of shell commands.

## Contribution

Contributions to Navi are very welcome! If you have a feature request, bug report, or proposal for code refactoring, please feel free to open an issue or submit a pull request.

This thing only supports Gemini for the moment, but it could be easily updated to support other models using Replicate, Ollama, etc. [Pull requests welcome](https://github.com/nitintf/navi/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc)!


## License

MIT
