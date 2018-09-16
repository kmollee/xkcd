package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"sync"
	"syscall"

	"github.com/kmollee/xkcd"
)

var (
	logInfo = log.New(ioutil.Discard, "INFO: ", log.Ldate|log.Ltime)
	logErr  = log.New(os.Stderr, "Error: ", log.Ldate|log.Ltime)
)

// create file
func createFile(dirPath, filename string) (*os.File, error) {

	dir, err := os.Stat(dirPath)
	if err != nil {
		return nil, err
	}
	if !dir.IsDir() {
		return nil, fmt.Errorf("could not save, the path is not directory")
	}

	d, err := filepath.Abs(dirPath)
	if err != nil {
		return nil, err
	}

	filePath := path.Join(d, filename)
	f, err := os.Create(filePath)
	return f, err
}

// save comic []bytes to file
func saveImg(c *xkcd.Comic, dirPath string) error {
	f, err := createFile(dirPath, c.GetFilename())
	if err != nil {
		return err
	}

	err = c.SaveTo(f)
	return err
}

func main() {
	ID := flag.Int("id", -1, "comic <ID>")
	all := flag.Bool("all", false, "get all comic")
	last := flag.Bool("last", false, "get last comic")
	worker := flag.Int("work", 3, "worker to download")
	out := flag.String("out", "", "output direcotry path to save; default download to current diretory")
	verbose := flag.Bool("v", false, "verbose")
	flag.Parse()

	if *verbose {
		logInfo.SetOutput(os.Stdout)
	}

	if *worker < 1 {
		log.Fatalf("worker could not lower than 1")
	}

	outDir := *out
	if *out == "" {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		outDir = pwd
	}

	ctx := context.Background()
	// trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(ctx)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		cancel()
	}()

	switch {
	case *last:
		comic, err := xkcd.FetchLast()
		logInfo.Printf("Start download comic id %d\n", comic.ID)
		if err != nil {
			log.Fatal(err)
		}

		if err := saveImg(comic, outDir); err != nil {
			log.Fatal(err)
		}

	case *all:
		lastComic, err := xkcd.FetchLast()
		if err != nil {
			log.Fatal(err)
		}

		comicCh := make(chan int, *worker)
		var wg sync.WaitGroup

		for i := 1; i <= *worker; i++ {
			wg.Add(1)
			go func(ctx context.Context, epics <-chan int, outDir string) {
				defer wg.Done()
				comic := xkcd.NewComic()
				for {
					select {
					case <-ctx.Done():
						logInfo.Println("get interupt.. Stoping...")
						return
					case epic, _ := <-epics:
						logInfo.Printf("Start download comic id %d\n", epic)

						if err := comic.Update(epic); err != nil {
							logErr.Printf("\tcomic id %d fail: %v\n", epic, err)
							continue
						}
						if err := saveImg(comic, outDir); err != nil {
							logErr.Printf("\tcomic id %d fail save: %v\n", epic, err)
							continue
						}
					}
				}

			}(ctx, comicCh, outDir)
		}
		go func(ctx context.Context, ch chan<- int) {
			defer close(ch)
			for i := 1; i <= lastComic.ID; i++ {
				select {
				case <-ctx.Done():
					logInfo.Println("cancel input....")
					return
				default:
					ch <- i
				}
			}

		}(ctx, comicCh)

		wg.Wait()
		logInfo.Println("Complete all job....")
	case *ID > 0:
		// single comic
		comic := xkcd.NewComic()
		logInfo.Printf("Start download comic id %d\n", *ID)
		if err := comic.Update(*ID); err != nil {
			log.Fatal(err)
		}
		if err := saveImg(comic, outDir); err != nil {
			log.Fatal(err)
		}
	default:
		flag.PrintDefaults()
		os.Exit(0)
	}

}
