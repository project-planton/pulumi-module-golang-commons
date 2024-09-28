package fieldsextractor

import (
	"github.com/plantoncloud/project-planton/apis/zzgo/cloud/planton/apis/code2cloud/v1/aws/awslambda"
	"google.golang.org/protobuf/proto"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestExtractApiResourceSpecField_Success tests the success scenario where fields are valid and messages are correctly structured.
func TestExtractApiResourceSpecField_Success(t *testing.T) {
	// Create mock input
	mockSpec := &awslambda.AwsLambdaSpec{}
	mockTarget := &awslambda.AwsLambda{Spec: mockSpec}
	mockStackInput := &awslambda.AwsLambdaStackInput{Target: mockTarget}

	// Call the function under test
	result, err := ExtractApiResourceSpecField(mockStackInput)

	// Assert no error and check result is as expected
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, proto.Equal((*result).Interface().(proto.Message), mockSpec), "Expected and actual protobuf messages are not equal")
}

// TestExtractApiResourceSpecField_NilInput tests the scenario where stackInput is nil.
func TestExtractApiResourceSpecField_NilInput(t *testing.T) {
	// Call the function under test with nil input
	result, err := ExtractApiResourceSpecField(nil)

	// Assert that an error is returned and result is nil
	assert.Nil(t, result)
	assert.EqualError(t, err, "stack-input is nil")
}

// TestExtractApiResourceSpecField_InvalidTargetField tests the scenario where the target field is missing or invalid.
func TestExtractApiResourceSpecField_InvalidTargetField(t *testing.T) {
	// Create a mock input without a valid Target field
	mockStackInput := &awslambda.AwsLambdaStackInput{Target: nil}

	// Call the function under test
	result, err := ExtractApiResourceSpecField(mockStackInput)

	// Assert that an error is returned
	assert.Nil(t, result)
	assert.EqualError(t, err, "Field target is nil in stack-input")
}

// TestExtractApiResourceSpecField_InvalidSpecField tests the scenario where the spec field in the target is missing or invalid.
func TestExtractApiResourceSpecField_InvalidSpecField(t *testing.T) {
	// Create a mock input with an invalid Spec field
	mockTarget := &awslambda.AwsLambda{Spec: nil}
	mockStackInput := &awslambda.AwsLambdaStackInput{Target: mockTarget}

	// Call the function under test
	result, err := ExtractApiResourceSpecField(mockStackInput)

	// Assert that an error is returned
	assert.Nil(t, result)
	assert.EqualError(t, err, "Field spec is nil in target")
}
