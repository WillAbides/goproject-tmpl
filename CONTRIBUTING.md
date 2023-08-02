# Contributing to goproject-tmpl

Your contributions are welcome here. Feel free to open issues and pull requests.
If you have a non-trivial change, you may want to open an issue before spending
much time coding, so we can discuss whether the change will be a good fit for
goproject-tmpl. But don't let that stop you from coding. Just be aware that
while all changes are welcome, not all will be merged.

## Releasing

Releases are automated
with [release-train](https://github.com/WillAbides/release-train). All PRs must
have a release label. See the release-train readme for more details.

## Scripts

goproject-tmpl uses a number of scripts to automate common tasks. They are found in the
`script` directory.

<!--- start script descriptions --->

### apply-template

script/apply-template updates the files in this repo based on the repository name.
This should only be run once, and it should be run by the apply-template GitHub Workflow.
It will be deleted after running.

### bindown

script/bindown runs bindown with the given arguments

### cibuild

script/cibuild is run by CI to test this project. It can also be run locally.

### fmt

script/fmt formats go code and shell scripts.

### generate

script/generate runs all generators for this repo.
`script/generate --check` checks that the generated files are up to date.

### goproject-tmpl

script/goproject-tmpl builds and runs the project with the given arguments.

### lint

script/lint runs linters on the project.

### new-from-template

script/new-from-template is for ad-hoc testing of this template.

### release-hook

script/release-hook is run by release-train as pre-tag-hook

### test

script/test runs tests on the project.

### update-docs

script/generate-readme updates documentation.
- For projects with binaries, it updates the usage output in README.md.
- Adds script descriptions to CONTRIBUTING.md.

### update-repo-settings

script/update-repo-settings updates the settings for the GitHub repository to the preferred settings.
This is only needed once per repository, so it deletes itself after running (actually it pushes
the change before updating the repository because the update disallows pushing changes to main).

This would normally be part of the one-time apply-template script, but it cannot be run
from GitHub Actions.

Settings changed:
- Projects: disabled
- Wiki: disabled
- Rebase merge: disabled
- Squash merge: disabled
- Always suggest updating pull request branches: enabled
- Automatically delete head branches: enabled
- Allow auto-merge: enabled
- Branch protection for "main":
  - Require a pull request before merging : enabled
  - Required status checks: "cibuild"

<!--- end script descriptions --->
