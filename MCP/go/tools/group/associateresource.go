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

func AssociateresourceHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		groupIdentifierVal, ok := args["groupIdentifier"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: groupIdentifier"), nil
		}
		groupIdentifier, ok := groupIdentifierVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: groupIdentifier"), nil
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
		url := fmt.Sprintf("%s/group/%s/associate", cfg.BaseURL, groupIdentifier)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
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
		var result models.AssociateResourceResponse
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

func CreateAssociateresourceTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_group_groupIdentifier_associate",
		mcp.WithDescription("<p>Associates a canary with a group. Using groups can help you with managing and automating your canaries, and you can also view aggregated run results and statistics for all canaries in a group. </p> <p>You must run this operation in the Region where the canary exists.</p>"),
		mcp.WithString("groupIdentifier", mcp.Required(), mcp.Description("Specifies the group. You can specify the group name, the ARN, or the group ID as the <code>GroupIdentifier</code>.")),
		mcp.WithString("ResourceArn", mcp.Required(), mcp.Description("Input parameter: The ARN of the canary that you want to associate with the specified group.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    AssociateresourceHandler(cfg),
	}
}
