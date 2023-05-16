// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package platform_application

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrequeue "github.com/aws-controllers-k8s/runtime/pkg/requeue"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/sns"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/sns-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.SNS{}
	_ = &svcapitypes.PlatformApplication{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
	_ = fmt.Sprintf("")
	_ = &ackrequeue.NoRequeue{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer func() {
		exit(err)
	}()
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromGetAttributesInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newGetAttributesRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.GetPlatformApplicationAttributesOutput
	resp, err = rm.sdkapi.GetPlatformApplicationAttributesWithContext(ctx, input)
	rm.metrics.RecordAPICall("GET_ATTRIBUTES", "GetPlatformApplicationAttributes", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "NotFound" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	ko.Spec.EventDeliveryFailure = resp.Attributes["EventDeliveryFailure"]
	ko.Spec.EventEndpointCreated = resp.Attributes["EventEndpointCreated"]
	ko.Spec.EventEndpointDeleted = resp.Attributes["EventEndpointDeleted"]
	ko.Spec.EventEndpointUpdated = resp.Attributes["EventEndpointUpdated"]
	ko.Spec.FailureFeedbackRoleARN = resp.Attributes["FailureFeedbackRoleArn"]
	ko.Spec.PlatformCredential = resp.Attributes["PlatformCredential"]
	ko.Spec.PlatformPrincipal = resp.Attributes["PlatformPrincipal"]
	ko.Spec.SuccessFeedbackRoleARN = resp.Attributes["SuccessFeedbackRoleArn"]
	ko.Spec.SuccessFeedbackSampleRate = resp.Attributes["SuccessFeedbackSampleRate"]

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromGetAtttributesInput returns true if there are any
// fields for the GetAttributes Input shape that are required by not present in
// the resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromGetAttributesInput(
	r *resource,
) bool {
	return (r.ko.Status.ACKResourceMetadata == nil || r.ko.Status.ACKResourceMetadata.ARN == nil)

}

// newGetAttributesRequestPayload returns SDK-specific struct for the HTTP
// request payload of the GetAttributes API call for the resource
func (rm *resourceManager) newGetAttributesRequestPayload(
	r *resource,
) (*svcsdk.GetPlatformApplicationAttributesInput, error) {
	res := &svcsdk.GetPlatformApplicationAttributesInput{}

	if r.ko.Status.ACKResourceMetadata != nil && r.ko.Status.ACKResourceMetadata.ARN != nil {
		res.SetPlatformApplicationArn(string(*r.ko.Status.ACKResourceMetadata.ARN))
	} else {
		res.SetPlatformApplicationArn(rm.ARNFromName(*r.ko.Spec.Name))
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer func() {
		exit(err)
	}()
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreatePlatformApplicationOutput
	_ = resp
	resp, err = rm.sdkapi.CreatePlatformApplicationWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreatePlatformApplication", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.PlatformApplicationArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.PlatformApplicationArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreatePlatformApplicationInput, error) {
	res := &svcsdk.CreatePlatformApplicationInput{}

	attrMap := map[string]*string{}
	if r.ko.Spec.EventDeliveryFailure != nil {
		attrMap["EventDeliveryFailure"] = r.ko.Spec.EventDeliveryFailure
	}
	if r.ko.Spec.EventEndpointCreated != nil {
		attrMap["EventEndpointCreated"] = r.ko.Spec.EventEndpointCreated
	}
	if r.ko.Spec.EventEndpointDeleted != nil {
		attrMap["EventEndpointDeleted"] = r.ko.Spec.EventEndpointDeleted
	}
	if r.ko.Spec.EventEndpointUpdated != nil {
		attrMap["EventEndpointUpdated"] = r.ko.Spec.EventEndpointUpdated
	}
	if r.ko.Spec.FailureFeedbackRoleARN != nil {
		attrMap["FailureFeedbackRoleArn"] = r.ko.Spec.FailureFeedbackRoleARN
	}
	if r.ko.Spec.PlatformCredential != nil {
		attrMap["PlatformCredential"] = r.ko.Spec.PlatformCredential
	}
	if r.ko.Spec.PlatformPrincipal != nil {
		attrMap["PlatformPrincipal"] = r.ko.Spec.PlatformPrincipal
	}
	if r.ko.Spec.SuccessFeedbackRoleARN != nil {
		attrMap["SuccessFeedbackRoleArn"] = r.ko.Spec.SuccessFeedbackRoleARN
	}
	if r.ko.Spec.SuccessFeedbackSampleRate != nil {
		attrMap["SuccessFeedbackSampleRate"] = r.ko.Spec.SuccessFeedbackSampleRate
	}
	if len(attrMap) > 0 {
		res.SetAttributes(attrMap)
	}
	if r.ko.Spec.Name != nil {
		res.SetName(*r.ko.Spec.Name)
	}
	if r.ko.Spec.Platform != nil {
		res.SetPlatform(*r.ko.Spec.Platform)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {
	var err error
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkUpdate")
	defer func() {
		exit(err)
	}()
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. And sdkUpdate should never be called if this is the
	// case, and it's an error in the generated code if it is...
	if rm.requiredFieldsMissingFromSetAttributesInput(desired) {
		panic("Required field in SetAttributes input shape missing!")
	}

	input, err := rm.newSetAttributesRequestPayload(desired)
	if err != nil {
		return nil, err
	}

	// NOTE(jaypipes): SetAttributes calls return a response but they don't
	// contain any useful information. Instead, below, we'll be returning a
	// DeepCopy of the supplied desired state, which should be fine because
	// that desired state has been constructed from a call to GetAttributes...
	_, respErr := rm.sdkapi.SetPlatformApplicationAttributesWithContext(ctx, input)
	rm.metrics.RecordAPICall("SET_ATTRIBUTES", "SetPlatformApplicationAttributes", respErr)
	if respErr != nil {
		if awsErr, ok := ackerr.AWSError(respErr); ok && awsErr.Code() == "NotFound" {
			// Technically, this means someone deleted the backend resource in
			// between the time we got a result back from sdkFind() and here...
			return nil, ackerr.NotFound
		}
		return nil, respErr
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()
	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromSetAtttributesInput returns true if there are any
// fields for the SetAttributes Input shape that are required by not present in
// the resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromSetAttributesInput(
	r *resource,
) bool {
	return (r.ko.Status.ACKResourceMetadata == nil || r.ko.Status.ACKResourceMetadata.ARN == nil)

}

// newSetAttributesRequestPayload returns SDK-specific struct for the HTTP
// request payload of the SetAttributes API call for the resource
func (rm *resourceManager) newSetAttributesRequestPayload(
	r *resource,
) (*svcsdk.SetPlatformApplicationAttributesInput, error) {
	res := &svcsdk.SetPlatformApplicationAttributesInput{}

	attrMap := map[string]*string{}
	if r.ko.Spec.EventDeliveryFailure != nil {
		attrMap["EventDeliveryFailure"] = r.ko.Spec.EventDeliveryFailure
	}
	if r.ko.Spec.EventEndpointCreated != nil {
		attrMap["EventEndpointCreated"] = r.ko.Spec.EventEndpointCreated
	}
	if r.ko.Spec.EventEndpointDeleted != nil {
		attrMap["EventEndpointDeleted"] = r.ko.Spec.EventEndpointDeleted
	}
	if r.ko.Spec.EventEndpointUpdated != nil {
		attrMap["EventEndpointUpdated"] = r.ko.Spec.EventEndpointUpdated
	}
	if r.ko.Spec.FailureFeedbackRoleARN != nil {
		attrMap["FailureFeedbackRoleArn"] = r.ko.Spec.FailureFeedbackRoleARN
	}
	if r.ko.Spec.PlatformCredential != nil {
		attrMap["PlatformCredential"] = r.ko.Spec.PlatformCredential
	}
	if r.ko.Spec.PlatformPrincipal != nil {
		attrMap["PlatformPrincipal"] = r.ko.Spec.PlatformPrincipal
	}
	if r.ko.Spec.SuccessFeedbackRoleARN != nil {
		attrMap["SuccessFeedbackRoleArn"] = r.ko.Spec.SuccessFeedbackRoleARN
	}
	if r.ko.Spec.SuccessFeedbackSampleRate != nil {
		attrMap["SuccessFeedbackSampleRate"] = r.ko.Spec.SuccessFeedbackSampleRate
	}
	res.SetAttributes(attrMap)
	if r.ko.Status.ACKResourceMetadata != nil && r.ko.Status.ACKResourceMetadata.ARN != nil {
		res.SetPlatformApplicationArn(string(*r.ko.Status.ACKResourceMetadata.ARN))
	} else {
		res.SetPlatformApplicationArn(rm.ARNFromName(*r.ko.Spec.Name))
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer func() {
		exit(err)
	}()
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeletePlatformApplicationOutput
	_ = resp
	resp, err = rm.sdkapi.DeletePlatformApplicationWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeletePlatformApplication", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeletePlatformApplicationInput, error) {
	res := &svcsdk.DeletePlatformApplicationInput{}

	if r.ko.Status.ACKResourceMetadata != nil && r.ko.Status.ACKResourceMetadata.ARN != nil {
		res.SetPlatformApplicationArn(string(*r.ko.Status.ACKResourceMetadata.ARN))
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.PlatformApplication,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.Region == nil {
		ko.Status.ACKResourceMetadata.Region = &rm.awsRegion
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}
	var termError *ackerr.TerminalError
	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	// No terminal_errors specified for this resource in generator config
	return false
}
