package utils

import (
	"net/url"
	"strconv"
)

// URLBuilder  一个简单的URL创建工具
type URLBuilder struct {
	Protocol string
	User     string
	Host     string
	Port     int
	Path     string
	Query    string
	Fragment string
}

// Init  ...
func (inst *URLBuilder) Init(u string) {

	x, err := url.Parse(u)
	if err != nil {
		return
	}

	port, _ := strconv.Atoi(x.Port())

	inst.Protocol = x.Scheme
	inst.User = x.User.String()
	inst.Host = x.Hostname()
	inst.Port = port
	inst.Path = x.Path
	inst.Query = x.RawQuery
	inst.Fragment = x.Fragment
}

// Create ...
func (inst *URLBuilder) Create() string {

	user := inst.User
	host := inst.Host
	port := inst.Port
	x := &url.URL{}

	if user != "" {
		x.User = url.User(user)
	}

	if port > 0 {
		x.Host = host + ":" + strconv.Itoa(port)
	} else {
		x.Host = host
	}

	x.Scheme = inst.Protocol
	x.Path = inst.Path
	x.Fragment = inst.Fragment
	x.RawQuery = inst.Query

	return x.String()
}
