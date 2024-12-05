setup:
	# Check mongodb in docker and run it
	docker start mongodb || docker run --name mongodb -p 27018:27017 -d mongodb/mongodb-community-server
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
