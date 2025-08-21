package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/kubernetes/mcp-server/config"
	"github.com/kubernetes/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Deleteflowcontrolapiserverv1beta3collectionflowschemaHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["continue"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("continue=%v", val))
		}
		if val, ok := args["dryRun"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dryRun=%v", val))
		}
		if val, ok := args["fieldSelector"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fieldSelector=%v", val))
		}
		if val, ok := args["gracePeriodSeconds"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("gracePeriodSeconds=%v", val))
		}
		if val, ok := args["labelSelector"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("labelSelector=%v", val))
		}
		if val, ok := args["limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit=%v", val))
		}
		if val, ok := args["orphanDependents"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("orphanDependents=%v", val))
		}
		if val, ok := args["propagationPolicy"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("propagationPolicy=%v", val))
		}
		if val, ok := args["resourceVersion"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("resourceVersion=%v", val))
		}
		if val, ok := args["resourceVersionMatch"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("resourceVersionMatch=%v", val))
		}
		if val, ok := args["sendInitialEvents"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sendInitialEvents=%v", val))
		}
		if val, ok := args["timeoutSeconds"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("timeoutSeconds=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/apis/flowcontrol.apiserver.k8s.io/v1beta3/flowschemas%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
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
		var result map[string]interface{}
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

func CreateDeleteflowcontrolapiserverv1beta3collectionflowschemaTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_apis_flowcontrol.apiserver.k8s.io_v1beta3_flowschemas",
		mcp.WithDescription("delete collection of FlowSchema"),
		mcp.WithObject("body", mcp.Description("")),
		mcp.WithString("continue", mcp.Description("The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the \"next key\".\n\nThis field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.")),
		mcp.WithString("dryRun", mcp.Description("When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed")),
		mcp.WithString("fieldSelector", mcp.Description("A selector to restrict the list of returned objects by their fields. Defaults to everything.")),
		mcp.WithString("gracePeriodSeconds", mcp.Description("The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.")),
		mcp.WithString("labelSelector", mcp.Description("A selector to restrict the list of returned objects by their labels. Defaults to everything.")),
		mcp.WithString("limit", mcp.Description("limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.\n\nThe server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.")),
		mcp.WithString("orphanDependents", mcp.Description("Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \"orphan\" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both.")),
		mcp.WithString("propagationPolicy", mcp.Description("Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground.")),
		mcp.WithString("resourceVersion", mcp.Description("resourceVersion sets a constraint on what resource versions a request may be served from. See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.\n\nDefaults to unset")),
		mcp.WithString("resourceVersionMatch", mcp.Description("resourceVersionMatch determines how resourceVersion is applied to list calls. It is highly recommended that resourceVersionMatch be set for list calls where resourceVersion is set See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.\n\nDefaults to unset")),
		mcp.WithString("sendInitialEvents", mcp.Description("`sendInitialEvents=true` may be set together with `watch=true`. In that case, the watch stream will begin with synthetic events to produce the current state of objects in the collection. Once all such events have been sent, a synthetic \"Bookmark\" event  will be sent. The bookmark will report the ResourceVersion (RV) corresponding to the set of objects, and be marked with `\"k8s.io/initial-events-end\": \"true\"` annotation. Afterwards, the watch stream will proceed as usual, sending watch events corresponding to changes (subsequent to the RV) to objects watched.\n\nWhen `sendInitialEvents` option is set, we require `resourceVersionMatch` option to also be set. The semantic of the watch request is as following: - `resourceVersionMatch` = NotOlderThan\n  is interpreted as \"data at least as new as the provided `resourceVersion`\"\n  and the bookmark event is send when the state is synced\n  to a `resourceVersion` at least as fresh as the one provided by the ListOptions.\n  If `resourceVersion` is unset, this is interpreted as \"consistent read\" and the\n  bookmark event is send when the state is synced at least to the moment\n  when request started being processed.\n- `resourceVersionMatch` set to any other value or unset\n  Invalid error is returned.\n\nDefaults to true if `resourceVersion=\"\"` or `resourceVersion=\"0\"` (for backward compatibility reasons) and to false otherwise.")),
		mcp.WithString("timeoutSeconds", mcp.Description("Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Deleteflowcontrolapiserverv1beta3collectionflowschemaHandler(cfg),
	}
}
