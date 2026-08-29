# LADR-003: Спецификация миграции базы данных через встроенный накат Goose

## Статус
Принято (Approved) — Август 2026 г.

## Контекст
Согласно архитектурным принципам проекта, для миграции базы данных используется встроенный (`in-code`) подход на базе библиотеки **Goose**. Миграции реализуются в виде Go-кода, выполняющего нативные SQL-операции в рамках транзакций.

## Спецификация DDL (Data First)

Реализуется двухтабличная схема локального кэша сотрудников. Текстовые поля снабжены `NOT NULL DEFAULT ''` для исключения указателей и `sql.NullString` в Go-коде репозитория.

### Таблица: `persons`
*   `id` (VARCHAR(50), Primary Key)
*   `external_id` (VARCHAR(100), Unique Index, id из внешней ERP)
*   `last_name` (VARCHAR(100))
*   `first_name` (VARCHAR(100))
*   `patronymic` (VARCHAR(100))
*   `birthday` (TIMESTAMP WITH TIME ZONE)
*   `department` (VARCHAR(100))
*   `position` (VARCHAR(100))
*   `status` (VARCHAR(100))
*   `avatar_url` (TEXT)
*   `active` (BOOLEAN, по умолчанию `FALSE`)
*   `deleted` (BOOLEAN, по умолчанию `FALSE`)
*   `created_at` (TIMESTAMP WITH TIME ZONE)
*   `updated_at` (TIMESTAMP WITH TIME ZONE) 

### Таблица: `profiles`
*   `id` (VARCHAR(50), Priary Key)
*   `user_id` (VARCHAR(100), Unique Index, отношение 1:1 к учетным записям, например в `tiny-auth-service`)
*   `person_id` (VARCHAR(50), Foreign Key на таблицу `persons`, ограничение `ON DELETE RESTRICT`)
*   `time_zone` (VARCHAR(50), по умолчанию `UTC`)
*   `lang` (VARCHAR(10), по умолчанию `ru`)
*   `active` (BOOLEAN, по умолчанию `FALSE`)
*   `deleted` (BOOLEAN, по умолчанию `FALSE`)
*   `created_at` (TIMESTAMP WITH TIME ZONE)
*   `updated_at` (TIMESTAMP WITH TIME ZONE)

## Последствия
*   Логика структуры БД зафиксирована в документации независимо от Go-кода мигратора.
*   При реализации `in-code` миграции Goose на Go разработчик использует SQL-блоки, жестко соответствующие этой спецификации.
