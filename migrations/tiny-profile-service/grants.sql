-- Manual usage

-- 1. Позволяет создавать новые объекты (таблицы, индексы, последовательности)
GRANT CREATE ON SCHEMA profile_db TO svc_profile;

-- 2. Позволяет видеть схему
GRANT USAGE ON SCHEMA profile_db TO svc_profile;

-- 3. Если таблицы уже созданы другим юзером (например, postgres),
-- нужно сделать test их владельцем или дать полные права:
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA profile_db TO svc_profile;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA profile_db TO svc_profile;

-- 4. Чтобы будущие объекты тоже были под контролем:
ALTER DEFAULT PRIVILEGES IN SCHEMA profile_db
    GRANT ALL PRIVILEGES ON TABLES TO svc_profile;
ALTER DEFAULT PRIVILEGES IN SCHEMA profile_db
    GRANT ALL PRIVILEGES ON SEQUENCES TO svc_profile;
