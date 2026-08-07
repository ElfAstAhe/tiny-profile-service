# LADR-003: Спецификация миграции базы данных через встроенный накат Goose

## Статус
Принято (Approved) — Август 2026 г.

## Контекст
Согласно архитектурным принципам проекта, для миграции базы данных используется встроенный (`in-code`) подход на базе библиотеки **Goose**. Миграции реализуются в виде Go-кода, выполняющего нативные SQL-операции в рамках транзакций.

## Спецификация DDL (Data First)

Реализуется двухтабличная схема локального кэша сотрудников. Текстовые поля снабжены `NOT NULL DEFAULT ''` для исключения указателей и `sql.NullString` в Go-коде репозитория.

### Таблица: `persons`
*   `id` (UUID, Primary Key)
*   `ext_id` (VARCHAR(100), Unique Index, опциональный табельный номер из внешней ERP)
*   `first_name` (VARCHAR(100))
*   `last_name` (VARCHAR(100))
*   `middle_name` (VARCHAR(100))
*   `avatar_url` (TEXT)
*   `created_at` (TIMESTAMP WITH TIME ZONE)

### Таблица: `profiles`
*   `user_id` (UUID, Primary Key, отношение 1:1 к учетным записям в `tiny-auth`)
*   `person_id` (UUID, Foreign Key на таблицу `persons`, ограничение `ON DELETE RESTRICT`)
*   `is_active` (BOOLEAN, по умолчанию `TRUE`)
*   `lang` (VARCHAR(10), по умолчанию `ru`)
*   `timezone` (VARCHAR(50), по умолчанию `UTC`)
*   `created_at` (TIMESTAMP WITH TIME ZONE)
*   `updated_at` (TIMESTAMP WITH TIME ZONE)

## Последствия
*   Логика структуры БД зафиксирована в документации независимо от Go-кода мигратора.
*   При реализации `in-code` миграции Goose на Go разработчик использует SQL-блоки, жестко соответствующие этой спецификации.
