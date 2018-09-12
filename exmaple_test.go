package xkcd_test

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/kmollee/xkcd"
)

func ExampleComic_Update() {
	firstComicID := 1
	c := xkcd.NewComic()
	if err := c.Update(firstComicID); err != nil {
		log.Fatalf("could not update comic: %v", err)
	}

	fmt.Println(c.ID)
	fmt.Println(c.Title)
	// Output:
	// 1
	// Barrel - Part 1
}

func ExampleComic_Update_updateError() {
	firstComicID := -1
	c := xkcd.NewComic()
	if err := c.Update(firstComicID); err != nil {
		fmt.Printf("%v", err)
	}

	// Output:
	// reponse is not ok: 404 Not Found
}

func ExampleComic_SaveTo_buffer() {
	firstComicID := 1
	c := xkcd.NewComic()

	if err := c.Update(firstComicID); err != nil {
		log.Fatalf("could not update comic: %v", err)
	}

	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)

	if err := c.SaveTo(writer); err != nil {
		log.Fatalf("could not save comic: %v", err)
	}
}

func ExampleComic_SaveTo_file() {
	firstComicID := 1
	c := xkcd.NewComic()

	if err := c.Update(firstComicID); err != nil {
		log.Fatalf("could not update comic: %v", err)
	}

	f, err := os.Create("somefile.png")
	if err != nil {
		log.Fatal(f)
	}

	if err := c.SaveTo(f); err != nil {
		log.Fatalf("could not save comic: %v", err)
	}
}
