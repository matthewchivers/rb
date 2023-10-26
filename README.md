# `rb` - The Repository Butler

## Description

Repo Butler (`rb`) is a Command-Line Interface (CLI) tool designed to manage code repositories based on a set of configured rules. Kiss manual organisation goodbye; let `rb` take the wheel.

## Features ✨

1. **Cloning of Repositories**:
   - Clone git repositories to configured directories using the `rb clone [git_repo]` command.

2. **Rule-Based Sorting**:
   - Define rules in the configuration file to manage the location and structure of cloned repositories.
   - Specify custom nesting patterns to organise repositories in a manner that suits your workflow (currently supports `host` and `owner` tags)

## Roadmap 🚀

1. **Opening Managed Repositories in an Editor**:
   - Develop a feature to open managed repositories directly in a specified editor.

## Usage 🎮

```bash
# Clone a single repository
rb clone [git_repo]
```

## Configuration 🛠️

Define rules in a `config.yaml` file located at `~/.config/rb/`. Each rule should specify a target directory, and can include a nesting pattern and matching criteria for repositories.

### Examples:

```yaml
# Example 1: Clone "https://github.com/matthewchivers/rb.git" to "~/code/github.com/matthewchivers/rb"
rules:
  - name: "example rule"
    directory: "~/code"
    nesting:
      pattern: "{host}/{owner}"
    match:
      repo:
        owner: "matthewchivers"

# Example 2: Default rule when no other rule matches
rules:
  - name: "default rule"
    directory: "~/code"
    default: true

# Example 3: Rule assessment (First-match basis)
# Any repo named `rb` will be cloned to `~/rb-project` as it's the first matching rule.
rules:
  - name: "rb rule"
    directory: "~/rb-project"
    match:
      repo:
        name: "rb"
  - name: "another rb rule"
    directory: "~/rb-newproject"
    match:
      repo:
        name: "rb"
```
Rules are evaluated on a first-match basis. In Example 3, any repo named rb will always match the first rule, hence it will be cloned to ~/rb-project and not ~/rb-newproject.

## Contributions 👨‍💻👩‍💻
Feel free to submit pull requests, issues, or feature requests. For major changes, kindly open an issue first.

## License 📜
This project is licensed under the MIT License.

## Let's Automate the Mundane! 🤖
Why click and drag when you can rb and brag? Make your life simpler, one repo at a time.