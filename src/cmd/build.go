package cmd

import (
	"os"

	"zzc/fall-script/src/module"
)

/*
* 构建，保留代码的结构，生成 .fsc和.fsm文件
* 将其放入 开始文件所在的 .build目录中
*
 */
func (c *Cmd) DoBuild() {
	file := c.Build
	l := module.NewLoader()
	l.Build = true
	l.BuildProgram(".", file)
}

func (c *Cmd) DoClean() {
	os.RemoveAll(".build")
}
