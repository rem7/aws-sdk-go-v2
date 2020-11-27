// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpAddPermission struct {
}

func (*validateOpAddPermission) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAddPermission) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AddPermissionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAddPermissionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCheckIfPhoneNumberIsOptedOut struct {
}

func (*validateOpCheckIfPhoneNumberIsOptedOut) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCheckIfPhoneNumberIsOptedOut) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CheckIfPhoneNumberIsOptedOutInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCheckIfPhoneNumberIsOptedOutInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpConfirmSubscription struct {
}

func (*validateOpConfirmSubscription) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpConfirmSubscription) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ConfirmSubscriptionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpConfirmSubscriptionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreatePlatformApplication struct {
}

func (*validateOpCreatePlatformApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreatePlatformApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreatePlatformApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreatePlatformApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreatePlatformEndpoint struct {
}

func (*validateOpCreatePlatformEndpoint) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreatePlatformEndpoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreatePlatformEndpointInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreatePlatformEndpointInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateTopic struct {
}

func (*validateOpCreateTopic) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateTopic) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateTopicInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateTopicInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteEndpoint struct {
}

func (*validateOpDeleteEndpoint) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteEndpoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteEndpointInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteEndpointInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeletePlatformApplication struct {
}

func (*validateOpDeletePlatformApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeletePlatformApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeletePlatformApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeletePlatformApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteTopic struct {
}

func (*validateOpDeleteTopic) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteTopic) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteTopicInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteTopicInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetEndpointAttributes struct {
}

func (*validateOpGetEndpointAttributes) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetEndpointAttributes) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetEndpointAttributesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetEndpointAttributesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetPlatformApplicationAttributes struct {
}

func (*validateOpGetPlatformApplicationAttributes) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetPlatformApplicationAttributes) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetPlatformApplicationAttributesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetPlatformApplicationAttributesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSubscriptionAttributes struct {
}

func (*validateOpGetSubscriptionAttributes) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSubscriptionAttributes) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSubscriptionAttributesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSubscriptionAttributesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetTopicAttributes struct {
}

func (*validateOpGetTopicAttributes) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetTopicAttributes) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetTopicAttributesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetTopicAttributesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListEndpointsByPlatformApplication struct {
}

func (*validateOpListEndpointsByPlatformApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListEndpointsByPlatformApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListEndpointsByPlatformApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListEndpointsByPlatformApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListSubscriptionsByTopic struct {
}

func (*validateOpListSubscriptionsByTopic) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListSubscriptionsByTopic) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListSubscriptionsByTopicInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListSubscriptionsByTopicInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpOptInPhoneNumber struct {
}

func (*validateOpOptInPhoneNumber) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpOptInPhoneNumber) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*OptInPhoneNumberInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpOptInPhoneNumberInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPublish struct {
}

func (*validateOpPublish) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPublish) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PublishInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPublishInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRemovePermission struct {
}

func (*validateOpRemovePermission) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRemovePermission) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RemovePermissionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRemovePermissionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSetEndpointAttributes struct {
}

func (*validateOpSetEndpointAttributes) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSetEndpointAttributes) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SetEndpointAttributesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSetEndpointAttributesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSetPlatformApplicationAttributes struct {
}

func (*validateOpSetPlatformApplicationAttributes) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSetPlatformApplicationAttributes) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SetPlatformApplicationAttributesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSetPlatformApplicationAttributesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSetSMSAttributes struct {
}

func (*validateOpSetSMSAttributes) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSetSMSAttributes) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SetSMSAttributesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSetSMSAttributesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSetSubscriptionAttributes struct {
}

func (*validateOpSetSubscriptionAttributes) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSetSubscriptionAttributes) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SetSubscriptionAttributesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSetSubscriptionAttributesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSetTopicAttributes struct {
}

func (*validateOpSetTopicAttributes) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSetTopicAttributes) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SetTopicAttributesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSetTopicAttributesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSubscribe struct {
}

func (*validateOpSubscribe) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSubscribe) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SubscribeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSubscribeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUnsubscribe struct {
}

func (*validateOpUnsubscribe) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUnsubscribe) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UnsubscribeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUnsubscribeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAddPermissionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAddPermission{}, middleware.After)
}

func addOpCheckIfPhoneNumberIsOptedOutValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCheckIfPhoneNumberIsOptedOut{}, middleware.After)
}

func addOpConfirmSubscriptionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpConfirmSubscription{}, middleware.After)
}

func addOpCreatePlatformApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreatePlatformApplication{}, middleware.After)
}

func addOpCreatePlatformEndpointValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreatePlatformEndpoint{}, middleware.After)
}

func addOpCreateTopicValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateTopic{}, middleware.After)
}

func addOpDeleteEndpointValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteEndpoint{}, middleware.After)
}

func addOpDeletePlatformApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeletePlatformApplication{}, middleware.After)
}

func addOpDeleteTopicValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteTopic{}, middleware.After)
}

func addOpGetEndpointAttributesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetEndpointAttributes{}, middleware.After)
}

func addOpGetPlatformApplicationAttributesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetPlatformApplicationAttributes{}, middleware.After)
}

func addOpGetSubscriptionAttributesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSubscriptionAttributes{}, middleware.After)
}

func addOpGetTopicAttributesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetTopicAttributes{}, middleware.After)
}

func addOpListEndpointsByPlatformApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListEndpointsByPlatformApplication{}, middleware.After)
}

func addOpListSubscriptionsByTopicValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListSubscriptionsByTopic{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpOptInPhoneNumberValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpOptInPhoneNumber{}, middleware.After)
}

func addOpPublishValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPublish{}, middleware.After)
}

func addOpRemovePermissionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRemovePermission{}, middleware.After)
}

func addOpSetEndpointAttributesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSetEndpointAttributes{}, middleware.After)
}

func addOpSetPlatformApplicationAttributesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSetPlatformApplicationAttributes{}, middleware.After)
}

func addOpSetSMSAttributesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSetSMSAttributes{}, middleware.After)
}

func addOpSetSubscriptionAttributesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSetSubscriptionAttributes{}, middleware.After)
}

func addOpSetTopicAttributesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSetTopicAttributes{}, middleware.After)
}

func addOpSubscribeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSubscribe{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUnsubscribeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUnsubscribe{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func validateMessageAttributeMap(v map[string]types.MessageAttributeValue) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MessageAttributeMap"}
	for key := range v {
		value := v[key]
		if err := validateMessageAttributeValue(&value); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%q]", key), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMessageAttributeValue(v *types.MessageAttributeValue) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MessageAttributeValue"}
	if v.DataType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTag(v *types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tag"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagList"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAddPermissionInput(v *AddPermissionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AddPermissionInput"}
	if v.TopicArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TopicArn"))
	}
	if v.Label == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Label"))
	}
	if v.AWSAccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AWSAccountId"))
	}
	if v.ActionName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ActionName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCheckIfPhoneNumberIsOptedOutInput(v *CheckIfPhoneNumberIsOptedOutInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CheckIfPhoneNumberIsOptedOutInput"}
	if v.PhoneNumber == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PhoneNumber"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpConfirmSubscriptionInput(v *ConfirmSubscriptionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ConfirmSubscriptionInput"}
	if v.TopicArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TopicArn"))
	}
	if v.Token == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Token"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreatePlatformApplicationInput(v *CreatePlatformApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreatePlatformApplicationInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Platform == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Platform"))
	}
	if v.Attributes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Attributes"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreatePlatformEndpointInput(v *CreatePlatformEndpointInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreatePlatformEndpointInput"}
	if v.PlatformApplicationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PlatformApplicationArn"))
	}
	if v.Token == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Token"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateTopicInput(v *CreateTopicInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateTopicInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteEndpointInput(v *DeleteEndpointInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteEndpointInput"}
	if v.EndpointArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndpointArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeletePlatformApplicationInput(v *DeletePlatformApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeletePlatformApplicationInput"}
	if v.PlatformApplicationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PlatformApplicationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteTopicInput(v *DeleteTopicInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteTopicInput"}
	if v.TopicArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TopicArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetEndpointAttributesInput(v *GetEndpointAttributesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetEndpointAttributesInput"}
	if v.EndpointArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndpointArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetPlatformApplicationAttributesInput(v *GetPlatformApplicationAttributesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetPlatformApplicationAttributesInput"}
	if v.PlatformApplicationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PlatformApplicationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSubscriptionAttributesInput(v *GetSubscriptionAttributesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSubscriptionAttributesInput"}
	if v.SubscriptionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubscriptionArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetTopicAttributesInput(v *GetTopicAttributesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetTopicAttributesInput"}
	if v.TopicArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TopicArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListEndpointsByPlatformApplicationInput(v *ListEndpointsByPlatformApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListEndpointsByPlatformApplicationInput"}
	if v.PlatformApplicationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PlatformApplicationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListSubscriptionsByTopicInput(v *ListSubscriptionsByTopicInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListSubscriptionsByTopicInput"}
	if v.TopicArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TopicArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpOptInPhoneNumberInput(v *OptInPhoneNumberInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "OptInPhoneNumberInput"}
	if v.PhoneNumber == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PhoneNumber"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPublishInput(v *PublishInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PublishInput"}
	if v.Message == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Message"))
	}
	if v.MessageAttributes != nil {
		if err := validateMessageAttributeMap(v.MessageAttributes); err != nil {
			invalidParams.AddNested("MessageAttributes", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRemovePermissionInput(v *RemovePermissionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RemovePermissionInput"}
	if v.TopicArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TopicArn"))
	}
	if v.Label == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Label"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSetEndpointAttributesInput(v *SetEndpointAttributesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SetEndpointAttributesInput"}
	if v.EndpointArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndpointArn"))
	}
	if v.Attributes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Attributes"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSetPlatformApplicationAttributesInput(v *SetPlatformApplicationAttributesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SetPlatformApplicationAttributesInput"}
	if v.PlatformApplicationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PlatformApplicationArn"))
	}
	if v.Attributes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Attributes"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSetSMSAttributesInput(v *SetSMSAttributesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SetSMSAttributesInput"}
	if v.Attributes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Attributes"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSetSubscriptionAttributesInput(v *SetSubscriptionAttributesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SetSubscriptionAttributesInput"}
	if v.SubscriptionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubscriptionArn"))
	}
	if v.AttributeName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AttributeName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSetTopicAttributesInput(v *SetTopicAttributesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SetTopicAttributesInput"}
	if v.TopicArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TopicArn"))
	}
	if v.AttributeName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AttributeName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSubscribeInput(v *SubscribeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SubscribeInput"}
	if v.TopicArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TopicArn"))
	}
	if v.Protocol == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Protocol"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	} else if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUnsubscribeInput(v *UnsubscribeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UnsubscribeInput"}
	if v.SubscriptionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubscriptionArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
