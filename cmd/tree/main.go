package main

import (
	"flag"
	"fmt"
	"go-tree-demo/pkg/dir"
	"log"
	"os"
	"strings"
)

func main() {
	path := flag.String("p", "", "Set the path to build a tree")
	depth := flag.Int("d", 5, "Provide custom depth for dir tree")
	help := flag.Bool("h", false, "Show this help documentation")

	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		*path = args[0]
	}

	if *help || *path == "" {
		flag.PrintDefaults()
		return
	}

	// hack to support "~" for Win OS
	if strings.HasPrefix(*path, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		*path = strings.Replace(*path, "~", homeDir, 1)
	}

	treeDir := dir.Scan(*path, *depth)
	treeStr := dir.AsTreeStr(treeDir)
	fmt.Printf("Dir tree for: `%s`\n", *path)
	fmt.Println(treeStr)
}
