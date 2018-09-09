# XKCD

this is a mini project, download xkcd comic


## example

download xkcd first comic: id 1

```go
package main

import (
    "github.com/kmollee/xkcd"
    "log"
)

func main(){
    targetComicID := 1

    comic := xkcd.NewComic()
    if err := comic.Update(targetComicID); err != nil {
        log.Fatal(err)
    }
    if err := comic.SaveTo("./"); err != nil {
        log.Fatal(err)
    }
}
```

## build

```sh
go get github.com/kmollee/xkcd
cd $GOPATH/github.com/kmollee/xkcd/cmd && go build -o xkcd
./xkcd
```

## how to use it

```
Usage of ./xkcd:
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

