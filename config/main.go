package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
)

var (
	help          = false
	gitRepository = ""
	outputDir     = ""
)

func main() {
	// parse args
	parse()
	if help || gitRepository == "" {
		flag.Usage()
		return
	}
	// use $HOME is outputDir is empty
	if outputDir == "" {
		if o, e := os.UserHomeDir(); e != nil {
			log.Fatalln(e)
		} else {
			outputDir = o
		}
	}
	// clone config to home dir if doesn't exist
	cmd := exec.Command("git", "clone", "--depth", "1", gitRepository, outputDir)
	log.Println(cmd.String())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
	cmd.Wait()
	// read config
}

func parse() {
	flag.BoolVar(&help, "h", false, "help")
	flag.StringVar(&gitRepository, "git", "", "eg: --git https://github.com/holicc/doom")
	flag.StringVar(&outputDir, "o", "", "default: $HOME")
	flag.Parse()
}
