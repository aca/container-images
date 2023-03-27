package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	// "k8s.io/apimachinery/pkg/api/errors"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"

	// appsv1 "k8s.io/api/apps/v1"
	eventsv1 "k8s.io/api/events/v1"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

func main() {
	log.SetOutput(os.Stdout)
	// use the current context in kubeconfig
	// config, err := clientcmd.BuildConfigFromFlags("", os.getenv("HOME") + "/.kube/config")
	// if err != nil {
	// 	panic(err.Error())
	// }
	//

	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	// if you want to change the loading rules (which files in which order), you can do so here

	configOverrides := &clientcmd.ConfigOverrides{}
	// if you want to change override values or bind them to flags, there are methods to help you

	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	config, err := kubeConfig.ClientConfig()
	if err != nil {
		// Do something
		panic(err)
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	sharedInformers := informers.NewSharedInformerFactory(clientset, time.Minute*10)
	eventInformer := sharedInformers.Events().V1().Events()

	// deploymentInformer := sharedInformers.Apps().V1().Deployments()
	//
	// deploymentInformer.Informer().AddEventHandler(
	// 	cache.ResourceEventHandlerFuncs{
	// 		AddFunc: func(obj interface{}) {
	// 		},
	// 		UpdateFunc: func(oldObj interface{}, newObj interface{}) {
	// 			e1 := oldObj.(*appsv1.Deployment)
	// 			e2 := newObj.(*appsv1.Deployment)
	// 			_ = e1
	// 			_ = e2
	// 			// log.Println(e1.GetName(), e2.GetName())
	// 			// log.Println(cmp.Diff(e1, e2))
	// 		},
	//
	// 		DeleteFunc: func(obj interface{}) {},
	// 	},
	// )

	eventInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				e := obj.(*eventsv1.Event)
				// b, err := json.MarshalIndent(e, "", "\t")
				// if err != nil {
				// 	log.Println(err)
				// 	return
				// }

				b, err := json.Marshal(e)
				if err != nil {
					log.Println(err)
					return
				}

				log.Println(string(b))
				// log.Printf("[ADD] %v", e.Message)
				// log.Printf("[ADD] %v", e.String())
				// log.Printf("[ADD] %v", e.ReportingInstance)

				// e, _ := obj.(*v1.Pod)
				// log.Printf("[ADD] %v", e.GetName())
			},
			UpdateFunc: func(oldObj interface{}, newObj interface{}) {},

			DeleteFunc: func(obj interface{}) {},
		},
	)

	{
		stop := make(chan struct{})
		sharedInformers.Start(stop)
	}

	{
		// stop := make(chan struct{})
		// deploymentInformer.Start(stop)
	}
	select {}
}
