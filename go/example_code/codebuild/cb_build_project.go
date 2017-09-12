/*
   Copyright 2010-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.

   This file is licensed under the Apache License, Version 2.0 (the "License").
   You may not use this file except in compliance with the License. A copy of
   the License is located at

    http://aws.amazon.com/apache2.0/

   This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
   CONDITIONS OF ANY KIND, either express or implied. See the License for the
   specific language governing permissions and limitations under the License.
*/

package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codebuild"
	"fmt"
	"os"
)

// Builds a CodeBuild project in the region configured in the shared config
func main() {
	// Requires one argument, the name of the project.
	if len(os.Args) != 2 {
		fmt.Println("Project name required!")
		os.Exit(1)
	}

	project := os.Args[1]

	// Initialize a session that the SDK will use to load configuration,
	// credentials, and region from the shared config file. (~/.aws/config).
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create CodeBuild service client
	svc := codebuild.New(sess)

	// Build the project
	_, err := svc.StartBuild(&codebuild.StartBuildInput{ProjectName: aws.String(project)})

	if err != nil {
		fmt.Println("Got error building project: ", err)
		os.Exit(1)
	}

	fmt.Printf("Started build for project %q\n", project)
}