include secrets.env

.EXPORT_ALL_VARIABLES:
AUTH0_CLIENT_ID ?= ${Auth0_Client_id}
AUTH0_DOMAIN ?= ${Autho0_Domain}
AUTH0_CLIENT_SECRET ?= ${Auth0_Client_secret}
# this needs to be updated to the correct url
AUTH0_CALLBACK_URL?=http://localhost:3000/api/auth/callback

SERVER_PORT ?= 3000

run:
	go run main.go