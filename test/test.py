import os

import httpx

url = os.getenv('url')
if url is None:
    url = 'http://127.0.0.1:8080'

print(f"{url=}")

c = httpx.Client(base_url=url)

code = 'code'
data = 'data'
msg = 'msg'


def test_getcode():
    p = {"email": "123@456"}
    c.delete("/auth", params=p)

    r = c.get("/auth", params=p)
    assert r.status_code == 200
    j = r.json()
    assert j['code'] == 200
    assert j['data'] is None
    assert isinstance(j['msg'], str)

    r = c.get("/auth", params=p)
    assert r.status_code == 200
    j = r.json()
    assert j['code'] == 400
    assert j['data'] is None
    assert isinstance(j['msg'], str)
