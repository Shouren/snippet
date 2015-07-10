#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

int
main(int argc, char *argv[]) {
    char *start = "80-433a";
    char target = '-';
    char *end = NULL;

    int val[2] = {0, 0};
    end = strchr(start, target);
    if (end == NULL) {
        printf("- not found");
        if (isdigit((int)*start) == 0)
            printf("nonumeric value in start\n");
        val[0] = atoi(start);
    }
    else {
        //*end = '\0';
        end++;
        if (isdigit((int)*start) == 0)
            printf("nonumeric value in end\n");
        val[0] = atoi(start);
        if (isdigit((int)*end) == 0)
            printf("nonumeric value in end\n");
        val[1] = atoi(end);
    }
    printf("%d\n", val[0]);
    printf("%d\n", val[1]);
}
