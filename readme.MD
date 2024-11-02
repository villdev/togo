## Togo
CLI to manage todos. Stored locally in json.

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

Move executable to an exported path like `/usr/local/bin` to make the command `togo` available globally.

<img width="600" alt="Screenshot 2024-10-31 at 7 11 19 AM" src="https://github.com/user-attachments/assets/3fb81d78-1139-4aa2-a3e1-05a75daf29f4">
