import os
import random
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
    assert r.content != b''
    assert isinstance(j := r.json(), dict)
    if j[code] != 200:
        log(j)
        raise AssertionError
    assert j[msg] == "请求成功"
    return j[data]


def _400(r: httpx.Response) -> Union[dict, list]:
    assert r.status_code == 200
    assert r.content != b''
    j = r.json()
    if j[code] != 400:
        log(j)
        raise AssertionError
    assert j[data] is None
    return j[msg]


def _500(r: httpx.Response) -> None:
    assert r.status_code == 200
    assert r.content != b''
    assert isinstance(j := r.json(), dict)
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
    err0 = {"password": ""}
    err1 = {"email": "5145615616515613261561561263156", "password": ""}
    err2 = {"email": "example@foo.bar", "password": ""}
    corr = {"email": "example@foo.bar", "password": "123456"}
    p = {"email": "example@foo.bar"}
    cd = "123456"

    r_err0 = _400(c.post("/auth", data=err0))
    assert r_err0 == '请提供邮箱！'

    r_err1 = _400(c.post("/auth", data=err1))
    assert r_err1 == "用户不存在！"

    r_err2 = _400(c.post("/auth", data=err2))
    assert r_err2 == "用户名与密码不匹配！"

    # r_corr = _200(c.post("/auth", params={code: cd}, data=corr))
    # log(r_corr)


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


def test_register():
    re0 = _400(c.post("/users"))
    assert re0 == "提供邮箱！"

    err1 = {"username": "1", "email": "example@foo.bar", "avatar": "", "info": "",
            "profile": "", "location": ""}
    re1 = _400(c.post("/users", params={code: "0"}, data=err1))
    assert re1 == "用户已注册!"

    err2 = {"username": "1", "email": random_password(), "avatar": "", "info": "",
            "profile": "", "location": ""}
    re2 = _400(c.post("/users", params={code: "0"}, data=err2))
    assert re2 == "提供密码！"

    corr = {"username": random_user_name(), "email": random_email(), "password": random_password(), "avatar": "",
            "info": "", "profile": "", "location": ""}
    rcor = _200(c.post("/users", data=corr))

    for key in corr.keys():
        assert key == "password" or rcor[key] == corr[key]


def test_userinfo():
    re = _400(c.get("/users/0"))
    assert re == "用户未注册!"

    co = _200(c.get("/users/24"))
    assert co == {
        'Uid': 24,
        'Username': 'alexzhangzhe',
        'Email': 'alexzhang@deu.wiki',
        'Hashed': 'sdsdsdssgdgfsd',
        'Info': '我shizhangzhe',
        'Profile': 'hhhhh',
        'Location': 'noname01',
        'Avatar': 'sdfsdd',
    }


if __name__ == "__main__":
    pass
