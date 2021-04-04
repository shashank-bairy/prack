# Prack (Project Rack)

A CLI project assistant for developers used to avoid repetitive typing of commands while opening a project.

## Glossary

- **Project alias**: A single single word used to identify project
- **Command block**: A set of commands to be executed in a terminal/prompt
- **Command block alias**: A short single world used to identify command block for a particular project

## Commands

- **init**: Creates prack.yaml file in project directory.
- **add**: Register the project using prack.yaml and the commands associated.
- **remove <project_alias>**: Deregister the project directory
- **open <project_alias> <cmd_block_alias>**: execute the command block of a project
- **list**: display list of registered projects

## prack.yaml file structre

```yaml
---
project:
  name: Project Name
  alias: project_alias
  description: Project Description
  tags:
    - tag_1
    - tag_2
  commands:
    - alias: command_alias_1
      commands:
        - command 1
        - command 2
        - command 3
    - alias: command_alias_2
      commands:
        - command 1
        - command 2
```

## Instructions:

1. Navigate inside your project root directory
2. Run `prack init` and above yaml template is created in your directory
3. Modify the yaml file according to your project requirements.
4. Now run `prack add` command to store project details.
5. You can run `prack open <project_alias> <commnad_block_alias>` to execute the required commands. (This can be done from any directory)
6. For any updation in commands, run `prack update` after making changes in prack.yaml.
7. Run `prack list` to view the list of projects and their aliases in prack database.
8. For removal of project from prack database, run `prack remove <project_alias>`.

## Sample prack.yaml file

```yaml
---
project:
  name: Todo List
  alias: todo
  description: A todo list application
  tags:
    - react
    - flask
  commands:
    - alias: frontend
      commands:
        - yarn --cwd /home/ba1ry/os/todolist/frontend start
        - code /home/ba1ry/os/todolist
    - alias: backend
      commands:
        - python3 /home/ba1ry/os/todolist/backend/server.py
```

## Limitation

This project was built with a view that few repetitive shell commands like `cd` could be avoided typing everytime by storing them in prack.yaml. But it was later discovered that `cd` and few other commands are not executable commands on the file system and are internal builtin commands of the shell. This further limits the use of this project.
