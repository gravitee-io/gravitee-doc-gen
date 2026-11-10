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
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/bootstrap/filehandlers"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/bootstrap/modules"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/bootstrap/plugin"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/examples"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("plugin an example configures", func() {
	var origDir string

	BeforeEach(func() {
		var err error
		origDir, err = os.Getwd()
		Expect(err).NotTo(HaveOccurred())
		Expect(os.Chdir("with-modules")).To(Succeed())

		bootstrap.Register(filehandlers.PropertiesFileHandler, filehandlers.PropertiesExt)
		bootstrap.Register(filehandlers.YamlFileHandler, filehandlers.YamlExt, filehandlers.YmlExt)
		bootstrap.RegisterPostProcessor("plugin", plugin.PostProcessor)
		bootstrap.RegisterPostProcessor("modules", modules.PostProcessor)
		Expect(bootstrap.Load(".")).To(Succeed())
	})

	AfterEach(func() {
		Expect(os.Chdir(origDir)).To(Succeed())
	})

	It("is the plugin of the repository when the example sets no module", func() {
		p, err := examples.PluginFor("")
		Expect(err).NotTo(HaveOccurred())
		Expect(p.ID).To(Equal("guardrails"))
	})

	It("is the plugin of the module the example sets, over the plugin of the repository", func() {
		p, err := examples.PluginFor("rule")
		Expect(err).NotTo(HaveOccurred())
		Expect(p.ID).To(Equal("guardrail-rule"))
		Expect(p.Title).To(Equal("Guardrail - Rule"))
	})

	It("fails for a module that is not declared", func() {
		_, err := examples.PluginFor("unknown")
		Expect(err).To(MatchError(ContainSubstring("module 'unknown' is not declared")))
	})

	It("fails for a module that builds no plugin", func() {
		_, err := examples.PluginFor("commons")
		Expect(err).To(MatchError(ContainSubstring("module 'commons' builds no plugin")))
	})
})
