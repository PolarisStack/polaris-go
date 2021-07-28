/**
 * Tencent is pleased to support the open source community by making polaris-go available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package config

import (
	"github.com/polarismesh/polaris-go/pkg/model"
	"github.com/polarismesh/polaris-go/pkg/plugin/common"
)

type SubscribeImpl struct {
	//上报插件链
	Type string `yaml:"type" json:"type"`
	// 插件相关配置
	Plugin PluginConfigs `yaml:"plugin" json:"plugin"`
}

func (s *SubscribeImpl) Init() {
	s.Plugin = PluginConfigs{}
	s.Plugin.Init(common.TypeSubScribe)
}

func (s *SubscribeImpl) SetDefault() {
	if s.Type == "" {
		s.Type = SubscribeLocalChannel
	}
	s.Plugin.SetDefault(common.TypeSubScribe)
}

//verify
func (s *SubscribeImpl) Verify() error {
	return nil
}

//设置插件
func (s *SubscribeImpl) SetType(typeStr string) {
	s.Type = typeStr
}

//获取插件
func (s *SubscribeImpl) GetType() string {
	return s.Type
}

//GetPluginNames
func (s *SubscribeImpl) GetPluginNames() model.HashSet {
	return nil
}

//GetPluginConfig
func (s *SubscribeImpl) GetPluginConfig(pluginName string) BaseConfig {
	if s.Plugin == nil {
		return nil
	}
	if v, ok := s.Plugin[pluginName]; ok {
		conf := v.(BaseConfig)
		return conf
	} else {
		return nil
	}
}

//SetPluginConfig
func (s *SubscribeImpl) SetPluginConfig(plugName string, value BaseConfig) error {
	return nil
}
