package config

import (
	"encoding/json"
	"github.com/xlvector/dlog"
	"io/ioutil"
)

type Redis struct {
	Host    string
	Timeout int64
}

type Captcha struct {
	Key      string
	AppId    string
	Username string
	Password string
}

type ES struct {
	Host  string
	Index string
}

type CookieTemplate struct {
	Site       string
	Path       string
	Domain     string
	Secure     bool
	HttpOnly   bool
	Persistent bool
	HostOnly   bool
	Tmpl       string
}

type Flume struct {
	Host string
	Port int
}

type Config struct {
	OutputRoot             string
	PythonExtractorService string
	Redis                  Redis
	Captcha                Captcha
	Templates              map[string]string
	ES                     ES
	Flume                  Flume
	CookieTemplateConfig   map[string]map[string]*CookieTemplate
	Buckets                map[string]string
	UploadApi              string
	SlackApi               string
}

func (p Config) HasRedis() bool {
	if len(p.Redis.Host) > 0 {
		return true
	}
	return false
}

func (p Config) HasFlume() bool {
	if len(p.Flume.Host) > 0 {
		return true
	}
	return false
}

var Instance Config

func Init(conf string) {
	b, err := ioutil.ReadFile(conf)
	if err != nil {
		dlog.Warn("fail to load config: %v", err)
	}
	err = json.Unmarshal(b, &Instance)
	if err != nil {
		dlog.Warn("fail to parse config: %v", err)
	}
	dlog.Info("OutputRoot: %s", Instance.OutputRoot)
	dlog.Info("Python Extactor service: %s", Instance.PythonExtractorService)
}

func GetCookieTemplate(tmpl string) map[string]*CookieTemplate {
	cookieTemplate := Instance.CookieTemplateConfig[tmpl]
	if resource, ok := cookieTemplate["_RESOURCE"]; ok {
		cookieTemplate = Instance.CookieTemplateConfig[resource.Tmpl]
	}
	return cookieTemplate
}
