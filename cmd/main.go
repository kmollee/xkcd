package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/kmollee/xkcd"
)

var (
	Info  = log.New(ioutil.Discard, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stdout, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func main() {
	comicID := flag.Int("id", -1, "comic <ID>")
	all := flag.Bool("all", false, "get all comic")
	last := flag.Bool("last", false, "get last comic")
	worker := flag.Int("work", 3, "worker to download")
	out := flag.String("out", "", "output direcotry path to save; default download to current diretory")
	verbose := flag.Bool("v", false, "verbose")
	flag.Parse()

	if *verbose {
		Info.SetOutput(os.Stdout)
	}

	outDir := *out
	if *out == "" {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		outDir = pwd
	}

	switch {
	case *last:
		comic, err := xkcd.FetchLast()
		Info.Printf("Start download comic id %d\n", comic.ID)
		if err != nil {
			log.Fatal(err)
		}
		comic.SaveTo(outDir)
	case *all:
		lastComic, err := xkcd.FetchLast()
		if err != nil {
			log.Fatal(err)
		}

		comicCh := make(chan int, *worker)
		var wg sync.WaitGroup

		for i := 1; i <= *worker; i++ {
			wg.Add(1)
			go func(job <-chan int) {
				comic := xkcd.NewComic()
				for comicID := range job {
					Info.Printf("Start download comic id %d\n", comicID)
					if err := comic.Update(comicID); err != nil {
						Error.Printf("\tcomic id %d fail: %v\n", comicID, err)
						continue
					}
					if err := comic.SaveTo(outDir); err != nil {
						Error.Printf("\tcomic id %d fail save: %v\n", comicID, err)
						continue
					}
				}
				wg.Done()
			}(comicCh)
		}
		for i := 1; i <= lastComic.ID; i++ {
			comicCh <- i
		}
		close(comicCh)
		wg.Wait()
		Info.Println("Complete all job....")
	case *comicID > 0:
		// single comic
		comic := xkcd.NewComic()
		Info.Printf("Start download comic id %d\n", comicID)
		if err := comic.Update(*comicID); err != nil {
			log.Fatal(err)
		}
		if err := comic.SaveTo(outDir); err != nil {
			log.Fatal(err)
		}
	default:
		flag.PrintDefaults()
		os.Exit(0)
	}

}
