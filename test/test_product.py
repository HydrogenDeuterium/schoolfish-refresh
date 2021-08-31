from test import c
from test.util import _200, _400, auth_verify, random_hex_str, random_location, random_price


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


def test_view_a_product():
    err0 = _400(c.get("/products/-1"))
    assert err0 == "货物不存在！"
    corr = _200(c.get("/products/1"))
    assert corr == {'info': '1',
                    'pid': 1,
                    'location': '南京',
                    'owner': 69,
                    'price': '1.20',
                    'title': '测试商品'}


def test_new_product():
    product = {'info': random_hex_str(200),
               'location': random_location(),
               'price': random_price(),
               'title': '测试商品' + random_hex_str(20)}
    corr = auth_verify(c.post, "/products", data={"product": product})
    del corr["owner"]
    del corr["pid"]
    assert corr == product


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
