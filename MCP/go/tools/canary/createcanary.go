package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/synthetics/mcp-server/config"
	"github.com/synthetics/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CreatecanaryHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/canary", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.CreateCanaryResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateCreatecanaryTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_canary",
		mcp.WithDescription("<p>Creates a canary. Canaries are scripts that monitor your endpoints and APIs from the outside-in. Canaries help you check the availability and latency of your web services and troubleshoot anomalies by investigating load time data, screenshots of the UI, logs, and metrics. You can set up a canary to run continuously or just once. </p> <p>Do not use <code>CreateCanary</code> to modify an existing canary. Use <a href="https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_UpdateCanary.html">UpdateCanary</a> instead.</p> <p>To create canaries, you must have the <code>CloudWatchSyntheticsFullAccess</code> policy. If you are creating a new IAM role for the canary, you also need the <code>iam:CreateRole</code>, <code>iam:CreatePolicy</code> and <code>iam:AttachRolePolicy</code> permissions. For more information, see <a href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Roles">Necessary Roles and Permissions</a>.</p> <p>Do not include secrets or proprietary information in your canary names. The canary name makes up part of the Amazon Resource Name (ARN) for the canary, and the ARN is included in outbound calls over the internet. For more information, see <a href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/servicelens_canaries_security.html">Security Considerations for Synthetics Canaries</a>.</p>"),
		mcp.WithString("Name", mcp.Required(), mcp.Description("Input parameter: <p>The name for this canary. Be sure to give it a descriptive name that distinguishes it from other canaries in your account.</p> <p>Do not include secrets or proprietary information in your canary names. The canary name makes up part of the canary ARN, and the ARN is included in outbound calls over the internet. For more information, see <a href=\"https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/servicelens_canaries_security.html\">Security Considerations for Synthetics Canaries</a>.</p>")),
		mcp.WithObject("Schedule", mcp.Required(), mcp.Description("Input parameter: This structure specifies how often a canary is to make runs and the date and time when it should stop making runs.")),
		mcp.WithNumber("SuccessRetentionPeriodInDays", mcp.Description("Input parameter: The number of days to retain data about successful runs of this canary. If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.")),
		mcp.WithObject("VpcConfig", mcp.Description("Input parameter: If this canary is to test an endpoint in a VPC, this structure contains information about the subnets and security groups of the VPC endpoint. For more information, see <a href=\"https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_VPC.html\"> Running a Canary in a VPC</a>.")),
		mcp.WithObject("ArtifactConfig", mcp.Description("Input parameter: A structure that contains the configuration for canary artifacts, including the encryption-at-rest settings for artifacts that the canary uploads to Amazon S3.")),
		mcp.WithString("ArtifactS3Location", mcp.Required(), mcp.Description("Input parameter: The location in Amazon S3 where Synthetics stores artifacts from the test runs of this canary. Artifacts include the log file, screenshots, and HAR files. The name of the S3 bucket can't include a period (.).")),
		mcp.WithNumber("FailureRetentionPeriodInDays", mcp.Description("Input parameter: The number of days to retain data about failed runs of this canary. If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.")),
		mcp.WithObject("RunConfig", mcp.Description("Input parameter: A structure that contains input information for a canary run.")),
		mcp.WithObject("Tags", mcp.Description("Input parameter: <p>A list of key-value pairs to associate with the canary. You can associate as many as 50 tags with a canary.</p> <p>Tags can help you organize and categorize your resources. You can also use them to scope user permissions, by granting a user permission to access or change only the resources that have certain tag values.</p>")),
		mcp.WithString("RuntimeVersion", mcp.Required(), mcp.Description("Input parameter: Specifies the runtime version to use for the canary. For a list of valid runtime versions and more information about runtime versions, see <a href=\"https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html\"> Canary Runtime Versions</a>.")),
		mcp.WithObject("Code", mcp.Required(), mcp.Description("Input parameter: Use this structure to input your script code for the canary. This structure contains the Lambda handler with the location where the canary should start running the script. If the script is stored in an S3 bucket, the bucket name, key, and version are also included. If the script was passed into the canary directly, the script code is contained in the value of <code>Zipfile</code>. ")),
		mcp.WithString("ExecutionRoleArn", mcp.Required(), mcp.Description("Input parameter: <p>The ARN of the IAM role to be used to run the canary. This role must already exist, and must include <code>lambda.amazonaws.com</code> as a principal in the trust policy. The role must also have the following permissions:</p> <ul> <li> <p> <code>s3:PutObject</code> </p> </li> <li> <p> <code>s3:GetBucketLocation</code> </p> </li> <li> <p> <code>s3:ListAllMyBuckets</code> </p> </li> <li> <p> <code>cloudwatch:PutMetricData</code> </p> </li> <li> <p> <code>logs:CreateLogGroup</code> </p> </li> <li> <p> <code>logs:CreateLogStream</code> </p> </li> <li> <p> <code>logs:PutLogEvents</code> </p> </li> </ul>")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreatecanaryHandler(cfg),
	}
}
