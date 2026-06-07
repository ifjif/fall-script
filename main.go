package main

import (
	"flag"

	"zzc/fall-script/src/cmd"
	"zzc/fall-script/src/repl"
)

func main() {
	cmd := cmd.Cmd{}
	flag.StringVar(&cmd.File, "f", "", "指定执行文件")
	flag.Parse()
	if cmd.File != "" {
		cmd.Execute()
	} else {
		repl.Repl()
	}
}
