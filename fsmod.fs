/*
 * 导入语法
 * import {a,b} from "./a.fs"
 * let a = 1 // 覆盖了导入的a, 不会使用a.fs中的a
 * import {a as aa, b as bb} from "./a.fs"  // 重命名
 * import * as math from "./math.fs"        // 命名空间导入
 *
 * let math = import("./math.fs")           // 动态导入 未来支持
 *
 *
 * 导出语法
 * export fn fib(xx) {xxx}
 * export let PI = 3.1415926
 *
 * export {  // 未来支持
 *  "a":"a",
 *  "b":"b"
 * }
 *
 *
 * import ast
 * {
 *   token import
 *   specifiers {
 *       local     string
 *       imported  string
 *   }
 *   source        string
 * }
 * 
 * export ast
 * {
 *   token export
 *   declaration Node
 * }
 *
 *
 * import 表中 {
 *  from      idx(u32)    Constant(string)
 *  imported  idx(u32)    Constant(string)
 *  local     idx(u32)    Constant(string)
 * }
 * export 表中 {
 *  name      idx(u32)    Constant(string)
 *  globalIdx idx(u32)    Global(object)
 * }
 */

global ref{from, imported, local}
为所有import分配global
将import from 加入constnjkjk
生成ref{将 imported加入const, local加入const} 加入import表

/*
 *
 * header
 * module {
 *  name
 *  global []
 *  metadata {
 *    imports []
 *    exports []
 *  }
 *  compiled_func
 * }
 *
 *
 */

