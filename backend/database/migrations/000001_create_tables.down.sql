ALTER TABLE IF EXISTS "Historico"
DROP CONSTRAINT IF EXISTS fk_historico_responsavel;

ALTER TABLE IF EXISTS "Historico"
DROP CONSTRAINT IF EXISTS fk_historico_equipamento;

ALTER TABLE IF EXISTS "Equipamento"
DROP CONSTRAINT IF EXISTS fk_equipamento_nf;

ALTER TABLE IF EXISTS "Equipamento"
DROP CONSTRAINT IF EXISTS fk_equipamento_type;

ALTER TABLE IF EXISTS "NF_Item"
DROP CONSTRAINT IF EXISTS fk_nf_item_id_type;

ALTER TABLE IF EXISTS "NF_Item"
DROP CONSTRAINT IF EXISTS fk_nf_item_nf_chave_acesso;

ALTER TABLE IF EXISTS "NF"
DROP CONSTRAINT IF EXISTS fk_nf_responsavel;


DROP INDEX IF EXISTS idx_nf_item_id_type;
DROP INDEX IF EXISTS idx_nf_item_nf_chave_acesso;
DROP INDEX IF EXISTS idx_historico_origem_destino;
DROP INDEX IF EXISTS idx_historico_equipamento_id;
DROP INDEX IF EXISTS idx_historico_data_movimentacao;
DROP INDEX IF EXISTS idx_equipamento_custodiante_type_id;
DROP INDEX IF EXISTS idx_equipamento_pon_serial;
DROP INDEX IF EXISTS idx_equipamento_mac_id;
DROP INDEX IF EXISTS idx_nf_data_emissao;
DROP INDEX IF EXISTS idx_user_cpf;


DROP TABLE IF EXISTS "Historico";
DROP TABLE IF EXISTS "Equipamento";
DROP TABLE IF EXISTS "NF_Item";
DROP TABLE IF EXISTS "NF";
DROP TABLE IF EXISTS "Type";
DROP TABLE IF EXISTS "Client";
DROP TABLE IF EXISTS "User";