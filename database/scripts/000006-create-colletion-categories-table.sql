create sequence ecosmart.seq_collection_categories;

create table ecosmart.collection_categories 
(
    id bigint primary key default nextval('ecosmart.seq_collection_categories'),
    name varchar(50) not null,
    description text,
    icon varchar(100),
    status     bpchar(1)                            not null default 'A' check ( status in ('A','I') ),
    created_at timestamp with time zone             not null default current_timestamp,
    updated_at timestamp with time zone             not null default current_timestamp,
    deleted_at timestamp with time zone
);

INSERT INTO ecosmart.collection_categories (name, description, icon)
VALUES 
('Resíduos Gerais', 'Coleta de resíduos gerais', 'assets/images/garbage-type-icons/default.png'),
('Vidro', 'Coleta de materiais de vidro', 'assets/images/garbage-type-icons/glass.png'),
('Eletrônicos', 'Coleta de equipamentos eletrônicos', 'assets/images/garbage-type-icons/eletronic.png');