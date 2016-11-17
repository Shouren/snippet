import functools


def test(a, b, c):
    print(a)
    print(b)
    print(c)


if __name__ == '__main__':
    test1 = functools.partial(test, 1, 2)
    test1(3)
