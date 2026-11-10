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
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/bootstrap/modules"
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

	When("the repository builds several plugins, one per module", func() {
		var origDir string

		BeforeEach(func() {
			var err error
			origDir, err = os.Getwd()
			Expect(err).NotTo(HaveOccurred())
			Expect(os.Chdir("with-modules")).To(Succeed())

			bootstrap.Register(bplugin.PropertiesFileHandler, bplugin.PropertiesExt)
			bootstrap.Register(bplugin.YamlFileHandler, bplugin.YamlExt, bplugin.YmlExt)
			bootstrap.RegisterPostProcessor("plugin", plugin.PostProcessor)
			bootstrap.RegisterPostProcessor("modules", modules.PostProcessor)
			Expect(bootstrap.Load(".")).To(Succeed())
		})

		AfterEach(func() {
			Expect(os.Chdir(origDir)).To(Succeed())
		})

		It("names the repository of the plugins after the repository property", func() {
			p, ok := bootstrap.GetData("plugin").(plugin.Plugin)
			Expect(ok).To(BeTrue())
			Expect(p.ID).To(Equal("guardrails"))
			Expect(p.Repository).To(Equal("gravitee-policy-guardrails"))
		})

		It("loads the plugin each module builds", func() {
			rule, ok := modules.Find("rule")
			Expect(ok).To(BeTrue())
			Expect(rule.Plugin).NotTo(BeNil())
			Expect(rule.Plugin.ID).To(Equal("guardrail-rule"))
			Expect(rule.Plugin.Title).To(Equal("Guardrail - Rule"))
			Expect(rule.Plugin.Repository).To(Equal("gravitee-policy-guardrail-rule"))
			Expect(rule.Plugin.ApiTypes).To(ConsistOf(plugin.HttpLlmProxyApiType))
		})

		It("reads the guardian phases of a module in declared order, once each, without unsupported ones", func() {
			rule, _ := modules.Find("rule")
			Expect(rule.Plugin.GuardianPhases).To(Equal([]plugin.FlowPhase{plugin.ResponsePhase, plugin.RequestPhase}))
		})

		It("loads no plugin for a module without plugin.properties", func() {
			commons, ok := modules.Find("commons")
			Expect(ok).To(BeTrue())
			Expect(commons.Plugin).To(BeNil())
		})

		It("finds no module of an unknown id", func() {
			_, ok := modules.Find("unknown")
			Expect(ok).To(BeFalse())
		})
	})

	When("a plugin declares no repository nor guardian phase", func() {
		var origDir string

		BeforeEach(func() {
			var err error
			origDir, err = os.Getwd()
			Expect(err).NotTo(HaveOccurred())
			Expect(os.Chdir("with-docgen-plugin")).To(Succeed())
			bootstrap.Register(bplugin.PropertiesFileHandler, bplugin.PropertiesExt)
			bootstrap.RegisterPostProcessor("plugin", plugin.PostProcessor)
			Expect(bootstrap.Load(".")).To(Succeed())
		})

		AfterEach(func() {
			Expect(os.Chdir(origDir)).To(Succeed())
		})

		It("names the repository gravitee-<type>-<id>, as the badges always did", func() {
			p, ok := bootstrap.GetData("plugin").(plugin.Plugin)
			Expect(ok).To(BeTrue())
			Expect(p.Repository).To(Equal("gravitee-policy-test-plugin"))
		})

		It("runs in no Guardian", func() {
			p, ok := bootstrap.GetData("plugin").(plugin.Plugin)
			Expect(ok).To(BeTrue())
			Expect(p.GuardianPhases).To(BeEmpty())
		})
	})
})
