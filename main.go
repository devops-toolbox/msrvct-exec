/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/devops-toolbox/msrvct-exec/cmd"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	cmd.Execute()
}
