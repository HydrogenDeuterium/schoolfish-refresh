from test import c, fake
from test.util import error, random_password, success, token_verify


def random_user():
    ret = {
        "username": fake.name(),
        "email": fake.email(),
        "password": fake.password(special_chars=False),
        "avatar": fake.image_url(),
        "info": fake.sentence(),
        "profile": "\n".join(fake.paragraph()),
        "location": fake.address(),
    }
    return ret


def test_register():
    re0 = error(c.post("/users"))
    assert re0 == "提供邮箱！"

    err1 = {"username": "1", "email": "example@foo.bar", "avatar": "", "info": "",
            "profile": "", "location": ""}
    re1 = error(c.post("/users", params={"code": "0"}, data=err1))
    assert re1 == "用户已注册!"

    err2 = {"username": "1", "email": random_password(), "avatar": "", "info": "",
            "profile": "", "location": ""}
    re2 = error(c.post("/users", params={"code": "0"}, data=err2))
    assert re2 == "提供密码！"

    d_corr = random_user()
    corr = success(c.post("/users", data=d_corr))

    del d_corr["password"]
    del corr["uid"]
    assert corr == d_corr


def test_userinfo():
    token = token_verify(c.get, "/users", )
    corr = success(c.get("/users", headers=token))
    assert corr == {
        'uid': 70,
        'username': '1',
        'email': 'example@foo.bar',
        'avatar': '',
        'info': '',
        'location': '',
        'profile': '',
    }
