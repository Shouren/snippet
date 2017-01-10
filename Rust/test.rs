struct Point {
    x: i32,
    y: i32,
}

fn main() {
    let mut point = Point {x: 0, y: 0};
    point.x = 6;
    point.y = 6;

    let point_r = point;
    println!("{}, {}", point_r.x, point_r.x);

    let mut v0 = 1;
    let v1 = v0;
    v0 = 2;
    println!("{}, {}", v0, v1)

}
