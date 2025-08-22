package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/synthetics/mcp-server/config"
	"github.com/synthetics/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetgroupHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/group/%s", cfg.BaseURL, groupIdentifier)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.GetGroupResponse
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

func CreateGetgroupTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_group_groupIdentifier",
		mcp.WithDescription("Returns information about one group. Groups are a global resource, so you can use this operation from any Region."),
		mcp.WithString("groupIdentifier", mcp.Required(), mcp.Description("Specifies the group to return information for. You can specify the group name, the ARN, or the group ID as the <code>GroupIdentifier</code>.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetgroupHandler(cfg),
	}
}
