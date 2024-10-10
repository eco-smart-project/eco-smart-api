create sequence ecosmart.seq_devices;

create table ecosmart.devices
(
  id         bigint primary key default nextval('ecosmart.seq_devices'),
  type       bpchar(1)                          not null check ( type in ('D','M') ),
  status     bpchar(1)                          not null default 'A' check ( status in ('A','I','D') ),
  token      text                               not null,
  os         text                               not null,
  model      text                               not null,
  push_token text,
  created_at timestamp with time zone           not null default current_timestamp,
  updated_at timestamp with time zone           not null default current_timestamp,
  deleted_at timestamp with time zone,
  created_by bigint references ecosmart.users(id) not null,
  updated_by bigint references ecosmart.users(id) not null
);

comment on column ecosmart.devices.type is 'D: Device, M: Mobile';
comment on column ecosmart.devices.status is 'A: Active, I: Inactive, D: Deleted';