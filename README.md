# GCLI - AI-Powered Git Commit Messages

[![Go Report Card](https://goreportcard.com/badge/github.com/misbahulhoq/gcm)](https://goreportcard.com/report/github.com/misbahulhoq/gcm)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**GCM** is a command-line interface (CLI) tool that uses the power of Google's Gemini AI to automatically generate clear, meaningful, and conventional commit messages for your Git repositories. Stop spending time trying to word the perfect commit message and let AI do it for you!

## Features

- **AI-Generated Commit Messages**: Leverages Google's Gemini to create high-quality commit messages based on your code changes.
- **Conventional Commits**: Generates messages that follow the [Conventional Commits](https://www.conventionalcommits.org/) specification.
- **Flexible Committing**: Commit all your changes at once or only the ones you've staged.
- **Interactive Confirmation**: Always asks for your approval before making a commit.
- **Easy Setup**: Simple API key configuration.

## Installation

1.  **Install Go**: Make sure you have Go (version 1.16 or later) installed on your system.
2.  **Install GCM**:
    ```sh
    go install github.com/misbahulhoq/gcm@latest
    ```
    This will install the `gcm` binary in your `$GOPATH/bin` directory. Make sure this directory is in your system's `PATH`.

## First-Time Setup: API Key

The first time you run `gcm`, it will prompt you for a Google Gemini API key.

1.  Get your free API key from [Google AI Studio](https://aistudio.google.com/app/apikey).
2.  Run any `gcm` command (e.g., `gcm commit`).
3.  Paste the key when prompted.

The key will be saved securely in a `.gcm_config` file in your home directory for future use.

Alternatively, you can set the `GEMINI_API_KEY` environment variable.

## Usage

### `gcm commit`

This is the primary command. It analyzes your file changes, generates a commit message, and prompts you to confirm the commit.

**Commit all changes (default):**

This will automatically stage all your unstaged changes and generate a commit message for everything.

```sh
gcm commit
```

**Commit only staged changes:**

If you prefer to manually stage your files with `git add`, you can use the `--staged` (or `-s`) flag.

```sh
git add <your-files>
gcm commit --staged
```

The tool will show you the proposed commit message and ask for confirmation before proceeding.

### `gcm healthcheck`

A simple utility to verify that you are in a Git repository.

```sh
gcm healthcheck
```

Aliases: `hc`, `check`

## Contributing

Contributions are welcome! If you have ideas for new features, bug fixes, or improvements, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
