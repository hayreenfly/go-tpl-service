#!/bin/bash

# Prompt for migration name
echo "Enter migration name:"
read MIGRATION_NAME

# Check if the user provided a name
if [ -z "$MIGRATION_NAME" ]; then
    echo "Migration name cannot be empty."
    exit 1
fi

# Run the migrate command
migrate create -ext sql -dir ./db/migrations "$MIGRATION_NAME"
