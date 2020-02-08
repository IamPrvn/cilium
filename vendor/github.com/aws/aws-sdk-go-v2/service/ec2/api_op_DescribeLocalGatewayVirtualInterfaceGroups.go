// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeLocalGatewayVirtualInterfaceGroupsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The IDs of the virtual interface groups.
	LocalGatewayVirtualInterfaceGroupIds []string `locationName:"LocalGatewayVirtualInterfaceGroupId" locationNameList:"item" type:"list"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token for the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeLocalGatewayVirtualInterfaceGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeLocalGatewayVirtualInterfaceGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeLocalGatewayVirtualInterfaceGroupsInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeLocalGatewayVirtualInterfaceGroupsOutput struct {
	_ struct{} `type:"structure"`

	// The virtual interface groups.
	LocalGatewayVirtualInterfaceGroups []LocalGatewayVirtualInterfaceGroup `locationName:"localGatewayVirtualInterfaceGroupSet" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeLocalGatewayVirtualInterfaceGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeLocalGatewayVirtualInterfaceGroups = "DescribeLocalGatewayVirtualInterfaceGroups"

// DescribeLocalGatewayVirtualInterfaceGroupsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the specified local gateway virtual interface groups.
//
//    // Example sending a request using DescribeLocalGatewayVirtualInterfaceGroupsRequest.
//    req := client.DescribeLocalGatewayVirtualInterfaceGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeLocalGatewayVirtualInterfaceGroups
func (c *Client) DescribeLocalGatewayVirtualInterfaceGroupsRequest(input *DescribeLocalGatewayVirtualInterfaceGroupsInput) DescribeLocalGatewayVirtualInterfaceGroupsRequest {
	op := &aws.Operation{
		Name:       opDescribeLocalGatewayVirtualInterfaceGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeLocalGatewayVirtualInterfaceGroupsInput{}
	}

	req := c.newRequest(op, input, &DescribeLocalGatewayVirtualInterfaceGroupsOutput{})
	return DescribeLocalGatewayVirtualInterfaceGroupsRequest{Request: req, Input: input, Copy: c.DescribeLocalGatewayVirtualInterfaceGroupsRequest}
}

// DescribeLocalGatewayVirtualInterfaceGroupsRequest is the request type for the
// DescribeLocalGatewayVirtualInterfaceGroups API operation.
type DescribeLocalGatewayVirtualInterfaceGroupsRequest struct {
	*aws.Request
	Input *DescribeLocalGatewayVirtualInterfaceGroupsInput
	Copy  func(*DescribeLocalGatewayVirtualInterfaceGroupsInput) DescribeLocalGatewayVirtualInterfaceGroupsRequest
}

// Send marshals and sends the DescribeLocalGatewayVirtualInterfaceGroups API request.
func (r DescribeLocalGatewayVirtualInterfaceGroupsRequest) Send(ctx context.Context) (*DescribeLocalGatewayVirtualInterfaceGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeLocalGatewayVirtualInterfaceGroupsResponse{
		DescribeLocalGatewayVirtualInterfaceGroupsOutput: r.Request.Data.(*DescribeLocalGatewayVirtualInterfaceGroupsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeLocalGatewayVirtualInterfaceGroupsResponse is the response type for the
// DescribeLocalGatewayVirtualInterfaceGroups API operation.
type DescribeLocalGatewayVirtualInterfaceGroupsResponse struct {
	*DescribeLocalGatewayVirtualInterfaceGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeLocalGatewayVirtualInterfaceGroups request.
func (r *DescribeLocalGatewayVirtualInterfaceGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}