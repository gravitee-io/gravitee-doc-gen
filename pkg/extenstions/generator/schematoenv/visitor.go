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

package schematoenv

import (
	"fmt"
	"strings"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/visitor"
)

type envSection struct {
	Title       string
	Description string
	Variables   []envVariable
}

func (s *envSection) AddVariable(v envVariable) {
	s.Variables = append(s.Variables, v)
}

type envVariable struct {
	visitor.Attribute
	Env string
	JVM string
}

type toEnvVisitor struct {
	inArray          bool
	currentSection   *envSection
	Sections         []*envSection
	indexPlaceholder string
	prefix           string
	envPaths         []string
	jvmPaths         []string
}

func (v *toEnvVisitor) OnObjectStart(object visitor.Object, level int) {
	if v.inArray {
		return
	}
	if level == 0 {
		v.currentSection = &envSection{
			Title:       object.Title,
			Description: object.Description,
			Variables:   make([]envVariable, 0),
		}
		v.Sections = append(v.Sections, v.currentSection)
	} else {
		v.addPaths(object.Name())
	}
}

func (v *toEnvVisitor) OnObjectEnd(_ visitor.Object, level int) {
	if level > 0 {
		v.removeLastPaths()
	}
}

func (v *toEnvVisitor) OnArrayStart(array visitor.Array, level int) {
	v.inArray = true
	if level == 0 {
		v.currentSection = &envSection{
			Variables: make([]envVariable, 0),
		}
		v.Sections = append(v.Sections, v.currentSection)
	} else {
		v.addPathsForArray(array.Name())
	}
}

func (v *toEnvVisitor) OnArrayItem(array visitor.Array, _ visitor.Value, _ int, last bool) {
	if last {
		attribute := visitor.NewAttribute("", nil)
		attribute.Type = array.ItemType
		attribute.Title = array.Title
		attribute.Description = array.Description
		v.currentSection.AddVariable(envVariable{
			Attribute: *attribute,
			Env:       v.getEnv(),
			JVM:       v.getJVM(),
		})
		v.removeLastPaths()
	}
}

func (v *toEnvVisitor) OnArrayEnd(visitor.Array, int) {
	v.inArray = false
}

func (v *toEnvVisitor) OnAttribute(attribute visitor.Attribute, _ int) {
	v.addPaths(attribute.Name())
	v.currentSection.AddVariable(envVariable{
		Attribute: attribute,
		Env:       v.getEnv(),
		JVM:       v.getJVM(),
	})
	v.removeLastPaths()
}

func (v *toEnvVisitor) addPaths(path string) {
	v.envPaths = append(v.envPaths, strings.ToUpper(path))
	v.jvmPaths = append(v.jvmPaths, strings.ToLower(path))
}

func (v *toEnvVisitor) addPathsForArray(path string) {
	v.envPaths = append(v.envPaths, v.formatArrayEnv(path))
	v.jvmPaths = append(v.jvmPaths, v.formatArrayJVM(path))
}

func (v *toEnvVisitor) formatArrayJVM(name string) string {
	return fmt.Sprintf("%s[%s]", strings.ToLower(name), v.indexPlaceholder)
}

func (v *toEnvVisitor) formatArrayEnv(name string) string {
	return strings.ToUpper(name) + "_" + v.indexPlaceholder
}

func (v *toEnvVisitor) removeLastPaths() {
	v.jvmPaths = v.jvmPaths[:len(v.jvmPaths)-1]
	v.envPaths = v.envPaths[:len(v.envPaths)-1]
}

func (v *toEnvVisitor) getJVM() string {
	return v.join(v.prefix, v.jvmPaths, ".")
}

func (v *toEnvVisitor) getEnv() string {
	return v.join(strings.ToUpper(v.prefix), v.envPaths, "_")
}

func (v *toEnvVisitor) join(prefix string, array []string, separator string) string {
	paths := make([]string, len(array)+1)
	paths[0] = prefix
	copy(paths[1:], array)
	return strings.Join(paths, separator)
}
