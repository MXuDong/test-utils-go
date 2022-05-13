# test-utils

提供了很多用于测试的工具。

需要添加 go 参数：-gcflags=-l

---

**注意，本工具是危险的，不可使用的，方法是直接修改内存的，如果稍有不慎，可能造成重要数据的丢失哦!**
任何数据丢失，本工具不负责！

**Note, this tool is dangerous, can not be used, the method is directly modify memory, if a little careless, may cause
the loss of important data oh!**

This tool is not responsible for any data loss!

---

## 思路来源

[https://bou.ke/blog/monkey-patching-in-go/](https://bou.ke/blog/monkey-patching-in-go/)
[https://blog.csdn.net/qq_42038407/article/details/124705040?spm=1001.2014.3001.5501](https://blog.csdn.net/qq_42038407/article/details/124705040?spm=1001.2014.3001.5501)

注意本仓库只是按照理论实现了一些方法

## 主要功能支持

### 时间控制

时间控制将会直接修改golang/runtime/time.go 中的 time.Now() 的方法，因此会暂停整个应用的时间流逝。

示例：
[freeze_time.go](./example/freeze_time.go)
  