#!/bin/bash

# Hardcoded pool UUID
POOL_UUID="ef2b70d8-43a7-4f3f-9312-9131348c3f93"

# Execute the curl commands
curl localhost:8080/pool/finish -X POST
curl localhost:8080/winnings/distribute -X POST -d "{\"pool_uuid\": \"$POOL_UUID\"}"
curl localhost:8080/pool/create -X POST
