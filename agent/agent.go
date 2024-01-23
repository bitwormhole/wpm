package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bitwormhole/wpm/common/objects/dto"
	"github.com/bitwormhole/wpm/common/objects/dxo"
	"github.com/bitwormhole/wpm/common/objects/vo"

	"github.com/starter-go/application"
	"github.com/starter-go/httpagent"
	"github.com/starter-go/vlog"
)

// WPMAgent ...
type WPMAgent struct {

	//starter:component

	AC   application.Context //starter:inject("context")
	HTTP httpagent.Clients   //starter:inject("#")

	ServerPort int //starter:inject("${server.http.port}")

	homeDirPath string
	queueID     dxo.IntentQueueID
}

func (inst *WPMAgent) _impl() application.Lifecycle {
	return inst
}

// Life ...
func (inst *WPMAgent) Life() *application.Life {
	return &application.Life{
		OnCreate: inst.init,
		OnLoop:   inst.loop,
		OnStart:  inst.start,
		OnStop:   inst.stop,
	}
}

func (inst *WPMAgent) init() error {
	wd, err := os.Getwd()
	if err == nil {
		inst.homeDirPath = wd
	} else {
		inst.homeDirPath = "/home/user1"
	}
	return nil
}

func (inst *WPMAgent) hasStopping() bool {
	return false
}

func (inst *WPMAgent) start() error {
	_, err := inst.open()
	return err
}

func (inst *WPMAgent) stop() error {
	_, err := inst.close()
	return err
}

func (inst *WPMAgent) loop() error {
	for {
		if inst.hasStopping() {
			break
		}
		inst.once()
	}
	return nil
}

func (inst *WPMAgent) once() {
	defer func() {
		x := recover()
		inst.handleErrorX(x)
	}()
	err := inst.run()
	inst.handleError(err)
}

func (inst *WPMAgent) handleError(err error) {
	if err == nil {
		return
	}
	vlog.Error(err.Error())
	time.Sleep(time.Second)
}

func (inst *WPMAgent) handleErrorX(x any) {

	if x == nil {
		return
	}

	err, ok := x.(error)
	if ok {
		inst.handleError(err)
		return
	}

	str, ok := x.(string)
	if ok {
		err = fmt.Errorf(str)
		inst.handleError(err)
		return
	}

	vlog.Warn("handleErrorX: unknown error type")
}

func (inst *WPMAgent) run() error {

	view, err := inst.fetch()
	if err != nil {
		return err
	}

	list := view.Queue.Intents
	for _, item := range list {
		err = inst.execute(item)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *WPMAgent) open() (*vo.IntentQueue, error) {
	body1 := new(vo.IntentQueue)
	body1.Queue.UserHomeDir = inst.homeDirPath
	req := new(httpagent.Request)
	req.Method = http.MethodPost
	req.URL = "/api/v1/intent-queues"
	body2, err := inst.send(req, body1)
	if err != nil {
		return nil, err
	}
	inst.queueID = body2.Queue.ID
	return body2, nil
}

func (inst *WPMAgent) close() (*vo.IntentQueue, error) {
	id := inst.queueID.String()
	req := new(httpagent.Request)
	req.Method = http.MethodDelete
	req.URL = "/api/v1/intent-queues/" + id
	return inst.send(req, nil)
}

func (inst *WPMAgent) fetch() (*vo.IntentQueue, error) {
	id := inst.queueID.String()
	body1 := new(vo.IntentQueue)
	req := new(httpagent.Request)
	req.Method = http.MethodPost
	req.URL = "/api/v1/intent-queues/" + id + "/fetch"
	return inst.send(req, body1)
}

func (inst *WPMAgent) send(req *httpagent.Request, body1 *vo.IntentQueue) (*vo.IntentQueue, error) {

	body2 := new(vo.IntentQueue)
	url := req.URL

	if strings.HasPrefix(url, "/") {
		port := strconv.Itoa(inst.ServerPort)
		url = "http://localhost:" + port + url
	}

	if body1 != nil {
		// inst.prepareRequestBody(body1)
		entity1 := httpagent.NewEntityWithJSON(body1)
		req.SetEntity(entity1)
	}

	req.Context = context.Background()
	req.URL = url

	client := inst.HTTP.GetClient()
	resp, err := client.Execute(req)
	if err != nil {
		return nil, err
	}

	code := resp.Status
	if code != http.StatusOK {
		return nil, fmt.Errorf("HTTP %s", resp.Message)
	}

	entity2, err := resp.GetEntity()
	if err != nil {
		return nil, err
	}

	err = entity2.ReadJSON(body2)
	if err != nil {
		return nil, err
	}

	return body2, nil
}

func (inst *WPMAgent) execute(item *dto.Intent) error {
	j, err := json.MarshalIndent(item, "", "\t")
	if err != nil {
		return err
	}
	vlog.Info("execute: %s", j)
	return nil
}
