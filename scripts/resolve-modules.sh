#!/usr/bin/env bash

# This is used by the linter action.
# Recursively finds all directories with a go.mod file and creates
# a GitHub Actions JSON output option.

set -o errexit

echo "Resolving modules in $(pwd)"

PROJECT_HOME=$(
	cd "$(dirname "${BASH_SOURCE[0]}")" &&
		cd .. &&
		pwd
)

source "${PROJECT_HOME}/scripts/util.sh"

FAILURE_FILE=${PROJECT_HOME}/scripts/.lintcheck_failures

all_modules=$(util::find_modules)
failing_modules=()
while IFS='' read -r line; do failing_modules+=("$line"); done < <(cat "$FAILURE_FILE")

echo "Ignored failing modules:"
echo "${failing_modules[*]}"
echo

PATHS=""

for mod in $all_modules; do
  echo "Checking module: $mod"
	util::array_contains "$mod" "${failing_modules[*]}" && in_failing=$? || in_failing=$?
	if [[ "$in_failing" -ne "0" ]]; then
		PATHS+=$(printf '{"workdir":"%s"},' ${mod})
	fi
done

echo "::set-output name=matrix::{\"include\":[${PATHS%?}]}"