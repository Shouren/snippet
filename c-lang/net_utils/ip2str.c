#include <stdio.h>
#include <stdlib.h>

char *
inet_ntoa(uint32_t ip, char *buf) {
    unsigned char *bytep;

    bytep = (unsigned char *) &ip;
    sprintf(buf, "%d.%d.%d.%d", bytep[0], bytep[1], bytep[2], bytep[3]);
    return buf;
}

int
main(int argc, char *argv[]) {
    uint32_t ip_1 = 0xc0a80001;
    uint32_t ip_2 = 0xc0a80002;
    char ip_addr_1[16];
    char ip_addr_2[16];

    printf("src:%s dst:%s\n",
            inet_ntoa(htonl(ip_1), (char *) &ip_addr_1),
            inet_ntoa(htonl(ip_2), (char *) &ip_addr_2));
    return 0;
}
