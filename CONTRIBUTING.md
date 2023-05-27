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
This should only need to be run once per repository, but can be run multiple times to override
any unwanted changes.

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
