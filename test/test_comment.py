import httpx

from test import fake, url
from test.util import auth_verify, token_verify, _200, _400

c = httpx.Client(base_url=f"{url}/comments", timeout=1000)


def random_comment():
    ret = {"text": fake.paragraph()}
    return ret


def test_comment_to_product():
    token = token_verify(c.post, "/products/str")
    err0 = _400(c.post("/products/str", headers=token))
    err1 = _400(c.post("/products/0", headers=token))
    assert err0 == err1 == "pid格式不正确！"

    comment = random_comment()
    pid = 2
    corr = _200(c.post(f"/products/{pid}", headers=token, data=comment))

    assert isinstance(corr["cid"], int)
    del corr["cid"]
    assert corr == comment | {'commentator': 70, 'product': pid, 'response_to': 0, }


def test_get_product_comment():
    result = [{'cid': 0, 'commentator': 70, 'product': 1, 'response_to': 0, 'text': ''},
              {'cid': 13,
               'commentator': 70,
               'product': 1,
               'response_to': 0,
               'text': '等级品牌单位的人朋友其实更新.详细详细系列单位密码问题.日期非常部门还是作者汽车工程.'},
              {'cid': 14,
               'commentator': 70,
               'product': 1,
               'response_to': 0,
               'text': '各种点击次数就是女人.显示部门发现发布如此出现有些.'},
              {'cid': 15,
               'commentator': 70,
               'product': 1,
               'response_to': 0,
               'text': '最后同时出现世界你的需要不同.单位历史行业部门工程制作非常.'},
              {'cid': 16,
               'commentator': 70,
               'product': 1,
               'response_to': 0,
               'text': '完成数据点击就是可能.'},
              {'cid': 17,
               'commentator': 70,
               'product': 1,
               'response_to': 0,
               'text': '查看手机完成孩子.电影非常就是而且.'},
              {'cid': 18,
               'commentator': 70,
               'product': 1,
               'response_to': 0,
               'text': '名称有些事情政府虽然发展.游戏内容系统只要具有进入还是.'},
              {'cid': 19,
               'commentator': 70,
               'product': 1,
               'response_to': 0,
               'text': '看到以上您的不同积分详细虽然.游戏之后一般这样法律那个.一起技术可能.空间美国如此.'},
              {'cid': 20,
               'commentator': 70,
               'product': 1,
               'response_to': 0,
               'text': '重要这是是一问题包括准备能力主要.在线希望这个登录如果一切拥有软件.人员介绍需要中国资源软件在线.'}]
    err = _400(c.get("/products/str"))
    assert err == "pid格式不正确！"
    corr = _200(c.get("/products/1"))
    assert corr == result


def test_get_comment_by_id():
    err = _400(c.get("/str"))
    assert err == "cid格式不正确！"
    corr = _200(c.get("/13"))
    assert corr == {'cid': 13,
                    'commentator': 70,
                    'product': 1,
                    'response_to': 0,
                    'text': '等级品牌单位的人朋友其实更新.详细详细系列单位密码问题.日期非常部门还是作者汽车工程.'}


def test_comment_update():
    corr = auth_verify(c.put, "/0")
    assert corr == "暂未实现！"
