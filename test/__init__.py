import os

import httpx

from faker import Faker

fake = Faker(locale="zh_CN")

url = os.getenv('url', default='http://127.0.0.1:8080')
print(f"{url=}")
c = httpx.Client(base_url=url)
