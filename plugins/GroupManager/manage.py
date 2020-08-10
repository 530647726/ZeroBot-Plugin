import re
import configparser

async def handling_group_messages(bot, msg, user_group, user_qq):
    if msg == "群管系统":
        await manage_system(bot, msg, user_group, user_qq)
    elif re.match(r"禁言(.+)\*(\d+)",msg):
        await ban(bot, msg, user_group, user_qq)
    elif re.match(r"解除禁言(.+)",msg):
        await ban_cancel(bot, msg, user_group, user_qq)
    elif msg == "全员禁言":
        await ban_all(bot, msg, user_group, user_qq)
    elif msg == "解除全员禁言":
        await ban_all_cancel(bot, msg, user_group, user_qq)
    elif re.match(r"警告(.+)",msg):
        await warm(bot, msg, user_group, user_qq)
    elif re.match(r"清空警告(.+)",msg):
        await warm_cancel(bot, msg, user_group, user_qq)
    elif re.match(r"踢出(.+)",msg):
        await kick(bot, msg, user_group, user_qq)
    elif re.match(r"退出群聊(.+)",msg):
        await leave(bot, msg, user_group, user_qq)
    elif re.match(r"修改群名片(.+)\*(.+)",msg):
        await card(bot, msg, user_group, user_qq)
    elif re.match(r"设置群头衔(.+)\*(.+)",msg):
        await title(bot, msg, user_group, user_qq)
    elif re.match(r"私聊转发(.+)\*(.+)",msg):
        await private_msg(bot, msg, user_group, user_qq)
    elif re.match(r"群聊转发(.+)\*(.+)",msg):
        await group_msg(bot, msg, user_group, user_qq)

async def handling_private_messages(bot, msg, user_group, user_qq):
    if msg == "群管系统":
        await manage_system(bot, msg, user_group, user_qq)
    elif re.match(r"退出群聊(.+)",msg):
        await leave(bot, msg, user_group, user_qq)
    elif re.match(r"私聊转发(.+)\*(.+)",msg):
        await private_msg(bot, msg, user_group, user_qq)
    elif re.match(r"群聊转发(.+)\*(.+)",msg):
        await group_msg(bot, msg, user_group, user_qq)

async def manage_system(bot, msg, user_group, user_qq):
    if user_group == "8888":
        message = "\
=======>群管系统<=======\n\
退出群聊[群号]\n\
私聊转发[QQ]/[@QQ]*[内容]\n\
群聊转发[QQ]/[@QQ]*[内容]\n\
======================="
        await bot.send_private_msg(user_id=user_qq, message=message)
    else:
        message = "\
=======>群管系统<=======\n\
全员禁言\n\
解除全员禁言\n\
禁言[QQ]/[@QQ]*[分钟]\n\
解除禁言[QQ]/[@QQ]\n\
警告[QQ]/[@QQ]\n\
清空警告[QQ]/[@QQ]\n\
踢出[QQ]/[@QQ]\n\
退出群聊[群号]\n\
修改群名片[QQ]/[@QQ]*[内容]\n\
设置群头衔[QQ]/[@QQ]*[内容]\n\
私聊转发[QQ]/[@QQ]*[内容]\n\
群聊转发[QQ]/[@QQ]*[内容]\n\
======================="
        await bot.send_group_msg(group_id=user_group, message=message)

async def ban(bot, msg, user_group, user_qq):
    content = re.findall(r"禁言(.+)\*(\d+)",msg)[0]
    ban_qq = content[0]
    ban_time = content[1]
    ban_time_m = int(ban_time)*60
    try:
        ban_qq = int(re.search(r'\d+', ban_qq).group())
    except:
        ban_qq = 12345678
    try:
        message = "[CQ:at,qq="+str(ban_qq)+"]你已经被禁言"+str(ban_time)+"分钟"
        await bot.set_group_ban(group_id=user_group, user_id=ban_qq, duration=ban_time_m)
        await bot.send_msg(group_id=user_group,message=message)
    except:
        message = "操作失败"
        await bot.send_msg(group_id=user_group,message=message)
        
async def ban_cancel(bot, msg, user_group, user_qq):
    content = re.findall(r"解除禁言(.+)",msg)[0]
    ban_qq = content
    try:
        ban_qq = int(re.search(r'\d+', ban_qq).group())
    except:
        ban_qq = 12345678
    try:
        message = "[CQ:at,qq="+str(ban_qq)+"]你已经被解除禁言"
        await bot.set_group_ban(group_id=user_group, user_id=ban_qq, duration=0)
        await bot.send_msg(group_id=user_group,message=message)
    except:
        message = "操作失败"
        await bot.send_msg(group_id=user_group,message=message)
        
async def ban_all(bot, msg, user_group, user_qq):
    try:
        message = "已开启全员禁言"
        await bot.set_group_whole_ban(group_id=user_group, enable="true")
        await bot.send_msg(group_id=user_group,message=message)
    except:
        message = "操作失败"
        await bot.send_msg(group_id=user_group,message=message)

async def ban_all_cancel(bot, msg, user_group, user_qq):
    try:
        message = "已解除全员禁言"
        await bot.set_group_whole_ban(group_id=user_group, enable="false")
        await bot.send_msg(group_id=user_group,message=message)
    except:
        message = "操作失败"
        await bot.send_msg(group_id=user_group,message=message)

