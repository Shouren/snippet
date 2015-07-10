#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <ctype.h>

int main(void)
{
    int arr[2] = {0, 0};
    int count = 0;
    char input[] = "800-433";
    printf("Parsing the input string '%s'\n", input);
    char *token = strtok(input, "-");
    while(token) {
        if (count >= 2) break;
        if (isdigit((int)token) == 0)
            printf("input include nonumeric value\n");
        arr[count] = atoi(token);
        token = strtok(NULL, "-");
        count += 1;
    }
    printf("%d\n", arr[0]);
    printf("%d\n", arr[1]);
    printf("%d\n", count);
}
