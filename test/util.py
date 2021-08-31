import random
from typing import Union

import httpx

from test import log, c

code = 'code'
data = 'data'
msg = 'msg'


def _200(r: httpx.Response) -> Union[dict, list, str]:
    assert r.status_code == 200
    assert r.content != b''
    assert isinstance(j := r.json(), dict)
    if j[code] != 200:
        log(j)
        raise AssertionError
    assert j[msg] == "请求成功"
    return j[data]


def _400(r: httpx.Response) -> Union[dict, list]:
    if r.status_code != 200:
        log(r.content)
        raise AssertionError
    assert r.content != b''
    j = r.json()
    if j[code] != 400:
        log(j)
        raise AssertionError
    assert j[data] is None
    return j[msg]


# def _500(r: httpx.Response) -> None:
#     assert r.status_code == 200
#     assert r.content != b''
#     assert isinstance(j := r.json(), dict)
#     if j[code] != 500:
#         log(j)
#         raise AssertionError
#     assert j[msg] == ""


def random_hex_str(n: int):
    return "".join(random.sample("0123456789abcdef", n))


def random_email():
    address = random_hex_str(random.randint(4, 10))
    domain = random_hex_str(random.randint(1, 4)) + random.choice([".edu", ".com", ".cn"])
    return f"{address}@{domain}"


def random_user_name():
    return random_hex_str(8)


def random_password():
    return random_hex_str(8)


def get_token():
    return _200(c.post("/auth", data={"email": "example@foo.bar", "password": "123456"}))


methods = Union[c.get, c.post, c.put, c.delete]


def token_verify(method: methods, url: str, **kwargs):
    err0 = _400(method(url, **kwargs))
    assert err0 == '请求头中auth为空'

    err1 = _400(method(url, headers={"Authorization": "Bearer 123456"}, **kwargs))
    assert err1 == "无效的Token"

    jwt: str = get_token()
    err2 = _400(method(url, headers={"Authorization": jwt}, **kwargs))
    assert err2 == "请求头中auth格式有误"

    jwt = "Bearer " + jwt
    return jwt


def auth_verify(method: methods, url: str, **kwargs):
    jwt = token_verify(method, url, **kwargs)
    corr = _200(method(url, headers={"Authorization": jwt}, **kwargs))
    return corr
