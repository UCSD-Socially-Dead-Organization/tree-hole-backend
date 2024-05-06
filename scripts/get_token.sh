#!/bin/bash

source config.env

curl --request POST --silent \
     --url "https://${AUTH0_DOMAIN}/oauth/token" \
     --header 'content-type: application/x-www-form-urlencoded' \
     --data grant_type=password \
     --data "username=${TEST_USERNAME}" \
     --data "password=${TEST_PASSWORD}" \
     --data "audience=${AUTH0_AUDIENCE}" \
     --data scope=read:sample \
     --data "client_id=${CLIENT_ID}" \
     --data "client_secret=${CLIENT_SECRET}" \
    | jq .access_token
