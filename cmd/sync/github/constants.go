// Copyright 2023 The Falco Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package github

const (
	commandName             = "github"
	commandShortDescription = "Synchronize Peribolos config on remote GitHub repositories via Pull Request"
	commandExample          = `
peribolos-syncer sync github --org=acme --team=app-maintainers
--peribolos-config-path=config/org.yaml --peribolos-config-repository=community --peribolos-config-git-ref=main
--owners-repository=app --owners-git-ref=main --owners-path=OWNERS
--github-username=bot --github-token-path=./bot_token
--git-author-name=bot --git-author-email="bot@acme.org"
--gpg-public-key=./bot.pub --gpg-private-key=./bot.asc
`

	syncerSignature     = "Autogenerated with [peribolos-syncer](https://github.com/falcosecurity/peribolos-syncer)."
	ownersDoc           = "https://docs.prow.k8s.io/docs/components/plugins/approve/approvers/#overview"
	peribolosConfigFile = "org.yaml"
)
