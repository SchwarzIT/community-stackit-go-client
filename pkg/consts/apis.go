package consts

// constants defining API paths
const (
	// Argus
	API_PATH_ARGUS                    = "/argus-service/v1/projects/%s"
	API_PATH_ARGUS_CREDENTIALS        = API_PATH_ARGUS_WITH_INSTANCE_ID + "/credentials"
	API_PATH_ARGUS_INSTANCES          = API_PATH_ARGUS + "/instances"
	API_PATH_ARGUS_WITH_INSTANCE_ID   = API_PATH_ARGUS_INSTANCES + "/%s"
	API_PATH_ARGUS_GRAFANA_CONFIGS    = API_PATH_ARGUS_WITH_INSTANCE_ID + "/grafana-configs"
	API_PATH_ARGUS_JOBS               = API_PATH_ARGUS_WITH_INSTANCE_ID + "/scrapeconfigs"
	API_PATH_ARGUS_JOBS_WITH_JOB_NAME = API_PATH_ARGUS_JOBS + "/%s"
	API_PATH_ARGUS_METRICS_RETENTION  = API_PATH_ARGUS_WITH_INSTANCE_ID + "/metrics-storage-retentions"
	API_PATH_ARGUS_PLANS              = API_PATH_ARGUS + "/plans"
	API_PATH_ARGUS_TRACES             = API_PATH_ARGUS_WITH_INSTANCE_ID + "/traces-configs"

	// Costs
	API_PATH_COSTS                     = "/costs-service/v1/costs/%s"
	API_PATH_COSTS_WITH_PARAMS         = API_PATH_COSTS + "?from=%s&to=%s&granularity=%v&depth=%v"
	API_PATH_COSTS_PROJECT             = "/costs-service/v1/costs/%s/projects/%s"
	API_PATH_COSTS_PROJECT_WITH_PARAMS = API_PATH_COSTS_PROJECT + "?from=%s&to=%s&granularity=%v&depth=%v"

	// Kubernetes
	API_PATH_SKE                 = "/ske/v1"
	API_PATH_SKE_PROJECTS        = API_PATH_SKE + "/projects/%s"
	API_PATH_SKE_CLUSTERS        = API_PATH_SKE_PROJECTS + "/clusters"
	API_PATH_SKE_WITH_CLUSTER_ID = API_PATH_SKE_CLUSTERS + "/%s"
	API_PATH_SKE_OPTIONS         = API_PATH_SKE + "/provider-options"

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

	// MongoDB Flex
	API_PATH_MONGO_DB_FLEX          = "/mongodb-service-api/v1/projects/%s"
	API_PATH_MONGO_DB_FLEX_VERSIONS = API_PATH_MONGO_DB_FLEX + "/versions"
	API_PATH_MONGO_DB_FLEX_FLAVORS  = API_PATH_MONGO_DB_FLEX + "/flavors"
	API_PATH_MONGO_DB_FLEX_STORAGE  = API_PATH_MONGO_DB_FLEX + "/storages/%s"

	// Resource Management
	API_PATH_RESOURCE_MANAGEMENT              = "/resource-management/v1"
	API_PATH_RESOURCE_MANAGEMENT_PROJECTS     = API_PATH_RESOURCE_MANAGEMENT + "/projects/%s"
	API_PATH_RESOURCE_MANAGEMENT_ORG_PROJECTS = API_PATH_RESOURCE_MANAGEMENT + "/organizations/%s/projects"
	API_PATH_RESOURCE_MANAGEMENT_ORG_PROJECT  = API_PATH_RESOURCE_MANAGEMENT_ORG_PROJECTS + "/%s"

	// Resource Manager v2
	API_PATH_RESOURCE_MANAGER_V2          = "/resource-manager/v2"
	API_PATH_RESOURCE_MANAGER_V2_PROJECTS = API_PATH_RESOURCE_MANAGER_V2 + "/projects"
	API_PATH_RESOURCE_MANAGER_V2_PROJECT  = API_PATH_RESOURCE_MANAGER_V2_PROJECTS + "/%s"
	API_PATH_RESOURCE_MANAGER_V2_ORG      = API_PATH_RESOURCE_MANAGER_V2 + "/organizations/%s"

	// Shadow Users
	API_PATH_SHADOW_USERS = "/ucp-shadow-user-management/v1/createcuaashadowuser/user"
)
