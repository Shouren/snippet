#include <stdio.h>
#include <stddef.h>

struct S {
    char c;
    double d;
};

int main(void) {
    struct S test;
    printf("Offset: %zu\n", offsetof(struct S, d));
    printf("Offset: %zu\n", ((char *)&test.d) - (char *)&test);
}
