# WPM 数据交换格式

WPM 数据交换格式是一个 properties 文件， 支持分段（类似 .git/config 文件）。

一个简单的示例如下：


    [location "11111111111111111111111111"]
        uri=https://a.c/b/c/d
        type=file
        isdir=false
        urn=urn:type:id:xxx
        content-type=

    [location "222222222222222222222222222"]
        uri=file:/a/x/y/z
        type=file
        isdir=false
        urn=urn:type:id:xxx
        content-type=

