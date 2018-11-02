#! /usr/bin/env python3

import re

from atlassian import Confluence
from six.moves.urllib.parse import urljoin
from bs4 import BeautifulSoup

WIKI_URL = "https://wiki.4paradigm.com"

confluence = Confluence(
        url=WIKI_URL,
        username="lijiangnan",
        password="Sweetydarling520")

ret = confluence.get_page_by_id("36220156", expand="body")
web_view_url = ret["_links"]["webui"]
url = urljoin(confluence.url, web_view_url)
resp = confluence._session.request(
        method="GET",
        url=url,
        auth=(confluence.username, confluence.password),
        timeout=confluence.timeout,
        verify=confluence.verify_ssl)

# p = re.compile(r"docker02:35000[a-zA-Z0-9./]+<")
p = re.compile(r"docker02:35000[-.:/a-zA-Z0-9]+")

soup = BeautifulSoup(resp.text, "html.parser")
table = soup.find(
        "table", attrs={"class": "wrapped relative-table confluenceTable"})
table_body = table.find("tbody")
rows = table_body.find_all("tr")
for row in rows:
    for match in p.findall(str(row)):
        print(match)
