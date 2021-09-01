import random
from typing import Union

import httpx

from test import c


code = 'code'
data = 'data'
msg = 'msg'


def _200(r: httpx.Response) -> Union[dict, list, str]:
    assert r.status_code == 200
    assert r.content != b''
    assert isinstance(j := r.json(), dict)
    if j[code] != 200:
        raise AssertionError(f"{j=}")
    assert j[msg] == "请求成功"
    return j[data]


def _400(r: httpx.Response) -> Union[dict, list]:
    if r.status_code != 200:
        raise AssertionError(f"{r.content=}\t maybe a real internal error happened.")
    assert r.content != b''
    j = r.json()
    if j[code] != 400:
        raise AssertionError(f"{j=}")
    assert j[data] is None
    return j[msg]


def random_hex_str(n: int):
    return "".join(random.choice("0123456789abcdef") for i in range(n))


# def random_email():
#     address = random_hex_str(random.randint(4, 10))
#     domain = random_hex_str(random.randint(1, 4)) + random.choice([".edu", ".com", ".cn"])
#     return f"{address}@{domain}"


# def random_location():
#     locations = "南京", "北京", "上海", "广州", "深圳", "武汉", "重庆", "西安", "吉林"
#     return random.choice(locations)

#
# def random_user_name():
#     return random_hex_str(8)


def random_password():
    return random_hex_str(8)


# def random_price():
#     return str(random.randint(100, 10000) / 100)


def get_token():
    return _200(c.post("/auth", data={"email": "example@foo.bar", "password": "123456"}))


methods = Union[c.get, c.post, c.put, c.delete]


def token_verify(method: methods, url: str, **kwargs):
    err0 = _400(method(url, **kwargs))
    assert err0 == '请求头中auth为空'

    err1 = _400(method(url, headers={"Authorization": "Bearer 123456"}, **kwargs))
    assert err1 == "无效的Token"

    token: str = get_token()
    err2 = _400(method(url, headers={"Authorization": token}, **kwargs))
    assert err2 == "请求头中auth格式有误"

    jwt = "Bearer " + token
    return {"Authorization": jwt}


def auth_verify(method: methods, url: str, **kwargs):
    token = token_verify(method, url, **kwargs)
    corr = _200(method(url, headers=token, **kwargs))
    return corr
