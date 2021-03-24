# Prack

A CLI project manager for developers.

## Use Case

| Tool name     | prack                                               |
| ------------- | --------------------------------------------------- |
| commands      | init, add, remove, list, recent, remove             |
| resources     | project                                             |
| resource-name | any random string to identify project               |
| parameters    | --number(-n): number of projects to display in list |
|               | --tag(-t): display list by tag                      |
|               | --info(-i): display name, alias & description       |

## Commands

- init: Initialize project and create prack.yaml file in project directory
- add: Register the project directory using prack.yaml
- remove <project_alias>: Deregister the project directory
- list: display list of registered projects
  - flags: number, tag
- recent: most recently used projects
  - flags: number
- <project_alias> <?cmd_block>: open the project
- recent: recently accessed projects

## YAML file structre

Generated: name, cd command

User input: alias, list of commands to execute for open and close, tags, info
