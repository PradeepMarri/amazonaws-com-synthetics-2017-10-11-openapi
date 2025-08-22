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

func TagresourceHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		resourceArnVal, ok := args["resourceArn"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: resourceArn"), nil
		}
		resourceArn, ok := resourceArnVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: resourceArn"), nil
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
		url := fmt.Sprintf("%s/tags/%s", cfg.BaseURL, resourceArn)
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
		var result models.TagResourceResponse
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

func CreateTagresourceTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_tags_resourceArn",
		mcp.WithDescription("<p>Assigns one or more tags (key-value pairs) to the specified canary or group. </p> <p>Tags can help you organize and categorize your resources. You can also use them to scope user permissions, by granting a user permission to access or change only resources with certain tag values.</p> <p>Tags don't have any semantic meaning to Amazon Web Services and are interpreted strictly as strings of characters.</p> <p>You can use the <code>TagResource</code> action with a resource that already has tags. If you specify a new tag key for the resource, this tag is appended to the list of tags associated with the resource. If you specify a tag key that is already associated with the resource, the new tag value that you specify replaces the previous value for that tag.</p> <p>You can associate as many as 50 tags with a canary or group.</p>"),
		mcp.WithString("resourceArn", mcp.Required(), mcp.Description("<p>The ARN of the canary or group that you're adding tags to.</p> <p>The ARN format of a canary is <code>arn:aws:synthetics:<i>Region</i>:<i>account-id</i>:canary:<i>canary-name</i> </code>.</p> <p>The ARN format of a group is <code>arn:aws:synthetics:<i>Region</i>:<i>account-id</i>:group:<i>group-name</i> </code> </p>")),
		mcp.WithObject("Tags", mcp.Required(), mcp.Description("Input parameter: The list of key-value pairs to associate with the resource.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    TagresourceHandler(cfg),
	}
}
