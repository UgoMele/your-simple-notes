# Your Simple Notes (ysn)

Just a simple cli notes manager

## Build

```bash
go build -o output-file-name main.go
```

## Usage

- `ysn` - Lists your most recently edited notes for quick access

- `ysn -a note-name` - creates a note in the default directory

- `ysn -a dir-name:note-name` - creates a note within a specific directory

## Directory structure

```
your-home/ysn-notes/
├── default/
│   ├── note1.md
│   └── note2.md
├── directory1/
│   ├── todo.md
│   └── ideas.md
└── directory2/
    └── list.md
```

License: [MIT](LICENSE)
