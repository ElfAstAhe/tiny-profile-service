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
