## png图片损坏识别方案
### 1、go Image/png/decode 源码
### 2、gm magick + crc32
### 3、gm magick + 开源工具pngcheck
### 4、imagemagick
## 提供单元测试
## 提供损坏png
## 结论，想要不解码验证png是否损坏，crc是否正确做不到。直接使用go的源码，不再重复造轮子
