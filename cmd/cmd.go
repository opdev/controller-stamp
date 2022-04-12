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
	"log"
	"os"
	"text/template"

	ct "github.com/opdev/controller-stamp/template"
)

// Execute is the primary entrypoint.
func Execute() int {
	i := ct.Input{
		CRDAPIImportAlias:         "demooperatorv1alpha1",
		CRDAPIImportPath:          "example.com/demo-operator/api/v1alpha1",
		CRDAPIGroup:               "demo.example.com",
		CRDKind:                   "MyCustomResource",
		CRDKindLower:              "mycustomresource",
		CRDResourcePlural:         "mycustomresources",
		SecondaryAPIImportPath:    "k8s.io/api/apps/v1",
		SecondaryAPIImportAlias:   "appsv1",
		SecondaryKind:             "Deployment",
		SecondaryKindLower:        "deployment",
		SecondaryResourceAPIGroup: "apps",
		SecondaryResourcePlural:   "deployments",
	}

	toRender, err := template.New("controller").Parse(ct.StandardControllerTemplate)
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
