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

func StopcanaryHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		nameVal, ok := args["name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: name"), nil
		}
		name, ok := nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: name"), nil
		}
		url := fmt.Sprintf("%s/canary/%s/stop", cfg.BaseURL, name)
		req, err := http.NewRequest("POST", url, nil)
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
		var result models.StopCanaryResponse
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

func CreateStopcanaryTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_canary_name_stop",
		mcp.WithDescription("<p>Stops the canary to prevent all future runs. If the canary is currently running,the run that is in progress completes on its own, publishes metrics, and uploads artifacts, but it is not recorded in Synthetics as a completed run.</p> <p>You can use <code>StartCanary</code> to start it running again with the canary’s current schedule at any point in the future. </p>"),
		mcp.WithString("name", mcp.Required(), mcp.Description("The name of the canary that you want to stop. To find the names of your canaries, use <a href=\"https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_DescribeCanaries.html\">ListCanaries</a>.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    StopcanaryHandler(cfg),
	}
}
