# Prack

## Use Case

| Tool name     | prack                                                      |
| ------------- | ---------------------------------------------------------- |
| commands      | init, add, remove, list, recent, open, close, remove, info |
| resources     | project                                                    |
| resource-name | any random string to identify project                      |
| parameters    | --number(-n): number of projects to display in list        |
|               | --tag(-t): display list by tag                             |

## Commands

- init: Initialize project and create prack.yaml file in project directory
- add: Register the project directory using prack.yaml
- remove: Deregister the project directory
- list: display list of registered projects
  - flags: number, tag
- recent: most recently used projects
  - flags: number
- open <project_name> or open <project_alias>: open the project
- close: close the project
- info: display info about the project
- recent: recently accessed projects

## YAML file structre

Generated: name, cd command

User input: alias, list of commands to execute for open and close, tags, info
