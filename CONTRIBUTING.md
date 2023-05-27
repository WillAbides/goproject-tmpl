# Contributing to goproject-tmpl

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

### release

script/release creates a new release. It is run by GitHub Actions on push to main.

### test

script/test runs tests on the project.

### update-docs

script/generate-readme updates README.md with the usage output of script/goproject-tmpl and
updates CONTRIBUTING.md with script documentation.

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

## Releasing

Releases are automated with GitHub Actions. The release workflow runs on every push to main and determines the version
to release based on the labels of the PRs that have been merged since the last release. The labels it looks for are:

| Label           | Change Level |
|-----------------|--------------|
| breaking        | major        |
| breaking change | major        |
| major           | major        |
| semver:major    | major        |
| bug             | patch        |
| enhancement     | minor        |
| minor           | minor        |
| semver:minor    | minor        |
| bug             | patch        |
| patch           | patch        |
| semver:patch    | patch        |
