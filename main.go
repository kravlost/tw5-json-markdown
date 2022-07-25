package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	flagIn  *string
	flagOut *string
)

func init() {

}

func main() {

	flagIn = flag.String("in", "", "JSON input file")
	flagOut = flag.String("out", "", "Markdown output file")

	flag.Parse()

	if flagIn == nil || *flagIn == "" {
		log.Fatalln("No -in flag passed.")
	}

	if flagOut == nil || *flagOut == "" {
		log.Fatalln("No -out flag passed.")
	}

	fmt.Println("TiddlyWiki 5 JSON to Markdown Converter")
	fmt.Println("---------------------------------------")
	fmt.Printf("In:  %s\n", *flagIn)
	fmt.Printf("Out: %s\n", *flagOut)

	if err := os.Mkdir(*flagOut, 0755); os.IsExist(err) {
		fmt.Println("Output folder exists.")
	} else if err != nil {
		fmt.Printf("Error creating output folder: %v.\n", err)
	} else {
		fmt.Println("Output folder created.")
	}

	json, err := os.ReadFile(*flagIn)

	if err != nil {
		log.Fatalf("Error reading JSON: %v", err)
	}

	err = ProcessJson(string(json), *flagOut)

	if err != nil {
		log.Fatalf("Conversion error: %v\n", err)
	}

	fmt.Println("Done.")
}
