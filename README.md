# treehole-backend

# Description
你的樹洞，我來填滿

# Get started
```
cp config.env.example config.env  # Get full version from jeffrey.wu
make dev

make clean


0.0.0.0:8001

```

## 這個 boiler plate 連 DB 方式

先讀 config
```
conf, err := env.SetupConfig()
```

用 config 建立 db connection

```
gorm, err = database.DBConnection(conf.DB.GetDSN(), conf)
```

然後 migration (init)
```
repository.InitDB(gorm)
```

這邊 gorm 定義成類似於 connection

## How to test
TODO


## How to run matcher job
```
$ go run matcher/matcher.go

1. Pulls all active users
2. Shuffle list
3. Match users 2 by 2.

Edge cases to handle:
- What if active users list is odd, there will be one person left out. Solution: Match him with dummy bot user.
- Can users be rematched again?
```


## Fetch auth token from Auth0
All api request needs bearer token from Auth0. Please use `get_token.sh` script to generate JWT token for testing
```
$ ./scripts/get_token.sh

"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6InFMbF8zUFdwYmE5OVA1VllLaG93UiJ9.eyJjcmVkcy1lbWFpbCI6InRlc3R1c2VyQGdtYWlsLmNvbSIsImNyZWRzLW5pY2tuYW1lIjoidGVzdHVzZXIiLCJjcmVkcy1uYW1lIjoidGVzdHVzZXJAZ21haWwuY29tIiwiY3JlZHMtcGljdHVyZSI6Imh0dHBzOi8vcy5ncmF2YXRhci5jb20vYXZhdGFyL2ExOGJmNzg2ZWZiNzZhM2Q1NmVlNjlhM2IzNDM5NTJhP3M9NDgwJnI9cGcmZD1odHRwcyUzQSUyRiUyRmNkbi5hdXRoMC5jb20lMkZhdmF0YXJzJTJGdGUucG5nIiwiaXNzIjoiaHR0cHM6Ly9kZXYtMTJ3eHhydWJ5cjRtN2tkZy51cy5hdXRoMC5jb20vIiwic3ViIjoiYXV0aDB8NjYyYzgzYTBjYjEwODJiOGVhN2YwZWZmIiwiYXVkIjoiaHR0cHM6Ly90cmVlLWhvbGUtYmFja2VuZCIsImlhdCI6MTcxNDM2NjM0NiwiZXhwIjoxNzE0NDUyNzQ2LCJndHkiOiJwYXNzd29yZCIsImF6cCI6IkxwRGx6VmhRd2MxQlZBc3hmMmNEUVU5MEtvQkxxd1JIIn0.BMr2Dv-u0F-WJGqprKNw4I6vZqi__EHKgQkU2Fw6_-IYJOsTuBEG8y8Qh0ylYZget4_9Z1guAQibsuiNmi3rVPFW_NSj1nSR5f23Xkmi3OMr_JxIiPH-b96RSwuIOjONrXFc0KN0zQFP5QVtw0RANTCnZENGwiqtnftqlkniH9q8RZUONnZTQFbJrntNvlaOhLoaZT2SpMp0qucM2zdS2eCPfiST7-PTLQ4YqeQu_bZuzfSSDtBHVoaTdoTqgGd55Qc_lR8OWVR5G2k0l9YGXETINZGmpR3vQ0D_trjBS8PzV-fnFPz59u864qWzBMkGzCNC2hBKSN-ptVYcoAyt0Q"
```

## Auth0 authentication process for testing

1. Get your token
```bash
curl --request POST \
  --url https://dev-12wxxrubyr4m7kdg.us.auth0.com/oauth/token \
  --header 'content-type: application/json' \
  --data '{"client_id":"WHAT IS YOUR CLIENT ID? CHECK MESSENGER","client_secret":"WHAT IS YOUR CLIENT SECRET? CHECK MESSENGER","audience":"https://tree-hole-backend","grant_type":"client_credentials"}'

```

2. Example of an API Call
```bash
curl --request GET \
--url http://localhost:8001/v1/auth/health \
--header 'authorization: Bearer WHAT IS YOUR TOKEN FROM THE LAST STEP'
```
