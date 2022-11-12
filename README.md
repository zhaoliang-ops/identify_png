## png图片损坏识别方案
-  1、go Image/png/decode 源码
-  2、gm magick + crc32
-  3、gm magick + 开源工具pngcheck
-  4、imagemagick
- 1. 提供单元测试
- 1. 提供损坏png
- [ ] 想要不解码验证png是否损坏，做不到。直接使用go的源码，不再重复造轮子
