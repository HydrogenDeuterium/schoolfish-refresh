from test import c
from test.util import _200, _400, auth_verify, random_hex_str, random_location, random_price, token_verify


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


def test_update_product():
    product = {'info': random_hex_str(200),
               'location': random_location(),
               'price': random_price(),
               'title': '测试商品' + random_hex_str(20)}
    token = token_verify(c.put, "/products/2", data={"product": product})

    err = _400(c.put("/products/8", headers=token, data={"product": product}))
    assert err == "商品不存在！"

    corr = _200(c.put("/products/2", headers=token, data={"product": product}))
    del corr["owner"]
    del corr["pid"]
    assert corr == product
