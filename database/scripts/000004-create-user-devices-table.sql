create sequence ecosmart.seq_user_devices;

create table ecosmart.user_devices
(
  id         bigint primary key default nextval('ecosmart.seq_user_devices'),
  user_id    bigint references ecosmart.users(id)   not null,
  device_id  bigint references ecosmart.devices(id) not null,
  created_at timestamp with time zone             not null default current_timestamp,
  updated_at timestamp with time zone             not null default current_timestamp,
  deleted_at timestamp with time zone
);