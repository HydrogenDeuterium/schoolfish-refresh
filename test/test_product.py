from test import c, fake
from test.util import success, error, token_verify


def test_product_get_by_page():
    result = [
        {'info': '1',
         'location': '南京',
         'owner': 69,
         'pid': 1,
         'price': '1.20',
         'title': '测试商品'},
        {
            'info': '52eeaae0b8b41086fe55bac4694a822655ab1d04d139fa5f12ccff3860b9bf5486744f1caf170ad731c73396eda332489'
                    '995688309bd13c676ff9173e8d226510dd9d80e8c0db48bcd7e472fd550014bfe3e80670e7121400a1'
                    '0a1df5e629568c10832ea',
            'location': '上海',
            'owner': 70,
            'pid': 2,
            'price': '23.09',
            'title': '测试商品38328db054bd93e35351'},
        {
            'info': '6acc5ce8f2083c813038344ce34a63d89ccad37fca930e04b1feecf6924fc5405368839c517a2'
                    '27f432efae3e31ba20d8128511c22a02ec2265124e7b5b727a6fda9512423314716e54994cf2d9a6'
                    '2dd858edfff10fd8f03bbae7ecce912b92ec17a62d1',
            'location': '武汉',
            'owner': 70,
            'pid': 3,
            'price': '1.20',
            'title': '测试商品d1c6f896eb39c071432b'},
        {
            'info': '41b220b62f33a565320903a52fda073e781e9a363188da91b7b4affa085d9a80e41d03525b7cf4fd307'
                    '8235febb625ef64f272fab4e7df30667d966bce815fead58ea5114aa2095d6f57cbd3786da1cfe1559a'
                    'b9df8e6e1021cc495170f8657838860167',
            'location': '上海',
            'owner': 70,
            'pid': 4,
            'price': '42.45',
            'title': '测试商品c8757ad8e9b51be9d8de'},
        {
            'info': 'ca7086e55feb639ff2aa4b3ba9ffe4f44a1dc3fc43ba51f2a593a94529a6b2a02772fcafb86c91ac'
                    '4e2e6fcb3e0706f6290213bef579dc43683e1ba38151552d424520681ee1b755228bd34eca3a9cb9'
                    'b48e24e656967a84c33cca63f5c188dc03e5edb9',
            'location': '北京',
            'owner': 70,
            'pid': 5,
            'price': '57.42',
            'title': '测试商品c4af509066356b31819a'},
        {
            'info': 'a87f9865085a12875ab2cc5c2c767247679006b18fe7f5f11b0ce0032bcbc45d49711ec0fc0bcdc'
                    '8c414aebb9cce733355336d1ab8f49f179c4607eccab1c66f20fa4e141b212e596f0e2e4917a84d'
                    'e7e9b6a8622f13f379c01f7b0ec5c2fdf583e08172',
            'location': '南京',
            'owner': 70,
            'pid': 6,
            'price': '40.02',
            'title': '测试商品dee8628b606165271d86'},
        {
            'info': 'da87464d7431a422f7e6e4e12469976d8edec2fd95f4b50c0cf99f760a8c5f2567b5b1743cd'
                    '50e2b0e98760ed8b285258f32bce5d44399487e54b60903bc58b938be510c075508f0ce084'
                    'c405fc291e310c2f61a9b39dcf772bb0f31c077bd54d5b1ce48',
            'location': '重庆',
            'owner': 70,
            'pid': 7,
            'price': '27.86',
            'title': '测试商品f6c86fb461c2618b739d'},
        {
            'info': '9e86592efc3cff84fa9b6a232a4dc654139d67abe152ea81510f229e358f37ae7d08e'
                    '2d660c2ffa05636fefeb90bdabc637e50e173da8855ba48b639182cefd350addfa62d'
                    '24b56a811e853a61b90a4e4ce789530fd2bbd6ed045c4ce00653c27d0f6594',
            'location': '南京',
            'owner': 70,
            'pid': 9,
            'price': '56.16',
            'title': '测试商品bc1254ec6552e35592b5'},
        {
            'info': 'cf6a0d5cd8d486125f1ea65992f41a6d6e74b964f9d9637cabd62b5ade75b77cf5'
                    'd350fdf9701a8a5e2cf5d9e0c3dd92d0be4d4169d196e7bb63d616ad86ea93443'
                    '7a078a67d1ae77e324f88b398b5214d55c3b5dbc202555ab11c7dcaf5b0571b45dd43',
            'location': '西安',
            'owner': 70,
            'pid': 10,
            'price': '40.22',
            'title': '测试商品aaefd29df71e9d76a449'},
        {
            'info': '43084e339b4da7c3f4f3d04ee846aad78a7d205ca525d7d1b6e94d1cd4eb41eacb9'
                    '42c150cacf5e324039e431c1a4e0042db7fccd19be1fe47b65c70ba6b9663b8bceb8d'
                    'dd27ebfc5e51f7b4dc612adfa7d0a9415c8a56ea9470b2d4866ad0c267ff29fc',
            'location': '西安',
            'owner': 70,
            'pid': 11,
            'price': '59.70',
            'title': '测试商品9ca375ae2ed1484b424f'}]
    co = success(c.get("/products", params={"page": 1}))
    assert co == result


def test_view_user_product():
    err0 = error(c.get("/products/users/-1"))
    assert err0 == "用户不存在！"
    corr = success(c.get("/products/users/69"))
    assert corr == [
        {'pid': 1, 'title': '测试商品', 'info': '1', 'price': '1.20', 'owner': 69, 'location': '南京'}
    ]


def test_view_a_product():
    err0 = error(c.get("/products/-1"))
    assert err0 == "货物不存在！"
    corr = success(c.get("/products/1"))
    assert corr == {'info': '1',
                    'pid': 1,
                    'location': '南京',
                    'owner': 69,
                    'price': '1.20',
                    'title': '测试商品'}


def random_product(s: str):
    ret = {
        'title': s + fake.paragraph(),
        'price': f"{fake.pyfloat(right_digits=2, min_value=1, max_value=200)}",
        'location': fake.address(),
        'info': "\n".join(fake.paragraphs(5)),
    }
    return ret


def test_new_product():
    product = random_product("测试新增商品")
    token = token_verify(c.post, "/products")
    corr = success(c.post("/products", headers=token, data={"product": product}))
    del corr["owner"]
    del corr["pid"]
    assert corr == product


def test_update_product():
    product = random_product("测试修改商品")
    token = token_verify(c.put, "/products/20")

    err = error(c.put("/products/8", headers=token, data={"product": product}))
    assert err == "商品不存在！"

    corr = success(c.put("/products/21", headers=token, data={"product": product}))
    del corr["owner"]
    del corr["pid"]
    assert corr == product


def test_delete_product():
    product = random_product('测试删除商品')
    token = token_verify(c.delete, "/products/2")

    to_del = success(c.post("/products", headers=token, data={"product": product}))
    pid = to_del["pid"]
    corr = success(c.delete(f"/products/{pid}", headers=token))
    assert float(corr["price"]) == float(product["price"])
    del corr["price"], product["price"]
    assert corr == product | {"owner": 70, "pid": pid}
