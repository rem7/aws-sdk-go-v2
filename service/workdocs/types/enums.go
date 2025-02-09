// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ActivityType string

// Enum values for ActivityType
const (
	ActivityTypeDocumentCheckedIn                      ActivityType = "DOCUMENT_CHECKED_IN"
	ActivityTypeDocumentCheckedOut                     ActivityType = "DOCUMENT_CHECKED_OUT"
	ActivityTypeDocumentRenamed                        ActivityType = "DOCUMENT_RENAMED"
	ActivityTypeDocumentVersionUploaded                ActivityType = "DOCUMENT_VERSION_UPLOADED"
	ActivityTypeDocumentVersionDeleted                 ActivityType = "DOCUMENT_VERSION_DELETED"
	ActivityTypeDocumentVersionViewed                  ActivityType = "DOCUMENT_VERSION_VIEWED"
	ActivityTypeDocumentVersionDownloaded              ActivityType = "DOCUMENT_VERSION_DOWNLOADED"
	ActivityTypeDocumentRecycled                       ActivityType = "DOCUMENT_RECYCLED"
	ActivityTypeDocumentRestored                       ActivityType = "DOCUMENT_RESTORED"
	ActivityTypeDocumentReverted                       ActivityType = "DOCUMENT_REVERTED"
	ActivityTypeDocumentShared                         ActivityType = "DOCUMENT_SHARED"
	ActivityTypeDocumentUnshared                       ActivityType = "DOCUMENT_UNSHARED"
	ActivityTypeDocumentSharePermissionChanged         ActivityType = "DOCUMENT_SHARE_PERMISSION_CHANGED"
	ActivityTypeDocumentShareableLinkCreated           ActivityType = "DOCUMENT_SHAREABLE_LINK_CREATED"
	ActivityTypeDocumentShareableLinkRemoved           ActivityType = "DOCUMENT_SHAREABLE_LINK_REMOVED"
	ActivityTypeDocumentShareableLinkPermissionChanged ActivityType = "DOCUMENT_SHAREABLE_LINK_PERMISSION_CHANGED"
	ActivityTypeDocumentMoved                          ActivityType = "DOCUMENT_MOVED"
	ActivityTypeDocumentCommentAdded                   ActivityType = "DOCUMENT_COMMENT_ADDED"
	ActivityTypeDocumentCommentDeleted                 ActivityType = "DOCUMENT_COMMENT_DELETED"
	ActivityTypeDocumentAnnotationAdded                ActivityType = "DOCUMENT_ANNOTATION_ADDED"
	ActivityTypeDocumentAnnotationDeleted              ActivityType = "DOCUMENT_ANNOTATION_DELETED"
	ActivityTypeFolderCreated                          ActivityType = "FOLDER_CREATED"
	ActivityTypeFolderDeleted                          ActivityType = "FOLDER_DELETED"
	ActivityTypeFolderRenamed                          ActivityType = "FOLDER_RENAMED"
	ActivityTypeFolderRecycled                         ActivityType = "FOLDER_RECYCLED"
	ActivityTypeFolderRestored                         ActivityType = "FOLDER_RESTORED"
	ActivityTypeFolderShared                           ActivityType = "FOLDER_SHARED"
	ActivityTypeFolderUnshared                         ActivityType = "FOLDER_UNSHARED"
	ActivityTypeFolderSharePermissionChanged           ActivityType = "FOLDER_SHARE_PERMISSION_CHANGED"
	ActivityTypeFolderShareableLinkCreated             ActivityType = "FOLDER_SHAREABLE_LINK_CREATED"
	ActivityTypeFolderShareableLinkRemoved             ActivityType = "FOLDER_SHAREABLE_LINK_REMOVED"
	ActivityTypeFolderShareableLinkPermissionChanged   ActivityType = "FOLDER_SHAREABLE_LINK_PERMISSION_CHANGED"
	ActivityTypeFolderMoved                            ActivityType = "FOLDER_MOVED"
)

