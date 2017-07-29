create SEQUENCE user_id_seq;

create table users (
    "id" text primary key not null default to_char(now() AT TIME ZONE 'cct', 'yymmdd') || trim(to_char(nextval('user_id_seq'), '00000')),
    tel text not null,
    password text not null,
    role int not null
);
