CREATE TABLE IF NOT EXISTS companies (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    address TEXT,
    website VARCHAR(255),
    phone VARCHAR(50),
    email VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create index for search functionality (since ListCompaniesRequest has search field)
CREATE INDEX idx_companies_name_description ON companies USING GIN (
    to_tsvector('english', name || ' ' || COALESCE(description, ''))
);

-- Create index for common queries
CREATE INDEX idx_companies_created_at ON companies(created_at);
