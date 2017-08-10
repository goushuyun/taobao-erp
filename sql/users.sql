create SEQUENCE user_id_seq;

create table users (
    "id" text primary key not null default  'u_' || to_char(now() AT TIME ZONE 'cct', 'yymmdd') || trim(to_char(nextval('user_id_seq'), '00000')),

    mobile text not null,
    password text not null,
    name text not null,
    role int not null default 512,

    create_at timestamptz not null default now(),
    update_at timestamptz not null default now()
);
