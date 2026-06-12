package main

import (
	"flag"

	"zzc/fall-script/src/cmd"
	"zzc/fall-script/src/repl"
)

func main() {
	cmd := cmd.Cmd{}
	flag.StringVar(&cmd.File, "f", "", "指定执行文件")
	flag.StringVar(&cmd.Engine, "engine", "vm", "指定执行引擎 -engine vm / -engine eval")
	flag.Parse()
	if cmd.File != "" {
		if cmd.Engine == "eval" {
			cmd.Execute()
		} else {
			cmd.Interprete()
		}
	} else {
		// repl.Repl()
		repl.ReplVM()
	}
}
