#include <stdio.h>

int *a[10];

int *
return_null() {
    return NULL;
}

int *
return_not_null(int *val) {
    return val;
}

int main(int argc, char* argv[]) {
    int b = 0;
    int *q;

    q = NULL;
    if (!q && !(q = return_null()))
        printf("One: true\n");

    q = NULL;
    if (!q && !(q = return_not_null(&b)))
        printf("Two: true\n");
}
