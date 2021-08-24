import os

import httpx

url = os.getenv('url')
if url is None:
    url = 'http://127.0.0.1:8080'

print(f"{url=}")

c = httpx.Client(base_url=url)

def test_getcode():
    pass
