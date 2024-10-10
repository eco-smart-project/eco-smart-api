create sequence ecosmart.seq_sessions;

create table ecosmart.sessions
(
  id         bigint primary key default nextval('ecosmart.seq_sessions'),
  user_id    bigint references ecosmart.users(id) not null,
  token      text                               not null,
  status     bpchar(1)                          not null default 'A' check ( status in ('A','I') ),
  expires_at timestamp with time zone           not null,
  created_at timestamp with time zone           not null default current_timestamp,
  updated_at timestamp with time zone           not null default current_timestamp
);