CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Tabela para Usuários do Sistema
CREATE TABLE IF NOT EXISTS "user" (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    cpf VARCHAR(14) UNIQUE NOT NULL,
    senha_hash VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL,
    ativo BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Tabela para Clientes
CREATE TABLE IF NOT EXISTS "client" (
    id SERIAL PRIMARY KEY,
    ativo BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Tabela para Tipos de Equipamento
CREATE TABLE IF NOT EXISTS "type" (
    id SERIAL PRIMARY KEY,
    marca VARCHAR(100) NOT NULL,
    modelo VARCHAR(100) NOT NULL,
    requer_mac BOOLEAN NOT NULL, -- TRUE se requer MAC, FALSE se quantitativo
    pon_mask VARCHAR(50),
    ativo BOOLEAN DEFAULT TRUE,
    minimo INT NOT NULL DEFAULT 0, -- 0 para produtos que não serão repostos 
    unidade_medida VARCHAR(10) -- 'un', 'metro', 'caixa'
);

-- Tabela para Notas Fiscais de Entrada
CREATE TABLE IF NOT EXISTS "nf" (
    chave_acesso VARCHAR(44) PRIMARY KEY UNIQUE NOT NULL,
    numero VARCHAR(20), -- Pode ser NULL
    data_emissao DATE NOT NULL,
    fornecedor VARCHAR(100), -- Pode ser NULL
    valor_total DECIMAL(10, 2),
    id_responsavel INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

--Detalha os itens de cada Nota Fiscal
CREATE TABLE IF NOT EXISTS "nf_item" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nf_chave_acesso VARCHAR(44) NOT NULL,
    id_type INT NOT NULL,
    quantidade INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Tabela para Equipamentos Individuais ou Itens Quantitativos
CREATE TABLE IF NOT EXISTS "equipamento" (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    mac_id VARCHAR(12) UNIQUE, -- MAC Address, nulo para quantitativos.
    pon_serial VARCHAR(50) UNIQUE, -- PON Serial, nulo para quantitativos ou equipamentos sem PON
    id_type INT NOT NULL,
    quantidade INT NOT NULL DEFAULT 1, -- 1 para equipamentos individuais
    custodiante_type VARCHAR(20) NOT NULL,
    custodiante_id VARCHAR(50),
    nf_chave_acesso VARCHAR(44), -- FK para NF (pode ser nulo)
    ativo BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Tabela para Histórico de Movimentações
CREATE TABLE IF NOT EXISTS "historico" (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    origem_type VARCHAR(20),
    origem_id VARCHAR(50),
    destino_type VARCHAR(20) NOT NULL,
    destino_id VARCHAR(50),
    id_equipamento UUID NOT NULL,
    quantidade INT NOT NULL DEFAULT 1, -- Quantidade movimentada, 1 para equipamentos individuais
    registro_do_chamado VARCHAR(50),
    motivo VARCHAR(100) NOT NULL,
    observacao TEXT,
    data_movimentacao TIMESTAMP WITH TIME ZONE NOT NULL,
    id_responsavel INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Adicionando chaves estrangeiras
ALTER TABLE "nf"
ADD CONSTRAINT fk_nf_responsavel
FOREIGN KEY (id_responsavel) REFERENCES "user"(id);

ALTER TABLE "nf_item"
ADD CONSTRAINT fk_nf_item_nf_chave_acesso
FOREIGN KEY (nf_chave_acesso) REFERENCES "nf"(chave_acesso);

ALTER TABLE "nf_item"
ADD CONSTRAINT fk_nf_item_id_type
FOREIGN KEY (id_type) REFERENCES "type"(id);

ALTER TABLE "equipamento"
ADD CONSTRAINT fk_equipamento_type
FOREIGN KEY (id_type) REFERENCES "type"(id);

ALTER TABLE "equipamento"
ADD CONSTRAINT fk_equipamento_nf
FOREIGN KEY (nf_chave_acesso) REFERENCES "nf"(chave_acesso);

ALTER TABLE "historico"
ADD CONSTRAINT fk_historico_equipamento
FOREIGN KEY (id_equipamento) REFERENCES "equipamento"(uuid);

ALTER TABLE "historico"
ADD CONSTRAINT fk_historico_responsavel
FOREIGN KEY (id_responsavel) REFERENCES "user"(id);

-- Adicionando index
CREATE INDEX IF NOT EXISTS idx_user_cpf ON "user"(cpf);
CREATE INDEX IF NOT EXISTS idx_nf_data_emissao ON "nf"(data_emissao);
CREATE INDEX IF NOT EXISTS idx_equipamento_mac_id ON "equipamento"(mac_id);
CREATE INDEX IF NOT EXISTS idx_equipamento_pon_serial ON "equipamento"(pon_serial);
CREATE INDEX IF NOT EXISTS idx_equipamento_custodiante_type_id ON "equipamento"(custodiante_type, custodiante_id);
CREATE INDEX IF NOT EXISTS idx_historico_data_movimentacao ON "historico"(data_movimentacao);
CREATE INDEX IF NOT EXISTS idx_historico_equipamento_id ON "historico"(id_equipamento);
CREATE INDEX IF NOT EXISTS idx_historico_origem_destino ON "historico"(origem_type, origem_id, destino_type, destino_id);
CREATE INDEX IF NOT EXISTS idx_nf_item_nf_chave_acesso ON "nf_item"(nf_chave_acesso);
CREATE INDEX IF NOT EXISTS idx_nf_item_id_type ON "nf_item"(id_type);