//
// Copyright 2022 Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package k8s

import (
	"github.com/redhat-appstudio/pvc-cleaner/pkg"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func getOusideClusterConfig(kubeConfig string) *rest.Config {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		panic(err.Error())
	}

	return config
}

func getInsideClusterConfig() *rest.Config {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	return config
}

func GetClusterConfig() *rest.Config {
	if pkg.IsOutSideClusterConfig() {
		kubeconfig := pkg.GetClusterConfigPath()
		return getOusideClusterConfig(kubeconfig)
	} else {
		return getInsideClusterConfig()
	}
}
