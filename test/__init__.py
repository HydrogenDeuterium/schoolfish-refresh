import os

import httpx

url = os.getenv('url', default='http://127.0.0.1:8080')
print(f"{url=}")
c = httpx.Client(base_url=url)
