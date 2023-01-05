package archivecontroller

import (
	k8upv1 "github.com/k8up-io/k8up/v2/api/v1"
	"github.com/k8up-io/k8up/v2/operator/locker"
	"github.com/k8up-io/k8up/v2/operator/reconciler"
	controllerruntime "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// +kubebuilder:rbac:groups=k8up.io,resources=archives,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=k8up.io,resources=archives/status;archives/finalizers,verbs=get;update;patch

// SetupWithManager configures the reconciler.
func SetupWithManager(mgr controllerruntime.Manager) error {
	name := "archive.k8up.io"
	r := reconciler.NewReconciler[*k8upv1.Archive, *k8upv1.ArchiveList](mgr.GetClient(), &ArchiveReconciler{
		Kube:   mgr.GetClient(),
		Locker: &locker.Locker{Kube: mgr.GetClient()},
	})
	return controllerruntime.NewControllerManagedBy(mgr).
		For(&k8upv1.Archive{}).
		Named(name).
		WithEventFilter(predicate.GenerationChangedPredicate{}).
		Complete(r)
}
