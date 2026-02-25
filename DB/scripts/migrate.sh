#!/usr/bin/env bash
set -euo pipefail

# ---------------------------------------------------------------------------
# migrate.sh – run all pending up-migrations against the target database.
#
# Usage:
#   DB_URL="postgres://user:password@host:5432/dbname?sslmode=disable" \
#     ./DB/scripts/migrate.sh
#
# Required environment variables:
#   DB_URL   – PostgreSQL connection string
#
# Optional environment variables:
#   MIGRATIONS_DIR – path to the migrations directory
#                    (default: directory relative to this script)
# ---------------------------------------------------------------------------

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
MIGRATIONS_DIR="${MIGRATIONS_DIR:-${SCRIPT_DIR}/../migrations}"

if [[ -z "${DB_URL:-}" ]]; then
  echo "ERROR: DB_URL environment variable is not set." >&2
  echo "  Example: export DB_URL=\"postgres://user:pass@localhost:5432/mydb?sslmode=disable\"" >&2
  exit 1
fi

if ! command -v migrate &>/dev/null; then
  echo "ERROR: 'migrate' CLI not found. Install it from https://github.com/golang-migrate/migrate" >&2
  exit 1
fi

echo "Running migrations from: ${MIGRATIONS_DIR}"
echo "Target database:         ${DB_URL%%\?*}"

migrate -path "${MIGRATIONS_DIR}" -database "${DB_URL}" up

echo "Migration completed successfully."
