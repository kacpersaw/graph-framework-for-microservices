package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// Client Error Codes.
type ClientErrorCode int

const (
	// Make sure to your error codes from only within the following range
	// 8-63
	// 79-125
	// 166-199
	// 243-255
	// DO NOT skip the sequence unless there is a valid reason, which
	// should be documented along with the const var.
	UNHANDLED_ERROR                    ClientErrorCode = 8
	INTERNAL_ERROR                     ClientErrorCode = 9
	DATAMODEL_DIRECTORY_NOT_FOUND      ClientErrorCode = 10
	DATAMODEL_BUILD_FAILED             ClientErrorCode = 11
	DOCKER_NOT_RUNNING                 ClientErrorCode = 12
	RUNTIME_INSTALL_FAILED             ClientErrorCode = 13
	RUNTIME_UNINSTALL_FAILED           ClientErrorCode = 14
	DATAMODEL_INSTALL_FAILED           ClientErrorCode = 15
	APPLICATION_DEPLOY_FAILED          ClientErrorCode = 16
	DATAMODEL_INIT_FAILED              ClientErrorCode = 17
	APPLICATION_PACKAGE_FAILED         ClientErrorCode = 18
	APPLICATION_PUBLISH_FAILED         ClientErrorCode = 19
	APPLICATION_RUN_FAILED             ClientErrorCode = 20
	APPLICATION_OPERATOR_CREATE_FAILED ClientErrorCode = 21
	CLI_UPGRADE_FAILED                 ClientErrorCode = 22
)

// ClientError defines error and information around it that are specific
// to this CLI client.
//
// Any error logged by the client should have the description of the error
// and the next steps that users of the client can check / follow to resolve
// and handle this error.
type ClientError struct {
	// Error code to be associated with this error.
	ErrorCode ClientErrorCode `json:"ErrorCode,omitempty"`

	// Detailed description about this error.
	Description string `json:"Description"`

	// What can the users of the client do when this error is encountered.
	// These can be series of things to check or if no mitigation is possible
	// can be a directive to reach out to support.
	WhatNext []string `json:"WhatNext"`

	// An underlying / custom error that should be wrapped around this specific error.
	// This allows us to propagate standard errors as part of the client error, so
	// the original error / msg is preserved and passed to the caller of the client.
	CustomError string `json:"CustomError,omitempty"`

	// If set to true, will terminate the client after error is printed out.
	fatal bool
}

