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
	"log"
	"os"
	"text/template"

	"github.com/opdev/controller-stamp/resource"
	ct "github.com/opdev/controller-stamp/template"
)

// Execute is the primary entrypoint.
func Execute() int {
	secondary, err := resource.Get("depsloment")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERR %s: %s", err, "falling back to a pod as the secondary resource.")
		secondary, _ = resource.Get("pod")
	}

	i := ct.Input{
		Primary: ct.ResourceData{
			APIImportAlias: "demooperatorv1alpha1",
			APIImportPath:  "example.com/demo-operator/api/v1alpha1",
			APIGroup:       "demo.example.com",
			Kind:           "MyCustomResource",
		},
		Secondary: secondary,
	}

	i.Introspect()

	toRender, err := template.New("controller").Parse(ct.StandardController)
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
