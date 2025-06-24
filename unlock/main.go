package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gen2brain/dlgs"
)

var sourceFilePath = flag.String("source", "", "source file path")
var destFilePath = flag.String("dest", "", "destination file path")

func main() {
	flag.Parse()
	sourceFilePath := *sourceFilePath
	destFilePath := *destFilePath
	err := os.Rename(sourceFilePath, destFilePath)
	if err != nil {
		dlgs.Error("错误", fmt.Sprintf("Rename Error:%s---->%s", sourceFilePath, destFilePath))
		os.Exit(1)
	}
}
