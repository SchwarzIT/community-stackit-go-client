package consts

// constants defining API paths
const (
	// Costs
	API_PATH_COSTS                     = "/costs-service/v1/costs/%s"
	API_PATH_COSTS_WITH_PARAMS         = API_PATH_COSTS + "?from=%s&to=%s&granularity=%v&depth=%v"
	API_PATH_COSTS_PROJECT             = "/costs-service/v1/costs/%s/projects/%s"
	API_PATH_COSTS_PROJECT_WITH_PARAMS = API_PATH_COSTS_PROJECT + "?from=%s&to=%s&granularity=%v&depth=%v"

	// Object Storage
	API_PATH_OBJECT_STORAGE                         = "/object-storage-api/v1"
	API_PATH_OBJECT_STORAGE_PROJECT                 = API_PATH_OBJECT_STORAGE + "/project/%s"
	API_PATH_OBJECT_STORAGE_PROJECT_FORCE_DELETE    = API_PATH_OBJECT_STORAGE + "/project-force-delete/%s"
	API_PATH_OBJECT_STORAGE_KEYS                    = API_PATH_OBJECT_STORAGE_PROJECT + "/access-keys"
	API_PATH_OBJECT_STORAGE_KEYS_WITH_PARAMS        = API_PATH_OBJECT_STORAGE_KEYS + "?credentials-group=%s"
	API_PATH_OBJECT_STORAGE_KEY                     = API_PATH_OBJECT_STORAGE_PROJECT + "/access-key"
	API_PATH_OBJECT_STORAGE_KEY_WITH_PARAMS         = API_PATH_OBJECT_STORAGE_KEY + "?credentials-group=%s"
	API_PATH_OBJECT_STORAGE_WITH_KEY_ID             = API_PATH_OBJECT_STORAGE_PROJECT + "/access-key/%s"
	API_PATH_OBJECT_STORAGE_WITH_KEY_ID_WITH_PARAMS = API_PATH_OBJECT_STORAGE_WITH_KEY_ID + "?credentials-group=%s"
	API_PATH_OBJECT_STORAGE_BUCKETS                 = API_PATH_OBJECT_STORAGE_PROJECT + "/buckets"
	API_PATH_OBJECT_STORAGE_BUCKET                  = API_PATH_OBJECT_STORAGE_PROJECT + "/bucket/%s"
	API_PATH_OBJECT_STORAGE_CREDENTIALS_CREATE      = API_PATH_OBJECT_STORAGE_PROJECT + "/credentials-group"
	API_PATH_OBJECT_STORAGE_CREDENTIALS_LIST        = API_PATH_OBJECT_STORAGE_PROJECT + "/credentials-groups"
	API_PATH_OBJECT_STORAGE_CREDENTIALS_DELETE      = API_PATH_OBJECT_STORAGE_PROJECT + "/credentials-group/%s"

	// Membership
	API_PATH_MEMBERSHIP                                   = "/membership/v1/projects/%s"
	API_PATH_MEMBERSHIP_ORG_PROJECT                       = "/membership/v1/organizations/%s/projects/%s"
	API_PATH_MEMBERSHIP_ROLES                             = API_PATH_MEMBERSHIP + "/roles"
	API_PATH_MEMBERSHIP_ROLE                              = API_PATH_MEMBERSHIP + "/roles/%s"
	API_PATH_MEMBERSHIP_ROLE_SERVICE_ACCOUNTS             = API_PATH_MEMBERSHIP_ROLE + "/service-accounts"
	API_PATH_MEMBERSHIP_ORG_PROJECT_ROLE_SERVICE_ACCOUNTS = API_PATH_MEMBERSHIP_ORG_PROJECT + "/roles/%s/service-accounts"
	API_PATH_MEMBERSHIP_ORG_PROJECT_ROLE_SERVICE_ACCOUNT  = API_PATH_MEMBERSHIP_ORG_PROJECT + "/roles/%s/service-accounts/%s"
	API_PATH_MEMBERSHIP_ROLES_DELETE                      = API_PATH_MEMBERSHIP + "/roles/delete"

	// Membership v2
	API_PATH_MEMBERSHIP_V2                            = "/membership/v2"
	API_PATH_MEMBERSHIP_V2_WITH_RESOURCE_ID           = API_PATH_MEMBERSHIP_V2 + "/%s"
	API_PATH_MEMBERSHIP_V2_WITH_RESOURCE_TYPE         = API_PATH_MEMBERSHIP_V2_WITH_RESOURCE_ID + "/%s"
	API_PATH_MEMBERSHIP_V2_MEMBERS                    = API_PATH_MEMBERSHIP_V2_WITH_RESOURCE_ID + "/members"
	API_PATH_MEMBERSHIP_V2_WITH_RESOURCE_TYPE_MEMBERS = API_PATH_MEMBERSHIP_V2_WITH_RESOURCE_TYPE + "/members"
	API_PATH_MEMBERSHIP_V2_REMOVE                     = API_PATH_MEMBERSHIP_V2_MEMBERS + "/remove"
	API_PATH_MEMBERSHIP_V2_VALIDATE                   = API_PATH_MEMBERSHIP_V2_MEMBERS + "/validate"
	API_PATH_MEMBERSHIP_V2_PERMISSIONS                = API_PATH_MEMBERSHIP_V2 + "/permissions"
	API_PATH_MEMBERSHIP_V2_USER_PERMISSIONS           = API_PATH_MEMBERSHIP_V2 + "/users/%s/permissions"
	API_PATH_MEMBERSHIP_V2_ROLES                      = API_PATH_MEMBERSHIP_V2_WITH_RESOURCE_ID + "/roles"
	API_PATH_MEMBERSHIP_V2_ROLES_WITH_RESOURCE_TYPE   = API_PATH_MEMBERSHIP_V2_WITH_RESOURCE_TYPE + "/roles"
)
