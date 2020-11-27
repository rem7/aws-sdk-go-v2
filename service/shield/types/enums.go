// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AttackLayer string

// Enum values for AttackLayer
const (
	AttackLayerNetwork     AttackLayer = "NETWORK"
	AttackLayerApplication AttackLayer = "APPLICATION"
)

// Values returns all known values for AttackLayer. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (AttackLayer) Values() []AttackLayer {
	return []AttackLayer{
		"NETWORK",
		"APPLICATION",
	}
}

type AttackPropertyIdentifier string

// Enum values for AttackPropertyIdentifier
const (
	AttackPropertyIdentifierDestinationUrl             AttackPropertyIdentifier = "DESTINATION_URL"
	AttackPropertyIdentifierReferrer                   AttackPropertyIdentifier = "REFERRER"
	AttackPropertyIdentifierSourceAsn                  AttackPropertyIdentifier = "SOURCE_ASN"
	AttackPropertyIdentifierSourceCountry              AttackPropertyIdentifier = "SOURCE_COUNTRY"
	AttackPropertyIdentifierSourceIpAddress            AttackPropertyIdentifier = "SOURCE_IP_ADDRESS"
	AttackPropertyIdentifierSourceUserAgent            AttackPropertyIdentifier = "SOURCE_USER_AGENT"
	AttackPropertyIdentifierWordpressPingbackReflector AttackPropertyIdentifier = "WORDPRESS_PINGBACK_REFLECTOR"
	AttackPropertyIdentifierWordpressPingbackSource    AttackPropertyIdentifier = "WORDPRESS_PINGBACK_SOURCE"
)

// Values returns all known values for AttackPropertyIdentifier. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AttackPropertyIdentifier) Values() []AttackPropertyIdentifier {
	return []AttackPropertyIdentifier{
		"DESTINATION_URL",
		"REFERRER",
		"SOURCE_ASN",
		"SOURCE_COUNTRY",
		"SOURCE_IP_ADDRESS",
		"SOURCE_USER_AGENT",
		"WORDPRESS_PINGBACK_REFLECTOR",
		"WORDPRESS_PINGBACK_SOURCE",
	}
}

type AutoRenew string

// Enum values for AutoRenew
const (
	AutoRenewEnabled  AutoRenew = "ENABLED"
	AutoRenewDisabled AutoRenew = "DISABLED"
)

// Values returns all known values for AutoRenew. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (AutoRenew) Values() []AutoRenew {
	return []AutoRenew{
		"ENABLED",
		"DISABLED",
	}
}

type ProactiveEngagementStatus string

// Enum values for ProactiveEngagementStatus
const (
	ProactiveEngagementStatusEnabled  ProactiveEngagementStatus = "ENABLED"
	ProactiveEngagementStatusDisabled ProactiveEngagementStatus = "DISABLED"
	ProactiveEngagementStatusPending  ProactiveEngagementStatus = "PENDING"
)

// Values returns all known values for ProactiveEngagementStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ProactiveEngagementStatus) Values() []ProactiveEngagementStatus {
	return []ProactiveEngagementStatus{
		"ENABLED",
		"DISABLED",
		"PENDING",
	}
}

type ProtectedResourceType string

// Enum values for ProtectedResourceType
const (
	ProtectedResourceTypeCloudfrontDistribution  ProtectedResourceType = "CLOUDFRONT_DISTRIBUTION"
	ProtectedResourceTypeRoute53HostedZone       ProtectedResourceType = "ROUTE_53_HOSTED_ZONE"
	ProtectedResourceTypeElasticIpAllocation     ProtectedResourceType = "ELASTIC_IP_ALLOCATION"
	ProtectedResourceTypeClassicLoadBalancer     ProtectedResourceType = "CLASSIC_LOAD_BALANCER"
	ProtectedResourceTypeApplicationLoadBalancer ProtectedResourceType = "APPLICATION_LOAD_BALANCER"
	ProtectedResourceTypeGlobalAccelerator       ProtectedResourceType = "GLOBAL_ACCELERATOR"
)

// Values returns all known values for ProtectedResourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ProtectedResourceType) Values() []ProtectedResourceType {
	return []ProtectedResourceType{
		"CLOUDFRONT_DISTRIBUTION",
		"ROUTE_53_HOSTED_ZONE",
		"ELASTIC_IP_ALLOCATION",
		"CLASSIC_LOAD_BALANCER",
		"APPLICATION_LOAD_BALANCER",
		"GLOBAL_ACCELERATOR",
	}
}

type ProtectionGroupAggregation string

// Enum values for ProtectionGroupAggregation
const (
	ProtectionGroupAggregationSum  ProtectionGroupAggregation = "SUM"
	ProtectionGroupAggregationMean ProtectionGroupAggregation = "MEAN"
	ProtectionGroupAggregationMax  ProtectionGroupAggregation = "MAX"
)

// Values returns all known values for ProtectionGroupAggregation. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ProtectionGroupAggregation) Values() []ProtectionGroupAggregation {
	return []ProtectionGroupAggregation{
		"SUM",
		"MEAN",
		"MAX",
	}
}

type ProtectionGroupPattern string

// Enum values for ProtectionGroupPattern
const (
	ProtectionGroupPatternAll            ProtectionGroupPattern = "ALL"
	ProtectionGroupPatternArbitrary      ProtectionGroupPattern = "ARBITRARY"
	ProtectionGroupPatternByResourceType ProtectionGroupPattern = "BY_RESOURCE_TYPE"
)

// Values returns all known values for ProtectionGroupPattern. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ProtectionGroupPattern) Values() []ProtectionGroupPattern {
	return []ProtectionGroupPattern{
		"ALL",
		"ARBITRARY",
		"BY_RESOURCE_TYPE",
	}
}

type SubResourceType string

// Enum values for SubResourceType
const (
	SubResourceTypeIp  SubResourceType = "IP"
	SubResourceTypeUrl SubResourceType = "URL"
)

// Values returns all known values for SubResourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SubResourceType) Values() []SubResourceType {
	return []SubResourceType{
		"IP",
		"URL",
	}
}

type SubscriptionState string

// Enum values for SubscriptionState
const (
	SubscriptionStateActive   SubscriptionState = "ACTIVE"
	SubscriptionStateInactive SubscriptionState = "INACTIVE"
)

// Values returns all known values for SubscriptionState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SubscriptionState) Values() []SubscriptionState {
	return []SubscriptionState{
		"ACTIVE",
		"INACTIVE",
	}
}

type Unit string

// Enum values for Unit
const (
	UnitBits     Unit = "BITS"
	UnitBytes    Unit = "BYTES"
	UnitPackets  Unit = "PACKETS"
	UnitRequests Unit = "REQUESTS"
)

// Values returns all known values for Unit. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Unit) Values() []Unit {
	return []Unit{
		"BITS",
		"BYTES",
		"PACKETS",
		"REQUESTS",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "FIELD_VALIDATION_FAILED"
	ValidationExceptionReasonOther                 ValidationExceptionReason = "OTHER"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"FIELD_VALIDATION_FAILED",
		"OTHER",
	}
}
