<html>
<head>
<title>test form input</title>
<body>
<form action="/login" method="post">
    <input type="checkbox" name="interest" value="football"> football
    <input type="checkbox" name="interest" value="basketball"> basketball
    <input type="checkbox" name="interest" value="tennis"> tennis
    用户名： <input type="text" name="username">
    密码:   <input type="password" name="password">
    <!-- 在模版里面增加了一个隐藏字段token，这个值我们通过MD5(时间戳)来获取唯一值，
    然后我们把这个值存储到服务器端session来控制 -->
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="登录">
</form>
</body>
</head>
</html>