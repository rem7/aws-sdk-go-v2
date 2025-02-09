// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Amazon S3 on Outposts Access Points simplify managing data access at scale for
// shared datasets in S3 on Outposts. S3 on Outposts uses endpoints to connect to
// Outposts buckets so that you can perform actions within your virtual private
// cloud (VPC). For more information, see  Accessing S3 on Outposts using VPC-only
// access points
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/WorkingWithS3Outposts.html)
// in the Amazon Simple Storage Service User Guide.
type Endpoint struct {

	// The type of connectivity used to access the Amazon S3 on Outposts endpoint.
	AccessType EndpointAccessType

	// The VPC CIDR committed by this endpoint.
	CidrBlock *string

	// The time the endpoint was created.
	CreationTime *time.Time

	// The ID of the customer-owned IPv4 address pool used for the endpoint.
	CustomerOwnedIpv4Pool *string

	// The Amazon Resource Name (ARN) of the endpoint.
	EndpointArn *string

	// The failure reason, if any, for a create or delete endpoint operation.
	FailedReason *FailedReason

	// The network interface of the endpoint.
	NetworkInterfaces []NetworkInterface

	// The ID of the Outposts.
	OutpostsId *string

	// The ID of the security group used for the endpoint.
	SecurityGroupId *string

	// The status of the endpoint.
	Status EndpointStatus

	// The ID of the subnet used for the endpoint.
	SubnetId *string

	// The ID of the VPC used for the endpoint.
	VpcId *string

	noSmithyDocumentSerde
}

// The failure reason, if any, for a create or delete endpoint operation.
type FailedReason struct {

	// The failure code, if any, for a create or delete endpoint operation.
	ErrorCode *string

	// Additional error details describing the endpoint failure and recommended action.
	Message *string

	noSmithyDocumentSerde
}

// The container for the network interface.
type NetworkInterface struct {

	// The ID for the network interface.
	NetworkInterfaceId *string

	noSmithyDocumentSerde
}

// Contains the details for the Outpost object.
type Outpost struct {

	// The Amazon S3 capacity of the outpost in bytes.
	CapacityInBytes int64

	// Specifies the unique Amazon Resource Name (ARN) for the outpost.
	OutpostArn *string

	// Specifies the unique identifier for the outpost.
	OutpostId *string

	// Returns the Amazon Web Services account ID of the outpost owner. Useful for
	// comparing owned versus shared outposts.
	OwnerId *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
