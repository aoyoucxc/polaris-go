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

package outlierdetection

import (
	"github.com/polarismesh/polaris-go/pkg/model"
	"github.com/polarismesh/polaris-go/pkg/plugin"
	"github.com/polarismesh/polaris-go/pkg/plugin/common"
)

//proxy of OutlierDetector
type Proxy struct {
	OutlierDetector
	engine model.Engine
}

//设置
func (p *Proxy) SetRealPlugin(plug plugin.Plugin, engine model.Engine) {
	p.OutlierDetector = plug.(OutlierDetector)
	p.engine = engine
}

//proxy OutlierDetector DetectInstance
func (p *Proxy) DetectInstance(inst model.Instance) (common.DetectResult, error) {
	result, err := p.OutlierDetector.DetectInstance(inst)
	return result, err
}

//注册proxy
func init() {
	plugin.RegisterPluginProxy(common.TypeOutlierDetector, &Proxy{})
}