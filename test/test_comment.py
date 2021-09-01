import httpx

from test import fake, url
from test.util import auth_verify

c = httpx.Client(base_url=f"{url}/comments")


def random_comment():
    ret = {"text": fake.paragraph()}
    return ret


def test_comment_update():
    corr = auth_verify(c.put, "/0")
    assert corr == "暂未实现！"
