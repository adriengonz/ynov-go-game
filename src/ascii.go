package main

import (
	"github.com/TheZoraiz/ascii-image-converter/aic_package"
	"time"
	"fmt"
)

func img(img_source string) {
	filePath := img_source
	flags := aic_package.DefaultFlags()
	flags.Colored = true
	asciiArt, err := aic_package.Convert(filePath, flags)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", asciiArt)
	time.Sleep(2 * time.Second)
}