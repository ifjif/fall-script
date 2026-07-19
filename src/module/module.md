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

# 跨模块struct解析流程

    不含 .fsm
    词法分析, 语法分析
    收集到所有imports和exports,struct,methods

    解析 exports
    将所有的 struct export加入到 exportMetas中
    当添加  exportMeta时，如果是struct 会额外记录
          exportStruct := map[string]*ast.StructDeclStmt{}
    这时加入的exportMeta是没有

        Methods    []string // struct AST
        FieldTotal int      // struct AST
        Fields     []int    // struct AST 所有字段的offset
     这些信息的
     如果是 ident，依然进行标记
     之后去处理 ident
        先从 program中找，是struct,加入
        然后 判断导出的 本文件中struct, 是否组合了 import中的
              加入 importStruct := map[string]bool{}
        存在对 import的组合，然后就去 import中找 相应的struct
        加入到：
            importStructMeta := map[string]*ExportMeta{}

    然后，处理所有 export 的 struct
    需要
      exports: 所有exportMeta(需要这个，是判断这export struct是否是重导出了，是不进行处理)
      exportStructs: 所有的 export structs
      structs: 当前文件中的所有structs
      methods: 当前文件中所有struct的方法
      importStructs: export struct中组合了外部导入的 struct

    ***********************************
    ** 注意：发现循环组合，直接panic **
    ***********************************

    遍历 exportStructs, 得到它的字段总数，每个字段的offset，它的所有方法名
      在exports中找，如果是重导出的，不进行处理
      不是，进行处理:
        计算字段总数和每个字段的offset
          遍历 它的fields
          如果不是组合，它的offset为offset,然后offset++
          是组合：
            先从本地找，是否是本地的struct,是，递归进行处理，得到total，offset = total
            不是本地，从importStruct中找，取FieldTotal, offset = offset+FieldTotal
          最后返回 offset, offsets
        然后收集方法名,只需要直接的
        从methods中拿到这个struct的methodDecl
          进行收集
      最后给exportMeta增加:
            FieldTotal
            Fields
            Methods
