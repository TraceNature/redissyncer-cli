# redissyncer-cli

The client cli tool of redissyncer facilitates the operation of migration tasks.

## Build

```shell script
git clone https://github.com/TraceNature/redissyncer-cli.git
cd redissyncer-cli
go build -o redissyncer-cli
```


## Function and Usage

* redissyncer-cli supports command line mode and interactive mode, "redissyncer-cli -i" enters interactive mode
* The client program is used by the redissyncer client program to create, start and stop, and monitor redis synchronization tasks. Please ensure that the server program is running normally before using this Cui
* The config.yml file is used to describe the basic configuration of the server link. The program reads .config.yml in the current directory by default, and the file name and path can also be customized

  ``` yaml  
  server: http://10.0.0.100:8080
  ```

* For services that enable user verification, first obtain the token through "redissyncer-cli login username password", and then write the token into the config file
  
  ``` yaml  
  server: http://10.0.0.100:8080
  token: 379F5E2BD55A4608B6A7557F0583CFC5
  ```


## Examples of interactive modes

   "redissyncer-cli -i" enters interactive mode

* Create Task
  
  * Use json plaintext to create tasks
  
   ```shell
   redissyncer-cli> task create '{"sourcePassword":"xxxxx","sourceRedisAddress":"10.0.1.100:6379","targetRedisAddress":"192.168.0.100:6379","targetPassword":"xxxxx","targetRedisVersion":4,"taskName":"firsttest"}';
   ```

  * Create task through json file
  
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

   For detailed configuration parameters, please refer to [API Document](api.md)

* Get task status
  * View all tasks

    ```shell
    redissyncer-cli> task status all;
    ```

  * View task status by task id

    ```shell
    redissyncer-cli> task status bytaskid 690DEF6222E34443884033B860CE01EC;
    ```

  * View task status by task name

    ```shell
    redissyncer-cli> task status byname taskname;
    ```


* Start task

   ```shell
   redissyncer-cli> task start 690DEF6222E34443884033B860CE01EC;
   ```

* Stop task

   ```shell
   redissyncer-cli> task stop 690DEF6222E34443884033B860CE01EC;
   ```

* Remove task by task id

   ```shell
   redissyncer-cli> task remove 690DEF6222E34443884033B860CE01EC；
   ```
