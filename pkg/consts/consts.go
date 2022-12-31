package consts

// constants used across client services
const (
	DEFAULT_AUTH_BASE_URL = "https://auth.01.idp.eu01.stackit.cloud/"
	DEFAULT_BASE_URL      = "https://api.stackit.cloud/"

	// schwarz specific constants
	SCHWARZ_ORGANIZATION_ID = "07a1ed91-2efb-42c2-9d00-e84ae71bce0d"
	SCHWARZ_CONTAINER_ID    = "schwarz-it-kg-WJACUK1"
	SCHWARZ_AUTH_ORIGIN     = "schwarz-federation"

	// resource types
	RESOURCE_TYPE_PROJECT = "project"
	RESOURCE_TYPE_ORG     = "organization"

	// Member Roles
	ROLE_PROJECT_OWNER   = "project.owner"
	ROLE_PROJECT_MEMBER  = "project.member"
	ROLE_PROJECT_ADMIN   = "project.admin"
	ROLE_PROJECT_AUDITOR = "project.auditor"

	// Project scopes
	PROJECT_SCOPE_PUBLIC  = "PUBLIC"
	PROJECT_SCOPE_PRIVATE = "PRIVATE"

	// Project lifecycle statuses
	PROJECT_STATUS_ACTIVE   = "ACTIVE"
	PROJECT_STATUS_CREATING = "CREATING"
	PROJECT_STATUS_DELETING = "DELETING"
	PROJECT_STATUS_INACTIVE = "INACTIVE"

	// granularity options; to be used for costs.GetProjectCosts()
	COSTS_GRANULARITY_NONE    = "none"
	COSTS_GRANULARITY_DAILY   = "daily"
	COSTS_GRANULARITY_WEEKLY  = "weekly"
	COSTS_GRANULARITY_MONTHLY = "monthly"
	COSTS_GRANULARITY_YEARLY  = "yearly"

	// depth options; to be used for costs.GetProjectCosts()
	COSTS_DEPTH_PROJECT = "project"
	COSTS_DEPTH_SERVICE = "service"
	COSTS_DEPTH_AUTO    = "auto"
)
