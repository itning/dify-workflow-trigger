<h3 align="center">Dify Workflow Trigger</h3>
<div align="center">

[![GitHub stars](https://img.shields.io/github/stars/itning/dify-workflow-trigger.svg?style=social&label=Stars)](https://github.com/itning/dify-workflow-trigger/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/itning/dify-workflow-trigger.svg?style=social&label=Fork)](https://github.com/itning/dify-workflow-trigger/network/members)
[![GitHub watchers](https://img.shields.io/github/watchers/itning/dify-workflow-trigger.svg?style=social&label=Watch)](https://github.com/itning/dify-workflow-trigger/watchers)
[![GitHub followers](https://img.shields.io/github/followers/itning.svg?style=social&label=Follow)](https://github.com/itning?tab=followers)


</div>

<div align="center">

[![GitHub issues](https://img.shields.io/github/issues/itning/dify-workflow-trigger.svg)](https://github.com/itning/dify-workflow-trigger/issues)
[![GitHub license](https://img.shields.io/github/license/itning/dify-workflow-trigger.svg)](https://github.com/itning/dify-workflow-trigger/blob/master/LICENSE)
[![GitHub last commit](https://img.shields.io/github/last-commit/itning/dify-workflow-trigger.svg)](https://github.com/itning/dify-workflow-trigger/commits)
[![GitHub repo size in bytes](https://img.shields.io/github/repo-size/itning/dify-workflow-trigger.svg)](https://github.com/itning/dify-workflow-trigger)
[![Hits](https://hitcount.itning.com?u=itning&r=dify-workflow-trigger)](https://github.com/itning/hit-count)

</div>

---

# 简介

支持定时任务CRON触发Dify工作流执行。

支持多个任务不同时间不同任务调用

![preview](./pic/preview.png)

# 使用

## 使用二进制文件

1. [下载最新版本](https://github.com/itning/dify-workflow-trigger/releases)
2. 执行
    ```shell
    ./dify-workflow-trigger -config config.json
    ```
   
## 使用Docker

```shell
docker run --name dify-workflow-trigger \
    -v /path/to/config.json:/app/config.json \
    -d itning/dify-workflow-trigger:latest
```

# 配置 config.json

```json
[
  {
    "name": "workflow-1",
    "cron": "*/30 * * * * ?",
    "url": "https://dify.itning.cn/v1/workflows/run",
    "token": "app-FZ8vjeH74tUBtRYNUjFx65aw",
    "body": {
      "inputs": {
      },
      "response_mode": "streaming",
      "user": "dify-workflow-trigger"
    }
  }
]
```