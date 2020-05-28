package install

import (
	"errors"
	"fmt"
	"time"

	pollwait "k8s.io/apimachinery/pkg/util/wait"

	"github.com/kudobuilder/kudo/pkg/apis/kudo/v1beta1"
	"github.com/kudobuilder/kudo/pkg/kudoctl/clog"
	"github.com/kudobuilder/kudo/pkg/kudoctl/util/kudo"
)

func installInstance(client *kudo.Client, instance *v1beta1.Instance) error {
	existingInstance, err := client.GetInstance(instance.Name, instance.Namespace)
	if err != nil {
		return fmt.Errorf("failed to verify existing instance: %v", err)
	}

	if existingInstance != nil {
		return clog.Errorf(
			"cannot install instance '%s' because an instance of that name already exists in namespace %s",
			instance.Name,
			instance.Namespace)
	}

	if _, err := client.InstallInstanceObjToCluster(instance, instance.Namespace); err != nil {
		return fmt.Errorf("failed to install instance %s: %v", instance.Name, err)
	}

	clog.Printf("instance.%s/%s created", instance.APIVersion, instance.Name)
	return nil
}

func waitForInstance(client *kudo.Client, instance *v1beta1.Instance, timeout time.Duration) error {
	err := client.WaitForInstance(instance.Name, instance.Namespace, nil, timeout)
	if errors.Is(err, pollwait.ErrWaitTimeout) {
		clog.Printf("timeout waiting for instance.%s/%s", instance.APIVersion, instance.Name)
	}
	if err != nil {
		return fmt.Errorf("failed to wait on instance %s: %v", instance.Name, err)
	}

	clog.Printf("instance.%s/%s ready", instance.APIVersion, instance.Name)
	return nil
}