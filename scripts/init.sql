-- Script de inicialização do banco de dados
-- Executado automaticamente na criação do container PostgreSQL

-- Criação de extensões úteis
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pg_trgm";

-- Configurações de timezone
SET timezone = 'UTC';

-- Criação de índices personalizados (se necessário)
-- Exemplo: CREATE INDEX IF NOT EXISTS idx_categories_name ON categories(name);

-- Log de inicialização
INSERT INTO pg_stat_statements_info (dealloc) VALUES (0) ON CONFLICT DO NOTHING;