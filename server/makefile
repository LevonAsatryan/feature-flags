setup:
	# Check postgres in docker and run it
	# Figure out how to remove the password from here
	docker start ff-postgres || docker run --name ff-postgres -e POSTGRES_PASSWORD=pass -p 5432:5432 -d postgres
	# Target to check and create .env if it doesn't exist
	@if [ ! -f .env ]; then \
		echo ".env file not found. Creating from .env.example..."; \
		cp .env.example .env; \
	else \
		echo ".env file already exists."; \
	fi
	@echo "Environment setup complete."

run:
	go run .
