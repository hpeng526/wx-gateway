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

0. [wechat-sandbox](https://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=sandbox/login) set up your sandbox (or just using an account with template message interface) record <b>appID appsecret</b>
0. add a new template and record its <b>template id</b>
    ```
        {{first.DATA}} from：{{send.DATA}} 内容：{{text.DATA}} 时间：{{time.DATA}} {{remark.DATA}}
    ```
0. set up redis-server
0. run wx-backend
    0. create a file name <b>backend_config.json</b> and add
        ```json
            {
              "app_id": "your appID",
              "app_secret": "your appsecret",
              "mq_address": "127.0.0.1:6379",
              "key": "gateway",
              "delay": 10
            }
        ```
        mq_address is your redis ip and port
    0. just run <b>wx-backend_linux_amd64</b>
0. run wx-gateway
    0. create a file name <b>backend_config.json</b> and add
        ```json
            {
              "server_address" : "ip:port",
              "database": "./gateway.sqlite",
              "mq_address": "127.0.0.1:6379",
              "key": "gateway"
            }
        ```
        server_address is your server ip and port
        mq_address is your redis ip and port
    0. just run <b>wx-gateway_linux_amd64</b>
    0. set up gateway.sqlite with table.sql and add your users (no initialization
 and user interface to add users, maybe later)
0. you can use nohup or screen to run this apps

## How to build

run `make` or `make gateway` `make backend` to build specify app

### Warning

To build other platform's gateway, you must set up your own cross C compiler for linux (because of github.com/mattn/go-sqlite3)

## Some Client Sample
- [gateway-client](https://github.com/hpeng526/gateway-client)
