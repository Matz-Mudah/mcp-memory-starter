-- Initialize pgvector extension for vector similarity search
CREATE EXTENSION IF NOT EXISTS vector;

-- Initialize Apache AGE extension for graph relationships
CREATE EXTENSION IF NOT EXISTS age;

-- Load AGE into search path
LOAD 'age';
SET search_path = ag_catalog, "$user", public;

-- Create memories table with vector support in public schema
CREATE TABLE IF NOT EXISTS public.memories (
    id BIGSERIAL PRIMARY KEY,
    text TEXT NOT NULL,
    embedding vector(768),  -- pgvector column (768 dimensions for embeddinggemma-300m)
    group_id TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for performance
-- NOTE: IVFFlat index requires ~1000+ rows to work properly. For small datasets, skip this index.
-- Uncomment after you have sufficient data:
-- CREATE INDEX IF NOT EXISTS idx_memories_embedding ON public.memories USING ivfflat (embedding vector_cosine_ops) WITH (lists = 100);

-- For now, let pgvector do sequential scans (works fine for small datasets)
CREATE INDEX IF NOT EXISTS idx_memories_group_id ON public.memories(group_id);
CREATE INDEX IF NOT EXISTS idx_memories_created_at ON public.memories(created_at DESC);

-- Create Apache AGE graph for memory relationships
SELECT create_graph('memory_graph');

-- Function to update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Trigger to auto-update updated_at
CREATE TRIGGER update_memories_updated_at
    BEFORE UPDATE ON public.memories
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Grant permissions
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO memoryuser;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO memoryuser;
GRANT USAGE ON SCHEMA ag_catalog TO memoryuser;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA ag_catalog TO memoryuser;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA ag_catalog TO memoryuser;
