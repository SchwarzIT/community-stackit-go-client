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
)