// Values returns all known values for ActivityType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ActivityType) Values() []ActivityType {
	return []ActivityType{
		"DOCUMENT_CHECKED_IN",
		"DOCUMENT_CHECKED_OUT",
		"DOCUMENT_RENAMED",
		"DOCUMENT_VERSION_UPLOADED",
		"DOCUMENT_VERSION_DELETED",
		"DOCUMENT_VERSION_VIEWED",
		"DOCUMENT_VERSION_DOWNLOADED",
		"DOCUMENT_RECYCLED",
		"DOCUMENT_RESTORED",
		"DOCUMENT_REVERTED",
		"DOCUMENT_SHARED",
		"DOCUMENT_UNSHARED",
		"DOCUMENT_SHARE_PERMISSION_CHANGED",
		"DOCUMENT_SHAREABLE_LINK_CREATED",
		"DOCUMENT_SHAREABLE_LINK_REMOVED",
		"DOCUMENT_SHAREABLE_LINK_PERMISSION_CHANGED",
		"DOCUMENT_MOVED",
		"DOCUMENT_COMMENT_ADDED",
		"DOCUMENT_COMMENT_DELETED",
		"DOCUMENT_ANNOTATION_ADDED",
		"DOCUMENT_ANNOTATION_DELETED",
		"FOLDER_CREATED",
		"FOLDER_DELETED",
		"FOLDER_RENAMED",
		"FOLDER_RECYCLED",
		"FOLDER_RESTORED",
		"FOLDER_SHARED",
		"FOLDER_UNSHARED",
		"FOLDER_SHARE_PERMISSION_CHANGED",
		"FOLDER_SHAREABLE_LINK_CREATED",
		"FOLDER_SHAREABLE_LINK_REMOVED",
		"FOLDER_SHAREABLE_LINK_PERMISSION_CHANGED",
		"FOLDER_MOVED",
	}
}

type AdditionalResponseFieldType string

// Enum values for AdditionalResponseFieldType
const (
	AdditionalResponseFieldTypeWeburl AdditionalResponseFieldType = "WEBURL"
)

// Values returns all known values for AdditionalResponseFieldType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (AdditionalResponseFieldType) Values() []AdditionalResponseFieldType {
	return []AdditionalResponseFieldType{
		"WEBURL",
	}
}

type BooleanEnumType string

// Enum values for BooleanEnumType
const (
	BooleanEnumTypeTrue  BooleanEnumType = "TRUE"
	BooleanEnumTypeFalse BooleanEnumType = "FALSE"
)

// Values returns all known values for BooleanEnumType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BooleanEnumType) Values() []BooleanEnumType {
	return []BooleanEnumType{
		"TRUE",
		"FALSE",
	}
}

type CommentStatusType string

// Enum values for CommentStatusType
const (
	CommentStatusTypeDraft     CommentStatusType = "DRAFT"
	CommentStatusTypePublished CommentStatusType = "PUBLISHED"
	CommentStatusTypeDeleted   CommentStatusType = "DELETED"
)

// Values returns all known values for CommentStatusType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CommentStatusType) Values() []CommentStatusType {
	return []CommentStatusType{
		"DRAFT",
		"PUBLISHED",
		"DELETED",
	}
}

type CommentVisibilityType string

// Enum values for CommentVisibilityType
const (
	CommentVisibilityTypePublic  CommentVisibilityType = "PUBLIC"
	CommentVisibilityTypePrivate CommentVisibilityType = "PRIVATE"
)

// Values returns all known values for CommentVisibilityType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CommentVisibilityType) Values() []CommentVisibilityType {
	return []CommentVisibilityType{
		"PUBLIC",
		"PRIVATE",
	}
}

type ContentCategoryType string

// Enum values for ContentCategoryType
const (
	ContentCategoryTypeImage        ContentCategoryType = "IMAGE"
	ContentCategoryTypeDocument     ContentCategoryType = "DOCUMENT"
	ContentCategoryTypePdf          ContentCategoryType = "PDF"
	ContentCategoryTypeSpreadsheet  ContentCategoryType = "SPREADSHEET"
	ContentCategoryTypePresentation ContentCategoryType = "PRESENTATION"
	ContentCategoryTypeAudio        ContentCategoryType = "AUDIO"
	ContentCategoryTypeVideo        ContentCategoryType = "VIDEO"
	ContentCategoryTypeSourceCode   ContentCategoryType = "SOURCE_CODE"
	ContentCategoryTypeOther        ContentCategoryType = "OTHER"
)

// Values returns all known values for ContentCategoryType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ContentCategoryType) Values() []ContentCategoryType {
	return []ContentCategoryType{
		"IMAGE",
		"DOCUMENT",
		"PDF",
		"SPREADSHEET",
		"PRESENTATION",
		"AUDIO",
		"VIDEO",
		"SOURCE_CODE",
		"OTHER",
	}
}

