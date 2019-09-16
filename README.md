# test-xorm

## prepare

install mysql
- start local mysql(port is 3308)
- create mysql database handerly named fruit

```
docker run --restart=always --name  some-mysql -p 3308:3306 -v ./mysql:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=1234 -d mysql:5.7.22
```

## test json

```bash
$ go test
```
