fn main() {
    let nums = vec![1, 2, 3, 4, 5];
    println!("{:?}", nums);

    let mut mut_nums = nums;

    mut_nums.push(6);

    for num in mut_nums{
        println!("{}", num);
    }
}
