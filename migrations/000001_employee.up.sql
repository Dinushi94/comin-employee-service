-- migrations/000001_create_initial_schema.up.sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Positions/Job Titles
CREATE TABLE positions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID NOT NULL,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    level INTEGER,
    department_id UUID,
    status VARCHAR(50) DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Employees
CREATE TABLE employees (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    organization_id UUID NOT NULL,
    position_id UUID REFERENCES positions(id),
    department_id UUID,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(50),
    date_of_birth DATE,
    hire_date DATE NOT NULL,
    termination_date DATE,
    employee_id VARCHAR(50),  -- Internal employee ID/number
    status VARCHAR(50) DEFAULT 'active',
    manager_id UUID REFERENCES employees(id),
    work_type VARCHAR(50) DEFAULT 'full-time',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(organization_id, employee_id)
);

-- Employee Contracts
CREATE TABLE employee_contracts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    employee_id UUID REFERENCES employees(id),
    start_date DATE NOT NULL,
    end_date DATE,
    contract_type VARCHAR(50) NOT NULL,
    salary_amount DECIMAL(12,2),
    currency VARCHAR(3),
    status VARCHAR(50) DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Employee Documents
CREATE TABLE employee_documents (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    employee_id UUID REFERENCES employees(id),
    document_type VARCHAR(50) NOT NULL,
    document_name VARCHAR(255) NOT NULL,
    document_url VARCHAR(1000) NOT NULL,
    expiry_date DATE,
    status VARCHAR(50) DEFAULT 'active',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes
CREATE INDEX idx_employees_organization ON employees(organization_id);
CREATE INDEX idx_employees_department ON employees(department_id);
CREATE INDEX idx_employees_position ON employees(position_id);
CREATE INDEX idx_employees_manager ON employees(manager_id);
CREATE INDEX idx_positions_organization ON positions(organization_id);
CREATE INDEX idx_positions_department ON positions(department_id);
CREATE INDEX idx_employee_contracts_employee ON employee_contracts(employee_id);
CREATE INDEX idx_employee_documents_employee ON employee_documents(employee_id);