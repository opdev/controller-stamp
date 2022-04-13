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

package cmd

import (
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/opdev/controller-stamp/resource"
	ct "github.com/opdev/controller-stamp/template"
)

// ExecuteFn is the primary entrypoint.
func ExecuteFn(userResource ct.ResourceData, userTmplSel, userSecondarySel string) int {
	secondary, err := resource.Get(userSecondarySel)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	tpl, err := ct.Get(userTmplSel)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}

	i := ct.Input{
		Primary:   userResource,
		Secondary: secondary,
	}

	i.Introspect()

	toRender, err := template.New("controller").Parse(tpl)
	if err != nil {
		log.Println("bah!")
		log.Fatal(err)
	}

	err = toRender.Execute(os.Stdout, i)
	if err != nil {
		log.Println("humbug")
		log.Fatal(err)
	}

	return 0
}
