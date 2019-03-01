## Get Started

Create Now account
=> https://zeit.co/now

Create Free mongodb
=> https://mlab.com/

Create Free Redis
=> https://redislabs.com/

```
// TODO: automate them
$ now secret add mongo-user "your-mongo-db-user"
$ now secret add mongo-pass "your-mongo-db-pass"
$ now secret add redis-pass "your-redis-pass"
```

Check secrets -> `$ now secret ls`

```
$ make deploy
```
