from test import c, fake
from test.util import _400, random_password, _200, auth_verify, random_hex_str


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

    d_corr = {"username": fake.name(), "email": fake.email(), "password": random_password(), "avatar": "",
              "info": fake.paragraph(), "profile": random_hex_str(60), "location": fake.address()}
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
