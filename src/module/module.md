# 跨模块宏解析流程

    不含 .fsm
       读源文件
       得到 []byte
       词法分析
       语法分析
       提取ast的imports和exports,得到不含imports和exports的program
         imports、exports、program

       解析出export对应的具体的ast
         先从exports中
         含let，named function，直接 生成exportMeta存到exportMetas
         含ident，先占时存储 [name]idx (idx为第几个export导出，从0开始)

         如果有ident
         先从 program中查找，找到，生成exportMeta存到exportMetas

         将 exportMetas加入到exportMetaRegister中

         如果还有ident
         从 它imports中 查找，只对匹配的ident的import进行查找
           得到这个import的exportMetas信息
           如果没有找到ident, panic
           找到，加入到当前 exportMetas, 并标记为 re-export

        将exportMetas序列化为 .fsm文件

    含 .fsm
        解析 .fsm文件
        得到它的 exportMetas信息

    解析出imports对应的具体的ast
      对于每个import,得到它的exportMetas
        先从exportMetaRegister找，然后.fsm中，最后源文件
      然后和specifier进行匹配
      找到后，进行re-export处理，因为re-export中不含ast,(此操作用来找到ast并设置, 递归处理)

    提取出importMetas中的 macro信息，并从import中删除
    找到exportMetas中的 macro信息, 并从export中删除
    提取出program中的 macro信息，并从program中删除
    最后，对program进行 宏展开
