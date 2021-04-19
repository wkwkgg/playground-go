## setup golang

#### tools
 - [anyenv](https://github.com/anyenv/anyenv)
 - [goenv](https://github.com/syndbg/goenv)

#### how to install (mac)

```
# install goenv
$ anyenv install goenv

# check
$ goenv -v
goenv 2.0.0beta11

# show available versions
$ goenv install -l
...

# install
$ goenv install 1.16.3

# check installed versions
$ goenv versions

$ goenv global 1.16.3
$ goenv rehash
$ go version
go version go1.16.3 darwin/amd64
```

