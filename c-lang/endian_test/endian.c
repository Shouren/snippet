#include <stdio.h>
#include <stdint.h>
#include <arpa/inet.h>

int main(int argc, char *argv[]) {
	uint16_t proto;
	proto = 0x0107;
	printf("Proto Code: %d\n", proto);
	proto = htonl(0x0107);
	printf("Proto Code: %d\n", proto);
}
