<html>
<head>
    <title>上传文件</title>
</head>
<body>
<!--
    application/x-www-form-urlencoded   表示在发送前编码所有字符（默认）
    multipart/form-data	  不对字符编码。在使用包含文件上传控件的表单时，必须使用该值。
    text/plain	  空格转换为 "+" 加号，但不对特殊字符编码。
-->

    <form enctype="multipart/form-data" action="/upload" method="post">
        <input type="file" name="uploadFile" />
        <input type="hidden" name="token" value="{{.}}"/>
        <input type="submit" name="UPLOAD"/>
    </form>
</body>
</html>