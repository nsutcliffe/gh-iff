# gh iff

A GitHub CLI extension to create multiple issues from a (CSV) file.

## Installation

```bash
gh extension install nsutcliffe/gh-issue-from-csv
```

## Usage

```bash
gh issue-from-csv --file <path-to-csv> --repo <owner/repo> [--header=false]
```

### Options

- `--file`: Path to the CSV file containing issue data (required)
- `--repo`: Repository in format owner/repo (required)
- `--header`: Whether the CSV file has a header row (default: true)

### CSV Format

The CSV file should have the following columns:

1. Title (required) - The issue title
2. Body (required) - The issue description/body
3. Labels (optional) - Comma-separated list of labels
4. Assignees (optional) - Comma-separated list of usernames to assign

Example CSV file:

```csv
Title,Body,Labels,Assignees
Fix login bug,Users cannot login on mobile devices,bug;high-priority,johndoe;janedoe
Add dark mode,Implement dark mode theme,enhancement,bobsmith
```

## Examples

Create issues from a CSV file with headers:
```bash
gh issue-from-csv --file issues.csv --repo octocat/Hello-World
```

Create issues from a CSV file without headers:
```bash
gh issue-from-csv --file issues.csv --repo octocat/Hello-World --header=false
```

## License

MIT 