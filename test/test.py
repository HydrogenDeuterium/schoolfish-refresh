import json
import os
from typing import Union

import httpx

url = os.getenv('url')
if url is None:
    url = 'http://127.0.0.1:8080'


def log(*args, **kwargs):
    print("\n\t>>>", *args, **kwargs)


log(f"{url=}")

c = httpx.Client(base_url=url)

code = 'code'
data = 'data'
msg = 'msg'


def _200(r: httpx.Response) -> Union[dict, list]:
    assert r.status_code == 200
    try:
        j = r.json()
    except json.JSONDecodeError as e:
        log(r.content)
        raise e

    assert isinstance(j, dict)
    log(j)
    if j[code] != 200:
        log(j)
        raise AssertionError
    assert j[msg] == "请求成功"
    return j[data]


def _400(r: httpx.Response) -> str:
    assert r.status_code == 200
    try:
        j = r.json()
    except json.JSONDecodeError as e:
        log(r.content)
        raise e
    log(j)
    if j[code] != 400:
        log(j)
        raise AssertionError
    assert j[data] is None
    return j[msg]


def _500(r: httpx.Response) -> None:
    assert r.status_code == 200
    try:
        j = r.json()
    except json.JSONDecodeError as e:
        log(r.content)
        raise e
    log(j)
    if j[code] != 500:
        log(j)
        raise AssertionError
    assert j[msg] == ""


def test_getcode():
    p = {"email": "123@456"}
    c.delete("/auth", params=p)

    d = _200(c.get("/auth", params=p))
    assert d is None
    m = _400(c.get("/auth", params=p))
    assert m == "获取验证码过于频繁！"


def test_login():
    err1 = {"email": "", "password": ""}
    err2 = {"email": "example@foo.bar", "password": ""}
    corr = {"email": "example@foo.bar", "password": "123456"}
    p = {"email": "example@foo.bar"}

    cd = "123456"
    r_err1 = _400(c.post("/auth", params={code: "0"}, data=err1))
    assert r_err1 == "用户不存在！"

    r_err2 = _400(c.post("/auth", params={code: cd}, data=err2))
    assert r_err2 == "用户名与密码不匹配！"

    r_corr = _200(c.post("/auth", params={code: cd}, data=corr))
    log(r_corr)


if __name__ == "__main__":
    pass
