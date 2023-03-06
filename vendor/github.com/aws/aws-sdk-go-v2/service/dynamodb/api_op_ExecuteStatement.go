// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation allows you to perform reads and singleton writes on data stored
// in DynamoDB, using PartiQL. For PartiQL reads (SELECT statement), if the total
// number of processed items exceeds the maximum dataset size limit of 1 MB, the
// read stops and results are returned to the user as a LastEvaluatedKey value to
// continue the read in a subsequent operation. If the filter criteria in WHERE
// clause does not match any data, the read will return an empty result set. A
// single SELECT statement response can return up to the maximum number of items
// (if using the Limit parameter) or a maximum of 1 MB of data (and then apply any
// filtering to the results using WHERE clause). If LastEvaluatedKey is present in
// the response, you need to paginate the result set. If NextToken is present, you
// need to paginate the result set and include NextToken.
func (c *Client) ExecuteStatement(ctx context.Context, params *ExecuteStatementInput, optFns ...func(*Options)) (*ExecuteStatementOutput, error) {
	if params == nil {
		params = &ExecuteStatementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExecuteStatement", params, optFns, c.addOperationExecuteStatementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExecuteStatementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExecuteStatementInput struct {

	// The PartiQL statement representing the operation to run.
	//
	// This member is required.
	Statement *string

	// The consistency of a read operation. If set to true, then a strongly consistent
	// read is used; otherwise, an eventually consistent read is used.
	ConsistentRead *bool

	// The maximum number of items to evaluate (not necessarily the number of matching
	// items). If DynamoDB processes the number of items up to the limit while
	// processing the results, it stops the operation and returns the matching values
	// up to that point, along with a key in LastEvaluatedKey to apply in a subsequent
	// operation so you can pick up where you left off. Also, if the processed dataset
	// size exceeds 1 MB before DynamoDB reaches this limit, it stops the operation and
	// returns the matching values up to the limit, and a key in LastEvaluatedKey to
	// apply in a subsequent operation to continue the operation.
	Limit *int32

	// Set this value to get remaining results, if NextToken was returned in the
	// statement response.
	NextToken *string

	// The parameters for the PartiQL statement, if any.
	Parameters []types.AttributeValue

	// Determines the level of detail about either provisioned or on-demand throughput
	// consumption that is returned in the response:
	//
	// * INDEXES - The response includes
	// the aggregate ConsumedCapacity for the operation, together with ConsumedCapacity
	// for each table and secondary index that was accessed. Note that some operations,
	// such as GetItem and BatchGetItem, do not access any indexes at all. In these
	// cases, specifying INDEXES will only return ConsumedCapacity information for
	// table(s).
	//
	// * TOTAL - The response includes only the aggregate ConsumedCapacity
	// for the operation.
	//
	// * NONE - No ConsumedCapacity details are included in the
	// response.
	ReturnConsumedCapacity types.ReturnConsumedCapacity

	noSmithyDocumentSerde
}

type ExecuteStatementOutput struct {

	// The capacity units consumed by an operation. The data returned includes the
	// total provisioned throughput consumed, along with statistics for the table and
	// any indexes involved in the operation. ConsumedCapacity is only returned if the
	// request asked for it. For more information, see Provisioned Throughput
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughputIntro.html)
	// in the Amazon DynamoDB Developer Guide.
	ConsumedCapacity *types.ConsumedCapacity

	// If a read operation was used, this property will contain the result of the read
	// operation; a map of attribute names and their values. For the write operations
	// this value will be empty.
	Items []map[string]types.AttributeValue

	// The primary key of the item where the operation stopped, inclusive of the
	// previous result set. Use this value to start a new operation, excluding this
	// value in the new request. If LastEvaluatedKey is empty, then the "last page" of
	// results has been processed and there is no more data to be retrieved. If
	// LastEvaluatedKey is not empty, it does not necessarily mean that there is more
	// data in the result set. The only way to know when you have reached the end of
	// the result set is when LastEvaluatedKey is empty.
	LastEvaluatedKey map[string]types.AttributeValue

	// If the response of a read request exceeds the response payload limit DynamoDB
	// will set this value in the response. If set, you can use that this value in the
	// subsequent request to get the remaining results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExecuteStatementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpExecuteStatement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpExecuteStatement{}, middleware.After)
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
	if err = addOpExecuteStatementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExecuteStatement(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addValidateResponseChecksum(stack, options); err != nil {
		return err
	}
	if err = addAcceptEncodingGzip(stack, options); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opExecuteStatement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "ExecuteStatement",
	}
}
