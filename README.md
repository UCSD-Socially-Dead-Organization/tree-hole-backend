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
