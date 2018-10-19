# XKCD

this is a mini project, download xkcd comic

## install

```
go get -v github.com/kmollee/xkcd/...
```

## how to use it

```sh
$ xkcdDownloader -h
Usage of xkcdDownloader:
  -all
        get all comic
  -id int
        comic <ID> (default -1)
  -last
        get last comic
  -out string
        output direcotry path to save; default download to current diretory
  -v    verbose
  -work int
        worker to download (default 3)

# download xkcd #1
$ xkcdDownloader -id 1 -out ./

# download all xkcd comic with 15 worker
$ xkcdDownloader -v -all -out ~/mycomic/xkcd/ -work 15

# download latest comic
$ xkcdDownloader -last -out ~/mycomic/xkcd/

```

## TODO

- update image's meta data

## Documentation

View Go Doc [online](https://godoc.org/github.com/kmollee/xkcd).

To view go docs locally, after installing the package run:

```
$ godoc -http=:6060
```

