# Sample of distributed system for votes collection

## first install

1. NSQ - messaging queue (after installing get drivers: go get github.com/bitly/go-nsq)
2. MongoDB - NoSQL DB (get drivers: go get github.com/mongodb/mongo-go-driver)

## run nsq

1. Use command: nsqlookupd
2. nsqd --lookupd-tcp-address=localhost:4160

## create new app on twitter

```
https://developer.twitter.com/en/apps
```
