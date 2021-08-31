from test import c
from test.util import _400, random_password, random_user_name, random_email, _200, auth_verify


def test_register():
    re0 = _400(c.post("/users"))
    assert re0 == "提供邮箱！"

    err1 = {"username": "1", "email": "example@foo.bar", "avatar": "", "info": "",
            "profile": "", "location": ""}
    re1 = _400(c.post("/users", params={"code": "0"}, data=err1))
    assert re1 == "用户已注册!"

    err2 = {"username": "1", "email": random_password(), "avatar": "", "info": "",
            "profile": "", "location": ""}
    re2 = _400(c.post("/users", params={"code": "0"}, data=err2))
    assert re2 == "提供密码！"

    d_corr = {"username": random_user_name(), "email": random_email(), "password": random_password(), "avatar": "",
              "info": "", "profile": "", "location": ""}
    corr = _200(c.post("/users", data=d_corr))

    del d_corr["password"]
    del corr["uid"]
    assert corr == d_corr


def test_userinfo():
    corr = auth_verify(c.get, "/users", )
    assert corr == {
        'uid': 70,
        'username': '1',
        'email': 'example@foo.bar',
        'avatar': '',
        'info': '',
        'location': '',
        'profile': '',
    }
