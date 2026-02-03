# git-ac 🤖

> **AI-Powered Git Commit Generator**
> automatically generates meaningful commit messages using the Google Gemini API.

[![npm version](https://img.shields.io/npm/v/git-ac-cli.svg)](https://www.npmjs.com/package/git-ac-cli)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**`git-ac`** (Auto Commit/AI Commit) is a CLI tool that integrates natively with Git. It analyzes your staged changes (`git diff --cached`), sends them to a lightweight AI model, and generates a conventional commit message for you.

## 📦 Installation

Install globally via NPM to access it anywhere:

```bash
npm install -g git-ac-cli
```

## 🚀 Usage

```bash
git-ac commit
```

This will automatically stage all your unstaged changes and generate a commit message for everything.

```bash
git-ac commit --staged
```

If you prefer to manually stage your files with `git add`, you can use the `--staged` (or `-s`) flag.

## 🚀 Features

- **Native Git Integration:** Works as a git subcommand (`git ac`).
- **Smart Analysis:** Understands your code changes, not just file names.
- **Conventional Commits:** Generates standard messages (e.g., `feat: allow user to login`, `fix: resolve null pointer`).
- **Fast:** Optimized for speed using Go.
- **Cross-Platform:** Works on Windows, Linux, and macOS.

---
