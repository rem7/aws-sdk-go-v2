// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about a patch baseline.
func (c *Client) GetPatchBaseline(ctx context.Context, params *GetPatchBaselineInput, optFns ...func(*Options)) (*GetPatchBaselineOutput, error) {
	if params == nil {
		params = &GetPatchBaselineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPatchBaseline", params, optFns, c.addOperationGetPatchBaselineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPatchBaselineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPatchBaselineInput struct {

	// The ID of the patch baseline to retrieve.
	//
	// This member is required.
	BaselineId *string
}

type GetPatchBaselineOutput struct {

	// A set of rules used to include patches in the baseline.
	ApprovalRules *types.PatchRuleGroup

	// A list of explicitly approved patches for the baseline.
	ApprovedPatches []string

	// Returns the specified compliance severity level for approved patches in the
	// patch baseline.
	ApprovedPatchesComplianceLevel types.PatchComplianceLevel

	// Indicates whether the list of approved patches includes non-security updates
	// that should be applied to the instances. The default value is false. Applies to
	// Linux instances only.
	ApprovedPatchesEnableNonSecurity bool

	// The ID of the retrieved patch baseline.
	BaselineId *string

	// The date the patch baseline was created.
	CreatedDate *time.Time

	// A description of the patch baseline.
	Description *string

	// A set of global filters used to exclude patches from the baseline.
	GlobalFilters *types.PatchFilterGroup

	// The date the patch baseline was last modified.
	ModifiedDate *time.Time

	// The name of the patch baseline.
	Name *string

	// Returns the operating system specified for the patch baseline.
	OperatingSystem types.OperatingSystem

	// Patch groups included in the patch baseline.
	PatchGroups []string

	// A list of explicitly rejected patches for the baseline.
	RejectedPatches []string

	// The action specified to take on patches included in the RejectedPatches list. A
	// patch can be allowed only if it is a dependency of another package, or blocked
	// entirely along with packages that include it as a dependency.
	RejectedPatchesAction types.PatchAction

	// Information about the patches to use to update the instances, including target
	// operating systems and source repositories. Applies to Linux instances only.
	Sources []types.PatchSource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationGetPatchBaselineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetPatchBaseline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPatchBaseline{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetPatchBaselineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPatchBaseline(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetPatchBaseline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetPatchBaseline",
	}
}
