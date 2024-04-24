# treehole-backend

# Description
你的樹洞，我來填滿

# Get started
```
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

# How to test
TODO


# Auth0 authentication process for testing

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