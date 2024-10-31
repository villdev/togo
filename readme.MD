## Togo
CLI to manage todos. Stored locally in json file.

### Usage
Clone repo and cd into it

```bash
go build
```

Commands
1. `./togo` list all todos
2. `./togo -add <title>` add a new todo (or `-a`)
3. `./togo -complete <1>` mark #1 task as completed (or `-c`)
4. `./togo -redo <2>` mark #2 task as undone (or `-r`)
5. `./togo -del <2>` delete #2 task (or `-d`)
6. `./togo -add todo1 -add todo2 -del 2` exec multi flags
