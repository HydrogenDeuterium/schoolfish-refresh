from test import c
from test.util import log, _200, _400, random_password, random_user_name, random_email, auth_verify


def test_getcode():
    p = {"email": "123@456"}
    c.delete("/auth", params=p)

    m = _400(c.get("/auth", params=p))
    assert m == "获取验证码过于频繁！"
    d = _200(c.get("/auth", params=p))
    assert d is None


def test_login():
    err0 = {"password": ""}
    err1 = {"email": "5145615616515613261561561263156", "password": ""}
    err2 = {"email": "example@foo.bar", "password": ""}
    corr = {"email": "example@foo.bar", "password": "123456"}
    cd = "123456"

    r_err0 = _400(c.post("/auth", data=err0))
    assert r_err0 == '请提供邮箱！'

    r_err1 = _400(c.post("/auth", data=err1))
    assert r_err1 == "用户不存在！"

    r_err2 = _400(c.post("/auth", data=err2))
    assert r_err2 == "用户名与密码不匹配！"

    r_corr = _200(c.post("/auth", params={"code": cd}, data=corr))
    log(r_corr)


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

    corr = {"username": random_user_name(), "email": random_email(), "password": random_password(), "avatar": "",
            "info": "", "profile": "", "location": ""}
    rcor = _200(c.post("/users", data=corr))

    for key in corr.keys():
        assert key == "password" or rcor[key] == corr[key]


def test_userinfo():
    jwt = auth_verify(c.get, "/users", )

    co = _200(c.get("/users", headers=jwt))
    assert co == {
        'Avatar': '',
        'Email': 'example@foo.bar',
        'Info': '',
        'Location': '',
        'Profile': '',
        'Uid': 70,
        'Username': '1'
    }


def test_product_get_all():
    co = _200(c.get("/products", params={"page": 1}))
    assert co == []


def test_product_new():
    return
    # err0 = _400(c.get("/products/user", params={"page": 1}))
    # assert err0 == '请求头中auth为空'
    # jwt: str = get_token()
    # print(jwt)
    # err1 = _400(c.get("/products/user", headers={"Authorization": jwt}, params={"page": 1}))
    # assert err1 == "请求头中auth格式有误"
    # err3 = _400(c.get("/products/user", headers={"Authorization": "Bearer 123456"}, params={"page": 1}))
    # assert err3 == "无效的Token"
    # jwt = "Bearer " + jwt
    # co = _200(c.get("/products/user", headers={"Authorization": jwt}, params={"page": 1}))
    # assert co == []


def test_view_user_product():
    pass


if __name__ == "__main__":
    pass
