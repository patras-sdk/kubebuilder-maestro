package plan

import (
	"encoding/json"
	"fmt"
	"github.com/kudobuilder/kudo/pkg/kudoctl/util/check"
	"github.com/pkg/errors"
	"time"

	kudov1alpha1 "github.com/kudobuilder/kudo/pkg/apis/kudo/v1alpha1"
	"github.com/spf13/cobra"
	"github.com/xlab/treeprint"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

// NewPlanHistoryCmd creates a command that shows the plan instory for an instance
func NewPlanHistoryCmd() *cobra.Command {
	listCmd := &cobra.Command{
		//Args: cobra.ExactArgs(1),
		Use:   "history",
		Short: "Lists history to a specific framework-version of an instance.",
		Long: `
	# View plan status
	kudoctl plan history <frameworkVersion> --instance=<instanceName>`,
		RunE: runHistory,
	}

	listCmd.Flags().StringVar(&instance, "instance", "", "The instance name.")
	listCmd.Flags().StringVar(&kubeConfig, "kubeconfig", "", "The file path to kubernetes configuration file; defaults to $HOME/.kube/config")
	listCmd.Flags().StringVar(&namespace, "namespace", "default", "The namespace where the operator watches for changes.")

	return listCmd
}

func runHistory(cmd *cobra.Command, args []string) error {

	instanceFlag, err := cmd.Flags().GetString("instance")
	if err != nil || instanceFlag == "" {
		return fmt.Errorf("Flag Error: Please set instance flag, e.g. \"--instance=<instanceName>\"")
	}

	configPath, err := check.KubeConfigLocationOrDefault(kubeConfig)
	if err != nil {
		return fmt.Errorf("error when getting default kubeconfig path: %+v", err)
	}
	kubeConfig = configPath
	if err := check.ValidateKubeConfigPath(kubeConfig); err != nil {
		return errors.WithMessage(err, "could not check kubeconfig path")
	}

	_, err = cmd.Flags().GetString("kubeconfig")
	// Todo: wrong flag
	if err != nil || instanceFlag == "" {
		return fmt.Errorf("flag Error: %v", err)
	}

	err = planHistory(args)
	if err != nil {
		return fmt.Errorf("client Error: %v", err)
	}
	return nil
}

func planHistory(args []string) error {

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		return err
	}

	//  Create a Dynamic Client to interface with CRDs.
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return err
	}

	planExecutionsGVR := schema.GroupVersionResource{
		Group:    "kudo.k8s.io",
		Version:  "v1alpha1",
		Resource: "planexecutions",
	}

	var labelSelector string
	if len(args) == 0 {
		fmt.Printf("History of all plan-executions for instance \"%s\" in namespace \"%s\":\n", instance, namespace)
		labelSelector = "instance=" + instance
	} else {
		fmt.Printf("History of plan-executions for instance \"%s\" in namespace \"%s\" to framework-version \"%s\":\n", instance, namespace, args[0])
		labelSelector = "framework-version=" + args[0] + ", instance=" + instance
	}

	instObj, err := dynamicClient.Resource(planExecutionsGVR).Namespace(namespace).List(metav1.ListOptions{
		LabelSelector: labelSelector,
	})
	if err != nil {
		return err
	}

	mInstObj, err := instObj.MarshalJSON()
	if err != nil {
		return err
	}

	planExecutionList := kudov1alpha1.PlanExecutionList{}

	err = json.Unmarshal(mInstObj, &planExecutionList)
	if err != nil {
		return err
	}

	tree := treeprint.New()

	if len(planExecutionList.Items) == 0 {
		fmt.Printf("No history found for \"%s\" in namespace \"%s\".\n", instance, namespace)
	} else {
		for _, i := range planExecutionList.Items {
			duration := time.Since(i.CreationTimestamp.Time)
			historyDisplay := fmt.Sprintf("%s (created %v ago)", i.Name, duration.Round(time.Second))
			tree.AddBranch(historyDisplay)
		}

		fmt.Println(tree.String())
	}

	return nil
}
