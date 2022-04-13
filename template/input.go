/*
Copyright © 2022 The OpDev Developers

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package template

import (
	"strings"

	pluralize "github.com/gertd/go-pluralize"
)

type Input struct {
	Primary   ResourceData
	Secondary ResourceData
}

// Introspect will fill in keys that can be filled in
// based on already provided values. Values that can
// be introspected are commented with "Introspectable".
func (i *Input) Introspect() {
	// Only introspect the lower if it's empty.
	if len(i.Primary.KindLower) == 0 {
		i.Primary.KindLower = strings.ToLower(i.Primary.Kind)
	}

	if len(i.Secondary.KindLower) == 0 {
		i.Secondary.KindLower = strings.ToLower(i.Secondary.Kind)
	}

	// Only introspect the plural if it's empty.
	if len(i.Primary.KindPlural) == 0 {
		i.Primary.KindPlural = pluralize.NewClient().Plural(i.Primary.KindLower)
	}

	if len(i.Secondary.KindPlural) == 0 {
		i.Secondary.KindPlural = pluralize.NewClient().Plural(i.Secondary.KindLower)
	}
}

type ResourceData struct {
	// The package where your API exists
	// e.g. example.com/demo/api/v1alpha1
	APIImportPath string
	// The import alias to use for your API import.
	// e.g. demov1alpha1
	APIImportAlias string
	// Your kind's APIGroup. No version.
	// e.g. demo.example.com
	APIGroup string
	// The title-cased string representation of your kind.
	// e.g. MyApp
	Kind string
	// The lower-case string representation of your kind.
	// Introspectable.
	// e.g. myapp.
	KindLower string
	// The lower-case and plural string representation of your kind.
	// Introspectable.
	// e.g. myapps
	KindPlural string
}
