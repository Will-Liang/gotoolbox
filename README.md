# gotoolbox

<div align=center>
    ![Go version](https://img.shields.io/badge/go-%3E%3Dv1.20-9cf)
    [![Release](https://img.shields.io/badge/release-2.2.5-green.svg)](https://github.com/Will-Liang/gotoolbox/releases)
    [![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/duke-git/lancet/blob/main/LICENSE)
</div>



This is a Go utility library.



## 安装

```
go get github.com/Will-Liang/gotoolbox@v0.0.7
```



## 文档

### <h2 id="index">目录</h2>

- [fileutils](#fileutils)
- [jsonutils](#jsonutils)
- [logutils](#logutils)
- [maputils](#maputils)
- [requestutils](#requestutils)
- [sliceutils](#sliceutils)
- [strutils](#strutils)
- [timeutils](#timeutils)
- [utils](#utils)

#### <h2 id="fileutils">1. fileutils 包实现文件和路径有关的函数</h2>


函数列表：

- **CheckPath**: 检查路径
- **CreateDir**: 创建目录
- **DeleteDirFiles**: 删除给定目录下的所有文件，不删除目录
- **ListFiles**: 遍历指定目录下的所有文件，返回文件路径和文件名称的切片

#### <h2 id="jsonutils">2. jsonutils 包实现json以及json与文件交互的有关函数</h2>


函数列表：

- **FileToJsonMap**: 读取标准json格式，返回map[string]interface{}
- **FileToJsonObject**: 读取标准josn格式,适用于`[`开头,`]`结尾
- **GetValueWithKeyFromFile**: 根据key从json文件中获得value
- **JsonSliceToLineFile**: 将结构体切片逐行写入到文件
- **JsonToFile**: 将json写入到文件
- **JsonToFileFormat**: 将json格式化后写入到文件
- **JsonToLineFile**: 将json追加到文件中的最后一行

#### <h2 id="logutils">3. logutils 包实现记录日志的有关函数</h2>



函数列表：

- **GetLog**: 获得一个日志信息
#### <h2 id="maputils">4. maputils 包实现记录日志的有关函数</h2>



函数列表：

- **ExistKey**: map是否存在该key



#### <h2 id="requestutils">5. requestutils包实现记录日志的有关函数</h2>



函数列表：

- **Get**: Get请求

#### <h2 id="sliceutils">6. sliceutils 包实现切片的有关函数</h2>



函数列表：

- **IsSlice**: 判断传入的类型是不是切片

#### <h2 id="strutils">7. strutils 包实现字符串的有关函数</h2>



函数列表：

- **IsAlphabet**: 判断字符串是否纯英文
- **TraditionalConvertSimple**: 中文繁体转简体
- **SimpleConvertTraditional**: 中文简体转繁体
- **StrToFloat**: 将字符串转成float
- **StrToInt**: 将字符串转成Int
- **SplitStrWithDelimiters**: 使用分隔符将字符串分成切片

#### <h2 id="timeutils">8. timeutils 包实现时间的有关函数</h2>



函数列表：

- **GetCurrentUnix**: 获取当前时间的时间戳，长度为13位(以毫秒为单位)
- **GetHourMin**: 获得时分


#### <h2 id="utils">9. utils 包实现一些基本工具函数</h2>



函数列表：

- **FillURLWithMap**: 将map中的各个参数填充到url中
- **RandomSleep**: 在两个整数中间随机随眠若干秒









