package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// DescribeCanariesResponse represents the DescribeCanariesResponse schema from the OpenAPI specification
type DescribeCanariesResponse struct {
	Canaries interface{} `json:"Canaries,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListGroupsRequest represents the ListGroupsRequest schema from the OpenAPI specification
type ListGroupsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// StartCanaryResponse represents the StartCanaryResponse schema from the OpenAPI specification
type StartCanaryResponse struct {
}

// VisualReferenceOutput represents the VisualReferenceOutput schema from the OpenAPI specification
type VisualReferenceOutput struct {
	Basecanaryrunid interface{} `json:"BaseCanaryRunId,omitempty"`
	Basescreenshots interface{} `json:"BaseScreenshots,omitempty"`
}

// CanaryLastRun represents the CanaryLastRun schema from the OpenAPI specification
type CanaryLastRun struct {
	Canaryname interface{} `json:"CanaryName,omitempty"`
	Lastrun interface{} `json:"LastRun,omitempty"`
}

// CanaryCodeInput represents the CanaryCodeInput schema from the OpenAPI specification
type CanaryCodeInput struct {
	S3bucket interface{} `json:"S3Bucket,omitempty"`
	S3key interface{} `json:"S3Key,omitempty"`
	S3version interface{} `json:"S3Version,omitempty"`
	Zipfile interface{} `json:"ZipFile,omitempty"`
	Handler interface{} `json:"Handler"`
}

// DescribeCanariesRequest represents the DescribeCanariesRequest schema from the OpenAPI specification
type DescribeCanariesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Names interface{} `json:"Names,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// RuntimeVersion represents the RuntimeVersion schema from the OpenAPI specification
type RuntimeVersion struct {
	Description interface{} `json:"Description,omitempty"`
	Releasedate interface{} `json:"ReleaseDate,omitempty"`
	Versionname interface{} `json:"VersionName,omitempty"`
	Deprecationdate interface{} `json:"DeprecationDate,omitempty"`
}

// GetCanaryRunsRequest represents the GetCanaryRunsRequest schema from the OpenAPI specification
type GetCanaryRunsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// EnvironmentVariablesMap represents the EnvironmentVariablesMap schema from the OpenAPI specification
type EnvironmentVariablesMap struct {
}

// S3EncryptionConfig represents the S3EncryptionConfig schema from the OpenAPI specification
type S3EncryptionConfig struct {
	Kmskeyarn interface{} `json:"KmsKeyArn,omitempty"`
	Encryptionmode interface{} `json:"EncryptionMode,omitempty"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// CreateGroupResponse represents the CreateGroupResponse schema from the OpenAPI specification
type CreateGroupResponse struct {
	Group interface{} `json:"Group,omitempty"`
}

// StopCanaryResponse represents the StopCanaryResponse schema from the OpenAPI specification
type StopCanaryResponse struct {
}

// CanaryScheduleOutput represents the CanaryScheduleOutput schema from the OpenAPI specification
type CanaryScheduleOutput struct {
	Durationinseconds interface{} `json:"DurationInSeconds,omitempty"`
	Expression interface{} `json:"Expression,omitempty"`
}

// CanaryScheduleInput represents the CanaryScheduleInput schema from the OpenAPI specification
type CanaryScheduleInput struct {
	Durationinseconds interface{} `json:"DurationInSeconds,omitempty"`
	Expression interface{} `json:"Expression"`
}

// GroupSummary represents the GroupSummary schema from the OpenAPI specification
type GroupSummary struct {
	Arn interface{} `json:"Arn,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// ArtifactConfigInput represents the ArtifactConfigInput schema from the OpenAPI specification
type ArtifactConfigInput struct {
	S3encryption interface{} `json:"S3Encryption,omitempty"`
}

// DisassociateResourceRequest represents the DisassociateResourceRequest schema from the OpenAPI specification
type DisassociateResourceRequest struct {
	Resourcearn interface{} `json:"ResourceArn"`
}

// GetGroupRequest represents the GetGroupRequest schema from the OpenAPI specification
type GetGroupRequest struct {
}

// CreateGroupRequest represents the CreateGroupRequest schema from the OpenAPI specification
type CreateGroupRequest struct {
	Name interface{} `json:"Name"`
	Tags interface{} `json:"Tags,omitempty"`
}

// CanaryCodeOutput represents the CanaryCodeOutput schema from the OpenAPI specification
type CanaryCodeOutput struct {
	Handler interface{} `json:"Handler,omitempty"`
	Sourcelocationarn interface{} `json:"SourceLocationArn,omitempty"`
}

// DeleteCanaryResponse represents the DeleteCanaryResponse schema from the OpenAPI specification
type DeleteCanaryResponse struct {
}

// DeleteGroupRequest represents the DeleteGroupRequest schema from the OpenAPI specification
type DeleteGroupRequest struct {
}

// CanaryRun represents the CanaryRun schema from the OpenAPI specification
type CanaryRun struct {
	Status interface{} `json:"Status,omitempty"`
	Timeline interface{} `json:"Timeline,omitempty"`
	Artifacts3location interface{} `json:"ArtifactS3Location,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// DescribeRuntimeVersionsResponse represents the DescribeRuntimeVersionsResponse schema from the OpenAPI specification
type DescribeRuntimeVersionsResponse struct {
	Runtimeversions interface{} `json:"RuntimeVersions,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// GetCanaryRequest represents the GetCanaryRequest schema from the OpenAPI specification
type GetCanaryRequest struct {
}

// UpdateCanaryRequest represents the UpdateCanaryRequest schema from the OpenAPI specification
type UpdateCanaryRequest struct {
	Visualreference interface{} `json:"VisualReference,omitempty"`
	Vpcconfig interface{} `json:"VpcConfig,omitempty"`
	Artifacts3location interface{} `json:"ArtifactS3Location,omitempty"`
	Schedule interface{} `json:"Schedule,omitempty"`
	Successretentionperiodindays interface{} `json:"SuccessRetentionPeriodInDays,omitempty"`
	Artifactconfig interface{} `json:"ArtifactConfig,omitempty"`
	Code interface{} `json:"Code,omitempty"`
	Executionrolearn interface{} `json:"ExecutionRoleArn,omitempty"`
	Runconfig interface{} `json:"RunConfig,omitempty"`
	Failureretentionperiodindays interface{} `json:"FailureRetentionPeriodInDays,omitempty"`
	Runtimeversion interface{} `json:"RuntimeVersion,omitempty"`
}

// ListGroupResourcesResponse represents the ListGroupResourcesResponse schema from the OpenAPI specification
type ListGroupResourcesResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Resources interface{} `json:"Resources,omitempty"`
}

// BaseScreenshot represents the BaseScreenshot schema from the OpenAPI specification
type BaseScreenshot struct {
	Ignorecoordinates interface{} `json:"IgnoreCoordinates,omitempty"`
	Screenshotname interface{} `json:"ScreenshotName"`
}

// DescribeCanariesLastRunRequest represents the DescribeCanariesLastRunRequest schema from the OpenAPI specification
type DescribeCanariesLastRunRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Names interface{} `json:"Names,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeCanariesLastRunResponse represents the DescribeCanariesLastRunResponse schema from the OpenAPI specification
type DescribeCanariesLastRunResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Canarieslastrun interface{} `json:"CanariesLastRun,omitempty"`
}

// GetGroupResponse represents the GetGroupResponse schema from the OpenAPI specification
type GetGroupResponse struct {
	Group interface{} `json:"Group,omitempty"`
}

// ArtifactConfigOutput represents the ArtifactConfigOutput schema from the OpenAPI specification
type ArtifactConfigOutput struct {
	S3encryption interface{} `json:"S3Encryption,omitempty"`
}

// Group represents the Group schema from the OpenAPI specification
type Group struct {
	Arn interface{} `json:"Arn,omitempty"`
	Createdtime interface{} `json:"CreatedTime,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// ListGroupsResponse represents the ListGroupsResponse schema from the OpenAPI specification
type ListGroupsResponse struct {
	Groups interface{} `json:"Groups,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// TagMap represents the TagMap schema from the OpenAPI specification
type TagMap struct {
}

// AssociateResourceResponse represents the AssociateResourceResponse schema from the OpenAPI specification
type AssociateResourceResponse struct {
}

// CanaryRunConfigInput represents the CanaryRunConfigInput schema from the OpenAPI specification
type CanaryRunConfigInput struct {
	Memoryinmb interface{} `json:"MemoryInMB,omitempty"`
	Timeoutinseconds interface{} `json:"TimeoutInSeconds,omitempty"`
	Activetracing interface{} `json:"ActiveTracing,omitempty"`
	Environmentvariables interface{} `json:"EnvironmentVariables,omitempty"`
}

// AssociateResourceRequest represents the AssociateResourceRequest schema from the OpenAPI specification
type AssociateResourceRequest struct {
	Resourcearn interface{} `json:"ResourceArn"`
}

// UpdateCanaryResponse represents the UpdateCanaryResponse schema from the OpenAPI specification
type UpdateCanaryResponse struct {
}

// DescribeRuntimeVersionsRequest represents the DescribeRuntimeVersionsRequest schema from the OpenAPI specification
type DescribeRuntimeVersionsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// VpcConfigOutput represents the VpcConfigOutput schema from the OpenAPI specification
type VpcConfigOutput struct {
	Subnetids interface{} `json:"SubnetIds,omitempty"`
	Vpcid interface{} `json:"VpcId,omitempty"`
	Securitygroupids interface{} `json:"SecurityGroupIds,omitempty"`
}

// VisualReferenceInput represents the VisualReferenceInput schema from the OpenAPI specification
type VisualReferenceInput struct {
	Basecanaryrunid interface{} `json:"BaseCanaryRunId"`
	Basescreenshots interface{} `json:"BaseScreenshots,omitempty"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
}

// Canary represents the Canary schema from the OpenAPI specification
type Canary struct {
	Vpcconfig VpcConfigOutput `json:"VpcConfig,omitempty"` // If this canary is to test an endpoint in a VPC, this structure contains information about the subnets and security groups of the VPC endpoint. For more information, see <a href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_VPC.html"> Running a Canary in a VPC</a>.
	Artifacts3location interface{} `json:"ArtifactS3Location,omitempty"`
	Runconfig CanaryRunConfigOutput `json:"RunConfig,omitempty"` // A structure that contains information about a canary run.
	Successretentionperiodindays interface{} `json:"SuccessRetentionPeriodInDays,omitempty"`
	Failureretentionperiodindays interface{} `json:"FailureRetentionPeriodInDays,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Code CanaryCodeOutput `json:"Code,omitempty"` // This structure contains information about the canary's Lambda handler and where its code is stored by CloudWatch Synthetics.
	Name interface{} `json:"Name,omitempty"`
	Artifactconfig interface{} `json:"ArtifactConfig,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Executionrolearn interface{} `json:"ExecutionRoleArn,omitempty"`
	Schedule interface{} `json:"Schedule,omitempty"`
	Enginearn interface{} `json:"EngineArn,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Timeline interface{} `json:"Timeline,omitempty"`
	Visualreference interface{} `json:"VisualReference,omitempty"`
	Runtimeversion interface{} `json:"RuntimeVersion,omitempty"`
}

// VpcConfigInput represents the VpcConfigInput schema from the OpenAPI specification
type VpcConfigInput struct {
	Securitygroupids interface{} `json:"SecurityGroupIds,omitempty"`
	Subnetids interface{} `json:"SubnetIds,omitempty"`
}

// StartCanaryRequest represents the StartCanaryRequest schema from the OpenAPI specification
type StartCanaryRequest struct {
}

// ListGroupResourcesRequest represents the ListGroupResourcesRequest schema from the OpenAPI specification
type ListGroupResourcesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CanaryRunStatus represents the CanaryRunStatus schema from the OpenAPI specification
type CanaryRunStatus struct {
	State interface{} `json:"State,omitempty"`
	Statereason interface{} `json:"StateReason,omitempty"`
	Statereasoncode interface{} `json:"StateReasonCode,omitempty"`
}

// ListAssociatedGroupsRequest represents the ListAssociatedGroupsRequest schema from the OpenAPI specification
type ListAssociatedGroupsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListAssociatedGroupsResponse represents the ListAssociatedGroupsResponse schema from the OpenAPI specification
type ListAssociatedGroupsResponse struct {
	Groups interface{} `json:"Groups,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CanaryRunConfigOutput represents the CanaryRunConfigOutput schema from the OpenAPI specification
type CanaryRunConfigOutput struct {
	Activetracing interface{} `json:"ActiveTracing,omitempty"`
	Memoryinmb interface{} `json:"MemoryInMB,omitempty"`
	Timeoutinseconds interface{} `json:"TimeoutInSeconds,omitempty"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Tags interface{} `json:"Tags"`
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"Tags,omitempty"`
}

// GetCanaryResponse represents the GetCanaryResponse schema from the OpenAPI specification
type GetCanaryResponse struct {
	Canary interface{} `json:"Canary,omitempty"`
}

// CreateCanaryResponse represents the CreateCanaryResponse schema from the OpenAPI specification
type CreateCanaryResponse struct {
	Canary interface{} `json:"Canary,omitempty"`
}

// CreateCanaryRequest represents the CreateCanaryRequest schema from the OpenAPI specification
type CreateCanaryRequest struct {
	Code interface{} `json:"Code"`
	Artifactconfig interface{} `json:"ArtifactConfig,omitempty"`
	Executionrolearn interface{} `json:"ExecutionRoleArn"`
	Runconfig interface{} `json:"RunConfig,omitempty"`
	Vpcconfig interface{} `json:"VpcConfig,omitempty"`
	Name interface{} `json:"Name"`
	Runtimeversion interface{} `json:"RuntimeVersion"`
	Schedule interface{} `json:"Schedule"`
	Tags interface{} `json:"Tags,omitempty"`
	Failureretentionperiodindays interface{} `json:"FailureRetentionPeriodInDays,omitempty"`
	Successretentionperiodindays interface{} `json:"SuccessRetentionPeriodInDays,omitempty"`
	Artifacts3location interface{} `json:"ArtifactS3Location"`
}

// StopCanaryRequest represents the StopCanaryRequest schema from the OpenAPI specification
type StopCanaryRequest struct {
}

// CanaryTimeline represents the CanaryTimeline schema from the OpenAPI specification
type CanaryTimeline struct {
	Created interface{} `json:"Created,omitempty"`
	Lastmodified interface{} `json:"LastModified,omitempty"`
	Laststarted interface{} `json:"LastStarted,omitempty"`
	Laststopped interface{} `json:"LastStopped,omitempty"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
}

// DeleteGroupResponse represents the DeleteGroupResponse schema from the OpenAPI specification
type DeleteGroupResponse struct {
}

// GetCanaryRunsResponse represents the GetCanaryRunsResponse schema from the OpenAPI specification
type GetCanaryRunsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Canaryruns interface{} `json:"CanaryRuns,omitempty"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// CanaryStatus represents the CanaryStatus schema from the OpenAPI specification
type CanaryStatus struct {
	Statereasoncode interface{} `json:"StateReasonCode,omitempty"`
	State interface{} `json:"State,omitempty"`
	Statereason interface{} `json:"StateReason,omitempty"`
}

// DisassociateResourceResponse represents the DisassociateResourceResponse schema from the OpenAPI specification
type DisassociateResourceResponse struct {
}

// CanaryRunTimeline represents the CanaryRunTimeline schema from the OpenAPI specification
type CanaryRunTimeline struct {
	Completed interface{} `json:"Completed,omitempty"`
	Started interface{} `json:"Started,omitempty"`
}

// DeleteCanaryRequest represents the DeleteCanaryRequest schema from the OpenAPI specification
type DeleteCanaryRequest struct {
}
