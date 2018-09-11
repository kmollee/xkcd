# XKCD

this is a mini project, download xkcd comic

## build

```
go install github.com/kmollee/xkcd/cmd/xkcdDownloader
```

## how to use it

```
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
```

## TODO

- update image's meta data

