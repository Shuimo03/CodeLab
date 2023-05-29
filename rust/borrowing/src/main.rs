#[allow(unused)]
fn calculate_length(s: &String) -> usize{
    return s.len()
}

#[allow(unused)]
fn change_string(s: &mut String){
    s.push_str(",world");
}

#[allow(unused)]
fn borrow_object(s: &String) {
}

fn main(){
    let mut s = String::from("hello, ");

    let p = &mut s;
    p.push_str("world");
  
}