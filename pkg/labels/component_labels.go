package labels

const (
	OperatorPrefix   = "operator.kyma-project.io"
	Separator        = "/"
	KymaName         = OperatorPrefix + Separator + "kyma-name"
	ManagedBy        = OperatorPrefix + Separator + "managed-by"
	LifecycleManager = "lifecycle-manager"
	OperatorName     = "module-manager"
	OwnedByLabel     = OperatorPrefix + Separator + "owned-by"
	WatchedByLabel   = OperatorPrefix + Separator + "watched-by"
)
