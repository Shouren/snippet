#include <stdlib.h>
#include <stdio.h>

#define PRINT_TEST(X) (\
        printf("%s\n", X))

int
main(int argc, char *argv[]) {
    PRINT_TEST("Hello World");
    return 0;
}
