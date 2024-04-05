# API Design

## User
```
POST /user
GET /user
GET /user/:id
PUT /user/:id
DELETE /user/:id  # use tombstone, mark as inactive.

{
    id
    username
    age
    profile_pic
    last_login
}
```

## Friend
```
POST /user/:id/friend
    or
POST /user/friend - generic add friend API

GET /user/:id/friend - Get all friends for given user
```

## Match
```
POST /match
GET /match
GET /match/:match_id

Match is append only. No update
```

## Post
```
POST /post
GET /post
GET /post/:id
PUT /post/:id
DELETE /post/:id  # tombstone, mark as unavaliable
```
