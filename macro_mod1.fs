#[macro(call)]
fn infix(a, b) {
    quote(
      unquote(a)
      +
      unquote(b)
    )
}

export infix

/*
 * 编译A时
 * 对A的每个import，拿到导入的信息，从目标文件中找
 * 在import文件中
 * 找导入的 export 的标识符导出
 * 先在其ast中找，是否能找到声明的标识符，然后从import中找
 * 如果找到符号，并且它是函数，而且A的import使用了，并且有宏属性，进行提取
 *
 */
