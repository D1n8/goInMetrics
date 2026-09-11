#!/usr/bin/sh

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
GO_PROJECT_MAIN_DIR="${SCRIPT_DIR}/backend"
BIN_DIR="${SCRIPT_DIR}/bin"

RED="\033[31m"
GREEN="\033[32m"
RESET="\033[0m"

fatal_exit() {
	timestamp=$(date +%D-%T)
	exit_message="$1"

	printf '[ %s ] | %b%s%b' "$timestamp" "$RED" "$exit_message" "$RESET"

	exit 1
}

if ! go version 2>/dev/null; then
	fatal_exit "Cannot find installed Go package."
fi

cd "$GO_PROJECT_MAIN_DIR"
if ! go build -o goinmetrics "./cmd/goinmetrics"; then
	fatal_exit "Backend build proccess error..."
fi

mkdir "${BIN_DIR}"
mv "${GO_PROJECT_MAIN_DIR}/goinmetrics" "${BIN_DIR}/goinmetrics"
chmod +x "${BIN_DIR}/goinmetrics"

