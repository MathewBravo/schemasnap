# SchemaSnap

**SchemaSnap** is a lightweight CLI tool to snapshot and track changes to database schemas over time.  
It helps developers and database administrators detect schema changes, generate diffs, and manage database evolution with ease.

---

## Features

- 📄 **Snapshot Schemas:** Quickly save the structure of a database as a file.
- 🔍 **Diff Snapshots:** Compare two snapshots and highlight added, removed, or modified schema elements.
- 🕒 **Timestamped Outputs:** Auto-generate snapshot filenames with timestamps.
- 🛠️ **Simple CLI Usage:** Minimal flags and clear commands.
- 🚀 **Fast and Expandable:** Start with PostgreSQL support, with plans for MySQL, SQLite, and MongoDB.

---

## Example Usage

Snapshot a database schema:

```bash
schemasnap snapshot postgres://user:password@localhost:5432/mydb --output ./snapshots
```

Diff two snapshots:

```bash
schemasnap diff ./snapshots/2025-04-10.sql ./snapshots/2025-04-17.sql
```

---

## Installation

_Coming soon._  
Once the first release is available, you will be able to install SchemaSnap easily via prebuilt binaries.

For now, clone and build locally:

```bash
git clone https://github.com/your-username/schemasnap.git
cd schemasnap
go build -o schemasnap ./cmd/schemasnap
```

---

## Roadmap

- [x] Snapshot PostgreSQL database schemas
- [x] Save snapshots with automatic timestamps
- [ ] Diff two schema snapshots
- [ ] Improve diff output formatting
- [ ] Add configuration file support (YAML/TOML)
- [ ] Support MySQL and SQLite
- [ ] CI/CD Integration (GitHub Actions plugin)
- [ ] Optional Web-based Visual Diff Viewer

---

## Contributing

Contributions are welcome!  
If you'd like to help, please open an issue, suggest an improvement, or submit a pull request.

---

## License

This project is licensed under the [MIT License](LICENSE).

---

## About

The goal is to make database schema management simpler, faster, and more reliable for developers and teams.

