## 免责
本工具仅用于个人学习用途，不得用于违法用途和盈利传播，否则造成的后果自负，请在下载后 24 小时内删除！

## 说明
此工具只能解密在当前电脑下可以进行正常打开的文件，若文件无法在当前电脑打开则无法解密。即要求电脑正常安装加密软件和登录账号，可以在没有 “特权解密” 权限下完成解密。

## 编译
```bash
cd main && go build -ldflags "-H windowsgui" -o ../dist/Acrobat.exe main.go
cd ../unlock && go build -ldflags "-H windowsgui" -o ../dist/Unlock.exe main.go

```

## 使用
目前仅针对 .pdf .docx .doc .xlsx .xls .pptx .ppt 格式可解密。

第一次运行系统可能会进行病毒扫描，导致解密时间较长。

步骤：
1. 把 Acrobat.exe 和 Unlock.exe 复制到需要解密的文件夹中
2. 双击 Acrobat.exe 会自动解密当前所在文件夹及子文件夹中所有文件（也可通过 `Acrobat.exe -d`  来指定遍历起始目录）

> 如果无法解密，可以将 Acrobat.exe 重命名为 wps.exe 或者 winword.exe 试试

**各公司加密策略不尽相同，不保证一定可用**

## 原理
亿赛通管理员会在服务端维护能打开加密文档的软件进程白名单，因此可以将本工具的名字改为与其相同的名字，通过模仿其打开行为来触发自解密流程以获取到明文数据，然后将明文数据复制到临时文件中，通过调用另一程序 Unlock.exe 来重命名为正常文件。

## 参考
[zhangchye/YiSaiTongUnLock](https://github.com/zhangchye/YiSaiTongUnLock)

[yltchina/YiSaiTong](https://github.com/yltchina/YiSaiTong)

