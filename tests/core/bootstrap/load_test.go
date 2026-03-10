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

package bootstrap

import (
	"os"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/bootstrap"
	bplugin "github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/bootstrap/filehandlers"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/bootstrap/plugin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("cloud object context methods", func() {
	When("loading an empty bootstrap file", func() {
		It("trigger no error", func() {
			Expect(bootstrap.Load("empty")).To(Succeed())
			Expect(bootstrap.GetExported()).To(HaveKeyWithValue("RootDir", "empty"))
		})
	})

	When("plugin.properties is missing but .docgen/plugin.properties fallback exists", func() {
		var origDir string

		BeforeEach(func() {
			var err error
			origDir, err = os.Getwd()
			Expect(err).NotTo(HaveOccurred())
			Expect(os.Chdir("with-docgen-plugin")).To(Succeed())

			bootstrap.Register(bplugin.PropertiesFileHandler, bplugin.PropertiesExt)
			bootstrap.Register(bplugin.YamlFileHandler, bplugin.YamlExt, bplugin.YmlExt)
			bootstrap.RegisterPostProcessor("plugin", plugin.PostProcessor)
		})

		AfterEach(func() {
			Expect(os.Chdir(origDir)).To(Succeed())
		})

		It("falls back to .docgen/plugin.properties and loads plugin metadata", func() {
			Expect(bootstrap.Load(".")).To(Succeed())
			p, ok := bootstrap.GetData("plugin").(plugin.Plugin)
			Expect(ok).To(BeTrue())
			Expect(p.ID).To(Equal("test-plugin"))
			Expect(p.Title).To(Equal("Test Plugin"))
			Expect(p.Type).To(Equal("policy"))
		})
	})
})
