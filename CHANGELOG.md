# 更新日志 


## [1.3.5](///compare/v1.3.4...v1.3.5) (2025-03-23)


### ✨ 新增特性

* 支持rz中断 (不支持sz中断) 70cbdce
* 转发策略新增默认运行已建立的tcp连接通过 1ca41e6, closes #12


### 📦  其他更新

* 自动生成文件更新 04131b4
* go代码自动格式化更新 420076c


### 🎉 样式相关

* 调整ip输入框为动态扩展 fad295f

## [1.3.4](///compare/v1.3.3...v1.3.4) (2024-11-05)


### 🐞 Bug 修复

* **custom:** 脚本升级脚本不生效 7281919
* **custom:** 修复更改转发数据参数时其他无关参数报错导致panic af5c77f


### 📦  其他更新

* 提交格式移除scope f392104
* **custom:** 提交格式调整 f038eab
* **release:** 1.3.4 ea3fb16


### 📝 文档更新

* **custom:** 文档错误修正 d87f965

## [1.3.3](///compare/v1.3.2...v1.3.3) (2024-10-23)


### ✨ 新增特性

* **custom:** 启动端口时根据修改数据库 1cc0bea
* **custom:** 自动检测主机防火墙类型 a752e20
* **custom:** ip完整dnat映射 7198019


### 📦  其他更新

* **release:** 1.3.3 0f925ea

## [1.3.2](///compare/v1.3.1...v1.3.2) (2024-10-11)


### 🐞 Bug 修复

* **custom:** 源地址转换位置错误 de02426
* **custom:** dnat目的地址修改不生效 eb009a2


### 📦  其他更新

* **release:** 1.3.2 3963365

## [1.3.1](///compare/v1.3.0...v1.3.1) (2024-10-11)


### ✨ 新增特性

* **custom:** 调整dnat让其更丰富灵活 399ce06


### 🐞 Bug 修复

* **custom:** 完善构建脚本中的一些问题 5b08372
* **custom:** 修复默认策略失效问题 d564248
* **custom:** 源地址转换编辑显示问题 3ed3a3d


### 📦  其他更新

* **release:** 1.3.1 7165a9e


### 📝 文档更新

* **custom:** 更新配置文件 5b26a4f
* **custom:** 修改文件 6feb69b


### 🎉 样式相关

* **custom:** 网卡选择优化，目的ip仅支持网卡配置的ip abae287

## [1.3.0](///compare/v1.2.1...v1.3.0) (2024-10-08)


### ✨ 新增特性

* **custom:** 支持iptables 5546b3b, closes #1


### ⏪ 代码回退

* **custom:** 移除版本更新通知 209292e


### 📦  其他更新

* **custom:** 杂项 7681b27
* **release:** 1.3.0 11ed797

## [1.2.1](///compare/v1.2.0...v1.2.1) (2024-09-29)


### ✨ 新增特性

* 增加docker部署方式 d41bf31
* **custom:** 新增配置文件定义链优先级 0344176
* **custom:** 新增页面开启数据转发功能 016d46a
* **custom:** 新增自动构建脚本 6edefe4, closes #4


### 🐞 Bug 修复

* **custom:** 删除comment参数 59263df
* **custom:** 新增启动执行策略初始化 f4cf72a
* **custom:** 修复webshell在容器中无法使用 a810999
* **custom:** web列表显示错误 f581eb0


### ⚡ 性能优化

* **custom:** 完善docker相关部署 eafc0b7


### 📦  其他更新

* **release:** 1.2.1 356e891


### 📝 文档更新

* **custom:** 修改readme错误 432a5a7


### ♻ 代码重构

* **custom:** \!重构本地策略 0710762
* **custom:** \!重构转发策略 cae501f
* **custom:** 移除旧策略代码 a7da0a0


### ✅ 测试用例

* **custom:** 完善构建脚本 0de214e

## [1.2.0](///compare/v1.1.0...v1.2.0) (2024-09-12)


### ✨ 新增特性

* **custom:** 新增转发策略 0407a3d
* **custom:** 新增转发流控功能 f1ab802
* **custom:** 新增dnat功能 f5bb1e0
* **custom:** 新增shell操作回放 cc5d6e1
* **custom:** 新增snat功能 d9c42a6


### 🐞 Bug 修复

* **custom:** 出站流控不生效 5a7eec1
* **custom:** 创建规则位置不生效 5833c43
* **custom:** 规则替换报错 1bb4793
* **custom:** 完善转发策略搜索功能 829e0ae
* **custom:** 修复不能使用rz sz命令问题 ffd6766
* **custom:** 修复转发策略bug 7cf4b23


### 📦  其他更新

* **custom:** 一些兼容性更新 ca57667
* **release:** 1.2.0 b8cf9f7


### 📝 文档更新

* **custom:** 更新readme文档 8f37edc


### 🎉 样式相关

* **custom:** 调整样式 95e2acc
* **custom:** 首页删除无效内容 46ee1dd
* **custom:** 修复一些错误提示 aa128de


### ✅ 测试用例

* **custom:** 移除多余文件 b65ba3a

## [1.1.0](///compare/726b5e80afe7cc584c6b1b2119213da7ea5dd54f...v1.1.0) (2024-09-04)


### ✨ 新增特性

* **custom:** 新版本发布 a698680


### 📦  其他更新

* 构建项目 726b5e8
* **release:** 1.1.0 dfcce0f