type DocumentSourceType string

// Enum values for DocumentSourceType
const (
	DocumentSourceTypeOriginal     DocumentSourceType = "ORIGINAL"
	DocumentSourceTypeWithComments DocumentSourceType = "WITH_COMMENTS"
)

// Values returns all known values for DocumentSourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DocumentSourceType) Values() []DocumentSourceType {
	return []DocumentSourceType{
		"ORIGINAL",
		"WITH_COMMENTS",
	}
}

type DocumentStatusType string

// Enum values for DocumentStatusType
const (
	DocumentStatusTypeInitialized DocumentStatusType = "INITIALIZED"
	DocumentStatusTypeActive      DocumentStatusType = "ACTIVE"
)

// Values returns all known values for DocumentStatusType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DocumentStatusType) Values() []DocumentStatusType {
	return []DocumentStatusType{
		"INITIALIZED",
		"ACTIVE",
	}
}

type DocumentThumbnailType string

// Enum values for DocumentThumbnailType
const (
	DocumentThumbnailTypeSmall   DocumentThumbnailType = "SMALL"
	DocumentThumbnailTypeSmallHq DocumentThumbnailType = "SMALL_HQ"
	DocumentThumbnailTypeLarge   DocumentThumbnailType = "LARGE"
)

// Values returns all known values for DocumentThumbnailType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DocumentThumbnailType) Values() []DocumentThumbnailType {
	return []DocumentThumbnailType{
		"SMALL",
		"SMALL_HQ",
		"LARGE",
	}
}

type DocumentVersionStatus string

// Enum values for DocumentVersionStatus
const (
	DocumentVersionStatusActive DocumentVersionStatus = "ACTIVE"
)

// Values returns all known values for DocumentVersionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DocumentVersionStatus) Values() []DocumentVersionStatus {
	return []DocumentVersionStatus{
		"ACTIVE",
	}
}

type FolderContentType string

// Enum values for FolderContentType
const (
	FolderContentTypeAll      FolderContentType = "ALL"
	FolderContentTypeDocument FolderContentType = "DOCUMENT"
	FolderContentTypeFolder   FolderContentType = "FOLDER"
)

// Values returns all known values for FolderContentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FolderContentType) Values() []FolderContentType {
	return []FolderContentType{
		"ALL",
		"DOCUMENT",
		"FOLDER",
	}
}

type LanguageCodeType string

// Enum values for LanguageCodeType
const (
	LanguageCodeTypeAr      LanguageCodeType = "AR"
	LanguageCodeTypeBg      LanguageCodeType = "BG"
	LanguageCodeTypeBn      LanguageCodeType = "BN"
	LanguageCodeTypeDa      LanguageCodeType = "DA"
	LanguageCodeTypeDe      LanguageCodeType = "DE"
	LanguageCodeTypeCs      LanguageCodeType = "CS"
	LanguageCodeTypeEl      LanguageCodeType = "EL"
	LanguageCodeTypeEn      LanguageCodeType = "EN"
	LanguageCodeTypeEs      LanguageCodeType = "ES"
	LanguageCodeTypeFa      LanguageCodeType = "FA"
	LanguageCodeTypeFi      LanguageCodeType = "FI"
	LanguageCodeTypeFr      LanguageCodeType = "FR"
	LanguageCodeTypeHi      LanguageCodeType = "HI"
	LanguageCodeTypeHu      LanguageCodeType = "HU"
	LanguageCodeTypeId      LanguageCodeType = "ID"
	LanguageCodeTypeIt      LanguageCodeType = "IT"
	LanguageCodeTypeJa      LanguageCodeType = "JA"
	LanguageCodeTypeKo      LanguageCodeType = "KO"
	LanguageCodeTypeLt      LanguageCodeType = "LT"
	LanguageCodeTypeLv      LanguageCodeType = "LV"
	LanguageCodeTypeNl      LanguageCodeType = "NL"
	LanguageCodeTypeNo      LanguageCodeType = "NO"
	LanguageCodeTypePt      LanguageCodeType = "PT"
	LanguageCodeTypeRo      LanguageCodeType = "RO"
	LanguageCodeTypeRu      LanguageCodeType = "RU"
	LanguageCodeTypeSv      LanguageCodeType = "SV"
	LanguageCodeTypeSw      LanguageCodeType = "SW"
	LanguageCodeTypeTh      LanguageCodeType = "TH"
	LanguageCodeTypeTr      LanguageCodeType = "TR"
	LanguageCodeTypeZh      LanguageCodeType = "ZH"
	LanguageCodeTypeDefault LanguageCodeType = "DEFAULT"
)

