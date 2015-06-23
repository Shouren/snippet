#include <stdint.h>
#include <stdio.h>
#include "main.h"

int main(int argc, char *argv[]) {
	struct ip_proto proto;
	proto.tcp = ip_tcp;
	proto.udp = 0x11;
	printf("tcp: %d\n", proto.tcp);
	printf("udp: %d\n", proto.udp);
}
