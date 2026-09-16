package config

const (
	ApplicationName = "tiny-profile-service.proto"
)

// FlagConfig - файл конфигурации
const FlagConfig = "config-path"

// App config flags
const (
	FlagAppNodeName           string = "node-name"
	FlagAppMaxListLimit       string = "app-max-list-limit"
	FlagAppTokenIssuer        string = "app-token-issuer"
	FlagAppCipherKey          string = "app-cipher-key"
	FlagAppAcceptTokenIssuers string = "app-accept-token-issuers"
)

// Auth config flags
const (
	FlagAuthJWTSecret          string = "auth-jwt-secret"
	FlagAuthJWTSigningMethod   string = "auth-jwt-signing-method"
	FlagAuthAccessTokenTTL     string = "auth-access-token-ttl"
	FlagAuthRefreshTokenTTL    string = "auth-refresh-token-ttl"
	FlagAuthRSAPrivateKeyPath  string = "auth-rsa-private-key-path"
	FlagAuthMasterPasswordSalt string = "auth-master-password-salt"
)

// DB config flags
const (
	FlagDBDSN             string = "db-dsn"
	FlagDBDriver          string = "db-driver"
	FlagDBMaxOpenConns    string = "db-max-open-conns"
	FlagDBMaxIdleConns    string = "db-max-idle-conns"
	FlagDBMaxIdleLifetime string = "db-max-idle-lifetime"
	FlagDBConnTimeout     string = "db-conn-timeout"
)

// gRPC config flags
const (
	FlagGRPCAddress          string = "grpc-address"
	FlagGRPCMaxConnIdle      string = "grpc-max-conn-idle"
	FlagGRPCMaxConnAge       string = "grpc-max-conn-age"
	FlagGRPCMaxConnAgeGrace  string = "grpc-max-conn-age-grace"
	FlagGRPCTimeout          string = "grpc-timeout"
	FlagGRPCKeepAliveTime    string = "grpc-keep-alive-time"
	FlagGRPCKeepAliveTimeout string = "grpc-keep-alive-timeout"
	FlagGRPCShutdownTimeout  string = "grpc-shutdown-timeout"
)

// http config flags
const (
	FlagHTTPAddress            string = "http-address"
	FlagHTTPReadTimeout        string = "http-read-timeout"
	FlagHTTPWriteTimeout       string = "http-write-timeout"
	FlagHTTPIdleTimeout        string = "http-idle-timeout"
	FlagHTTPShutdownTimeout    string = "http-shutdown-timeout"
	FlagHTTPPrivateKeyPath     string = "http-private-key-path"
	FlagHTTPCertificatePath    string = "http-certificate-path"
	FlagHTTPSecure             string = "http-secure"
	FlagHTTPMaxRequestBodySize string = "http-max-request-body-size"
)

// log config flags
const (
	FlagLogLevel  string = "log-level"
	FlagLogFormat string = "log-format"
)

// telemetry
const (
	FlagTelemetryEnabled          string = "telemetry-enabled"
	FlagTelemetryServiceName      string = "telemetry-service-name"
	FlagTelemetryExporterEndpoint string = "telemetry-exporter-endpoint"
	FlagTelemetrySampleRate       string = "telemetry-sample-rate"
	FlagTelemetryTimeout          string = "telemetry-timeout"
)

// EnvConfig - файл конфигурации
const EnvConfig string = "CONFIG_PATH"

// amqp connector
const (
	FlagAMQPConnectorURL             string = "amqp-connector-url"
	FlagAMQPConnectorUsername        string = "amqp-connector-username"
	FlagAMQPConnectorPassword        string = "amqp-connector-password"
	FlagAMQPConnectorConnectTimeout  string = "amqp-connector-connect-timeout"
	FlagAMQPConnectorWriteTimeout    string = "amqp-connector-write-timeout"
	FlagAMQPConnectorIdleTimeout     string = "amqp-connector-idle-timeout"
	FlagAMQPConnectorShutdownTimeout string = "amqp-connector-shutdown-timeout"
)

// Data-audit-client
const (
	FlagDataAuditClientBaseURL            string = "data-audit-client-base-url"
	FlagDataAuditClientTimeout            string = "data-audit-client-timeout"
	FlagDataAuditClientWorkerCount        string = "data-audit-client-worker-count"
	FlagDataAuditClientDataCapacity       string = "data-audit-client-data-capacity"
	FlagDataAuditClientCompleteProcessing string = "data-audit-client-complete-processing"
	FlagDataAuditClientShutdownTimeout    string = "data-audit-client-shutdown-timeout"
)

// Creds
const (
	FlagCredsUsername              string = "svc-creds-username"
	FlagCredsPassword              string = "svc-creds-password"
	FlagCredsScheduleInterval      string = "svc-creds-schedule-interval"
	FlagCredsErrorScheduleInterval string = "svc-creds-error-schedule-interval"
)

// app
const (
	keyAppNodeName           string = "app.node_name"
	keyAppMaxListLimit       string = "app.max_list_limit"
	keyAppTokenIssuer        string = "app.token_issuer"
	keyAppCipherKey          string = "app.cipher_key"
	keyAppAcceptTokenIssuers string = "app.accept_token_issuers"
)

// amqp connector
const (
	keyAMQPConnectorURL             string = "amqp_connector.url"
	keyAMQPConnectorUsername        string = "amqp_connector.username"
	keyAMQPConnectorPassword        string = "amqp_connector.password"
	keyAMQPConnectorConnectTimeout  string = "amqp_connector.connect_timeout"
	keyAMQPConnectorWriteTimeout    string = "amqp_connector.write_timeout"
	keyAMQPConnectorIdleTimeout     string = "amqp_connector.idle_timeout"
	keyAMQPConnectorShutdownTimeout string = "amqp_connector.shutdown_timeout"
)

// data audit client
const (
	keyDataAuditClientBaseURL            string = "data_audit_client.base_url"
	keyDataAuditClientTimeout            string = "data_audit_client.timeout"
	keyDataAuditClientWorkerCount        string = "data_audit_client.worker_count"
	keyDataAuditClientDataCapacity       string = "data_audit_client.data_capacity"
	keyDataAuditClientCompleteProcessing string = "data_audit_client.complete_processing"
	keyDataAuditClientShutdownTimeout    string = "data_audit_client.shutdown_timeout"
)

// service credentials
const (
	keySvcCredsUsername              string = "svc_creds.username"
	keySvcCredsPassword              string = "svc_creds.password"
	keySvcCredsScheduleInterval      string = "svc_creds.schedule_interval"
	keySvcCredsErrorScheduleInterval string = "svc_creds.error_schedule_interval"
)
