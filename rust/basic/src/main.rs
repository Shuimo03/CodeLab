
#[allow(unused)]
fn range_number(){
    for i in 1..=5{
        println!("{}",i);
    }
}

#[allow(unused)] //允许出现暂时没有被使用的函数
fn add_with_extra(x: i32, y: i32) -> i32{
    let x = x+1;
    let y = y+5;
    return x+y;
}

fn add_five(x: i32) -> i32{
    if x > 5 {
       return x-5;
    }

    x + 5
}





fn main(){
    let  res = add_five(6);
    println!("{}",res);
}



