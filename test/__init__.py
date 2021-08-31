import os

import httpx


def log(*args, **kwargs):
    print("\n\t>>>", *args, **kwargs)


url = os.getenv('url', default='http://127.0.0.1:8080')
log(f"{url=}")
c = httpx.Client(base_url=url)