async def warm(bot, msg, user_group, user_qq):
    content = re.findall(r"警告(.+)",msg)[0]
    warm_qq = content
    CONFIG_PATH = "./plugins/GroupManager/data/warm.ini"
    config = configparser.ConfigParser()
    config.read(CONFIG_PATH,encoding='UTF-8')
    try:
        warm_qq = int(re.search(r'\d+', warm_qq).group())
    except:
        warm_qq = 12345678
    try:
        warm_times = int(config[str(user_group)][str(warm_qq)])
    except:
        warm_times = 0
    if warm_times >= 5:
        try:
            await kick(bot, "踢出"+str(warm_qq), user_group, user_qq)
        except:
            message = "权限不足..."
            await bot.send_msg(group_id=user_group,message=message)
    else:
        try:
            warm_times = warm_times + 1
            if not config.has_section(str(user_group)) :
                config.add_section(str(user_group))
            config.set(str(user_group),str(warm_qq),str(warm_times))
            with open(CONFIG_PATH, 'w') as configfile:
                config.write(configfile)
            message = "[CQ:at,qq="+str(warm_qq)+"]你已经被警告"+str(warm_times)+"次\n超过5次则被踢出群聊"
            await bot.send_msg(group_id=user_group,message=message)
        except:
            message = "操作失败"
            await bot.send_msg(group_id=user_group,message=message)

async def warm_cancel(bot, msg, user_group, user_qq):
    content = re.findall(r"清空警告(.+)",msg)[0]
    warm_qq = content
    CONFIG_PATH = "./plugins/GroupManager/data/warm.ini"
    config = configparser.ConfigParser()
    config.read(CONFIG_PATH,encoding='UTF-8')
    try:
        warm_qq = int(re.search(r'\d+', warm_qq).group())
    except:
        warm_qq = 12345678
    try:
        config.remove_option(str(user_group),str(warm_qq))
        with open(CONFIG_PATH, 'w') as configfile:
            config.write(configfile)
        message = "警告次数已清空"
        await bot.send_msg(group_id=user_group,message=message)
    except:
        message = "暂无数据..."
        await bot.send_msg(group_id=user_group,message=message)
    
async def kick(bot, msg, user_group, user_qq):
    content = re.findall(r"踢出(.+)",msg)[0]
    kick_qq = content
    try:
        kick_qq = int(re.search(r'\d+', kick_qq).group())
    except:
        kick_qq = 12345678
    try:
        message = str(kick_qq)+"已经被踢出，请大家引以为戒"
        await bot.set_group_kick(group_id=user_group, user_id=kick_qq)
        await bot.send_msg(group_id=user_group,message=message)
    except:
        message = "操作失败"
        await bot.send_msg(group_id=user_group,message=message)
        
async def leave(bot, msg, user_group, user_qq):
    content = re.findall(r"退出群聊(.+)",msg)[0]
    leave_group = content
    try:
        leave_group = int(re.search(r'\d+', leave_group).group())
    except:
        leave_group = 12345678
    try:
        message = "已退出群聊"+str(leave_group)
        await bot.set_group_leave(group_id=leave_group)
    except:
        message = "操作失败"
    if user_group == "8888":
        await bot.send_private_msg(user_id=user_qq, message=message)
    else:
        await bot.send_group_msg(group_id=user_group, message=message)

async def card(bot, msg, user_group, user_qq):
    content = re.findall(r"修改群名片(.+)\*(.+)",msg)[0]
    card_qq = content[0]
    card_name = content[1]
    try:
        card_qq = int(re.search(r'\d+', card_qq).group())
    except:
        card_qq = 12345678
    try:
        await bot.set_group_card(group_id=user_group, user_id=card_qq, card=card_name)
    except:
        message = "操作失败"
        await bot.send_msg(group_id=user_group,message=message)

async def title(bot, msg, user_group, user_qq):
    content = re.findall(r"设置群头衔(.+)\*(.+)",msg)[0]
    title_qq = content[0]
    title_name = content[1]
    try:
        title_qq = int(re.search(r'\d+', title_qq).group())
    except:
        title_qq = 12345678

    try:
        #set_group_special_title成功也会报错
        await bot.set_group_special_title(group_id=user_group,user_id=title_qq,special_title=title_name)
    except:
        #随便一句
        message = "操作失败"
        

async def private_msg(bot, msg, user_group, user_qq):
    content = re.findall(r"私聊转发(.+)\*(.+)",msg)[0]
    private_msg_qq = content[0]
    private_msg = content[1]
    try:
        private_msg_qq = int(re.search(r'\d+', private_msg_qq).group())
    except:
        private_msg_qq = 12345678
    try:
        message = "转发成功！"
        await bot.send_private_msg(user_id=private_msg_qq, message=private_msg)
    except:
        message = "操作失败"
    if user_group == "8888":
        await bot.send_private_msg(user_id=user_qq, message=message)
    else:
        await bot.send_group_msg(group_id=user_group, message=message)

async def group_msg(bot, msg, user_group, user_qq):
    content = re.findall(r"群聊转发(.+)\*(.+)",msg)[0]
    group_msg_group = content[0]
    group_msg = content[1]
    try:
        group_msg_group = int(re.search(r'\d+', group_msg_group).group())
    except:
        group_msg_group = 12345678
    try:
        message = "转发成功！"
        await bot.send_group_msg(group_id=group_msg_group, message=group_msg)
    except:
        message = "操作失败"
    if user_group == "8888":
        await bot.send_private_msg(user_id=user_qq, message=message)
    else:
        await bot.send_group_msg(group_id=user_group, message=message)

