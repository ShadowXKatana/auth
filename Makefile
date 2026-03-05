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

.PHONY: migrate migrate-down migrate-info migrate-clean react-init

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

# ---------------------------------------------------------------------------
# React init template
# ---------------------------------------------------------------------------

## Copy React (Vite) template to a new project directory
## Usage: make react-init DEST=FE/REACT/my-new-app
react-init:
ifndef DEST
	$(error DEST is required. Usage: make react-init DEST=FE/REACT/my-new-app)
endif
	@echo "Scaffolding React project → $(DEST) …"
	@mkdir -p $(DEST)
	@cp -R FE/REACT/init/. $(DEST)/
	@echo "Done! cd $(DEST) && npm install && npm run dev"
