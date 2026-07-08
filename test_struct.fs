/*
 * struct St1 {
 *   a:
 *   b:
 *   c:
 * }
 *
 * struct St2 {
 *  St1     // 组合
 *  d:      // 内部字段
 *  e:
 * }
 *
 * fn (s1 St1) m1() {   // St1的方法，第一个参数为s1
 *
 * }
 *
 * fn (s2 St2) m2() {   // St2的方法，第一个参数为s2
 *
 * }
 *
 * let st1 = St1 {      // St1 结构体值
 *  a: 1,
 *  b: 2,
 *  c: 3
 * }
 * 
 * let st2 = St2 {
 *  St1: st1,         // 显示指定 st1 给 内部的 St1组合字段
 *  d:1,
 *  e:2
 * }
 *
 * st1.a      // 访问 St1 字段
 * st1.m1()   // 访问 St1 方法
 *
 * st2.d      // 访问 St2 字段
 * st2.m2()   // 访问 St2 方法
 * st2.a      // 访问 St1 字段
 * st2.m1()   // 访问 St1 方法
 *
 * AST:
 *  StructDeclared {
 *    Name
 *    Fileds
 *  }
 *
 *  FiledDeclared{
 *    Name
 *    IsEmbed  // true 表示它是一个组合
 *  }
 *
 *  MethodDeclared{
 *    StructName
 *    FnExpr
 *  }
 *
 *  StructLiteral {
 *    Tag: Ident
 *    Elements[
 *       {key:ident, value: expr}
 *    ]
 *  }
 *
 *  MemberExpr {
 *    Visitor
 *    Member
 *  }
 *
 *
 */
