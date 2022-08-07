# douyin

由 base 改进过来的后端抖音项目

项目作者： [whyy1](https://github.com/whyy1) </br>

### Quickly Start

pre-dependencies

```bash
# mac os
brew install golang-migrate
```

## 项目使用技术栈

- GIN
- GORM
- MYSQL
- 阿里云 OSS</br>
  演示文档地址:[演示文档](https://aio03fkuce.feishu.cn/file/boxcnwasGHVQCXHQX1HMH5tYnNd)</br>

## 项目说明

项目结构
</br>![项目结构截图](https://y1-image.oss-cn-beijing.aliyuncs.com/image/%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_20220613154430.png)

- config 中包含了一些数据库配置以及使用到阿里云 OSS 的配置。
- Controller 层中进行参数校验，负责具体的业务模块流程的控制。
- Service 层通过 Controller 传递的参数的进行业务逻辑的处理，然后调用 Dao 进行数据持久化。
- Dao 层主要是做数据持久层的工作，与数据库有关的操作都在这里。
  \*util 中存放了 OSS 对象存储的一些工具类
  项目细节说明
