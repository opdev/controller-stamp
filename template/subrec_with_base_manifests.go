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

var SubreconcilerControllerWithBaseManifests = `package controllers

import (
	"fmt"
	"os"
	"context"

	"github.com/imdario/mergo"
	{{ .Primary.APIImportAlias }} "{{ .Primary.APIImportPath }}"
	subrec "github.com/opdev/subreconciler"
	{{ .Secondary.APIImportAlias }} "{{ .Secondary.APIImportPath }}"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/yaml"
)

// {{ .Primary.Kind }}{{ .Secondary.Kind }}Reconciler reconciles the deployment resource.
type {{ .Primary.Kind }}{{ .Secondary.Kind }}Reconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups={{ .Primary.APIGroup }},resources={{ .Primary.KindPlural }},verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups={{ .Primary.APIGroup }},resources={{ .Primary.KindPlural }}/status,verbs=get;update;patch
//+kubebuilder:rbac:groups={{ .Primary.APIGroup }},resources={{ .Primary.KindPlural }}/finalizers,verbs=update
//+kubebuilder:rbac:groups={{ .Secondary.APIGroup }},resources={{ .Secondary.KindPlural }},verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups={{ .Secondary.APIGroup }},resources={{ .Secondary.KindPlural }}/finalizers,verbs=update

// Reconcile will ensure that the Kubernetes {{ .Secondary.Kind }} for {{ .Primary.Kind }}
// reaches the desired state.
func (r *{{ .Primary.Kind }}{{ .Secondary.Kind }}Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)
	l.Info("{{ .Secondary.KindLower }} reconciliation initiated.")
	defer l.Info("{{ .Secondary.KindLower }} reconciliation complete.")
	instanceKey := req.NamespacedName

	// Get the {{ .Primary.Kind }} instance to make sure it still exists.
	var instance {{ .Primary.APIImportAlias }}.{{ .Primary.Kind }}
	err := r.Client.Get(ctx, instanceKey, &instance)

	if apierrors.IsNotFound(err) {
		return subrec.Evaluate(subrec.DoNotRequeue())
	}

	if err != nil {
		return subrec.Evaluate(subrec.RequeueWithError(err))
	}

	/* 
		TODO:
		This template allows you to read a base resource from disk
		and the mutate that resource based on your primary resource's
		spec.

		This primary resource needs to exist on the filesystem where
		your controller runs. If you are building a container for
		your controller, make sure you've copied your bases
		into the container at build time, or plan to volume mount them
		at runtime.
		
		Feel free to delete this comment once you've made 
		your adjustments.
	*/ 

	// Read the base manifest for this resource from file.
	var new {{ .Secondary.APIImportAlias }}.{{ .Secondary.Kind }}
	resourceBaseManifestPath := "manifests/base-{{ .Secondary.KindLower }}.yml"
	base, err := os.ReadFile(resourceBaseManifestPath)
	if err != nil {
		// We encountered an error trying to read the base from disk.
		// This is unlikely to resolve itself, so do not requeue.
		l.Error(err, "could not read base manifest from disk", "path", resourceBaseManifestPath)
		return subrec.Evaluate(subrec.DoNotRequeue())
	}

	err = yaml.Unmarshal(base, &new)
	if err != nil {
		// We encountered an error trying to marshal the base on disk to the proper type.
		// This is unlikely to resolve itself, so do not requeue.
		l.Error(err, "base manifest was not for the expected resource", "kind", "{{ .Secondary.KindLower }}")
		return subrec.Evaluate(subrec.DoNotRequeue())
	}

	// TODO: modify your {{ .Secondary.KindLower }} based on your {{ .Primary.KindLower }} spec. We modify
	// the .metadata.name value here as an example.
	new.ObjectMeta.Name = fmt.Sprintf("%s-%s", instance.GetName(), "{{ .Secondary.KindLower }}")

	err = ctrl.SetControllerReference(&instance, &new, r.Scheme)
	if err != nil {
		return subrec.Evaluate(subrec.RequeueWithError(err))
	}

	// If the {{ .Secondary.KindLower }} exists, get it and patch it
	var existing {{ .Secondary.APIImportAlias }}.{{ .Secondary.Kind }}
	err = r.Client.Get(ctx, client.ObjectKeyFromObject(&new), &existing)

	if apierrors.IsNotFound(err) {
		// create the resource because it does not exist.
		l.Info("creating resource", new.Kind, new.Name)
		if err := r.Client.Create(ctx, &new); err != nil {
			return subrec.Evaluate(subrec.RequeueWithError(err))
		}
	}

	if err != nil {
		return subrec.Evaluate(subrec.RequeueWithError(err))
	}

	l.Info("updating resources if necessary", existing.Kind, existing.GetName())
	patchDiff := client.MergeFrom(existing.DeepCopy())
	if err = mergo.Merge(&existing, new, mergo.WithOverride); err != nil {
		return subrec.Evaluate(subrec.RequeueWithError(err))
	}

	if err = r.Patch(ctx, &existing, patchDiff); err != nil {
		return subrec.Evaluate(subrec.RequeueWithError(err))
	}

	return subrec.Evaluate(subrec.DoNotRequeue()) // success
}

// SetupWithManager sets up the controller with the Manager.
func (r *{{ .Primary.Kind }}{{ .Secondary.Kind }}Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&{{ .Primary.APIImportAlias }}.{{ .Primary.Kind }}{}).
		Owns(&{{ .Secondary.APIImportAlias }}.{{ .Secondary.Kind }}{}).
		Complete(r)
}
`
