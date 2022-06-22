# douyin
由base改进过来的后端抖音项目

项目作者： [whyy1](https://github.com/whyy1) </br>

项目使用技术栈
-------
* GIN
* GORM
* MYSQL
* 阿里云OSS</br>
演示文档地址:[演示文档](https://aio03fkuce.feishu.cn/file/boxcnwasGHVQCXHQX1HMH5tYnNd)</br>

项目说明
-------
项目结构
</br>![项目结构截图](https://y1-image.oss-cn-beijing.aliyuncs.com/image/%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_20220613154430.png)
* config中包含了一些数据库配置以及使用到阿里云OSS的配置。
* Controller层中进行参数校验，负责具体的业务模块流程的控制。
* Service层通过Controller传递的参数的进行业务逻辑的处理，然后调用Dao进行数据持久化。
* Dao层主要是做数据持久层的工作，与数据库有关的操作都在这里。
*util中存放了OSS对象存储的一些工具类
项目细节说明
