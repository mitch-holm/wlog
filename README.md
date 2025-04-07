# wlog

A terminal-based work logging application that helps you track and manage your tasks with importance levels.

## Features

- Log tasks with importance levels (1-5)
- Filter tasks by importance and time period
- Simple JSON-based storage
- Global command-line access
- Easy to use and extend

## Installation

### Prerequisites

- Go 1.16 or later
- Git

### Building from Source

1. Clone the repository:
```bash
git clone https://github.com/mholm/wlog.git
cd wlog
```

2. Build and install:
```bash
./build.sh
```

The script will:
- Build the application
- Install it to `~/bin`
- Add `~/bin` to your PATH if needed

## Usage

### Adding Tasks

```bash
# Add a task with default importance (3)
wlog add "Regular standup meeting"

# Add a high importance task
wlog add "Critical bug fix" -i 5

# Add a low importance task
wlog add "Update documentation" -i 2
```

### Listing Tasks

```bash
# List all tasks
wlog list

# List only important tasks (importance 4-5)
wlog list -m 4

# List tasks from this week
wlog list -p week

# List top 5 tasks
wlog list -t 5
```

### Getting Help

```bash
# Show detailed help
wlog help

# Show help for specific commands
wlog add --help
wlog list --help
```

## Data Storage

Tasks are stored in `~/.wlog/tasks.json`. The file is automatically created when you add your first task.

## Contributing

We welcome contributions! Here's how you can help:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Setup

1. Fork and clone the repository
2. Install dependencies:
```bash
go mod download
```

3. Make your changes
4. Run tests (when available)
5. Build and test locally:
```bash
./build.sh
```

### Project Structure

```
wlog/
├── cmd/
│   └── wlog/
│       └── main.go         # Application entry point
├── internal/
│   ├── cli/               # Command-line interface
│   ├── storage/           # Data storage
│   └── task/              # Task data structure
├── build.sh               # Build and install script
└── README.md              # This file
```

### Guidelines

- Follow Go best practices and style guidelines
- Write clear commit messages
- Add tests for new features
- Update documentation as needed
- Keep the code simple and maintainable

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [spf13](https://github.com/spf13) - For various Go utilities
