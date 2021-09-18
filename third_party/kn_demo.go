package third_party

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//dynamicfake "knative.dev/client/pkg/dynamic/fake"
	"knative.dev/client/pkg/kn/commands"
	servinglib "knative.dev/client/pkg/serving"
	eventingV1 "knative.dev/eventing/pkg/apis/eventing/v1"
	messagingV1 "knative.dev/eventing/pkg/apis/messaging/v1"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"
)

func KnConfig() {
	// get a service
	p := commands.KnParams{}
	p.Initialize()
	ctx := context.TODO()

	client, _ := p.NewServingClient("default")
	service, _ := client.GetService(ctx, "helloworld-go")
	fmt.Println(service.GetName())

	// list services
	serviceList, _ := client.ListServices(ctx)
	for _, v := range serviceList.Items {
		fmt.Println(v.GetName())
	}

	ch := messagingV1.Channel{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{},
		Spec:       messagingV1.ChannelSpec{},
		Status:     messagingV1.ChannelStatus{},
	}
	fmt.Println(ch)

	broker := eventingV1.Broker{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{},
		Spec:       eventingV1.BrokerSpec{},
		Status:     eventingV1.BrokerStatus{},
	}
	fmt.Println(broker)

	// create a service
	var svcInstance = &servingv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "guotuo-sdk-test3",
			Namespace: "default",
		},
	}
	srv := servingv1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "",
			APIVersion: "",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:                       "",
			GenerateName:               "",
			Namespace:                  "",
			SelfLink:                   "",
			UID:                        "",
			ResourceVersion:            "",
			Generation:                 0,
			CreationTimestamp:          metav1.Time{},
			DeletionTimestamp:          nil,
			DeletionGracePeriodSeconds: nil,
			Labels:                     nil,
			Annotations:                nil,
			OwnerReferences:            nil,
			Finalizers:                 nil,
			ClusterName:                "",
			ManagedFields:              nil,
		},
		Spec: servingv1.ServiceSpec{
			ConfigurationSpec: servingv1.ConfigurationSpec{},
			RouteSpec:         servingv1.RouteSpec{},
		},
		Status: servingv1.ServiceStatus{},
	}
	fmt.Println(srv)

	svcInstance.Spec.Template = servingv1.RevisionTemplateSpec{
		Spec: servingv1.RevisionSpec{},
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				servinglib.UserImageAnnotationKey: "",
			},
		},
	}

	// svcInstance.Spec.Template.Spec.PodSpec.Containers = []corev1.Container{{
	// 	Image: "docker.io/lijiawang/helloworld-go:v1",
	// 	Name:  "hwg",
	// }}

	svcInstance.Spec.Template.Spec.Containers = []corev1.Container{{Image: "guotuo1024/knative-web-demo:version-1.0.0"}}

	// servinglib.UpdateImage(svcInstance.Spec.Template, "docker.io/guotuo1024/knative-web-demo:v1")

	err := client.CreateService(ctx, svcInstance)
	if err != nil {
		fmt.Println(err)
	}
	// Update
	targetService, _ := client.GetService(ctx, "guotuo-sdk-test3")
	fmt.Println("Will update service " + targetService.GetName())
	//servinglib.UpdateImage(&targetService.Spec.Template, "guotuo1024/knative-web-demo:version-2.0.0")
	client.UpdateService(ctx, targetService)
}
