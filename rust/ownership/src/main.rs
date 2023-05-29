fn takes_ownership(some_string: String) -> String{
    let mut world = String::from("hello");
    world.push_str(&some_string);
    world
}

fn makes_copy(some_integer: i32){
    println!("{}",some_integer);
}

fn main() {
  let s = String::from(" word");
  let world:String =takes_ownership(s);
  println!("{}",world);
  let x =5;
  makes_copy(x);
}
