/*
Copyright 2019 The Kubernetes Authors.

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

package aggregator

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	regenerationCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "aggregator_openapi_v2_regeneration_count",
			Help: "Counter of OpenAPI v2 spec regeneration count broken down by causing APIService name and reason.",
		},
		[]string{"apiservice", "reason"},
	)
	regenerationDurationGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "aggregator_openapi_v2_regeneration_duration",
			Help: "Gauge of OpenAPI v2 spec regeneration duration in seconds.",
		},
		[]string{"reason"},
	)
)

func init() {
	prometheus.MustRegister(regenerationCounter)
	prometheus.MustRegister(regenerationDurationGauge)
}
