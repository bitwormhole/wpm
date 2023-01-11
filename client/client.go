package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// Client ...
type Client struct {
	markup.Component `class:"life"`

	IntentHandler IntentHandler `inject:"#IntentHandler"`

	Protocol string `inject:"${wpm.server.protocol}"`
	Host     string `inject:"${wpm.server.host}"`
	Port     int    `inject:"${wpm.server.port}"`

	pipeID dxo.PipeID
}

func (inst *Client) _Impl() application.LifeRegistry {
	return inst
}

// GetLifeRegistration ...
func (inst *Client) GetLifeRegistration() *application.LifeRegistration {
	return &application.LifeRegistration{
		OnInit:  inst.onInit,
		OnStart: inst.onStart,
		Looper:  inst,
	}
}

func (inst *Client) onInit() error {
	return nil
}

func (inst *Client) onStart() error {
	return inst.doPost()
}

// Loop ...
func (inst *Client) Loop() error {
	const timeout = 1000 * 30
	for {
		err := inst.doGet(timeout)
		if err != nil {
			vlog.Warn(err)
			time.Sleep(time.Second * 3)
		}
	}
	return nil
}

func (inst *Client) getWebURL(url string) string {

	pro := inst.Protocol
	host := inst.Host
	port := inst.Port

	if pro == "" {
		pro = "http"
	}

	if host == "" {
		host = "localhost"
	}

	if port <= 0 {
		port = 8080
	}

	const f = "%v://%v:%v%v"
	return fmt.Sprintf(f, pro, host, port, url)
}

func (inst *Client) getPipeURL(id dxo.PipeID, timeout int) string {

	url := inst.getWebURL("/api/v1/pipe")

	if id > 0 {
		idstr := fmt.Sprintf("%v", id)
		url = url + "/" + idstr
	}

	if timeout > 0 {
		q := fmt.Sprintf("timeout=%v", timeout)
		url = url + "?" + q
	}

	return url
}

func (inst *Client) doPost() error {

	body1obj := &vo.Pipe{}

	body1reader := inst.prepareRequestBody(body1obj)
	url := inst.getPipeURL(0, 0)
	req, err := http.NewRequest(http.MethodPost, url, body1reader)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body2obj, err := inst.handleResponse(resp)
	if err != nil {
		return err
	}

	return inst.handlePostResult(body2obj)
}

func (inst *Client) doGet(timeout int) error {

	url := inst.getPipeURL(inst.pipeID, timeout)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	body2obj, err := inst.handleResponse(resp)
	if err != nil {
		return err
	}

	return inst.handleGetResult(body2obj)
}

func (inst *Client) prepareRequestBody(o *vo.Pipe) io.Reader {
	if o == nil {
		o = &vo.Pipe{}
	}
	buffer := &bytes.Buffer{}
	data, err := json.Marshal(o)
	if err == nil {
		buffer.Write(data)
	}
	return buffer
}

func (inst *Client) handlePostResult(o *vo.Pipe) error {
	all := o.Pipes
	count := 0
	for _, p := range all {
		pid := p.ID
		if pid > 0 {
			count++
		}
		inst.pipeID = pid
	}
	if count < 1 {
		return fmt.Errorf("no pipe id in result")
	}
	return nil
}

func (inst *Client) handleGetResult(o *vo.Pipe) error {

	defer func() {
		x := recover()
		if x != nil {
			vlog.Error(x)
		}
	}()

	all := o.Pipes
	for _, p := range all {
		if p.ID == inst.pipeID {
			i := p.Intent
			err := inst.IntentHandler.HandleIntent(i)
			if err != nil {
				vlog.Error(err)
			}
		}
	}

	return nil
}

func (inst *Client) handleResponse(resp *http.Response) (*vo.Pipe, error) {

	dst := &vo.Pipe{}
	src := resp.Body
	status := resp.Status
	code := resp.StatusCode

	if src == nil {
		return nil, fmt.Errorf("bad HTTP response: no body")
	}

	defer func() {
		src.Close()
	}()

	if code != http.StatusOK {
		return nil, fmt.Errorf("HTTP %v", status)
	}

	data, err := ioutil.ReadAll(src)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, dst)
	if err != nil {
		return nil, err
	}

	return dst, nil
}
