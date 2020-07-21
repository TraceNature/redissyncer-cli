package httpquerry

import (
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//var logger = globalzap.GetLogger()

const CreateTaskPath = "/api/v2/createtask"
const StartTaskPath = "/api/v2/starttask"
const StopTaskPath = "/api/v2/stoptask"
const RemoveTaskPath = "/api/v2/removetask"
const ListTasksPath = "/api/v2/listtasks"

type Request struct {
	Server string
	Api    string
	Body   string
}

func (r Request) ExecRequest() (string, error) {
	client := &http.Client{}
	client.Timeout = 30 * time.Second

	req, err := http.NewRequest("POST", r.Server+r.Api, strings.NewReader(r.Body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, resperr := client.Do(req)

	if resperr != nil {
		return "", resperr
	}
	defer resp.Body.Close()

	body, readerr := ioutil.ReadAll(resp.Body)
	if readerr != nil {
		return "", readerr
	}

	var dat map[string]interface{}
	json.Unmarshal(body, &dat)
	bodystr, jsonerr := json.MarshalIndent(dat, "", " ")
	if jsonerr != nil {
		return "", jsonerr
	}
	return string(bodystr), nil
}

//创建同步任务
func CreateTask(syncserver string, createjson string) ([]string, error) {
	createreq := &Request{
		Server: syncserver,
		Api:    CreateTaskPath,
		Body:   createjson,
	}

	resp, err := createreq.ExecRequest()
	if err != nil {
		return nil, err
	}
	taskids := gjson.Get(resp, "data").Array()
	if len(taskids) == 0 {
		return nil, errors.New("task create faile")
	}
	taskidsstrarray := []string{}
	for _, v := range taskids {
		taskidsstrarray = append(taskidsstrarray, gjson.Get(v.String(), "taskId").String())
	}

	return taskidsstrarray, nil

}

//Start task
func StartTask(syncserver string, taskid string) (string, error) {
	jsonmap := make(map[string]interface{})
	jsonmap["taskid"] = taskid
	startjson, err := json.Marshal(jsonmap)
	if err != nil {
		return "", err
	}
	startreq := &Request{
		Server: syncserver,
		Api:    StartTaskPath,
		Body:   string(startjson),
	}
	return startreq.ExecRequest()

}

//Stop task by task ids
func StopTaskByIds(syncserver string, ids []string) (string, error) {
	jsonmap := make(map[string]interface{})

	jsonmap["taskids"] = ids
	stopjsonStr, err := json.Marshal(jsonmap)
	if err != nil {
		return "", err
	}
	stopreq := &Request{
		Server: syncserver,
		Api:    StopTaskPath,
		Body:   string(stopjsonStr),
	}
	return stopreq.ExecRequest()

}

//Remove task by name
func RemoveTaskByName(syncserver string, taskname string) (string, error) {
	jsonmap := make(map[string]interface{})

	taskids, err := GetSameTaskNameIds(syncserver, taskname)
	if err != nil {
		return "", err
	}

	if len(taskids) == 0 {
		return "", errors.New("no taskid")
	}

	jsonmap["taskids"] = taskids
	stopjsonStr, err := json.Marshal(jsonmap)
	if err != nil {
		return "", err
	}
	stopreq := &Request{
		Server: syncserver,
		Api:    StopTaskPath,
		Body:   string(stopjsonStr),
	}
	stopreq.ExecRequest()

	removereq := &Request{
		Server: syncserver,
		Api:    RemoveTaskPath,
		Body:   string(stopjsonStr),
	}

	return removereq.ExecRequest()

}

//获取同步任务状态
func GetTaskStatus(syncserver string, ids []string) (map[string]string, error) {
	jsonmap := make(map[string]interface{})

	jsonmap["regulation"] = "byids"
	jsonmap["taskids"] = ids

	listtaskjsonStr, err := json.Marshal(jsonmap)
	if err != nil {
		return nil, err
	}
	listreq := &Request{
		Server: syncserver,
		Api:    ListTasksPath,
		Body:   string(listtaskjsonStr),
	}
	listresp, err := listreq.ExecRequest()
	taskarray := gjson.Get(listresp, "data").Array()

	if len(taskarray) == 0 {
		return nil, errors.New("No status return")
	}

	statusmap := make(map[string]string)

	for _, v := range taskarray {
		id := gjson.Get(v.String(), "taskId").String()
		status := gjson.Get(v.String(), "status").String()
		statusmap[id] = status
	}

	return statusmap, nil
}

// @title    GetSameTaskNameIds
// @description   获取同名任务列表
// @auth      Jsw             时间（2020/7/1   10:57 ）
// @param     syncserver        string         "redissyncer ip:port"
// @param    taskname        string         "任务名称"
// @return    taskids        []string         "任务id数组"
func GetSameTaskNameIds(syncserver string, taskname string) ([]string, error) {

	existstaskids := []string{}
	listjsonmap := make(map[string]interface{})
	listjsonmap["regulation"] = "bynames"
	listjsonmap["tasknames"] = strings.Split(taskname, ",")
	listjsonStr, err := json.Marshal(listjsonmap)
	if err != nil {
		return nil, err
	}
	listtaskreq := &Request{
		Server: syncserver,
		Api:    ListTasksPath,
		Body:   string(listjsonStr),
	}
	listresp, err := listtaskreq.ExecRequest()
	if err != nil {
		return nil, err
	}
	tasklist := gjson.Get(listresp, "data").Array()

	if len(tasklist) > 0 {
		for _, v := range tasklist {
			existstaskids = append(existstaskids, gjson.Get(v.String(), "taskId").String())
		}
	}
	return existstaskids, nil
}
