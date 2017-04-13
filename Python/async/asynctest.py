# /usr/bin/env python3

import signal
import sys

import asyncio
import aiohttp

loop = asyncio.get_event_loop()
client = aiohttp.ClientSession(loop=loop)

async def get_status_code(client, url):
    async with client.get(url) as response:
        return await print_status(response.status)

async def print_status(status):
    print(status)


def signal_handler(signal, frame):
    loop.stop()
    client.close()
    sys.exit(0)


signal.signal(signal.SIGINT, signal_handler)

asyncio.ensure_future(get_status_code(client, 'https://www.google.com/generate_204'))

loop.run_forever()
