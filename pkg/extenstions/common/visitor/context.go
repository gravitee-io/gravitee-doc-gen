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

package visitor

type VisitContext struct {
	oneOfs              []OneOfDescriptor
	queueNodes          bool
	autoDefaultBooleans bool
	nodeStack           *NodeStack
	oneOfFilter         OneOfFilter
}

type OneOfFilter struct {
	Path           []string       `json:"path"`
	Discriminators map[string]any `json:"discriminators"`
	Index          int            `json:"index"`
}

type OneOfDescriptor struct {
	ParentTitle string
	Present     bool
	Specs       []DiscriminatorSpec
}

func NewVisitContext(queueNodes bool, autoDefaultBooleans bool) *VisitContext {
	return &VisitContext{
		oneOfs:              make([]OneOfDescriptor, 0),
		queueNodes:          queueNodes,
		autoDefaultBooleans: autoDefaultBooleans,
	}
}

func (v *VisitContext) WithStack(root *Object) *VisitContext {
	v.nodeStack = NewNodeStack(root)
	return v
}

func (v *VisitContext) WithOneOfFilter(filter OneOfFilter) *VisitContext {
	v.oneOfFilter = filter
	return v
}

func (v *VisitContext) IsAutoDefaultBooleans() bool {
	return v.autoDefaultBooleans
}

func (v *VisitContext) NodeStack() *NodeStack {
	return v.nodeStack
}

func (v *VisitContext) IsQueueNodes() bool {
	return v.queueNodes
}

func (v *VisitContext) PushOneOf(oneOf OneOfDescriptor) {
	if !oneOf.IsZero() {
		v.oneOfs = append(v.oneOfs, oneOf)
	} else {
		panic("push one-of that is zero")
	}
}

func (v *VisitContext) PopOneOf() {
	v.oneOfs = v.oneOfs[:len(v.oneOfs)-1]
}

func (v *VisitContext) PeekOneOf() OneOfDescriptor {
	if len(v.oneOfs) == 0 {
		return OneOfDescriptor{}
	}
	return v.oneOfs[len(v.oneOfs)-1]
}

func (v *VisitContext) OneOfFilter() OneOfFilter {
	return v.oneOfFilter
}

func (o OneOfDescriptor) IsZero() bool {
	return o.ParentTitle == "" && !o.Present && len(o.Specs) == 0
}

func (o OneOfDescriptor) IsDiscriminator(property string) bool {
	for _, spec := range o.Specs {
		if spec.Property == property {
			return true
		}
	}
	return false
}

func (f OneOfFilter) IsZero() bool {
	return len(f.Path) == 0 && len(f.Discriminators) == 0 && f.Index == 0
}
