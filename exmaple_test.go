package xkcd_test

import (
	"fmt"
	"log"

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

func ExampleComic_SaveTo() {
	firstComicID := 1
	c := xkcd.NewComic()
	if err := c.Update(firstComicID); err != nil {
		log.Fatalf("could not update comic: %v", err)
	}
	saveDirectory := "./"
	if err := c.SaveTo(saveDirectory); err != nil {
		log.Fatalf("could not save comic: %v", err)
	}
}

func ExampleComic_SaveTo_errPermissionDenied() {
	// try to save image to direcotry that user don't permission
	firstComicID := 1
	c := xkcd.NewComic()
	if err := c.Update(firstComicID); err != nil {
		log.Fatalf("could not update comic: %v", err)
	}
	saveDirectory := "/"
	if err := c.SaveTo(saveDirectory); err != nil {
		fmt.Printf("%v\n", err)
	}

	// Output:
	// open /1_2006_1_1.png: permission denied
}
