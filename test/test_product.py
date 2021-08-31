from test import c
from test.util import _200, _400


def test_product_get_by_page():
    co = _200(c.get("/products", params={"page": 1}))
    assert co == []


def test_view_user_product():
    err0 = _400(c.get("/products/users/-1"))
    assert err0 == "用户不存在！"
    corr = _200(c.get("/products/users/69"))
    assert corr == [
        {'pid': 1, 'title': '测试商品', 'info': '1', 'price': '1.20', 'owner': 69, 'location': '南京'}
    ]


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
