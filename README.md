大家好，dtm最终跟原公司谈下来了知识产权转让，现已恢复维护，请大家访问[dtm](https://github.com/dtm-labs/dtf)。中间给大家带来的不便，敬请谅解！

![license](https://img.shields.io/github/license/dtm-labs/dtm)
![Build Status](https://github.com/dtm-labs/dtm/actions/workflows/tests.yml/badge.svg?branch=main)
[![codecov](https://codecov.io/gh/dtm-labs/dtm/branch/main/graph/badge.svg?token=UKKEYQLP3F)](https://codecov.io/gh/dtm-labs/dtm)
[![Go Report Card](https://goreportcard.com/badge/github.com/dtm-labs/dtm)](https://goreportcard.com/report/github.com/dtm-labs/dtm)
[![Go Reference](https://pkg.go.dev/badge/github.com/dtm-labs/dtm.svg)](https://pkg.go.dev/github.com/dtm-labs/dtm)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge-flat.svg)](https://github.com/avelino/awesome-go#database)

English | [简体中文](https://github.com/dtm-labs/dtm/blob/main/helper/README-cn.md)

# Distributed Transactions Manager

## What is DTM

DTM is a distributed transaction framework which provides cross-service eventual data consistency. It provides saga, tcc, xa, 2-phase message strategies for a variety of application scenarios. It also supports multiple languages and multiple store engine to form up a transaction as following:

<img alt="function-picture" src="https://en.dtm.pub/assets/function.7d5618f8.png" height=250 />

## Who's using DTM (partial)

[Tencent](https://www.tencent.com/)

[Ivydad](https://ivydad.com)

[Eglass](https://epeijing.cn)


## Features

* Language-agnostic
  - Suit for companies with multiple-language stacks.
    Easy to write bindings for Go, Python, PHP, Node.js, Ruby, and other languages.

* Support for multiple distributed transaction solutions
  - TCC, SAGA, XA, 2-phases message.

* Extremely easy to adapt
  - Support HTTP and gRPC, provide easy-to-use programming interfaces, lower substantially the barrier of getting started with distributed transactions. Newcomers can adapt quickly.

* Easy to use
  - Relieving developers from worrying about suspension, null compensation, idempotent transaction, and other tricky problems, the framework layer handles them all.

* Easy to deploy, easy to extend
  - DTM depends only on MySQL/Redis, easy to deploy, cluster, and scale horizontally.


## [Cook Book](https://en.dtm.pub)

## Quick start

### run dtm

``` bash
git clone https://github.com/dtm-labs/dtm && cd dtm
go run main.go
```

### Start an example
Suppose we want to perform an inter-bank transfer. The operations of transfer out (TransOut) and transfer in (TransIn) are coded in separate micro-services.

Here is an example to illustrate a solution of dtm to this problem:

``` bash
git clone https://github.com/dtm-labs/dtmcli-go-sample && cd dtmcli-go-sample
go run main.go
```

## Code

### Use
``` go
  // business micro-service address
  const qsBusi = "http://localhost:8081/api/busi_saga"
  // The address where DtmServer serves DTM, which is a url
  DtmServer := "http://localhost:36789/api/dtmsvr"
  req := &gin.H{"amount": 30} // micro-service payload
	// DtmServer is the address of DTM micro-service
	saga := dtmcli.NewSaga(DtmServer, dtmcli.MustGenGid(DtmServer)).
		// add a TransOut subtraction，forward operation with url: qsBusi+"/TransOut", reverse compensation operation with url: qsBusi+"/TransOutCom"
		Add(qsBusi+"/TransOut", qsBusi+"/TransOutCom", req).
		// add a TransIn subtraction, forward operation with url: qsBusi+"/TransIn", reverse compensation operation with url: qsBusi+"/TransInCom"
		Add(qsBusi+"/TransIn", qsBusi+"/TransInCom", req)
	// submit the created saga transaction，dtm ensures all subtractions either complete or get revoked
	err := saga.Submit()
```

When the above code runs, we can see in the console that services TransOut, TransIn has been called.

#### Timing diagram
A timing diagram for a successfully completed SAGA transaction would be as follows:

<img alt="saga-success" src="https://en.dtm.pub/assets/saga_normal.59a75c01.jpg" height=450/>

#### Rollback upon failure
If any forward operation fails, DTM invokes the corresponding compensating operation of each sub-transaction to roll back, after which the transaction is successfully rolled back.

Let's purposely fail the forward operation of the second sub-transaction and watch what happens

``` go
app.POST(qsBusiAPI+"/TransIn", func(c *gin.Context) {
  log.Printf("TransIn")
  // c.JSON(200, "")
  c.JSON(409, "") // Status 409 for Failure. Won't be retried
})
```

The timing diagram for the intended failure is as follows:

<img alt="saga-failed" src="https://en.dtm.pub/assets/saga_rollback.7989c866.jpg" height=550>

## More examples

Refer to [dtm-examples](https://github.com/dtm-labs/dtm-examples).

## Slack

You can join the [DTM slack channel here](https://join.slack.com/t/dtm-w6k9662/shared_invite/zt-vkrph4k1-eFqEFnMkbmlXqfUo5GWHWw).

## Give a star! ⭐

If you think this project is good, or helpful to you, please give a star!
