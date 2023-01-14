package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
	"github.com/bitwormhole/wpm/server/data/dxo"
	"github.com/bitwormhole/wpm/server/service"
	"github.com/bitwormhole/wpm/server/web/dto"
	"github.com/bitwormhole/wpm/server/web/vo"
)

// Client ...
type Client struct {
	markup.Component `class:"life"`

	IntentHandler   IntentHandler           `inject:"#IntentHandler"`
	PlatformService service.PlatformService `inject:"#PlatformService"`

	Protocol string `inject:"${wpm.server.protocol}"`
	Host     string `inject:"${wpm.server.host}"`
	Port     int    `inject:"${wpm.server.port}"`

	stopping bool
	pipeID   dxo.PipeID
	pipeInfo dto.PipeInfo
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
	err := inst.doPost()
	if err != nil {
		return err
	}
	return inst.logWebGuiURL()
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
		if inst.stopping {
			break
		}
	}
	return nil
}

func (inst *Client) getWebURL(path string, query map[string]string) string {

	scheme := inst.Protocol
	host := inst.Host
	port := inst.Port

	if scheme == "" {
		scheme = "http"
	}

	if host == "" {
		host = "localhost"
	}

	if port <= 0 {
		port = 8080
	}

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	builder := strings.Builder{}
	builder.WriteString(scheme)
	builder.WriteString("://")
	builder.WriteString(host)
	builder.WriteString(":")
	builder.WriteString(strconv.Itoa(port))
	builder.WriteString(path)

	// query
	keys := make([]string, 0)
	for key := range query {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	sep := "?"
	for _, key := range keys {
		value := query[key]
		builder.WriteString(sep)
		builder.WriteString(key)
		builder.WriteString("=")
		builder.WriteString(value)
		sep = "&"
	}

	return builder.String()
}

func (inst *Client) getPipeURL(id dxo.PipeID, info *dto.PipeInfo, timeout int) string {

	path := "/api/v1/pipe"
	query := make(map[string]string)

	if id > 0 {
		idstr := fmt.Sprintf("%v", id)
		path = path + "/" + idstr
	}

	if timeout > 0 {
		query["timeout"] = strconv.Itoa(timeout)
	}

	if info != nil {
		dsid := fmt.Sprintf("%v", info.DesktopSessionID)
		uuid := fmt.Sprintf("%v", info.UUID)
		query["dsid"] = dsid
		query["uuid"] = uuid
	}

	return inst.getWebURL(path, query)
}

func (inst *Client) prepareNewPipe() (*dto.PipeInfo, error) {
	info, err := inst.PlatformService.GetInfo(nil)
	if err != nil {
		return nil, err
	}
	body1dto := &dto.PipeInfo{}
	body1dto.DesktopSessionHome = info.Home
	body1dto.DesktopSessionUser = info.User
	return body1dto, nil
}

func (inst *Client) logWebGuiURL() error {

	pid := inst.pipeID
	dsid := inst.pipeInfo.DesktopSessionID

	query := make(map[string]string)

	query["pipe"] = fmt.Sprintf("%v", pid)
	query["dsid"] = fmt.Sprintf("%v", dsid)

	url := inst.getWebURL("/", query)
	vlog.Info("webgui.url=" + url)
	return nil
}

func (inst *Client) doPost() error {

	body1info, err := inst.prepareNewPipe()
	if err != nil {
		return err
	}
	body1vo := &vo.Pipe{}
	body1vo.Pipes = append(body1vo.Pipes, body1info)

	body1reader := inst.prepareRequestBody(body1vo)
	url := inst.getPipeURL(0, nil, 0)
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

	url := inst.getPipeURL(inst.pipeID, &inst.pipeInfo, timeout)
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
	for _, pi := range all {
		pid := pi.ID
		if pid > 0 {
			inst.pipeID = pid
			inst.pipeInfo = *pi
			count++
		}
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

	all := o.Packets
	for _, p := range all {
		i := p.Intent
		err := inst.IntentHandler.HandleIntent(i)
		if err != nil {
			vlog.Error(err)
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
