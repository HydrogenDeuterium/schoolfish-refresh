import os
from typing import Union

import httpx

url = os.getenv('url')
if url is None:
    url = 'http://127.0.0.1:8080'

print(f"{url=}")

c = httpx.Client(base_url=url)

code = 'code'
data = 'data'
msg = 'msg'


def verify_200(r: httpx.Response) -> Union[dict, list]:
    assert r.status_code == 200
    assert isinstance(j := r.json(), dict)
    print(j)
    assert j[code] == 200
    assert j[msg] == ""
    return j[data]


def verify_400(r: httpx.Response) -> str:
    assert r.status_code == 200
    assert isinstance(j := r.json(), dict)
    print(j)
    assert j[code] == 400
    assert j[data] is None
    return j[msg]


def verify_500(r: httpx.Response) -> None:
    assert r.status_code == 200
    assert isinstance(j := r.json(), dict)
    print(j)
    assert j[code] == 500
    assert j[msg] == ""


def test_getcode():
    p = {"email": "123@456"}
    c.delete("/auth", params=p)

    d = verify_200(c.get("/auth", params=p))
    print("code=", d)
    m = verify_400(c.get("/auth", params=p))
    assert m == ""


if __name__ == "__main__":
    pass
