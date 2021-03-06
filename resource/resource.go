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

package resource

import (
	"fmt"

	"github.com/opdev/controller-stamp/template"
)

func Get(resource string) (template.ResourceData, error) {
	switch resource {
	case "deployment", "deploy":
		return deployment, nil
	case "pod", "po":
		return pod, nil
	case "secret":
		return secret, nil
	case "configmap":
		return configmap, nil
	case "persistentvolumeclaim", "pvc":
		return pvc, nil
	case "job":
		return job, nil
	case "cronjob":
		return cronjob, nil
	default:
		return pod, fmt.Errorf(
			"woah there! I've got no clue what a \"%s\" resource is, pal",
			resource,
		)
	}
}

var deployment = template.ResourceData{
	APIImportPath:  "k8s.io/api/apps/v1",
	APIImportAlias: "appsv1",
	APIGroup:       "apps",
	Kind:           "Deployment",
	KindLower:      "deployment",
	KindPlural:     "deployments",
}

var pod = template.ResourceData{
	APIImportPath:  "k8s.io/api/core/v1",
	APIImportAlias: "corev1",
	APIGroup:       "core",
	Kind:           "Pod",
	KindLower:      "pod",
	KindPlural:     "Pods",
}

var secret = template.ResourceData{
	APIImportPath:  "k8s.io/api/core/v1",
	APIImportAlias: "corev1",
	APIGroup:       "core",
	Kind:           "Secret",
	KindLower:      "secret",
	KindPlural:     "secrets",
}

var configmap = template.ResourceData{
	APIImportPath:  "k8s.io/api/core/v1",
	APIImportAlias: "corev1",
	APIGroup:       "core",
	Kind:           "ConfigMap",
	KindLower:      "configmap",
	KindPlural:     "configmaps",
}

var pvc = template.ResourceData{
	APIImportPath:  "k8s.io/api/core/v1",
	APIImportAlias: "corev1",
	APIGroup:       "core",
	Kind:           "PersistentVolumeClaim",
	KindLower:      "persistentvolumeclaim",
	KindPlural:     "persistentvolumeclaims",
}

var job = template.ResourceData{
	APIImportPath:  "k8s.io/api/batch/v1",
	APIImportAlias: "batchv1",
	APIGroup:       "batch",
	Kind:           "Job",
	KindLower:      "job",
	KindPlural:     "jobs",
}

var cronjob = template.ResourceData{
	APIImportPath:  "k8s.io/api/batch/v1",
	APIImportAlias: "batchv1",
	APIGroup:       "batch",
	Kind:           "CronJob",
	KindLower:      "cronjob",
	KindPlural:     "cronjobs",
}
