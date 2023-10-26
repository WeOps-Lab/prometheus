# 支持从外部URL获取Target和Rule的配置

# 动态规则

新增了EXTRA_RULES_URL环境变量，用于指定额外的Rule配置文件URL，用于动态更新规则

```
[
  {
    "id": "1",
    "group": "weops",
    "alert": "服务异常",
    "expr": "status==0",
    "for": 60,
    "keepFiringFor": 60,
    "labels": [
      {
        "key": "pruducer",
        "value": "weops"
      }
    ],
    "annotations": [
      {
        "key": "pruducer",
        "value": "weops"
      }
    ]
  }
]
```

EXTRA_RULES_URL指向目标URL地址即可

# 动态Target

新增了EXTRA_SCRAPE_TARGET环境变量，用于动态获取Prometheus主动Scrape的对象，响应示例如下：

```
[
  {
    "id": "1",
    "name": "example",
    "metricsPath": "/metrics",
    "labels": [
      
    ],
    "scheme": "http",
    "interval": 30,
    "timeout": 10,
    "describe": "Example scrape group",
    "authUser": "",
    "authPassword": "",
    "authToken": "",
    "scrapeTarget": [
      {
        "id": "1",
        "target": "example.com:8080",
        "labels": [
          {
            "key": "1",
            "value": "2"
          }
        ],
        "groupId": "1",
        "describe": "Example scrape target"
      },
      {
        "id": "2",
        "target": "example.com:8081",
        "labels": [
          {
            "key": "1",
            "value": "2"
          }
        ],
        "groupId": "1",
        "describe": "Another example scrape target"
      }
    ]
  }
]
```
