/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2018 Red Hat, Inc.
 */

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ksimon1/cpu-model-nfd-plugin/pkg/collector"
)

func main() {
	domCapabilitiesFilePath := flag.String("domCapabilitiesFile", "/etc/kubernetes/node-feature-discovery/source.d/virsh_domcapabilities.xml", "virsh domcapabilities file")
	flag.Parse()

	modelBlackList := os.Getenv("CPU_MODEL_BLACK_LIST")
	cpuModelBlackList := map[string]bool{}
	if modelBlackList != "" {
		for _, model := range strings.Split(modelBlackList, " ") {
			cpuModelBlackList[strings.ToLower(model)] = true
		}
	}

	result, err := collector.CollectData(*domCapabilitiesFilePath, cpuModelBlackList)
	if err != nil {
		os.Exit(1)
	}

	for _, cpu := range result {
		fmt.Println("/cpu-model-" + cpu)
	}
}
