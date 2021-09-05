# liveuser

微信小程序的服务器后台。

删除过期的 access_token

    sqlite3 /root/.microsqlstore/io_goeasy.db "delete from auth where key like 'tokens/%' and expiry < datetime('now','localtime')"

删除过期的 log 和 trace

```
/opt/goeasy/log/goeasy.log
/root/.microtrace/trace.csv
{
    missingok
    daily
    copytruncate
    rotate 7
    notifempty
}
```

门户页面放在

    /portal

用户数据库放在

    /root/.liveuser/liveuser.db
