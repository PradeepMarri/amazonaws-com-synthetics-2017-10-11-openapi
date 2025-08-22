package main

import (
	"github.com/synthetics/mcp-server/config"
	"github.com/synthetics/mcp-server/models"
	tools_canary "github.com/synthetics/mcp-server/tools/canary"
	tools_group "github.com/synthetics/mcp-server/tools/group"
	tools_tags "github.com/synthetics/mcp-server/tools/tags"
	tools_resource "github.com/synthetics/mcp-server/tools/resource"
	tools_canaries "github.com/synthetics/mcp-server/tools/canaries"
	tools_groups "github.com/synthetics/mcp-server/tools/groups"
	tools_runtime_versions "github.com/synthetics/mcp-server/tools/runtime_versions"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_canary.CreateDeletecanaryTool(cfg),
		tools_canary.CreateGetcanaryTool(cfg),
		tools_canary.CreateUpdatecanaryTool(cfg),
		tools_canary.CreateGetcanaryrunsTool(cfg),
		tools_canary.CreateStopcanaryTool(cfg),
		tools_group.CreateGetgroupTool(cfg),
		tools_group.CreateDeletegroupTool(cfg),
		tools_tags.CreateListtagsforresourceTool(cfg),
		tools_tags.CreateTagresourceTool(cfg),
		tools_tags.CreateUntagresourceTool(cfg),
		tools_group.CreateDisassociateresourceTool(cfg),
		tools_group.CreateListgroupresourcesTool(cfg),
		tools_resource.CreateListassociatedgroupsTool(cfg),
		tools_group.CreateAssociateresourceTool(cfg),
		tools_canaries.CreateDescribecanariesTool(cfg),
		tools_canary.CreateCreatecanaryTool(cfg),
		tools_canary.CreateStartcanaryTool(cfg),
		tools_groups.CreateListgroupsTool(cfg),
		tools_canaries.CreateDescribecanarieslastrunTool(cfg),
		tools_group.CreateCreategroupTool(cfg),
		tools_runtime_versions.CreateDescriberuntimeversionsTool(cfg),
	}
}
