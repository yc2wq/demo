# tKeel 物联网Paas平台

[![codecov](https://codecov.io/gh/xujielong/demo/branch/master/graph/badge.svg?token=MR6NSOHHA9)](https://codecov.io/gh/xujielong/demo)

----

## tKeel 是什么

> [English](README.md) | 中文

[tKeel](https://docs.qingcloud.com/iot/) 是一个连接开发者、制造商及终端用户的开放物联网Paas平台。

tKeel 的发音是 ['tki։l]，它本意是船的龙骨，意指我们做的是物联网核心框架，目标是提供高扩展、插件化的开源的 IoT Platfor平台，建设功能丰富、面向行业的商业 Paas、SaaS 服务。

tKeel 目前最新的版本为 0.1.0，所有版本 100% 开源，关于 tKeel 更详细的介绍与说明请参阅 [什么是 tKeel](docs/introduction/what-is-tkeel.md)。


## 快速体验

使用体验账号 `demo1 / Demo123` 登录 [Demo 环境](https://iot.console.qingcloud.com/)，该账号仅授予了 view 权限，建议自行安装体验完整的管理功能。

## 架构

tKeel 采用了前后端分离的架构设计，后端的各个功能组件可通过 REST API 对接外部系统，详见 [架构说明](docs/introduction/architecture.md)。本仓库仅包含后端代码，前端代码参考 [Console 项目](https://github.com/xujielong/console)。

![Architecture](docs/images/architecture.png)

## 核心功能

|功能 |介绍 |
| --- | ---|
| 数据接入 | 提供海量设备进行快速数据接入能力 |
| 设备管理 | 提供设备管理能力 |
| 安全机制 | 提供插件隔离、用户角色管理、颁发设备Token等安全机制 |
| 插件管理 | 提供可插拔的基础插件管理功能 |
| 多租户管理 | 提供基于角色的细粒度多租户统一认证，提供多层级的权限管理 |
| 丰富的可观察性功能 | 提供基于多租户的日志查询与日志收集，支持应用层级的告警与通知 |

以上功能说明详见 [产品功能](docs/introduction/features.md)

## 最新发布

tKeel 0.1.0 将于 2021 年 9 月 30 日正式上线，点击 [Release Notes For 0.1.0](docs/release/release-v010.md) 查看 0.1.0 版本的更新详情。

## 安装 0.1.0

### 快速入门

[快速入门系列](docs/introduction/quick-start.md) 提供了快速安装与入门示例，供初次安装体验参考。

### 基于 Linux 安装 tKeel

- [在 Ubuntu 安装 tKeel](docs/introduction/install-tkeel-on-ubuntu.md)
- [在 CentOS 安装 tKeel](docs/introduction/install-tkeel-on-centos.md)

### 基于 Windows 安装 tKeel

- [在 Windows 安装 tKeel](docs/introduction/install-tkeel-on-windows.md)

## 技术社区

[tKeel 社区](docs/community/README.md) 包含所有社区的信息，包括如何开发，兴趣小组(SIG)等。比如[开发指南](docs/community/contribution/design-proposal-template.md) 详细说明了如何从源码编译、tKeel 的 GitHub 工作流、如何贡献代码以及如何测试等。

- [Bug 与建议反馈（GitHub Issue）](https://github.com/xujielong/demo/issues)


