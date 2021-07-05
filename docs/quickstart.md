# redissyncer-cli

[English](quickstart_en.md)

redissyncer 的客户端cli工具，方便迁移任务操作。

## 构建

```shell script
git clone https://github.com/TraceNature/redissyncer-cli.git
cd redissyncer-cli
go build -o redissyncer-cli
```


## 功能与使用方法

* redissyncer-cli支持命令行模式和交互模式，"redissyncer-cli -i"进入交互模式
* 该客户端程序为redissyncer客户端程序用与创建、启停、监控redis同步任务，在使用本Cui之前请确保服务端程序正常运行
* .config.yaml文件用于描述服务器链接的基本配置，程序默认读取当前目录下的 .config.yaml，也可自定义文件名称及路径

  ``` yaml  
  syncserver: http://10.0.0.100:8080
  ```

* 对于开启用户校验的服务，先通过"redissyncer-cli login username password" 获取token，然后将 token 写入config文件
  
  ``` yaml  
  syncserver: http://10.0.0.100:8080
  token: 379F5E2BD55A4608B6A7557F0583CFC5
  ```


## 交互模式举例

   "redissyncer-cli -i"进入交互模式

* 创建任务
  * json明文创建任务
  
   ```shell
   redissyncer-cli> task create '{"sourcePassword":"xxxxx","sourceRedisAddress":"10.0.1.100:6379","targetRedisAddress":"192.168.0.100:6379","targetPassword":"xxxxx","targetRedisVersion":4,"taskName":"firsttest"}';
   ```

  * 通过json文件创建任务
    createtask.json文件
  
   ```json
   {
       "dbNum":{
           "1":"1"
       },
       "sourcePassword":"xxxxxx",
       "sourceRedisAddress":"10.0.1.100:6379",
       "targetRedisAddress":"192.168.0.100:6379",
       "targetPassword":"xxxxxx",
       "targetRedisVersion":4,
       "taskName":"testtask",
       "autostart":true,
       "afresh":true,
       "batchSize":100
   }
   ```

    ```shell
   redissyncer-cli> task create source ./createtask.json;
   ```

   详细配置参数详见[API文档](api.md)


* 查看任务
  * 查看全部任务

    ```shell
    redissyncer-cli> task status all;
    ```

  * 查看通过任务id查看任务状态

    ```shell
    redissyncer-cli> task status bytaskid 690DEF6222E34443884033B860CE01EC;
    ```

  * 查看通过任务名称查看任务状态

    ```shell
    redissyncer-cli> task status byname taskname;
    ```


* 启动任务

   ```shell
   redissyncer-cli> task start 690DEF6222E34443884033B860CE01EC;
   ```

* 停止任务

   ```shell
   redissyncer-cli> task stop 690DEF6222E34443884033B860CE01EC;
   ```

* 通过任务名删除任务

   ```
   redissyncer-cli> task remove 690DEF6222E34443884033B860CE01EC；
   ```
