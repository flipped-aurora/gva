# 介绍
gva-ctl:一个友好的终端工具，方便使用 [gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin) 项目

# install

```go
go get github.com/flipped-aurora/gva
``` 

# 贡献指南
- Issue
    - issues 可提交bug和建议
    - 在提交 issue 之前，请搜索相关内容是否已被提出
- Pull Request
    - 请先 fork 一份到自己的项目下，不要直接在仓库下建分支。
    - commit 信息要以[文件名]: 描述信息 的形式填写，例如 README.md: fix xxx bug。
    - 确保 PR 是提交到 develop 分支，而不是 master 分支。
    - 如果是修复 bug，请在 PR 中给出描述信息。
    - 合并代码需要两名维护人员参与：一人进行 review 后 approve，另一人再次 review，通过后即可合并
- 版本信息
    - v0.0.1
        - 支持mysql(完美支持)与postgresql(慎用)
        
# 使用指南
- clone或者download [gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin) 项目
```shell script
git clone https://github.com/flipped-aurora/gin-vue-admin.git
```
- 进入server项目
```shell script
cd server
```
- 按需修改 [server](https://github.com/flipped-aurora/gin-vue-admin/tree/master/server) 的 [config.yaml](https://github.com/flipped-aurora/gin-vue-admin/blob/master/server/config.yaml) 文件
- 修改完成后使用 [gva](#install) 终端工具进行生成表和初始数据
```shell script
gva initdb
# 当你看以下文本输出就代表初始化成功了
[Mysql]-->初始化数据表成功
[Mysql]-->初始化数据成功
```