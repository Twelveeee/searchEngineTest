# 搜索引擎调研

Meilisearch vs Typesense vs Algolia 

这三个引擎进行实际的调研，

他们的区别 ：
[meilisearch](https://www.meilisearch.com/docs/learn/what_is_meilisearch/comparison_to_alternatives) 、 [typesense](  https://typesense.org/typesense-vs-algolia-vs-elasticsearch-vs-meilisearch/)



官网

Meilisearch ：https://www.meilisearch.com/

Typesense：https://typesense.org/

Algolia：https://www.algolia.com/



Algolia 为第三方托管的服务，没办法自己搭建，所以没有压测

# 背景

服务器为一台4C8G的机器
数据量大概为5000条左右document

关注的指标，性能和响应时间，中文搜索能力

# 流程

建立测试环境

导入测试数据

进行 10次查询，取查询耗时平均值

进行一分钟内进行不限次数查询，记录总查询次数，平均查询耗时，最高查询耗时，最低查询耗时

进行中文关联性查询，手动评分



# 环境说明

测试机器一台，4C8G，cpu为intel N100

数据为 3000条

meilisearch 1.5.0

typesense 0.25.1

# 开始

## Install

### Pre-compiled executables

Get them [here](https://github.com/Twelveeee/searchEngineTest/releases).

## Useage manual

```bash
NAME:
   searchEngineTest - Twelveeee

USAGE:
   searchEngineTest [global options] command [command options] [arguments...]

VERSION:
   v0.0.1 development

DESCRIPTION:
   搜索引擎测试

COMMANDS:
   MeillSearchImport, mi        init MeillSearch data
   MeillSearchSearch, ms        search MeillSearch data once
   MeillSearchCreateIndex, mci  create MeillSearch index
   MeillSearchDeleteIndex, mdi  delete MeillSearch index
   MeillSearchTest, mt          test MeillSearch
   TypeSenseImport, ti          init TypeSense data
   TypeSenseSearch, ts          search TypeSense data once
   TypeSenseCreateIndex, tci    create TypeSense index
   TypeSenseDeleteIndex, tdi    delete TypeSense index
   TypeSenseTest, tt            test TypeSense
   help, h                      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config value, -c value  config file path (default: "./config.yaml")
   --help, -h                show help
   --version, -v             print the version

COPYRIGHT:
   (c) 2023 Twelveeee @ Twelveeee
```



### config

```yaml
# 导入数据存放文件未知
DataFile: "./data/data.json"
MeillSearch:
    Host:  "http://localhost:7700"
    APIKey: "aSampleMasterKey"
    IndexName: "article"
Typesense:
    Host:  "http://localhost:8108"
    APIKey: "xvRbulP1P0Rw3h9ZFuT8yQH0sc35JLLj9SkwGPCyDbrkPDIp"
    IndexName: "article"
# 测试
TestRate:
    # 每秒次数
    PerSecond: 100
    # 持续时间
    Duration: 60
```



### data demo

这里提供了10条数据，实际测试数据大约为3000条

```json
[
    {"rid": "09b983b090264286865e3bcef616cb8a", "tags": "教程 Python", "author": "michaelliao", "author_avatar": "https://img.hellogithub.com/github_avatar/470058.png!small", "name": "awesome-python-webapp", "full_name": "michaelliao/awesome-python-webapp", "title": "Python 小白的入门实战教程", "description": "小白的Python入门教程实战篇：网站+iOS App源码", "summary": "廖老师的免费 Python 入门教程，实践部分的代码。", "lang_color": "#24292e", "primary_lang": "Other", "stars": 2151, "stars_str": "2.2k", "publish_at": "2014-12-28 13:13:43", "has_chinese": true, "is_active": false},
    {"rid": "61be747e0e14438c804bed7db016d6b5", "tags": "Web 应用 Tornado Python 归档", "author": "phith0n", "author_avatar": "https://img.hellogithub.com/github_avatar/5711185.png!small", "name": "Minos", "full_name": "phith0n/Minos", "title": "基于 Tornado 的简约社区系统", "description": "一个基于Tornado/mongodb/redis的社区系统。", "summary": "一个基于 Tornado+MongoDB+Redis 构建的社区系统。", "lang_color": "#f1e05a", "primary_lang": "JavaScript", "stars": 681, "stars_str": "681", "publish_at": "2019-08-10 22:41:31", "has_chinese": true, "is_active": false},
    {"rid": "411d211476a8475dbafa75dc223d27e0", "tags": "数据分析 Python", "author": "waditu", "author_avatar": "https://img.hellogithub.com/github_avatar/10395504.jpeg!small", "name": "tushare", "full_name": "waditu/tushare", "title": "金融数据分析的 Python 工具包", "description": "TuShare is a utility for crawling historical data of China stocks", "summary": "这是一个免费、开源的 Python 财经数据接口包，实现了对股票、期货等金融数据，从数据采集、清洗加工到数据存储过程。", "lang_color": "#3572A5", "primary_lang": "Python", "stars": 12388, "stars_str": "12.4k", "publish_at": "2020-03-04 22:36:33", "has_chinese": true, "is_active": false},
    {"rid": "2631e0aa82db48aebcea0f0c7968d1b9", "tags": "Python", "author": "wong2", "author_avatar": "https://img.hellogithub.com/github_avatar/321947.jpeg!small", "name": "beijing_bus", "full_name": "wong2/beijing_bus", "title": "北京实时公交查询工具", "description": "北京实时公交 for Python", "summary": "该项目可以查询北京公交到达某站还需多久。", "lang_color": "#3572A5", "primary_lang": "Python", "stars": 373, "stars_str": "373", "publish_at": "2018-01-30 12:08:22", "has_chinese": true, "is_active": false},
    {"rid": "fc76d2b8c6574257b0256850ea309050", "tags": "阿里 TypeScript React JavaScript", "author": "ant-design", "author_avatar": "https://img.hellogithub.com/github_avatar/12101536.png!small", "name": "ant-design", "full_name": "ant-design/ant-design", "title": "一套企业级 UI 设计语言和 React 组件库", "description": "An enterprise-class UI design language and React UI library", "summary": "该项目是阿里开源的一套开箱即用的 React 组件库，视觉风格现代化，可用于快速构建企业级的中、后台管理系统。", "lang_color": "#2b7489", "primary_lang": "TypeScript", "stars": 88053, "stars_str": "88.1k", "publish_at": "2023-10-25 23:19:16", "has_chinese": true, "is_active": true},
    {"rid": "eed9fd7cf0a844cca6a7b62d06d07fec", "tags": "Lua Nginx Python 归档", "author": "alexazhou", "author_avatar": "https://img.hellogithub.com/github_avatar/9353779.jpeg!small", "name": "VeryNginx", "full_name": "alexazhou/VeryNginx", "title": "一个功能强大且友好的 Nginx 扩展项目", "description": " A very powerful and friendly  nginx base on lua-nginx-module( openresty ) which provide WAF, Control Panel, and Dashboards. ", "summary": "基于 lua_nginx_module(openrestry) 开发，实现了高级的防火墙、访问统计和 Web 界面等功能的 Nginx 扩展程序。", "lang_color": "#000080", "primary_lang": "Lua", "stars": 5916, "stars_str": "5.9k", "publish_at": "2019-10-26 23:19:02", "has_chinese": true, "is_active": false},
    {"rid": "d50eb59a63c74ae2b9d3f841abbad86f", "tags": "JavaScript 归档", "author": "disjukr", "author_avatar": "https://img.hellogithub.com/github_avatar/690661.png!small", "name": "activate-power-mode", "full_name": "disjukr/activate-power-mode", "title": "爆炸输入效果", "description": "Activate POWER MODE anywhere", "summary": "采用 JavaScript 实现的炫酷输入效果。", "lang_color": "#f1e05a", "primary_lang": "JavaScript", "stars": 416, "stars_str": "416", "publish_at": "2017-04-29 16:52:14", "has_chinese": false, "is_active": false},
    {"rid": "bbe2e588690446ccba3a06fcd2ea944d", "tags": "CSS", "author": "sofish", "author_avatar": "https://img.hellogithub.com/github_avatar/153183.png!small", "name": "typo.css", "full_name": "sofish/typo.css", "title": "用于构建适合中文阅读网页的 CSS", "description": "中文网页重设与排版：一致化浏览器排版效果，构建最适合中文阅读的网页排版", "summary": "该项目提供一致化浏览器排版效果的同时，构建最适合中文阅读的网页排版，支持桌面和移动平台。", "lang_color": "#e34c26", "primary_lang": "HTML", "stars": 4442, "stars_str": "4.4k", "publish_at": "2020-07-17 11:15:31", "has_chinese": true, "is_active": false},
    {"rid": "4949f0ae191a4967a20f120ffaac6798", "tags": "书籍 Python", "author": "yidao620c", "author_avatar": "https://img.hellogithub.com/github_avatar/1510785.jpeg!small", "name": "python3-cookbook", "full_name": "yidao620c/python3-cookbook", "title": "《Python Cookbook》 中文版", "description": "《Python Cookbook》 3rd Edition Translation", "summary": "该书是 Python3 的工具书，里面包含了 Python3 各种实用的知识点和代码片段，适合想深入学习 Python 和想提高编程水平的 Python 爱好者。", "lang_color": "#DA5B0B", "primary_lang": "Jupyter Notebook", "stars": 11041, "stars_str": "11k", "publish_at": "2023-07-02 00:15:02", "has_chinese": true, "is_active": false}
]
```



## Meilisearch 

https://www.meilisearch.com/docs/learn/getting_started/quick_start

### 搭建

```bash
# Install Meilisearch ， 跳转到github了，需要魔法
curl -L https://install.meilisearch.com | sh

# Test if it works
./meilisearch --version

# Move to user env
sudo mv meilisearch /usr/local/bin/

# Launch Meilisearch
meilisearch --master-key="aSampleMasterKey"
```

### 数据导入

```bash
$ ./searchEngineTest mci
create index success 
run done

$ ./searchEngineTest mi
add MeillSearch status:enqueued TaskUID: 9
add MeillSearch result status:processing 
add MeillSearch status:enqueued TaskUID: 10
add MeillSearch result status:enqueued 
add MeillSearch status:enqueued TaskUID: 11
add MeillSearch result status:enqueued 
add MeillSearch status:enqueued TaskUID: 12
add MeillSearch result status:enqueued 
add MeillSearch status:enqueued TaskUID: 13
add MeillSearch result status:enqueued 
add MeillSearch status:enqueued TaskUID: 14
add MeillSearch result status:enqueued 
run done
```



### 搜索

```bash
$ ./searchEngineTest ms -q "Redis"
TotalHits: 55 
rid: f18f16d7dd944a36971f8746b7079c87 name: redis 
rid: 025ed78aac4c4dc79518bd99a1c4e835 name: cachecloud 
rid: f0cba6ce547c49e2b401b490c035dd3a name: redisbook 
rid: 152adb02e91a4f9aa7f1f094732ddf00 name: redis-3.0-annotated 
rid: 7aeda21ce8aa48f2aad621364834d8c0 name: kvrocks 
rid: 224c17d258284da18b23c982a3dea902 name: godis 
rid: f951d9abf7db495e99cbc666a70d8ce7 name: pottery 
rid: d45d81b60d0748329648bcd1858beb88 name: lettuce-core 
rid: b6839796762b4511b65ffea1e683f9cd name: RedisInsight 
rid: c6e0beb5ee8d486188c16e7761e3feda name: RedisDesktopManager 
run don
```



### 压力测试

```bash
$ ./searchEngineTest mt
start pressure test, duration:1m0s , rate: 100/s
...
5999, response:{"hits":[{"Rid":"f4164844fb7549f7b579d59d19550ca4","Tags":"Elasticsearch CLI Python 测试","Author" 
6000, response:{"hits":[{"Rid":"f4164844fb7549f7b579d59d19550ca4","Tags":"Elasticsearch CLI Python 测试","Author" 


 ======= report ======= 

Requests      [total, rate, throughput]  6000, 100.02, 99.95
Duration      [total, attack, wait]      1m0.028549913s, 59.990403058s, 38.146855ms
Latencies     [mean, 50, 95, 99, max]    18.23064ms, 5.688964ms, 50.404274ms, 52.483802ms, 60.476845ms
Bytes In      [total, mean]              39942001, 6657.00
Bytes Out     [total, mean]              582000, 97.00
Success       [ratio]                    100.00%
Status Codes  [code:count]               200:6000  
Error Set:
run done 
```

测试结果

一分钟测试6000下，平均耗时 18.23ms，P99 52.48ms，最大 60.47ms

## Typesense

https://typesense.org/docs/guide/install-typesense.html#mac-binary

### 搭建

```bash
curl -O https://dl.typesense.org/releases/0.25.1/typesense-server-0.25.1-amd64.deb
sudo apt install ./typesense-server-0.25.1-amd64.deb

# Test if it works
curl http://localhost:8108/health
# {"ok":true}

# config api key 
#/usr/bin/typesense-server --config=/etc/typesense/typesense-server.ini
vim /etc/typesense/typesense-server.ini
 
```

### 数据导入

```bash
$ ./searchEngineTest tci
create index success 
run done

$ ./searchEngineTest ti
...
import stauts true document 
run done 
```



### 搜索

```bash
$ ./searchEngineTest ts -q "Redis"
TotalHits: 45 
rid: f18f16d7dd944a36971f8746b7079c87 name: redis 
rid: 152adb02e91a4f9aa7f1f094732ddf00 name: redis-3.0-annotated 
rid: c1bd9c95e62645feb520cf87dce52d90 name: redis-faina 
rid: 7fcaf011245141b9b08aa7145351370c name: redis-tui 
rid: 74fd12bb23e0468c8d5809be908c4266 name: redis-memory-analyzer 
rid: 9fdcb81f29764574997a219bc6a12a4a name: AnotherRedisDesktopManager 
rid: 9ef7be9658c74ba0bf5fafcaaf2843c5 name: dragonfly 
rid: 025ed78aac4c4dc79518bd99a1c4e835 name: cachecloud 
rid: 470b8584a2bb4eef91b64be9e74efcda name: haipproxy 
rid: d212adabfe5043b481ea82bb5964577b name: pika 
run done 
```



### 压力测试

```bash
 ./searchEngineTest tt
start pressure test, duration:1m0s , rate: 100/s
...
6000, response:{"facet_counts":[],"found":3,"hits":[{"document":{"Author":"apache","Author_avatar":"https://img.hel 


 ======= report ======= 

Requests      [total, rate, throughput]  6000, 100.02, 100.01
Duration      [total, attack, wait]      59.996613429s, 59.989559216s, 7.054213ms
Latencies     [mean, 50, 95, 99, max]    6.850387ms, 7.021459ms, 7.38063ms, 8.126713ms, 18.582676ms
Bytes In      [total, mean]              29604004, 4934.00
Bytes Out     [total, mean]              0, 0.00
Success       [ratio]                    100.00%
Status Codes  [code:count]               200:6000  
Error Set:
run done 
```

测试结果

一分钟测试6000下，平均耗时 6.85ms，P99 8.12 ms，最大 18.58ms











