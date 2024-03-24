# Navi: AI-Powered Shell Guide

"Navi" means "guide" in Hindi, and that's exactly what this tool aims to be. Navi is a command-line tool that utilizes the Gemini API to leverage artificial intelligence in generating shell commands based on your prompts. It simplifies your workflow by understanding your intent and providing the necessary commands to achieve your tasks.


## Commands

### navi
The `navi` command takes a string as an argument and generates a shell command based on that string.

Example:

```shell
$ navi "list all files`

$ >> ls -a
```

### navi explain
The `navi explain` command also takes a string (shell command) as an argument, and explains the functionality of a given shell command.

Example:

```shell
$ navi explain "ls -a`

$ Explanation >>

the command explanation
```