var wellKnownErrors = map[ClientErrorCode]ClientError{
	UNHANDLED_ERROR: {
		Description: "unhandled error",
		WhatNext: []string{
			"The error is unhandled and unresolvable. Reach out to support",
		},
		fatal: true,
	},
	INTERNAL_ERROR: {
		Description: "error while processing internal logic",
		WhatNext: []string{
			"error is internal to the implementation and environment. Reach out to support",
		},
		fatal: false,
	},
	DATAMODEL_DIRECTORY_NOT_FOUND: {
		Description: "unable to find datamodel directory",
		WhatNext: []string{
			"ensure that the command is being executed from application top directory",
			"run with --debug option to get detailed logging on the build",
		},
		fatal: true,
	},
	DATAMODEL_BUILD_FAILED: {
		Description: "datamodel build failed",
		WhatNext: []string{
			"run with --debug option to get detailed logging on the build",
			"verify dsl for code and syntax errors",
		},
		fatal: true,
	},
	DOCKER_NOT_RUNNING: {
		Description: "docker daemon not running",
		WhatNext: []string{
			"run with --debug option to get detailed logging on the build",
			"docker daemon is running on the host",
			"you have permissions to access docker process",
		},
		fatal: true,
	},
	RUNTIME_INSTALL_FAILED: {
		Description: "runtime installation failed",
		WhatNext: []string{
			"kubectl command line is installed and available",
			"kubernetes cluster is reachable through kubectl command line",
		},
		fatal: true,
	},
	RUNTIME_UNINSTALL_FAILED: {
		Description: "runtime installation failed",
		WhatNext: []string{
			"kubectl command line is installed and available",
			"kubernetes cluster is reachable through kubectl command line",
		},
		fatal: true,
	},
	DATAMODEL_INSTALL_FAILED: {
		Description: "datamodel installation failed",
		WhatNext: []string{
			"verify that kubectl command line is installed and available",
			"verify that kubernetes cluster is reachable through kubectl command line",
			"verify that runtime is successfully intalled on the kubernetes cluster",
		},
		fatal: true,
	},
	APPLICATION_DEPLOY_FAILED: {
		Description: "application deploy failed",
		WhatNext: []string{
			"kubectl command line is installed and available",
			"kubernetes cluster is reachable through kubectl command line",
			"depolyment image is pushed and available in the image registry",
		},
		fatal: true,
	},
	DATAMODEL_INIT_FAILED: {
		Description: "datamodel init failed",
		WhatNext: []string{
			"user has write permissions on the disk to be able to create files and directories",
		},
		fatal: true,
	},
	APPLICATION_PACKAGE_FAILED: {
		Description: "application package failed",
		WhatNext: []string{
			"the user has write permissions on the disk to be able to create files and directories",
		},
		fatal: true,
	},
	APPLICATION_PUBLISH_FAILED: {
		Description: "application publish failed",
		WhatNext: []string{
			"user has write permissions on the disk to be able to create files and directories",
			"the user has permissions to publish the image to image registry",
		},
		fatal: true,
	},
	APPLICATION_RUN_FAILED: {
		Description: "application publish failed",
		WhatNext: []string{
			"kubectl command line is installed and available",
			"kubernetes cluster is reachable through kubectl command line",
			"depolyment image is pushed and available in the image registry",
		},
		fatal: true,
	},
	APPLICATION_OPERATOR_CREATE_FAILED: {
		Description: "application operator create failed",
		WhatNext: []string{
			"the user has write permissions on the disk to be able to create files and directories",
		},
		fatal: true,
	},
	CLI_UPGRADE_FAILED: {
		Description: "cli upgrade failed",
		WhatNext: []string{
			"run with --debug option to get detailed logging on the build",
			"verify that the requested version is available for upgrade",
			"verify permissions and access to cli repo",
		},
		fatal: true,
	},
}

// GetError returns an client error object given a predefined and well known
// error code.
//
// If the error code is unknown, then an error object of type UNHANDLED_ERROR
// is returned to the caller.
func GetError(code ClientErrorCode) ClientError {
	if val, ok := wellKnownErrors[code]; ok {
		val.ErrorCode = code
		return val
	}
	return GetError(UNHANDLED_ERROR)
}

// GetCustomError returns an client error object given a predefined and well known
// error code and a custom error that needs to be propagated into the client error.
//
// If the error code is unknown, then an error object of type UNHANDLED_ERROR
// is returned to the caller.
func GetCustomError(code ClientErrorCode, customErr error) ClientError {
	if val, ok := wellKnownErrors[code]; ok {
		val.CustomError = customErr.Error()
		val.ErrorCode = code
		return val
	}
	return GetError(UNHANDLED_ERROR)
}

// Error returns an error string associated with a client error.
func (e ClientError) Error() string {
	ret, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		fmt.Printf("FATAL: marshalling failed for error with code %+v", e.ErrorCode)
		os.Exit(1)
	}
	return string(ret)
}

// Print prints an error string associated with a client error on the console.
func (e ClientError) Print() ClientError {
	fmt.Println(e.Error())
	return e
}

// ExitIfFatalOrReturn determines if an error is fatal and it it is, terminate the client
// execution. If error is determined to be non-fatal, it returns back the original error
// without any change for subsequent eval.
func (e ClientError) ExitIfFatalOrReturn() ClientError {
	if e.fatal {
		os.Exit(1)
	}
	return e
}
