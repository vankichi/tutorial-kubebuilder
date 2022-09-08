/*
Copyright 2022.

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

package controllers

import (
	"context"
	"encoding/json"
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	benchmarkv1 "github.com/vankichi/tutorial-kubebuilder/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// BenchmarkOperatorReconciler reconciles a BenchmarkOperator object
type BenchmarkOperatorReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=benchmark.vankichi.github.io,resources=benchmarkoperators,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=benchmark.vankichi.github.io,resources=benchmarkoperators/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=benchmark.vankichi.github.io,resources=benchmarkoperators/finalizers,verbs=update
//+kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=events,verbs=create;update;patch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the BenchmarkOperator object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.1/pkg/reconcile
func (r *BenchmarkOperatorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// Get
	var bO benchmarkv1.BenchmarkOperator
	err := r.Get(ctx, req.NamespacedName, &bO)
	if errors.IsNotFound(err) {
		return ctrl.Result{}, nil
	}
	if err != nil {
		logger.Error(err, "unable to get BenchmarkOperator", "name", req.NamespacedName)
		return ctrl.Result{}, err
	}
	logger.Info("BenchmarkOperator:\n", "spec", bO)

	// List
	var bOL benchmarkv1.BenchmarkOperatorList
	err = r.List(ctx, &bOL, &client.ListOptions{
		Namespace: "default",
		// LabelSelector: labels.SelectorFromSet(map[string]string{"app": "sample"}),
	})

	if errors.IsNotFound(err) {
		return ctrl.Result{}, nil
	}
	if err != nil {
		logger.Error(err, "unable to get BenchmarkOperatorList", "name", req.NamespacedName)
		return ctrl.Result{}, err
	}
	for _, b := range bOL.Items {
		logger.Info("BenchmarkOperatorList:\n", "spec", b.Spec)
	}

	if !bO.ObjectMeta.DeletionTimestamp.IsZero() {
		return ctrl.Result{}, nil
	}

	// Apply Configmap
	err = r.reconcileConfigMap(ctx, bO)
	if err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *BenchmarkOperatorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	benchmarkv1.AddToScheme(r.Scheme)
	return ctrl.NewControllerManagedBy(mgr).
		For(&benchmarkv1.BenchmarkOperator{}).
		Complete(r)
}

func (r *BenchmarkOperatorReconciler) Reconcile_createOrUpdate(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	svc := &corev1.Service{}
	svc.SetNamespace("default")
	svc.SetName("sample")

	op, err := ctrl.CreateOrUpdate(ctx, r.Client, svc, func() error {
		svc.Spec.Type = corev1.ServiceTypeClusterIP
		svc.Spec.Selector = map[string]string{"app": "nginx"}
		svc.Spec.Ports = []corev1.ServicePort{
			{
				Name:       "http",
				Protocol:   corev1.ProtocolTCP,
				Port:       80,
				TargetPort: intstr.FromInt(80),
			},
		}
		return nil
	})
	if err != nil {
		return ctrl.Result{}, err
	}
	if op != controllerutil.OperationResultNone {
		fmt.Printf("Deployment %s\n", op)
	}
	return ctrl.Result{}, nil
}

func (r *BenchmarkOperatorReconciler) Reconcile_deleteWithPreConditions(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var deploy appsv1.Deployment
	err := r.Get(ctx, client.ObjectKey{Namespace: "default", Name: "sample"}, &deploy)
	if err != nil {
		return ctrl.Result{}, err
	}
	uid := deploy.GetUID()
	resourceVersion := deploy.GetResourceVersion()
	cond := metav1.Preconditions{
		UID:             &uid,
		ResourceVersion: &resourceVersion,
	}
	err = r.Delete(ctx, &deploy, &client.DeleteOptions{
		Preconditions: &cond,
	})
	return ctrl.Result{}, err
}

func (r *BenchmarkOperatorReconciler) Reconcile_deleteAllOfDeployment(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	err := r.DeleteAllOf(ctx, &appsv1.Deployment{}, client.InNamespace("default"))
	return ctrl.Result{}, err
}

func (r *BenchmarkOperatorReconciler) reconcileConfigMap(ctx context.Context, bO benchmarkv1.BenchmarkOperator) error {
	logger := log.FromContext(ctx)

	cm := &corev1.ConfigMap{}
	cm.SetNamespace(bO.Namespace)
	cm.SetName("benchmark-" + bO.Name)

	op, err := ctrl.CreateOrUpdate(ctx, r.Client, cm, func() error {
		if cm.Data == nil {
			cm.Data = make(map[string]string)
		}
		for _, content := range bO.Spec.Jobs {
			b, _ := json.Marshal(content)
			cm.Data[content.JobType] = string(b)

		}
		return nil
	})

	if err != nil {
		logger.Error(err, "unable to create or update ConfigMap")
		return err
	}
	if op != controllerutil.OperationResultNone {
		logger.Info("reconcile ConfigMap successfully", "op", op)
	}
	return nil
}
