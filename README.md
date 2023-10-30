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

To configure `rb`, place a `config.yaml` file in `~/.config/rb/`. This file defines the rules for cloning and organising repositories. For a comprehensive guide on crafting your configuration, please refer to [/examples/config.md](/examples/config.md).

## Contributions 👨‍💻👩‍💻
Feel free to submit pull requests, issues, or feature requests. For major changes, kindly open an issue first.

## License 📜
This project is licensed under the MIT License.

## Let's Automate the Mundane! 🤖
Why click and drag when you can rb and brag? Make your life simpler, one repo at a time.