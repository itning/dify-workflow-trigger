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

# Introduction

Supports triggering Dify workflow execution via CRON scheduled tasks.

Supports multiple tasks with different schedules invoking different workflows.

![preview](./pic/preview.png)

# Usage

## Using Binary File

1. [Download the latest version](https://github.com/itning/dify-workflow-trigger/releases)
2. Execute
    ```shell
    ./dify-workflow-trigger -config config.json
    ```

## Using Docker

```shell
docker run --name dify-workflow-trigger \
    -v /path/to/config.json:/app/config.json \
    -d itning/dify-workflow-trigger:latest
```

# Configuration `config.json`

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