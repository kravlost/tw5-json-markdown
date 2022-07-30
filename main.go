package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

var (
	flagWiki   *string
	flagOut    *string
	flagFilter *string
)

func init() {

}

func main() {

	// tiddlywiki "NodeJS WikiFolder" --output "temp output folder" --render "." "temp file.json" "text/plain" '$:/core/templates/exporters/JsonFile' "exportFilter" "[!is[system]]"

	flagWiki = flag.String("wiki", "", "TiddlyWiki folder")
	flagOut = flag.String("out", "", "Markdown output file")
	flagFilter = flag.String("filter", "[!is[system]]", "Selection filter")

	flag.Parse()

	if flagWiki == nil || *flagWiki == "" {
		log.Fatalln("No -wiki flag passed.")
	}

	if flagOut == nil || *flagOut == "" {
		log.Fatalln("No -out flag passed.")
	}

	fmt.Println("TiddlyWiki 5 NodeJS to Markdown Converter")
	fmt.Println("-----------------------------------------")
	fmt.Printf("Wiki:   %s\n", *flagWiki)
	fmt.Printf("Out:    %s\n", *flagOut)
	fmt.Printf("Filter: %s\n", *flagFilter)

	if err := os.Mkdir(*flagOut, 0755); os.IsExist(err) {
		fmt.Println("Output folder exists.")
	} else if err != nil {
		fmt.Printf("Error creating output folder: %v.\n", err)
	} else {
		fmt.Println("Output folder created.")
	}

	tempDir := os.TempDir()
	tempFile, err := os.CreateTemp(tempDir, "tw5-json")
	tempFileBase := path.Base(tempFile.Name())

	if err != nil {
		log.Fatalf("Error creating temp file: %v\n", tempFile)
	}

	fmt.Printf("Temp:   %s\n", tempFile.Name())

	tempFile.Close()

	//defer os.Remove(tempFile.Name())

	cmd := exec.Command("tiddlywiki", *flagWiki, "--output", tempDir, "--render", ".", tempFileBase, "text/plain", "$:/core/templates/exporters/JsonFile", "exportFilter", *flagFilter)
	fmt.Println(cmd.String())
	_, err = cmd.Output()

	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			log.Fatalln("failed executing:", err)
		case *exec.ExitError:
			log.Fatalln("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	json, err := os.ReadFile(tempFile.Name())

	if err != nil {
		log.Fatalf("Error reading JSON: %v", err)
	}

	err = ProcessJson(string(json), *flagOut)

	if err != nil {
		log.Fatalf("Conversion error: %v\n", err)
	}

	fmt.Println("Done.")
}
