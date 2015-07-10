#include <stdio.h>
#include <arpa/inet.h>

int
main(int argc, char *argv[]) {
    unsigned short val;
    unsigned short start;
    unsigned short end;

    val = htons(8080);

    start = htons(80);
    end = htons(433);

    if (ntohs(start) <= ntohs(val) && ntohs(val) <= ntohs(end))
        printf("%d\n", ntohs(val));

    return 0;
}