// Values returns all known values for LanguageCodeType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LanguageCodeType) Values() []LanguageCodeType {
	return []LanguageCodeType{
		"AR",
		"BG",
		"BN",
		"DA",
		"DE",
		"CS",
		"EL",
		"EN",
		"ES",
		"FA",
		"FI",
		"FR",
		"HI",
		"HU",
		"ID",
		"IT",
		"JA",
		"KO",
		"LT",
		"LV",
		"NL",
		"NO",
		"PT",
		"RO",
		"RU",
		"SV",
		"SW",
		"TH",
		"TR",
		"ZH",
		"DEFAULT",
	}
}

type LocaleType string

// Enum values for LocaleType
const (
	LocaleTypeEn      LocaleType = "en"
	LocaleTypeFr      LocaleType = "fr"
	LocaleTypeKo      LocaleType = "ko"
	LocaleTypeDe      LocaleType = "de"
	LocaleTypeEs      LocaleType = "es"
	LocaleTypeJa      LocaleType = "ja"
	LocaleTypeRu      LocaleType = "ru"
	LocaleTypeZhCn    LocaleType = "zh_CN"
	LocaleTypeZhTw    LocaleType = "zh_TW"
	LocaleTypePtBr    LocaleType = "pt_BR"
	LocaleTypeDefault LocaleType = "default"
)

// Values returns all known values for LocaleType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (LocaleType) Values() []LocaleType {
	return []LocaleType{
		"en",
		"fr",
		"ko",
		"de",
		"es",
		"ja",
		"ru",
		"zh_CN",
		"zh_TW",
		"pt_BR",
		"default",
	}
}

type OrderByFieldType string

// Enum values for OrderByFieldType
const (
	OrderByFieldTypeRelevance         OrderByFieldType = "RELEVANCE"
	OrderByFieldTypeName              OrderByFieldType = "NAME"
	OrderByFieldTypeSize              OrderByFieldType = "SIZE"
	OrderByFieldTypeCreatedTimestamp  OrderByFieldType = "CREATED_TIMESTAMP"
	OrderByFieldTypeModifiedTimestamp OrderByFieldType = "MODIFIED_TIMESTAMP"
)

// Values returns all known values for OrderByFieldType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OrderByFieldType) Values() []OrderByFieldType {
	return []OrderByFieldType{
		"RELEVANCE",
		"NAME",
		"SIZE",
		"CREATED_TIMESTAMP",
		"MODIFIED_TIMESTAMP",
	}
}

type OrderType string

// Enum values for OrderType
const (
	OrderTypeAscending  OrderType = "ASCENDING"
	OrderTypeDescending OrderType = "DESCENDING"
)

// Values returns all known values for OrderType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (OrderType) Values() []OrderType {
	return []OrderType{
		"ASCENDING",
		"DESCENDING",
	}
}

type PrincipalRoleType string

// Enum values for PrincipalRoleType
const (
	PrincipalRoleTypeViewer      PrincipalRoleType = "VIEWER"
	PrincipalRoleTypeContributor PrincipalRoleType = "CONTRIBUTOR"
	PrincipalRoleTypeOwner       PrincipalRoleType = "OWNER"
	PrincipalRoleTypeCoowner     PrincipalRoleType = "COOWNER"
)

// Values returns all known values for PrincipalRoleType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PrincipalRoleType) Values() []PrincipalRoleType {
	return []PrincipalRoleType{
		"VIEWER",
		"CONTRIBUTOR",
		"OWNER",
		"COOWNER",
	}
}

type PrincipalType string

// Enum values for PrincipalType
const (
	PrincipalTypeUser         PrincipalType = "USER"
	PrincipalTypeGroup        PrincipalType = "GROUP"
	PrincipalTypeInvite       PrincipalType = "INVITE"
	PrincipalTypeAnonymous    PrincipalType = "ANONYMOUS"
	PrincipalTypeOrganization PrincipalType = "ORGANIZATION"
)

// Values returns all known values for PrincipalType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PrincipalType) Values() []PrincipalType {
	return []PrincipalType{
		"USER",
		"GROUP",
		"INVITE",
		"ANONYMOUS",
		"ORGANIZATION",
	}
}

