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
	flag.BoolVar(&cmd.Dump, "c", false, "编译为字节码")
	flag.StringVar(&cmd.Build, "build", "", "进行构建，指定起始文件")
	flag.BoolVar(&cmd.Clean, "clean", false, "清理构建产物")
	flag.Parse()

	if cmd.Clean {
		cmd.DoClean()
	} else if cmd.Build != "" {
		cmd.DoBuild()
	} else if cmd.File != "" {
		if cmd.Engine == "eval" {
			cmd.Execute()
		} else {
			cmd.Interprete()
		}
	} else {
		if cmd.Engine == "eval" {
			repl.Repl()
		} else {
			repl.ReplVM()
		}
	}
}
