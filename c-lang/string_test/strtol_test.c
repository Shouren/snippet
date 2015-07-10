#include <stdlib.h>
#include <stdio.h>
#include <string.h>

int
main(int argc, char *argv[]) {
    char *str = "80-433";
    char *end = NULL;

    end = strchr(str, '-');
    if (end == NULL)
        return 0;
    strtol(str, &end, 10)
}