type ResourceCollectionType string

// Enum values for ResourceCollectionType
const (
	ResourceCollectionTypeSharedWithMe ResourceCollectionType = "SHARED_WITH_ME"
)

// Values returns all known values for ResourceCollectionType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceCollectionType) Values() []ResourceCollectionType {
	return []ResourceCollectionType{
		"SHARED_WITH_ME",
	}
}

type ResourceSortType string

// Enum values for ResourceSortType
const (
	ResourceSortTypeDate ResourceSortType = "DATE"
	ResourceSortTypeName ResourceSortType = "NAME"
)

// Values returns all known values for ResourceSortType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceSortType) Values() []ResourceSortType {
	return []ResourceSortType{
		"DATE",
		"NAME",
	}
}

type ResourceStateType string

// Enum values for ResourceStateType
const (
	ResourceStateTypeActive    ResourceStateType = "ACTIVE"
	ResourceStateTypeRestoring ResourceStateType = "RESTORING"
	ResourceStateTypeRecycling ResourceStateType = "RECYCLING"
	ResourceStateTypeRecycled  ResourceStateType = "RECYCLED"
)

// Values returns all known values for ResourceStateType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceStateType) Values() []ResourceStateType {
	return []ResourceStateType{
		"ACTIVE",
		"RESTORING",
		"RECYCLING",
		"RECYCLED",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeFolder   ResourceType = "FOLDER"
	ResourceTypeDocument ResourceType = "DOCUMENT"
)

// Values returns all known values for ResourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"FOLDER",
		"DOCUMENT",
	}
}

type ResponseItemType string

// Enum values for ResponseItemType
const (
	ResponseItemTypeDocument        ResponseItemType = "DOCUMENT"
	ResponseItemTypeFolder          ResponseItemType = "FOLDER"
	ResponseItemTypeComment         ResponseItemType = "COMMENT"
	ResponseItemTypeDocumentVersion ResponseItemType = "DOCUMENT_VERSION"
)

// Values returns all known values for ResponseItemType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResponseItemType) Values() []ResponseItemType {
	return []ResponseItemType{
		"DOCUMENT",
		"FOLDER",
		"COMMENT",
		"DOCUMENT_VERSION",
	}
}

type RolePermissionType string

// Enum values for RolePermissionType
const (
	RolePermissionTypeDirect    RolePermissionType = "DIRECT"
	RolePermissionTypeInherited RolePermissionType = "INHERITED"
)

// Values returns all known values for RolePermissionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RolePermissionType) Values() []RolePermissionType {
	return []RolePermissionType{
		"DIRECT",
		"INHERITED",
	}
}

type RoleType string

// Enum values for RoleType
const (
	RoleTypeViewer      RoleType = "VIEWER"
	RoleTypeContributor RoleType = "CONTRIBUTOR"
	RoleTypeOwner       RoleType = "OWNER"
	RoleTypeCoowner     RoleType = "COOWNER"
)

// Values returns all known values for RoleType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (RoleType) Values() []RoleType {
	return []RoleType{
		"VIEWER",
		"CONTRIBUTOR",
		"OWNER",
		"COOWNER",
	}
}

type SearchCollectionType string

// Enum values for SearchCollectionType
const (
	SearchCollectionTypeOwned        SearchCollectionType = "OWNED"
	SearchCollectionTypeSharedWithMe SearchCollectionType = "SHARED_WITH_ME"
)

// Values returns all known values for SearchCollectionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SearchCollectionType) Values() []SearchCollectionType {
	return []SearchCollectionType{
		"OWNED",
		"SHARED_WITH_ME",
	}
}

type SearchQueryScopeType string

// Enum values for SearchQueryScopeType
const (
	SearchQueryScopeTypeName    SearchQueryScopeType = "NAME"
	SearchQueryScopeTypeContent SearchQueryScopeType = "CONTENT"
)

// Values returns all known values for SearchQueryScopeType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SearchQueryScopeType) Values() []SearchQueryScopeType {
	return []SearchQueryScopeType{
		"NAME",
		"CONTENT",
	}
}

type SearchResourceType string

// Enum values for SearchResourceType
const (
	SearchResourceTypeFolder          SearchResourceType = "FOLDER"
	SearchResourceTypeDocument        SearchResourceType = "DOCUMENT"
	SearchResourceTypeComment         SearchResourceType = "COMMENT"
	SearchResourceTypeDocumentVersion SearchResourceType = "DOCUMENT_VERSION"
)

