# SchemaSnap# SchemaSnap

**SchemaSnap** is a lightweight CLI tool to snapshot and track changes to database schemas over time.  
It helps developers and database administrators detect schema changes, generate diffs, and manage database evolution with ease.

---

## Features

- 📄 **Snapshot Schemas:** Quickly save the structure of a database as a file (**placeholder implemented**).
- 🔍 **Diff Snapshots:** Compare two snapshots and highlight added, removed, or modified schema elements (**placeholder implemented**).
- 🛠️ **Interactive Init Command:** Set up database connection details through a TUI (Terminal User Interface) wizard.
- 🕒 **Timestamped Outputs:** (Planned for future releases.)
- 🚀 **Fast and Expandable:** Built in Go with a modular architecture. PostgreSQL first, other DBs coming soon.

---

## Example Usage

Snapshot a database schema (currently prints "Taking a snapshot"):

```bash
schemasnap my_database
```

Diff two snapshots (currently prints "Doing a diff"):

```bash
schemasnap diff snapshot1.sql snapshot2.sql
```

Initialize configuration (creates a `schemasnap` config folder and lets you input DB details):

```bash
schemasnap init
```

Show version:

```bash
schemasnap version
```

---

## Installation

_Coming soon._  
For now, clone and build locally:

```bash
git clone https://github.com/your-username/schemasnap.git
cd schemasnap
go build -o schemasnap ./cmd/schemasnap
```

---

## Current Status

✅ Root CLI with Cobra  
✅ Version command  
✅ Snapshot command (placeholder)  
✅ Diff command (placeholder)  
✅ Interactive `init` command using Bubbletea and Bubbles  
✅ Config directory creation  
🛠️ Snapshot actual schema structure (In Progress)  
🛠️ Real diffing of schema files (In Progress)

---

## Roadmap

- [x] Basic CLI with commands (snapshot, diff, init, version)
- [x] Interactive TUI for DB connection setup
- [ ] Implement PostgreSQL schema dumping
- [ ] Implement diff logic between two snapshot files
- [ ] Save snapshots with automatic timestamps
- [ ] Improve diff output formatting
- [ ] Add configuration file support (YAML/TOML)
- [ ] Add MySQL and SQLite support
- [ ] CI/CD Integration (GitHub Actions plugin)
- [ ] Optional Web-based Visual Diff Viewer

---

## Contributing

Contributions are welcome!  
If you'd like to help, please open an issue, suggest an improvement, or submit a pull request.


---

## About

The goal is to make database schema management simpler, faster, and more reliable for developers and teams.

---

### 📌 Quick Note

> **This project is in an early development phase.**  
> Commands are scaffolded but actual snapshot/diff functionality is still under construction.

