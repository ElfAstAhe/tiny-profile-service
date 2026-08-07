# LADR-002: Техническая спецификация контрактов и схем данных tiny-profile

## Статус
Принято (Approved) — Август 2026 г.

## Контекст
На основании концептуальных решений, принятых в [ADR-001](adr-001-tiny-profile-architecture.md), технические спецификации вынесены в отдельные независимые файлы контрактов для автоматизации линтинга и упрощения кодогенерации на базе `go-service-template`.

## Спецификация контрактов

### 1. Синхронный gRPC Контракт (Protobuf)
Спецификация серверных заглушек находится по пути `api/proto/tiny/profile/v1/profile.proto`.

```protobuf
syntax = "proto3";
package tiny.profile.v1;

service ProfileService {
  rpc GetProfile (GetProfileRequest) returns (GetProfileResponse);
  rpc DeactivateProfile (DeactivateProfileRequest) returns (DeactivateProfileResponse);
  rpc DeactivatePerson (DeactivatePersonRequest) returns (DeactivatePersonResponse);
}

message GetProfileRequest { string user_id = 1; }
message GetProfileResponse {
  string user_id = 1;
  string person_id = 2;
  string first_name = 3;
  string last_name = 4;
  string middle_name = 5;
  string avatar_url = 6;
  string lang = 7;
  string timezone = 8;
  bool is_active = 9;
}

message DeactivateProfileRequest { string user_id = 1; string reason = 2; }
message DeactivateProfileResponse { bool success = 1; }

message DeactivatePersonRequest { string person_id = 1; string reason = 2; }
message DeactivatePersonResponse { bool success = 1; }
```

### 2. Асинхронные события (AMQP JSON)
Полная схема сообщений брокера зафиксирована в формате AsyncAPI по пути `docs/api/asyncapi.yaml`.

#### Тема (Subject): `profile.events.v1.deactivated`
```json
{
  "user_id": "8f2d59b4-7d5a-43eb-8f72-91d83501a357",
  "reason": "SECURITY_COMPROMISE",
  "timestamp": "2026-08-07T17:10:00Z"
}
```

#### Тема (Subject): `person.events.v1.deactivated`
```json
{
  "person_id": "41a10de3-b452-4759-99bb-cf85c7f8a111",
  "user_ids": ["8f2d59b4-7d5a-43eb-8f72-91d83501a357"],
  "reason": "EMPLOYEE_TERMINATION",
  "timestamp": "2026-08-07T17:10:00Z"
}
```