// Values returns all known values for SearchResourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SearchResourceType) Values() []SearchResourceType {
	return []SearchResourceType{
		"FOLDER",
		"DOCUMENT",
		"COMMENT",
		"DOCUMENT_VERSION",
	}
}

type ShareStatusType string

// Enum values for ShareStatusType
const (
	ShareStatusTypeSuccess ShareStatusType = "SUCCESS"
	ShareStatusTypeFailure ShareStatusType = "FAILURE"
)

// Values returns all known values for ShareStatusType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ShareStatusType) Values() []ShareStatusType {
	return []ShareStatusType{
		"SUCCESS",
		"FAILURE",
	}
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAsc  SortOrder = "ASC"
	SortOrderDesc SortOrder = "DESC"
)

// Values returns all known values for SortOrder. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (SortOrder) Values() []SortOrder {
	return []SortOrder{
		"ASC",
		"DESC",
	}
}

type StorageType string

// Enum values for StorageType
const (
	StorageTypeUnlimited StorageType = "UNLIMITED"
	StorageTypeQuota     StorageType = "QUOTA"
)

// Values returns all known values for StorageType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (StorageType) Values() []StorageType {
	return []StorageType{
		"UNLIMITED",
		"QUOTA",
	}
}

type SubscriptionProtocolType string

// Enum values for SubscriptionProtocolType
const (
	SubscriptionProtocolTypeHttps SubscriptionProtocolType = "HTTPS"
	SubscriptionProtocolTypeSqs   SubscriptionProtocolType = "SQS"
)

// Values returns all known values for SubscriptionProtocolType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SubscriptionProtocolType) Values() []SubscriptionProtocolType {
	return []SubscriptionProtocolType{
		"HTTPS",
		"SQS",
	}
}

type SubscriptionType string

// Enum values for SubscriptionType
const (
	SubscriptionTypeAll SubscriptionType = "ALL"
)

// Values returns all known values for SubscriptionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SubscriptionType) Values() []SubscriptionType {
	return []SubscriptionType{
		"ALL",
	}
}

type UserFilterType string

// Enum values for UserFilterType
const (
	UserFilterTypeAll           UserFilterType = "ALL"
	UserFilterTypeActivePending UserFilterType = "ACTIVE_PENDING"
)

// Values returns all known values for UserFilterType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UserFilterType) Values() []UserFilterType {
	return []UserFilterType{
		"ALL",
		"ACTIVE_PENDING",
	}
}

type UserSortType string

// Enum values for UserSortType
const (
	UserSortTypeUserName     UserSortType = "USER_NAME"
	UserSortTypeFullName     UserSortType = "FULL_NAME"
	UserSortTypeStorageLimit UserSortType = "STORAGE_LIMIT"
	UserSortTypeUserStatus   UserSortType = "USER_STATUS"
	UserSortTypeStorageUsed  UserSortType = "STORAGE_USED"
)

// Values returns all known values for UserSortType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (UserSortType) Values() []UserSortType {
	return []UserSortType{
		"USER_NAME",
		"FULL_NAME",
		"STORAGE_LIMIT",
		"USER_STATUS",
		"STORAGE_USED",
	}
}

type UserStatusType string

// Enum values for UserStatusType
const (
	UserStatusTypeActive   UserStatusType = "ACTIVE"
	UserStatusTypeInactive UserStatusType = "INACTIVE"
	UserStatusTypePending  UserStatusType = "PENDING"
)

// Values returns all known values for UserStatusType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UserStatusType) Values() []UserStatusType {
	return []UserStatusType{
		"ACTIVE",
		"INACTIVE",
		"PENDING",
	}
}

type UserType string

// Enum values for UserType
const (
	UserTypeUser           UserType = "USER"
	UserTypeAdmin          UserType = "ADMIN"
	UserTypePoweruser      UserType = "POWERUSER"
	UserTypeMinimaluser    UserType = "MINIMALUSER"
	UserTypeWorkspacesuser UserType = "WORKSPACESUSER"
)

// Values returns all known values for UserType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (UserType) Values() []UserType {
	return []UserType{
		"USER",
		"ADMIN",
		"POWERUSER",
		"MINIMALUSER",
		"WORKSPACESUSER",
	}
}
