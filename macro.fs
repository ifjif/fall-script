#[macro(call)]
fn name() {

}

#[macro(attr)]
fn router(attrs, ast) {

 quote(
  fn() {
    puts("xx")
    let reuslt = wrapcall(ast) //callast: fn(){ast.block}()
    puts("xxx")
  }
 ) 
}



#[router()]
fn aaa(a,b) {


let result = fn(){a;b}
}

let a = name()// name 是macro，进行展开

// 函数调用，如果是宏，进行展开
// 除此，如果函数定义，有宏信息，对函数定义进行处理
