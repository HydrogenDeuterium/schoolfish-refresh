from test import c
from test.util import _200


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