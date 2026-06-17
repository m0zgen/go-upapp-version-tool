# go-upapp-version-tool

A small CLI utility for automatically increasing the patch version in version.txt.

The tool reads the current version from version.txt, increments the patch number, and writes the new version back to the same file.

Example:

`1.1.1`

after running the tool becomes:

`1.1.2`

## Usage

Create a version.txt file:

`echo "1.1.1" > version.txt`

Run the tool:

`./go-upapp-version-tool`

Example output:

```bash
──────────────────────────────
 Update Releaser Tool v1.1.1
──────────────────────────────
Current version: 1.1.1
New version: 1.1.2
```

## Version format

The version must use the following format:

`major.minor.patch`

Example:

`0.1.0`
`1.2.3`
`2.0.15`
`v1.0.0`

Increasing examples:

`1.1.1` -> `1.1.2`
`0.0.9` -> `0.0.10`
`v1.2.3` -> `v1.2.4`

The tool supports versions with or without the `v` prefix.

## Notes

This tool updates only the local version.txt file.

It does not:

- create Git tags
- commit changes
- push to GitHub
- run GoReleaser
- update go.mod

These steps can be added separately in a Makefile or CI pipeline.

## Example Makefile usage

```bash
APP := go-upapp-version-tool

build:
	go build -o $(APP) .

up-version:
	./$(APP)

release: up-version
	@echo "Release version updated"
```