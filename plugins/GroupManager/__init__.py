# -*- coding: utf-8 -*-

from nonebot import *
from nonebot import permission as perm
from . import manage
from .config import *


bot = get_bot()

@bot.on_message("group")
async def manage_group_entrance(context):
    msg = str(context["message"])
    user_group = context["group_id"]
    user_qq = context["user_id"]
    if str(user_qq) in PERMISSION:
        await manage.handling_group_messages(bot, msg, user_group, user_qq)

@bot.on_message("private")
async def manage_private_entrance(context):
    msg = str(context["message"])
    user_group = "8888"
    user_qq = context["user_id"]
    if str(user_qq) in PERMISSION:
        await manage.handling_private_messages(bot, msg, user_group, user_qq)
