// Copyright 2015-2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package flags

import (
	"fmt"
	"os"

	"github.com/aws/amazon-ecs-cli/ecs-cli/modules/utils"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// Flag names used by the cli.
const (
	// Configure
	AccessKeyFlag           = "access-key"
	SecretKeyFlag           = "secret-key"
	SessionTokenFlag        = "session-token"
	RegionFlag              = "region"
	EndpointFlag            = "endpoint"
	AwsRegionEnvVar         = "AWS_REGION"
	AwsDefaultRegionEnvVar  = "AWS_DEFAULT_REGION"
	AwsDefaultProfileEnvVar = "AWS_DEFAULT_PROFILE"
	ProfileFlag             = "profile"
	ClusterFlag             = "cluster"
	ClusterEnvVar           = "ECS_CLUSTER"
	VerboseFlag             = "verbose"
	ClusterConfigFlag       = "cluster-config"
	ECSProfileFlag          = "ecs-profile"
	ProfileNameFlag         = "profile-name"
	ConfigNameFlag          = "config-name"
	AWSProfileFlag          = "aws-profile"
	ECSProfileEnvVar        = "ECS_PROFILE"
	AWSProfileEnvVar        = "AWS_PROFILE"
	AWSAccessKeyEnvVar      = "AWS_ACCESS_KEY_ID"
	AWSSecretKeyEnvVar      = "AWS_SECRET_ACCESS_KEY"

	// logs
	TaskIDFlag         = "task-id"
	TaskDefinitionFlag = "task-def"
	FollowLogsFlag     = "follow"
	FilterPatternFlag  = "filter-pattern"
	SinceFlag          = "since"
	StartTimeFlag      = "start-time"
	EndTimeFlag        = "end-time"
	TimeStampsFlag     = "timestamps"
	CreateLogsFlag     = "create-log-groups"

	// Service Discovery
	PrivateDNSNamespaceNameFlag                 = "private-dns-namespace"
	PrivateDNSNamespaceIDFlag                   = "private-dns-namespace-id"
	PublicDNSNamespaceIDFlag                    = "public-dns-namespace-id"
	PublicDNSNamespaceNameFlag                  = "public-dns-namespace"
	EnableServiceDiscoveryFlag                  = "enable-service-discovery"
	DNSTypeFlag                                 = "dns-type"
	DNSTTLFlag                                  = "dns-ttl"
	ServiceDiscoveryContainerNameFlag           = "sd-container-name"
	ServiceDiscoveryContainerPortFlag           = "sd-container-port"
	HealthcheckCustomConfigFailureThresholdFlag = "healthcheck-custom-config-failure-threshold"
	DeletePrivateNamespaceFlag                  = "delete-namespace"
	UpdateServiceDiscoveryFlag                  = "update-service-discovery"

	ComposeProjectNamePrefixFlag         = "compose-project-name-prefix"
	ComposeProjectNamePrefixDefaultValue = "ecscompose-"
	ComposeServiceNamePrefixFlag         = "compose-service-name-prefix"
	ComposeServiceNamePrefixDefaultValue = ComposeProjectNamePrefixDefaultValue + "service-"
	CFNStackNameFlag                     = "cfn-stack-name"
	CFNStackNamePrefixDefaultValue       = utils.ECSCLIResourcePrefix

	LaunchTypeFlag         = "launch-type"
	DefaultLaunchTypeFlag  = "default-launch-type"
	SchedulingStrategyFlag = "scheduling-strategy"

	//attribute-checker
	ContainerInstancesFlag = "container-instances"

	// Cluster
	AsgMaxSizeFlag                  = "size"
	VpcAzFlag                       = "azs"
	SecurityGroupFlag               = "security-group"
	SourceCidrFlag                  = "cidr"
	EcsPortFlag                     = "port"
	SubnetIdsFlag                   = "subnets"
	VpcIdFlag                       = "vpc"
	InstanceTypeFlag                = "instance-type"
	SpotPriceFlag                   = "spot-price"
	InstanceRoleFlag                = "instance-role"
	ImageIdFlag                     = "image-id"
	KeypairNameFlag                 = "keypair"
	CapabilityIAMFlag               = "capability-iam"
	NoAutoAssignPublicIPAddressFlag = "no-associate-public-ip-address"
	ForceFlag                       = "force"
	EmptyFlag                       = "empty"
	UserDataFlag                    = "extra-user-data"

	// Image
	RegistryIdFlag = "registry-id"
	TaggedFlag     = "tagged"
	UntaggedFlag   = "untagged"
	UseFIPSFlag    = "use-fips" // TODO: repurpose to use more generally with other services/workflows

	// Compose
	ProjectNameFlag           = "project-name"
	ComposeFileNameFlag       = "file"
	TaskRoleArnFlag           = "task-role-arn"
	ECSParamsFileNameFlag     = "ecs-params"
	ForceUpdateFlag           = "force-update"
	RegistryCredsFileNameFlag = "registry-creds"

	// Compose Service
	CreateServiceCommandName                = "create"
	DeploymentMaxPercentDefaultValue        = 200
	DeploymentMaxPercentFlag                = "deployment-max-percent"
	DeploymentMinHealthyPercentDefaultValue = 100
	DeploymentMinHealthyPercentFlag         = "deployment-min-healthy-percent"
	TargetGroupArnFlag                      = "target-group-arn"
	ContainerNameFlag                       = "container-name"
	ContainerPortFlag                       = "container-port"
	LoadBalancerNameFlag                    = "load-balancer-name"
	HealthCheckGracePeriodFlag              = "health-check-grace-period"
	RoleFlag                                = "role"
	ComposeServiceTimeOutFlag               = "timeout"
	ForceDeploymentFlag                     = "force-deployment"
  EnableExecuteCommandFlag                = "enable-execute-command"
	TargetGroupsFlag                        = "target-groups"

	// Registry Creds
	UpdateExistingSecretsFlag = "update-existing-secrets"
	RoleNameFlag              = "role-name"
	NoRoleFlag                = "no-role"
	NoOutputFileFlag          = "no-output-file"
	OutputDirFlag             = "output-dir"

	DesiredTaskStatus = "desired-status"

	ResourceTagsFlag          = "tags"
	DisableECSManagedTagsFlag = "disable-ecs-managed-tags"

	// Local
	TaskDefinitionFile    = "task-def-file"
	TaskDefinitionRemote  = "task-def-remote"
	TaskDefinitionCompose = "task-def-compose"
	ComposeOverride       = "override"
	Output                = "output"
	JSON                  = "json"
	All                   = "all"
	UseRole               = "use-role"
)

func OptRegionFlag() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name: RegionFlag + ", r",
			Usage: fmt.Sprintf(
				"[Optional] Specifies the AWS region to use. Defaults to the region configured using the configure command",
			),
		},
	}
}

