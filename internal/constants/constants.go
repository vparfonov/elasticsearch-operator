package constants

const (
	Kibana                      = "kibana"
	Elasticsearch               = "elasticsearch"
	ProxyName                   = "cluster"
	OAuthName                   = "cluster"
	TrustedCABundleKey          = "ca-bundle.crt"
	TrustedCABundleMountDir     = "/etc/pki/ca-trust/extracted/pem/"
	TrustedCABundleMountFile    = "tls-ca-bundle.pem"
	InjectTrustedCABundleLabel  = "config.openshift.io/inject-trusted-cabundle"
	TrustedCABundleHashName     = "logging.openshift.io/hash"
	KibanaTrustedCAName         = "kibana-trusted-ca-bundle"
	SecretHashPrefix            = "logging.openshift.io/"
	ElasticsearchDefaultImage   = "quay.io/openshift-logging/elasticsearch6:6.8.1"
	ProxyDefaultImage           = "quay.io/openshift-logging/elasticsearch-proxy:1.0"
	CuratorDefaultImage         = "quay.io/openshift-logging/curator5:5.8.1"
	TheoreticalShardMaxSizeInMB = 40960

	// OcpTemplatePrefix is the prefix all operator generated templates
	OcpTemplatePrefix = "ocp-gen"

	SecurityIndex = ".security"

	EOCertManagementLabel = "logging.openshift.io/elasticsearch-cert-management"
	EOComponentCertPrefix = "logging.openshift.io/elasticsearch-cert."

	ConsoleDashboardLabel          = "console.openshift.io/dashboard"
	LoggingHashLabel               = "logging.openshift.io/hash"
	ElasticsearchDashboardFileName = "openshift-elasticsearch.json"

	// K8s recommended label names: https://kubernetes.io/docs/concepts/overview/working-with-objects/common-labels/
	LabelK8sName      = "app.kubernetes.io/name"       // The name of the application (string)
	LabelK8sInstance  = "app.kubernetes.io/instance"   // A unique name identifying the instance of an application (string)
	LabelK8sVersion   = "app.kubernetes.io/version"    // The current version of the application (e.g., a semantic version, revision hash, etc.) (string)
	LabelK8sComponent = "app.kubernetes.io/component"  // The component within the architecture (string)
	LabelK8sPartOf    = "app.kubernetes.io/part-of"    // The name of a higher level application this one is part of (string)
	LabelK8sManagedBy = "app.kubernetes.io/managed-by" // The tool being used to manage the operation of an application (string)
	LabelK8sCreatedBy = "app.kubernetes.io/created-by" // The controller/user who created this resource (string)

	RedhatClusterLogging        = "redhat-cluster-logging"
	RedhatElasticsearchOperator = "redhat-elasticsearch-operator"
)

var (
	ReconcileForGlobalProxyList = []string{KibanaTrustedCAName}
	ExpectedSecretKeys          = []string{
		"admin-ca",
		"admin-cert",
		"admin-key",
		"elasticsearch.crt",
		"elasticsearch.key",
		"logging-es.crt",
		"logging-es.key",
	}
)
