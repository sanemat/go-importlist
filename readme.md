# importlist

List imported tools.

## Overview

`import-list -z tools.go` returns url list from tools.go for go install, split by NULL.

## Limitation

Not many options.

## Usage

```
import-list -z tools.go | xargs -0 -P 4 -I {} go install {}
```

## One more thing

Use [x-go-install](https://github.com/sanemat/go-xgoinstall/)

```
import-list -z tools.go | x-go-install -0
```

## Motivation

`cat tools.go | grep _ | awk -F'"' '{print $$2}'` is only for unix users.

## Install

### ghg

`ghg get sanemat/go-importlist`

### go install

`go install https://github.com/sanemat/go-importlist/cmd/import-list`


## Design

[design](./design.md)

## Changelog

[chagelog](./changelog.md) by [git-chglog](https://github.com/git-chglog/git-chglog)

## License

Copyright 2019 Matt (Sanemat) (Murahashi Kenichi)
[Apache License Version 2.0](./license.txt)

## Credits

[credits](./credits.txt) by [gocredits](https://github.com/Songmu/gocredits/)
