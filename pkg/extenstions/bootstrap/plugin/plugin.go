package plugin

import (
	"errors"
	"fmt"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/bootstrap"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
	"github.com/rickar/props"
	"os"
	"path"
	"strings"
)

const flowPhaseProxyKey = "http_proxy"
const flowPhaseMessageKey = "http_message"
const flowPhaseNativeKafkaKey = "native_kafka"
const legacyProxyPhaseTypeKey = "proxy"
const legacyMessagePhaseTypeKey = "message"

type FlowPhase int
type ApiType int

const (
	UnknownPhase   FlowPhase = iota
	RequestPhase   FlowPhase = iota
	ResponsePhase  FlowPhase = iota
	InteractPhase  FlowPhase = iota
	ConnectPhase   FlowPhase = iota
	PublishPhase   FlowPhase = iota
	SubscribePhase FlowPhase = iota
)

const (
	UnknownApiType     ApiType = iota
	ProxyApiType       ApiType = iota
	MessageApiType     ApiType = iota
	NativeKafkaApiType ApiType = iota
)

func NewFlowPhase(str string) FlowPhase {
	if str == "" {
		return UnknownPhase
	}
	switch strings.ToLower(str) {
	case "request":
		return RequestPhase
	case "response":
		return ResponsePhase
	case "subscribe", "message_response":
		return SubscribePhase
	case "publish", "message_request":
		return PublishPhase
	case "interact":
		return InteractPhase
	case "connect":
		return ConnectPhase
	}
	return UnknownPhase
}

func (p FlowPhase) String() string {
	switch p {
	case UnknownPhase:
		goto unknown
	case RequestPhase:
		return "request"
	case ResponsePhase:
		return "response"
	case SubscribePhase:
		return "subscribe"
	case PublishPhase:
		return "publish"
	case InteractPhase:
		return "interact"
	case ConnectPhase:
		return "connect"
	}
unknown:
	return "unknown"
}

func (t ApiType) String() string {
	switch t {
	case UnknownApiType:
		goto unknown
	case ProxyApiType:
		return "proxy"
	case MessageApiType:
		return "message"
	case NativeKafkaApiType:
		return "native kafka"
	}
unknown:
	return "unknown"
}

type Plugin struct {
	Id         string
	Type       string
	Title      string
	FlowPhases []FlowPhase
	ApiTypes   []ApiType
}

func (p Plugin) String() string {
	return fmt.Sprintf("Plugin{Id: %s, Type: %s, Title: %s}", p.Id, p.Type, p.Title)
}

func PluginPostProcessor(data any) (any, error) {
	properties := data.(*props.Properties)
	plugin := Plugin{
		Id:         properties.GetDefault("id", ""),
		Type:       properties.GetDefault("type", ""),
		Title:      properties.GetDefault("description", ""),
		FlowPhases: extractPhases(properties),
		ApiTypes:   extractApiTypes(properties),
	}
	return plugin, plugin.Validate()
}

func (p Plugin) Validate() error {
	if p.Type == "" {
		return errors.New("plugin type is required")
	}
	if p.Title == "" {
		return errors.New("plugin description is required")
	}
	if p.Id == "" {
		return errors.New("plugin id is required")
	}
	return nil
}

func extractPhases(properties *props.Properties) []FlowPhase {
	set := util.Set{}
	for _, key := range []string{flowPhaseProxyKey, flowPhaseMessageKey, flowPhaseNativeKafkaKey, legacyMessagePhaseTypeKey, legacyProxyPhaseTypeKey} {
		if proxy, ok := properties.Get(key); ok {
			for _, token := range strings.Split(proxy, ",") {
				set.Add(NewFlowPhase(strings.TrimSpace(token)))
			}
		}
	}
	return util.ToSlice[FlowPhase](set)
}
func extractApiTypes(properties *props.Properties) []ApiType {
	types := make([]ApiType, 0)

	for _, key := range []string{flowPhaseProxyKey, legacyProxyPhaseTypeKey} {
		if _, ok := properties.Get(key); ok {
			types = append(types, ProxyApiType)
		}
	}
	for _, key := range []string{flowPhaseMessageKey, legacyMessagePhaseTypeKey} {
		if _, ok := properties.Get(key); ok {
			types = append(types, MessageApiType)
		}
	}

	if _, ok := properties.Get(flowPhaseNativeKafkaKey); ok {
		types = append(types, NativeKafkaApiType)
	}

	return types
}

func PluginRelatedFile(filename string) (string, error) {
	plugin := bootstrap.GetData("plugin").(Plugin)
	rootDir := bootstrap.GetData(bootstrap.RootDirDataKey).(string)
	specificConfig := path.Join(rootDir, plugin.Type, plugin.Id, filename)
	defaultConfig := path.Join(rootDir, plugin.Type, filename)
	if _, err := os.Stat(specificConfig); err == nil {
		return specificConfig, nil
	}
	if _, err := os.Stat(defaultConfig); err == nil {
		return defaultConfig, nil
	}
	return "", errors.New(fmt.Sprintf("plugin related file not found. filename: %s, plugin: %s", filename, plugin))
}
