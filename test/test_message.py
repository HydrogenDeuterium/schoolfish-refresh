import httpx

from test import fake, url
from test.util import token_verify

c = httpx.Client(base_url=f"{url}/message", timeout=1000)


def random_message():
    return fake.sentance()


def test_send_message():
    token = token_verify(c.post, f"/{71}")
