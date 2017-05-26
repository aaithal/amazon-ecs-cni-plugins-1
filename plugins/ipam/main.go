// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/amazon-ecs-cni-plugins/pkg/logger"
	"github.com/aws/amazon-ecs-cni-plugins/pkg/version"
	"github.com/aws/amazon-ecs-cni-plugins/plugins/ipam/commands"
	"github.com/aws/amazon-ecs-cni-plugins/plugins/ipam/version/cnispec"
	log "github.com/cihub/seelog"
	"github.com/containernetworking/cni/pkg/skel"
)

const (
	defaultLogFilePath = "/var/log/ecs/ecs-cni-ipam-plugin.log"
)

func main() {
	defer log.Flush()
	logger.SetupLogger(logger.GetLogFileLocation(defaultLogFilePath))

	var printVersion bool
	flag.BoolVar(&printVersion, "version", false, "prints version and exits")
	flag.Parse()

	if printVersion {
		if err := printVersionInfo(); err != nil {
			os.Stderr.WriteString(
				fmt.Sprintf("Error getting version string: %s", err.Error()))
		}
		return
	}

	skel.PluginMain(commands.Add, commands.Del, cnispec.GetSpecVersionSupported())
}

func printVersionInfo() error {
	versionInfo, err := version.String()
	if err != nil {
		return err
	}
	fmt.Println(versionInfo)
	return nil
}
