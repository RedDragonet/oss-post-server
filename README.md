# oss-post-server
simple aliyun oss file post http server
简单的阿里云oss文件上传http服务

主要目的方便其他软件通过http上传图片到oss，如图床服务

请求参数
```
endpoint
secret
key
bucket
domain => 文件访问域名
filename_type => 默认为原文件名称, random 时间名称
```
