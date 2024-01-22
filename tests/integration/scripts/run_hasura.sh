#!/bin/bash

# Function to check if Hasura is up and running
# check_hasura() {
#   curl --silent --fail http://localhost:8080/healthz
# }

# Function to track a table using the Hasura metadata API
track_table() {
  local db_name=$1
  local schema_name=$2
  local table_name=$3

  curl --silent --fail -X POST http://localhost:8080/v1/metadata \
       -H "Content-Type: application/json" \
       -H "X-Hasura-Admin-Secret: your-admin-secret" \
       --data-raw '{
         "type": "pg_track_table",
         "args": {
           "source": "'"${db_name}"'",
           "schema": "'"${schema_name}"'",
           "name": "'"${table_name}"'"
         }
       }'
}

sleep 5

graphql-engine serve &
pid=$!

sleep 20
# Wait for Hasura to be up (max 30 attempts, waiting 10 seconds between each)


echo "Hasura is up. Tracking tables..."

# Track tables using the API
# Replace 'postgres', 'public', and 'articles' with your actual database name, schema, and table name
track_table "postgres" "public" "Regions"
track_table "postgres" "public" "Projects"

echo "Tables tracked successfully."

# Keep the container running
wait $pid