package helloworld

import (
	"context"
	"fmt"
	v1v1alpha1 "github.com/epmd-edp/go-go-operator-sdk/pkg/apis/v1/v1alpha1"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileHelloWorld{
		client: mgr.GetClient(),
		scheme: mgr.GetScheme(),
	}
}

func add(mgr manager.Manager, r reconcile.Reconciler) error {
	c, err := controller.New("helloworld-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	return c.Watch(&source.Kind{Type: &v1v1alpha1.HelloWorld{}}, &handler.EnqueueRequestForObject{})
}

var _ reconcile.Reconciler = &ReconcileHelloWorld{}

type ReconcileHelloWorld struct {
	client client.Client
	scheme *runtime.Scheme
}

func (r *ReconcileHelloWorld) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	instance := &v1v1alpha1.HelloWorld{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if k8sErrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}
	defer r.updateStatus(instance)

	instance.Status.TargetMessage = fmt.Sprintf("Source message has been retrieved: %v", instance.Spec.SourceMessage)

	return reconcile.Result{}, err
}

func (r *ReconcileHelloWorld) updateStatus(hw *v1v1alpha1.HelloWorld) {
	err := r.client.Status().Update(context.TODO(), hw)
	if err != nil {
		_ = r.client.Update(context.TODO(), hw)
	}
}
