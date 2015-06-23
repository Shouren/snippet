#include <stdio.h>
#include <string.h>

struct list_head {
    int idx;
    struct list *pre;
    struct list *next;
};

struct test {
    struct list_head first;
    struct list_head second;
    struct list_head third;
};

enum words {
    a,
    b,
    c,
    d,
};

int main(int argc, char * argv[]) {
    struct test p;
    struct test * q;
    char * a = "123";

    p.first.idx = 1;

    q = &p;

    printf("%d\n", argc);
    printf("%s\n", a);
    printf("Res: %u\n", p.first.idx);
    printf("Res: %u\n", q->first.idx);

    //struct list_head * ptr;
    //unsigned long res;
    //res = (unsigned long) (&((struct test*) 0)->second);
    //q = ((struct test *) ((char *) (ptr) - (unsigned long) (&((struct test*) 0)->second)));
    //printf("Res is %d\n", res);
    //printf("Size of list is %d\n", sizeof(*ptr));
    //printf("Address of q is 0x%x\n", q);
}
