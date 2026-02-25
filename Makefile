# ---------------------------------------------------------------------------
# Flyway migration targets
#
# Required environment variables (can be set in a .env file or exported):
#   FLYWAY_URL      – JDBC URL, e.g. jdbc:postgresql://localhost:5432/mydb
#   FLYWAY_USER     – database user
#   FLYWAY_PASSWORD – database password
#
# Optional:
#   FLYWAY_LOCATIONS – defaults to filesystem:DB/migrations
# ---------------------------------------------------------------------------

FLYWAY          ?= flyway
FLYWAY_LOCATIONS ?= filesystem:DB/migrations

.PHONY: migrate migrate-down migrate-info migrate-clean

## Apply all pending versioned migrations
migrate:
	$(FLYWAY) \
		-url="$(FLYWAY_URL)" \
		-user="$(FLYWAY_USER)" \
		-password="$(FLYWAY_PASSWORD)" \
		-locations="$(FLYWAY_LOCATIONS)" \
		migrate

## Undo the last applied versioned migration (requires Flyway Teams/Enterprise)
migrate-down:
	$(FLYWAY) \
		-url="$(FLYWAY_URL)" \
		-user="$(FLYWAY_USER)" \
		-password="$(FLYWAY_PASSWORD)" \
		-locations="$(FLYWAY_LOCATIONS)" \
		undo

## Print the status of all migrations
migrate-info:
	$(FLYWAY) \
		-url="$(FLYWAY_URL)" \
		-user="$(FLYWAY_USER)" \
		-password="$(FLYWAY_PASSWORD)" \
		-locations="$(FLYWAY_LOCATIONS)" \
		info

## Drop all objects in the configured schemas (USE WITH CAUTION)
migrate-clean:
	$(FLYWAY) \
		-url="$(FLYWAY_URL)" \
		-user="$(FLYWAY_USER)" \
		-password="$(FLYWAY_PASSWORD)" \
		-locations="$(FLYWAY_LOCATIONS)" \
		-cleanDisabled=false \
		clean
