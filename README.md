api-gateway
====

[![Build Status](https://travis-ci.org/hpeng526/wx-gateway.svg?branch=master)](https://travis-ci.org/hpeng526/wx-gateway)
[![Go Report Card](https://goreportcard.com/badge/github.com/hpeng526/wx-gateway)](https://goreportcard.com/report/github.com/hpeng526/wx-gateway)
### Dependency

- github.com/hpeng526/wx
- github.com/garyburd/redigo
- github.com/mattn/go-sqlite3

## What is api-gateway ?

It is a interface using [WeChat template message api]('https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1433751277') to send custom message to the user or user group.

## How to use ?

1. [wechat-sandbox](https://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=sandbox/login) set up your sandbox (or just using an account with template message interface) record <b>appID appsecret</b>
2. add a new template and record its <b>template id</b>

```
{{first.DATA}} from：{{send.DATA}} 内容：{{text.DATA}} 时间：{{time.DATA}} {{remark.DATA}}
```

3. set up redis-server
4. run wx-backend
    1. create a file name <b>backend_config.json</b> (mq_address is your redis ip and port) and run <b>wx-backend_linux_amd64</b>

    ```json
    {
    "app_id": "your appID",
    "app_secret": "your appsecret",
    "mq_address": "127.0.0.1:6379",
    "key": "gateway",
    "delay": 10
    }
    ```

5. run wx-gateway
    1. create a file name <b>backend_config.json</b> (server_address is your server ip and port, mq_address is your redis ip and port) and run <b>wx-gateway_linux_amd64</b>

    ```json
    {
    "server_address" : "ip:port",
    "database": "./gateway.sqlite",
    "mq_address": "127.0.0.1:6379",
    "key": "gateway"
    }
    ```

    2. set up gateway.sqlite with table.sql and add your users (no initialization
 and user interface to add users, maybe later)
6. you can use nohup or screen to run this apps

## How to build

run `make` or `make gateway` `make backend` to build specify app

### Warning

To build other platform's gateway, you must set up your own cross C compiler for linux (because of github.com/mattn/go-sqlite3)

## Some Client Sample
- [gateway-client](https://github.com/hpeng526/gateway-client)
