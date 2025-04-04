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

[中文](https://github.com/itning/dify-workflow-trigger/blob/main/README-CN.md)

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

[![Docker Pulls](https://img.shields.io/docker/pulls/itning/dify-workflow-trigger.svg?style=flat&label=pulls&logo=docker)](https://hub.docker.com/r/itning/dify-workflow-trigger/tags?page=1&ordering=last_updated)

```shell
docker run --name dify-workflow-trigger \
    -v /path/to/config.json:/app/config.json \
    -d itning/dify-workflow-trigger:latest
```

# Configuration `config.json`

```json
[
   {
      "name": "test1",
      "cron": "TZ=Asia/Shanghai 0 0 18 * * *",
      "url": "http://your-dify-domain/api/v1/workflows/run",
      "token": "app-FZ8vjeH74tUBtRYNUjFx65aw",
      "body": {
         "inputs": {
         },
         "response_mode": "streaming",
         "user": "dify-workflow-trigger"
      }
   },
   {
      "name": "test2",
      "cron": "TZ=Asia/Shanghai 0 30 8 * * *",
      "url": "http://your-dify-domain/v1/workflows/run",
      "token": "app-LoFN2hiaYCMfTIhN8yCDmWmo",
      "body": {
         "inputs": {
         },
         "response_mode": "streaming",
         "user": "dify-workflow-trigger"
      }
   }
]
```
cron use [go-co-op/gocron](https://pkg.go.dev/github.com/go-co-op/gocron/v2#CronJob) 

support seconds level scheduling: Seconds Minutes Hours Day-of-Month Month Day-of-Week

example:
- `TZ=Asia/Shanghai 0 0 18 * * *` run at 18:00 every day (timezone: Asia/Shanghai)
- `TZ=Asia/Shanghai 0 30 8 * * *` run at 18:30 every day (timezone: Asia/Shanghai)
- `0 0 18 * * *` run at 18:00 every day (default timezone)
- `*/5 * * * * *` run every 5 seconds (default timezone)