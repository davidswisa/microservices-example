#!/bin/bash

curl --location --request POST 'http://localhost:8080/reservations' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 1,
    "date": "string",
    "name": "string",
    "hour": 1,
    "party": 2
}'

