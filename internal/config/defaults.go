package config

import (
	"time"

	conf "github.com/ElfAstAhe/go-service-template/pkg/config"
	"github.com/spf13/viper"
)

// app
const (
	defaultAppEnv       conf.AppEnv = conf.AppEnvDevelopment
	defaultAppNodeName  string      = ApplicationName
	defaultMaxListLimit int         = 100
	defaultTokenIssuer  string      = "tiny-auth-service"
)

// amqp connector
const (
	defaultAMQPConnectorUsername string = "svc-profile"
	defaultAMQPConnectorPassword string = "test"
)

// creds
const (
	defaultCredsScheduleInterval      time.Duration = 5 * time.Minute
	defaultCredsErrorScheduleInterval time.Duration = 3 * time.Second
)

// audit client
const (
	defaultAuditClientTimeout            time.Duration = 5 * time.Second
	defaultAuditClientWorkerCount        int           = 4
	defaultAuditClientDataCapacity       int           = 10000
	defaultAuditClientCompleteProcessing bool          = true
	defaultAuditClientShutdownTimeout    time.Duration = 15 * time.Second
)

//goland:noinspection DuplicatedCode
func applyDefaults(v *viper.Viper) {
	// App
	v.SetDefault(conf.KeyAppEnv, defaultAppEnv)
	v.SetDefault(conf.KeyAppInitTimeout, conf.DefaultAppInitTimeout)
	v.SetDefault(conf.KeyAppStopTimeout, conf.DefaultAppStopTimeout)
	v.SetDefault(conf.KeyAppCloseTimeout, conf.DefaultAppCloseTimeout)
	v.SetDefault(keyAppNodeName, defaultAppNodeName)
	v.SetDefault(keyAppMaxListLimit, defaultMaxListLimit)
	v.SetDefault(keyAppTokenIssuer, defaultTokenIssuer)
	// Auth
	v.SetDefault(conf.KeyAuthJWTSigningMethod, conf.DefaultAuthSigningMethod)
	v.SetDefault(conf.KeyAuthAccessTokenTTL, conf.DefaultAuthAccessTokenTTL)
	v.SetDefault(conf.KeyAuthRefreshTokenTTL, conf.DefaultAuthRefreshTokenTTL)
	// HTTP
	v.SetDefault(conf.KeyHTTPAddress, conf.DefaultHTTPAddress)
	v.SetDefault(conf.KeyHTTPReadTimeout, conf.DefaultHTTPReadTimeout)
	v.SetDefault(conf.KeyHTTPWriteTimeout, conf.DefaultHTTPWriteTimeout)
	v.SetDefault(conf.KeyHTTPIdleTimeout, conf.DefaultHTTPIdleTimeout)
	v.SetDefault(conf.KeyHTTPShutdownTimeout, conf.DefaultHTTPShutdownTimeout)
	v.SetDefault(conf.KeyHTTPSecure, conf.DefaultHTTPSecure)
	v.SetDefault(conf.KeyHTTPMaxRequestBodySize, conf.DefaultHTTPMaxRequestBodySize)
	// gRPC
	v.SetDefault(conf.KeyGRPCAddress, conf.DefaultGRPCAddress)
	v.SetDefault(conf.KeyGRPCMaxConnIdle, conf.DefaultGRPCMaxConnIdle)
	v.SetDefault(conf.KeyGRPCMaxConnAge, conf.DefaultGRPCMaxConnAge)
	v.SetDefault(conf.KeyGRPCMaxConnAgeGrace, conf.DefaultGRPCMaxConnAgeGrace)
	v.SetDefault(conf.KeyGRPCTimeout, conf.DefaultGRPCTimeout)
	v.SetDefault(conf.KeyGRPCKeepAliveTime, conf.DefaultGRPCKeepAliveTime)
	v.SetDefault(conf.KeyGRPCKeepAliveTimeout, conf.DefaultGRPCKeepAliveTimeout)
	v.SetDefault(conf.KeyGRPCShutdownTimeout, conf.DefaultGRPCShutdownTimeout)
	// DB
	v.SetDefault(conf.KeyDBDriver, conf.DefaultDBDriver)
	v.SetDefault(conf.KeyDBDSN, conf.DefaultDBDSN)
	v.SetDefault(conf.KeyDBMaxOpenConns, conf.DefaultDBMaxOpenConns)
	v.SetDefault(conf.KeyDBMaxIdleConns, conf.DefaultDBMaxIdleConns)
	v.SetDefault(conf.KeyDBConnMaxIdleLifetime, conf.DefaultDBConnMaxIdleLifetime)
	v.SetDefault(conf.KeyDBConnTimeout, conf.DefaultDBConnTimeout)
	// Log
	v.SetDefault(conf.KeyLogLevel, conf.DefaultLogLevel)
	v.SetDefault(conf.KeyLogFormat, conf.DefaultLogFormat)
	// Telemetry
	v.SetDefault(conf.KeyTelemetryEnabled, conf.DefaultTelemetryEnabled)
	v.SetDefault(conf.KeyTelemetryExporterEndpoint, conf.DefaultTelemetryExporterEndpoint)
	v.SetDefault(conf.KeyTelemetrySampleRate, conf.DefaultTelemetrySampleRate)
	v.SetDefault(conf.KeyTelemetryTimeout, conf.DefaultTelemetryTimeout)
	// amqp connector
	v.SetDefault(keyAMQPConnectorURL, conf.DefaultAMQPConnectorURL)
	v.SetDefault(keyAMQPConnectorUsername, defaultAMQPConnectorUsername)
	v.SetDefault(keyAMQPConnectorPassword, defaultAMQPConnectorPassword)
	v.SetDefault(keyAMQPConnectorConnectTimeout, conf.DefaultAMQPSenderConnectTimeout)
	v.SetDefault(keyAMQPConnectorWriteTimeout, conf.DefaultAMQPConnectorWriteTimeout)
	v.SetDefault(keyAMQPConnectorIdleTimeout, conf.DefaultAMQPConnectorIdleTimeout)
	v.SetDefault(keyAMQPConnectorShutdownTimeout, conf.DefaultAMQPConnectorShutdownTimeout)
	// data-audit-client
	v.SetDefault(keyDataAuditClientTimeout, defaultAuditClientTimeout)
	v.SetDefault(keyDataAuditClientWorkerCount, defaultAuditClientWorkerCount)
	v.SetDefault(keyDataAuditClientDataCapacity, defaultAuditClientDataCapacity)
	v.SetDefault(keyDataAuditClientCompleteProcessing, defaultAuditClientCompleteProcessing)
	v.SetDefault(keyDataAuditClientShutdownTimeout, defaultAuditClientShutdownTimeout)
	// creds
	v.SetDefault(keySvcCredsScheduleInterval, defaultCredsScheduleInterval)
	v.SetDefault(keySvcCredsErrorScheduleInterval, defaultCredsErrorScheduleInterval)
}