func OptECSProfileFlag() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   ECSProfileFlag,
			EnvVar: ECSProfileEnvVar,
			Usage: fmt.Sprintf(
				"[Optional] Specifies the name of the ECS profile configuration to use. Defaults to the default profile configuration.",
			),
		},
	}
}

func OptAWSProfileFlag() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   AWSProfileFlag,
			EnvVar: AWSProfileEnvVar,
			Usage: fmt.Sprintf(
				"[Optional] Use the AWS credentials from an existing named profile in ~/.aws/credentials.",
			),
		},
	}
}

func OptClusterConfigFlag() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name: ClusterConfigFlag,
			Usage: fmt.Sprintf(
				"[Optional] Specifies the name of the ECS cluster configuration to use. Defaults to the default cluster configuration.",
			),
		},
	}
}

// OptionalRegionAndProfileFlags provides these flags:
// - OptRegionFlag inline overrides region
// - OptClusterConfigFlag specifies the cluster profile to read from config
// - OptECSProfileEnvVar specifies the credentials profile to read from the config
// - OptAWSProfileFlag specifies the AWS Profile to use for credential information
func OptionalRegionAndProfileFlags() []cli.Flag {
	return AppendFlags(OptRegionFlag(), OptECSProfileFlag(), OptAWSProfileFlag(), OptClusterConfigFlag())
}

