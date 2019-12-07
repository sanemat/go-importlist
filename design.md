# import-list

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
