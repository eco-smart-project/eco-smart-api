create sequence ecosmart.seq_collection_points;

create table ecosmart.collection_points 
(
    id bigint primary key default nextval('ecosmart.seq_collection_points'),
    position_latitude decimal(9, 6) not null,
    position_longitude decimal(9, 6) not null,
    title varchar(100) not null,
    description text,
    icon varchar(100),
    status     bpchar(1)                            not null default 'A' check ( status in ('A','I') ),
    category_id int references ecosmart.collection_categories(id) on delete set null,
    created_at timestamp with time zone             not null default current_timestamp,
    updated_at timestamp with time zone             not null default current_timestamp,
    deleted_at timestamp with time zone
);

INSERT INTO ecosmart.collection_points (position_latitude, position_longitude, title, description, icon, category_id)
VALUES
(-23.55052, -46.63331, 'Ponto de Coleta 1', 'Coleta de resíduos gerais', 'assets/images/garbage-type-icons/default.png',
 (SELECT id FROM ecosmart.collection_categories WHERE name = 'Resíduos Gerais')),
(-23.55103, -46.64048, 'Ponto de Coleta 2', 'Coleta de metais e vidros', 'assets/images/garbage-type-icons/glass.png',
 (SELECT id FROM ecosmart.collection_categories WHERE name = 'Vidro')),
(-23.55255, -46.64395, 'Ponto de Coleta 3', 'Coleta de eletrônicos', 'assets/images/garbage-type-icons/eletronic.png',
 (SELECT id FROM ecosmart.collection_categories WHERE name = 'Eletrônicos'));