# go-redis

GoでRedisを使う

## Redisの使い方

versionの確認

```sh
$ redis-cli -v
redis-cli 6.2.7
```

Redisに接続

--rawでマルチバイトに対応する

```sh
redis-cli --raw
```

setで値を登録して、getで値を取得する

```
127.0.0.1:6379> keys *

127.0.0.1:6379> set id:1 red
OK
127.0.0.1:6379> set id:2 blue
OK
127.0.0.1:6379> keys *
id:2
id:1
127.0.0.1:6379> get id:1
red
127.0.0.1:6379> flushall
OK
127.0.0.1:6379> keys *

```
