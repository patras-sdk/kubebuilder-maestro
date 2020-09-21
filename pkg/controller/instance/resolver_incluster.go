package instance

import (
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client"

	kudoapi "github.com/kudobuilder/kudo/pkg/apis/kudo/v1beta1"
	"github.com/kudobuilder/kudo/pkg/kudoctl/packages"
	"github.com/kudobuilder/kudo/pkg/util/kudo"
)

// InClusterResolver is a server-side package resolver for packages that are already installed in the cluster. It is a simpler
// version of the client-side pkg/kudoctl/packages/resolver/resolver_incluster.go. The client-side version would search
// the installed OperatorVersions and try to resolve any valid combination of the operator name and its app and operator versions,
// same as we would search in the repository.
// This resolver is only used to make sure that all the dependencies of an operator exist and that referenced operator versions
// are installed and uniquely identifiable by the passed operator name, appVersion and operatorVersion parameters
// (see pkg/apis/kudo/v1beta1/operatorversion_types_helpers.go::OperatorVersionName method).
// Also, note that unlike other resolvers, the resulting 'packages.Package' struct does not contain package 'packages.Files'
// (we don't have the original files) and doesn't have an Instance resource because multiple Instances of the same
// Operator/OperatorVersion  can exist.
type InClusterResolver struct {
	c  client.Client
	ns string
}

func (r InClusterResolver) Resolve(name string, appVersion string, operatorVersion string) (*packages.Package, error) {
	ovList, err := kudoapi.ListOperatorVersions(r.ns, r.c)
	if err != nil {
		return nil, fmt.Errorf("failed to list operator versions in namespace %q: %v", r.ns, err)
	}

	if len(ovList.Items) == 0 {
		return nil, fmt.Errorf("failed to find any operator version in namespace %s", r.ns)
	}

	// Put all items into a new list to be sortable
	newOvList := kudo.SortableOperatorList{}
	for _, ovFromList := range ovList.Items {
		ovFromList := ovFromList
		newOvList = append(newOvList, &ovFromList)
	}

	// Only consider OVs for the given name
	newOvList = newOvList.FilterByName(name)

	// Sort items
	newOvList.Sort()

	// Find first matching OV
	ov := newOvList.FindFirstMatch(name, operatorVersion, appVersion).(*kudoapi.OperatorVersion) // nolint:errcheck

	if ov == nil {
		return nil, fmt.Errorf("failed to resolve operator version in namespace %q for name %q, version %q, appVersion %q", r.ns, name, operatorVersion, appVersion)
	}

	o, err := kudoapi.GetOperator(name, r.ns, r.c)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve operator %s/%s", r.ns, name)
	}

	return &packages.Package{
		Resources: &packages.Resources{
			Operator:        o,
			OperatorVersion: ov,
			Instance:        nil,
		},
		Files: nil,
	}, nil
}
