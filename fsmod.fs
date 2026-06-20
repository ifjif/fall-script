/*
 * 导入语法
 * import {a,b} from "./a.fs"
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
 *   name        string
 *   declaration Node
 * }
 *
 */

