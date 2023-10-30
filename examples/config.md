# Config Files

## Example Files for `rb` - The Repository Butler

This document contains explanations of configuration files to help you get started with `rb`. The focus here is on utilising clone rules for various scenarios.

### Table of Contents

1. [Single Clone Rule](#single-clone-rule)
2. [Multiple Clone Rules](#multiple-clone-rules)

---

## Dissection of an `rb` Config File

Here, we break down the various components of the example `rb` config file to better understand its structure and functionality.

### Config Example

```yaml
directories:
  - name: "my-projects"
    path: "/home/my/projects"

clone_rules:
  - name: "my-first-rule"
    default: true
    directory_name: "my-projects"
    nesting_pattern: "{host}/{owner}"
    match:
      repository:
        name: "repo-name"
        owner: "my-github-account"
        host: "github.com"
```

### Sections

#### `directories`

- `name`: Specifies a unique identifier for the directory. In this case, it's "my-projects". Used as a reference to a directory later.
- `path`: Provides the full system path where the repositories will be cloned. Here, it's "/home/my/projects".

#### `clone_rules`

- `name`: Uniquely identifies the clone rule. Here, the name is "my-first-rule". Used when manually specifying rules (`--rule=`)
- `default`: Indicates whether this is the default rule. If set to `true`, it will be applied in the absence of any matching rule.
- `directory_name`: Refers back to the `name` of the directory defined in `directories`. It tells `rb` where to place cloned repos based on this rule.
- `nesting_pattern`: (_optional_) Determines the folder structure under which the repository will be cloned. Utilises placeholders like `{host}` and `{owner}`.
  > Only supply a pattern if wishing to nest within the specified directory.
- `match`
  - `repository`: Specifies filtering criteria for repositories. Further explanation [below](#single-clone-rule).

With the config above, a repository (url: `https://github.com/my-github-account/repo-name.git`) that matches the rule will be cloned into `/home/my/projects/github.com/my-github-account/repo-name`.


---

### Single Clone Rule

#### Single clone rule.

```yaml
directories:
  - name: "go-projects"
    path: "/home/user/go"

clone_rules:
  - name: "go-repos"
    default: true
    directory_name: "go-projects"
    nesting_pattern: "{host}/{owner}"
    match:
      repository:
        name: "example"
        owner: "me"
        host: "github.com"
```

This will clone any repository named `example` by owner `me` on `github.com` into `/home/user/go/github.com/me/`.

#### Single Clone Rule - Simple Rule
An example illustrating a single clone rule matching just one repo element

```yaml
directories:
  - name: "home"
    path: "/home/user"

clone_rules:
  - name: "personal-repos"
    default: true
    directory_name: "home"
    match:
      repository:
        owner: "me"
```

This will clone any repository with owner `me` on into `/home/user/me/`.

---

### Multiple Clone Rules

#### `multiple-clone-rules-config.yaml`

This example demonstrates how to set up multiple clone rules.

```yaml
directories:
  - name: "serious-dir"
    path: "/home/user/projects"
  - name: "hobby-dir"
    path: "/home/user/playground"
  - name: "home"
    path: "/home/user"

clone_rules:
  - name: "docs"
    directory_name: "home"
    match:
      repository:
        owner: "foo"
  - name: "serious"
    directory_name: "serious-dir"
    match:
      repository:
        owner: "bar" # Same as rule below
  - name: "hobby"
    directory_name: "hobby-dir"
    default: true
    match:
      repository:
        owner: "bar" # Same as rule above
```

With multiple rules, `rb` applies the first matching rule. If a repository is owned by `bar`, it goes into `/home/user/projects` because "serious" is the first matching rule, even if another rule has `default: true`.
