[![Build Status](https://travis-ci.org/plutov/playgo.svg?branch=master)](https://travis-ci.org/plutov/playgo)
[![Chat](https://img.shields.io/badge/gitter-dev_chat-46bc99.svg)](https://gitter.im/plutov/playgo)

Usually when we share a runnable Go code we do: copy code, open [Go Playground](https://play.golang.org/), paste code, click Share.

So `playgo` does it for you: open link in the browser and also copy it to clipboard.

### Installation and Usage

```
go get -u github.com/plutov/playgo/cmd/playgo
playgo helloworld.go
https://play.golang.org/p/v3rrZLwEUC (copied to clipboard)
```
