// Copyright 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vpnseedserver

import (
	"fmt"

	fluentbitv1alpha2 "github.com/fluent/fluent-operator/v2/apis/fluentbit/v1alpha2"
	fluentbitv1alpha2filter "github.com/fluent/fluent-operator/v2/apis/fluentbit/v1alpha2/plugins/filter"
	fluentbitv1alpha2parser "github.com/fluent/fluent-operator/v2/apis/fluentbit/v1alpha2/plugins/parser"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"

	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	"github.com/gardener/gardener/pkg/component"
)

const (
	containerNameVPNSeedServer = "vpn-seed-server"
	containerNameEnvoyProxy    = "envoy-proxy"
)

// CentralLoggingConfiguration returns a fluent-bit parser and filter for the kube-apiserver logs.
func CentralLoggingConfiguration() (component.CentralLoggingConfig, error) {
	return component.CentralLoggingConfig{
		Filters: generateClusterFilters(),
		Parsers: generateClusterParsers(),
	}, nil
}

func generateClusterFilters() []*fluentbitv1alpha2.ClusterFilter {
	return []*fluentbitv1alpha2.ClusterFilter{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:   fmt.Sprintf("%s--%s", v1beta1constants.DeploymentNameVPNSeedServer, containerNameVPNSeedServer),
				Labels: map[string]string{v1beta1constants.LabelKeyCustomLoggingResource: v1beta1constants.LabelValueCustomLoggingResource},
			},
			Spec: fluentbitv1alpha2.FilterSpec{
				Match: fmt.Sprintf("kubernetes.*%s*%s*", v1beta1constants.DeploymentNameVPNSeedServer, containerNameVPNSeedServer),
				FilterItems: []fluentbitv1alpha2.FilterItem{
					{
						Parser: &fluentbitv1alpha2filter.Parser{
							KeyName:     "log",
							Parser:      containerNameVPNSeedServer + "-parser",
							ReserveData: pointer.Bool(true),
						},
					},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:   fmt.Sprintf("%s--%s", v1beta1constants.DeploymentNameVPNSeedServer, containerNameEnvoyProxy),
				Labels: map[string]string{v1beta1constants.LabelKeyCustomLoggingResource: v1beta1constants.LabelValueCustomLoggingResource},
			},
			Spec: fluentbitv1alpha2.FilterSpec{
				Match: fmt.Sprintf("kubernetes.*%s*%s*", v1beta1constants.DeploymentNameVPNSeedServer, containerNameEnvoyProxy),
				FilterItems: []fluentbitv1alpha2.FilterItem{
					{
						Parser: &fluentbitv1alpha2filter.Parser{
							KeyName:     "log",
							Parser:      containerNameEnvoyProxy + "-parser",
							ReserveData: pointer.Bool(true),
						},
					},
				},
			},
		},
	}
}

func generateClusterParsers() []*fluentbitv1alpha2.ClusterParser {
	return []*fluentbitv1alpha2.ClusterParser{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:   containerNameVPNSeedServer + "-parser",
				Labels: map[string]string{v1beta1constants.LabelKeyCustomLoggingResource: v1beta1constants.LabelValueCustomLoggingResource},
			},
			Spec: fluentbitv1alpha2.ParserSpec{
				Regex: &fluentbitv1alpha2parser.Regex{
					Regex:      "^(?<time>\\d{4}-\\d{2}-\\d{2}\\s+[^\\s]+)\\s+(?<log>.*)$",
					TimeKey:    "time",
					TimeFormat: "%Y-%m-%d %H:%M:%S",
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:   containerNameEnvoyProxy + "-parser",
				Labels: map[string]string{v1beta1constants.LabelKeyCustomLoggingResource: v1beta1constants.LabelValueCustomLoggingResource},
			},
			Spec: fluentbitv1alpha2.ParserSpec{
				Regex: &fluentbitv1alpha2parser.Regex{
					Regex:      "^\\[(?<time>\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}\\.\\d+\\S+?)]\\s+(?<log>.*)$",
					TimeKey:    "time",
					TimeFormat: "%Y-%m-%dT%H:%M:%S.%L%z",
				},
			},
		},
	}
}
