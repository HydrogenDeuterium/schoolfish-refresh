import httpx

from test import fake, url
from test.util import token_verify, success, error

c = httpx.Client(base_url=f"{url}/message", timeout=1000)


def random_message():
    return {"text": fake.sentence()}


def test_get_all_messages():
    token = token_verify(c.get, f"")
    corr = success(c.get(f"", headers=token))
    assert corr == [
        {'from': 70, 'mid': 1, 'text': '', 'to': 71},
        {'from': 70, 'mid': 2, 'text': '', 'to': 72},
        {'from': 70, 'mid': 3, 'text': '', 'to': 71},
        {'from': 70, 'mid': 4, 'text': '合作到了完全文章任何非常发布.', 'to': 71},
        {'from': 70, 'mid': 5, 'text': '关于你们不能朋友支持信息.', 'to': 71},
        {'from': 70, 'mid': 6, 'text': '不能能够设备经营进入.', 'to': 71},
        {'from': 70, 'mid': 7, 'text': '政府推荐推荐问题他们就是.', 'to': 71},
        {'from': 70, 'mid': 8, 'text': '电子不要之后用户密码注册大家比较.', 'to': 71},
        {'from': 70, 'mid': 9, 'text': '开发表示业务自己免费正在应该都是.', 'to': 71},
        {'from': 70, 'mid': 10, 'text': '手机经营的人法律密码商品也是.', 'to': 72}]


def test_send_message():
    token = token_verify(c.post, f"/{72}")
    message = random_message()
    err0 = error(c.post(f"/str", data=message, headers=token))
    err1 = error(c.post(f"/0", data=message, headers=token))
    err2 = error(c.post(f"/70", data=message, headers=token))
    assert err0 == err1 == err2 == "不正确的uid！"
    corr = success(c.post(f"/{72}", data=message, headers=token))
    assert isinstance(corr["mid"], int)
    assert corr['mid'] > 0
    del corr['mid']
    assert corr == message | {"from": 70, "to": 72}


def test_get_messages():
    token = token_verify(c.get, f"/{71}")
    corr = success(c.get(f"/{71}", headers=token))
    assert corr == [
        {'from': 70, 'mid': 1, 'text': '', 'to': 71},
        {'from': 70, 'mid': 3, 'text': '', 'to': 71},
        {'from': 70, 'mid': 4, 'text': '合作到了完全文章任何非常发布.', 'to': 71},
        {'from': 70, 'mid': 5, 'text': '关于你们不能朋友支持信息.', 'to': 71},
        {'from': 70, 'mid': 6, 'text': '不能能够设备经营进入.', 'to': 71},
        {'from': 70, 'mid': 7, 'text': '政府推荐推荐问题他们就是.', 'to': 71},
        {'from': 70, 'mid': 8, 'text': '电子不要之后用户密码注册大家比较.', 'to': 71},
        {'from': 70, 'mid': 9, 'text': '开发表示业务自己免费正在应该都是.', 'to': 71},
        {'from': 70, 'mid': 21, 'text': '具有音乐出现完成.', 'to': 70},
        {'from': 70, 'mid': 22, 'text': '社会一些作为.', 'to': 70}]