// OptionalClusterFlag inline overrides cluster
func OptionalClusterFlag() cli.Flag {
	return cli.StringFlag{
		Name: ClusterFlag + ", c",
		Usage: fmt.Sprintf(
			"[Optional] Specifies the ECS cluster name to use. Defaults to the cluster configured using the configure command",
		),
	}
}

// OptionalConfigFlags returns the concatenation of OptionalRegionAndProfileFlags and OptionalClusterFlag
func OptionalConfigFlags() []cli.Flag {
	return append(OptionalRegionAndProfileFlags(), OptionalClusterFlag())
}

// OptionalLaunchTypeFlag allows users to specify the launch type for their task/service/cluster
func OptionalLaunchTypeFlag() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name: LaunchTypeFlag,
			Usage: fmt.Sprintf(
				"[Optional] Specifies the launch type. Options: EC2 or FARGATE. Overrides the default launch type stored in your cluster configuration. Defaults to EC2 if a cluster configuration is not used.",
			),
		},
	}
}

// OptionalSchedulingStrategyFlag allows users to specify the scheduling strategy for their task/service/cluster
func OptionalSchedulingStrategyFlag() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name: SchedulingStrategyFlag,
			Usage: fmt.Sprintf(
				"[Optional] Specifies the scheduling strategy type. Options: REPLICA (default) or DAEMON.",
			),
		},
	}
}

// OptionalCreateLogsFlag allows users to specify the launch type for their task/service/cluster
func OptionalCreateLogsFlag() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name: CreateLogsFlag,
			Usage: fmt.Sprintf(
				"[Optional] Create the CloudWatch log groups specified in your compose file(s).",
			),
		},
	}
}

// OptionalForceUpdateFlag allows users to force an update of running tasks on compose up.
func OptionalForceUpdateFlag() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  ForceUpdateFlag + ",u",
			Usage: "[Optional] Forces update of task or service with current run parameters",
		},
	}
}

func DebugFlag() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  VerboseFlag + ",debug",
			Usage: "[Optional] Increase the verbosity of command output to aid in diagnostics.",
		},
	}
}

// OptionalDesiredStatusFlag allows users to filter tasks returned by the ps commands
func OptionalDesiredStatusFlag() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name: DesiredTaskStatus,
			Usage: fmt.Sprintf(
				"[Optional] Filter tasks by desired status. Valid values: %s or %s",
				ecs.DesiredStatusRunning,
				ecs.DesiredStatusStopped,
			),
		},
	}
}

// UsageErrorFactory Returns a usage error function for the specified command
func UsageErrorFactory(command string) func(*cli.Context, error, bool) error {
	return func(c *cli.Context, err error, isSubcommand bool) error {
		if err != nil {
			logrus.Error(err)
		}
		err = cli.ShowCommandHelp(c, command)
		if err != nil {
			logrus.Debug(err)
		}
		os.Exit(1)
		return err
	}
}

func CFNResourceFlags() []string {
	return []string{
		AsgMaxSizeFlag,
		VpcAzFlag,
		SecurityGroupFlag,
		SourceCidrFlag,
		EcsPortFlag,
		SubnetIdsFlag,
		VpcIdFlag,
		InstanceTypeFlag,
		InstanceRoleFlag,
		ImageIdFlag,
		KeypairNameFlag,
		SpotPriceFlag,
	}
}

// AppendFlags appends a series of lists of flags
func AppendFlags(flags ...[]cli.Flag) []cli.Flag {
	var allFlags []cli.Flag
	for _, set := range flags {
		allFlags = append(allFlags, set...)
	}
	return allFlags
}
