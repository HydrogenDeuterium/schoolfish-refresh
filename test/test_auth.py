import re

from test import c
from test.util import error, success


def test_getcode():
    p = {"email": "123@456"}

    c.delete("/auth", params=p)
    c.get("/auth", params=p)
    m = error(c.get("/auth", params=p))
    assert m == "获取验证码过于频繁！"

    c.delete("/auth", params=p)
    d = success(c.get("/auth", params=p))
    assert d is None


def test_login():
    err0 = {"password": ""}

    r_err0 = error(c.post("/auth", data=err0))
    assert r_err0 == '请提供邮箱！'

    err1 = {"email": "5145615616515613261561561263156", "password": ""}
    r_err1 = error(c.post("/auth", data=err1))
    assert r_err1 == "用户不存在！"

    err2 = {"email": "example@foo.bar", "password": ""}
    r_err2 = error(c.post("/auth", data=err2))
    assert r_err2 == "用户名与密码不匹配！"

    corr = {"email": "example@foo.bar", "password": "123456"}
    r_corr = success(c.post("/auth", params={"code": "123456"}, data=corr))
    assert re.match(r"(.*)\.(.*)\.(.*)", r_corr)
