// Copyright (C) 2015 The Gravitee team (http://gravitee.io)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plugin

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/bootstrap"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
	"github.com/rickar/props"
)

const flowPhaseProxyKey = "http_proxy"
const flowPhaseMessageKey = "http_message"
const flowPhaseNativeKafkaKey = "native_kafka"
const flowPhaseHttpMcpProxy = "http_mcp_proxy"
const flowPhaseHttpLlmProxy = "http_llm_proxy"
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
	UnknownApiType      ApiType = iota
	ProxyApiType        ApiType = iota
	MessageApiType      ApiType = iota
	NativeKafkaApiType  ApiType = iota
	HttpMcpProxyApiType ApiType = iota
	HttpLlmProxyApiType ApiType = iota
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
	case HttpMcpProxyApiType:
		return "mcp proxy"
	case HttpLlmProxyApiType:
		return "llm proxy"
	}
unknown:
	return "unknown"
}

type Plugin struct {
	ID         string
	Type       string
	Title      string
	FlowPhases []FlowPhase
	ApiTypes   []ApiType
}

func (p Plugin) String() string {
	return fmt.Sprintf("Plugin{ID: %s, Type: %s, Title: %s}", p.ID, p.Type, p.Title)
}

func PostProcessor(data any) (any, error) {
	properties := util.As[*props.Properties](data)
	plugin := Plugin{
		ID:         properties.GetDefault("id", ""),
		Type:       properties.GetDefault("type", ""),
		Title:      properties.GetDefault("name", ""),
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
		return errors.New("plugin name is required")
	}
	if p.ID == "" {
		return errors.New("plugin id is required")
	}
	return nil
}

func extractPhases(properties *props.Properties) []FlowPhase {
	set := util.NewSet()
	for _, key := range []string{
		flowPhaseProxyKey,
		flowPhaseMessageKey,
		flowPhaseNativeKafkaKey,
		flowPhaseHttpMcpProxy,
		flowPhaseHttpLlmProxy,
		legacyMessagePhaseTypeKey,
		legacyProxyPhaseTypeKey} {
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

	if _, ok := properties.Get(flowPhaseHttpMcpProxy); ok {
		types = append(types, HttpMcpProxyApiType)
	}

	if _, ok := properties.Get(flowPhaseHttpLlmProxy); ok {
		types = append(types, HttpMcpProxyApiType)
	}

	return types
}

func RelativeFile(filename string) (string, error) {
	plugin, _ := bootstrap.GetData("plugin").(Plugin)
	rootDir, _ := bootstrap.GetData(bootstrap.RootDirDataKey).(string)
	specificConfig := path.Join(rootDir, plugin.Type, plugin.ID, filename)
	defaultConfig := path.Join(rootDir, plugin.Type, filename)
	if _, err := os.Stat(specificConfig); err == nil {
		return specificConfig, nil
	}
	if _, err := os.Stat(defaultConfig); err == nil {
		return defaultConfig, nil
	}
	return "", fmt.Errorf("plugin related file not found. filename: %s, plugin: %s", filename, plugin)
}
