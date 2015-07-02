#include <stdio.h>

int *a[10];

int main(int argc, char* argv[]) {
    int b = 2;
    int *q;

    q = &b;
    if (!(q = NULL))
        printf("!q is true\n");
}
