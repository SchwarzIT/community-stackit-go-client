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
	API_PATH_MONGO_DB_FLEX           = "/mongodb/v1/projects/%s"
	API_PATH_MONGO_DB_FLEX_FLAVORS   = API_PATH_MONGO_DB_FLEX + "/flavors"
	API_PATH_MONGO_DB_FLEX_INSTANCES = API_PATH_MONGO_DB_FLEX + "/instances"
	API_PATH_MONGO_DB_FLEX_INSTANCE  = API_PATH_MONGO_DB_FLEX_INSTANCES + "/%s"
	API_PATH_MONGO_DB_FLEX_STORAGES  = API_PATH_MONGO_DB_FLEX + "/storages/%s"
	API_PATH_MONGO_DB_FLEX_VERSIONS  = API_PATH_MONGO_DB_FLEX + "/versions"
	API_PATH_MONGO_DB_FLEX_USERS     = API_PATH_MONGO_DB_FLEX_INSTANCE + "/users"
	API_PATH_MONGO_DB_FLEX_USER      = API_PATH_MONGO_DB_FLEX_USERS + "/%s"

	// Postgres Flex
	API_PATH_POSTGRES_FLEX                  = "/postgres/v1/projects/%s"
	API_PATH_POSTGRES_FLEX_VERSIONS         = API_PATH_POSTGRES_FLEX + "/versions"
	API_PATH_POSTGRES_FLEX_STORAGES         = API_PATH_POSTGRES_FLEX + "/storages/%s"
	API_PATH_POSTGRES_FLEX_FLAVORS          = API_PATH_POSTGRES_FLEX + "/flavors"
	API_PATH_POSTGRES_FLEX_INSTANCES        = API_PATH_POSTGRES_FLEX + "/instances"
	API_PATH_POSTGRES_FLEX_INSTANCE         = API_PATH_POSTGRES_FLEX_INSTANCES + "/%s"
	API_PATH_POSTGRES_FLEX_INSTANCE_BACKUPS = API_PATH_POSTGRES_FLEX_INSTANCE + "/backups"
	API_PATH_POSTGRES_FLEX_INSTANCE_BACKUP  = API_PATH_POSTGRES_FLEX_INSTANCE_BACKUPS + "/%s"
	API_PATH_POSTGRES_FLEX_USERS            = API_PATH_POSTGRES_FLEX_INSTANCE + "/users"
	API_PATH_POSTGRES_FLEX_USER             = API_PATH_POSTGRES_FLEX_USERS + "/%s"

	// Resource Management v1
	API_PATH_RESOURCE_MANAGEMENT              = "/resource-management/v1"
	API_PATH_RESOURCE_MANAGEMENT_PROJECTS     = API_PATH_RESOURCE_MANAGEMENT + "/projects/%s"
	API_PATH_RESOURCE_MANAGEMENT_ORG_PROJECTS = API_PATH_RESOURCE_MANAGEMENT + "/organizations/%s/projects"
	API_PATH_RESOURCE_MANAGEMENT_ORG_PROJECT  = API_PATH_RESOURCE_MANAGEMENT_ORG_PROJECTS + "/%s"

	// Resource Management v2
	API_PATH_RESOURCE_MANAGEMENT_V2          = "/resource-management/v2"
	API_PATH_RESOURCE_MANAGEMENT_V2_PROJECTS = API_PATH_RESOURCE_MANAGEMENT_V2 + "/projects"
	API_PATH_RESOURCE_MANAGEMENT_V2_PROJECT  = API_PATH_RESOURCE_MANAGEMENT_V2_PROJECTS + "/%s"
	API_PATH_RESOURCE_MANAGEMENT_V2_ORG      = API_PATH_RESOURCE_MANAGEMENT_V2 + "/organizations/%s"

	// Data Services Access (DSA)
	API_PATH_DSA             = "/v1/projects/%s"
	API_PATH_DSA_INSTANCES   = API_PATH_DSA + "/instances"
	API_PATH_DSA_INSTANCE    = API_PATH_DSA_INSTANCES + "/%s"
	API_PATH_DSA_OFFERINGS   = API_PATH_DSA + "/offerings"
	API_PATH_DSA_CREDENTIALS = API_PATH_DSA_INSTANCE + "/credentials"
	API_PATH_DSA_CREDENTIAL  = API_PATH_DSA_CREDENTIALS + "/%s"

	API_BASEURL_DSA_ELASTICSEARCH = "https://elasticsearch.api.eu01.stackit.cloud"
	API_BASEURL_DSA_LOGME         = "https://logme.api.eu01.stackit.cloud"
	API_BASEURL_DSA_MARIADB       = "https://mariadb.api.eu01.stackit.cloud"
	API_BASEURL_DSA_POSTGRES      = "https://postgresql.api.eu01.stackit.cloud"
	API_BASEURL_DSA_RABBITMQ      = "https://rabbitmq.api.eu01.stackit.cloud"
	API_BASEURL_DSA_REDIS         = "https://redis.api.eu01.stackit.cloud"
)
