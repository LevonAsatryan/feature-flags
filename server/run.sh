#!/bin/bash
# create or run the docker image
if [ ! -f ".env" ]; then
	echo "Please make sure to edit the .env file!"
	echo "Copying the default .env file"
	cp "./.env.example" "./.env"
fi
echo "Initializing database"
docker compose up -d

SERVICE_NAME="postgres"
TIMEOUT=60

echo "waiting for service '$SERVICE_NAME' to be healthy..."

elapsed=0

until [ "$(docker inspect --format='{{.State.Health.Status}}' $(docker compose ps -q $SERVICE_NAME))" == "healthy" ]; do
	sleep 2
	elapsed=$((elapsed + 2))
	if [ $elapsed -ge $TIMEOUT ]; then
		echo "Timeout waiting for $SERVICE_NAME to become healthy"
		exit 1
	fi
done

echo "$SERVICE_NAME is healthy after $elapsed s. Trying to run the BE..."

echo "done"
echo "running the BE service"
go run .
exit 0
