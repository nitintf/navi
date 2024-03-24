# Navi: AI-Powered Shell Guide

Navi is a command-line tool that utilizes artificial intelligence to generate shell commands based on your prompts. It simplifies your workflow by understanding your intent and providing the necessary commands to achieve your tasks.


## Commands

### navi
The `navi` command takes a string as an argument and generates a shell command based on that string.

Example:

```shell
navi "list all files`
```
This might output `ls -a`

### navi explain
The `navi explain` command also takes a string (shell command) as an argument, and explains the functionality of a given shell command.

Example:

```shell
navi explain "ls -a"
```
