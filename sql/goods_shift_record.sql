
create table goods_shift_record (
    id UUID         PRIMARY KEY         NOT NULL        DEFAULT gen_random_uuid(),
    goods_id        UUID                not null ,
    location_id     UUID                not null ,
    warehouse       text                not null ,
    shelf           text                not null ,
    floor           text                not null ,
    user_id         text                not null ,
    stock           int                 not null ,
    operate_type    text                not null ,     --load or unload
    create_at       timestamptz not null default now(),
    update_at       timestamptz not null default now()
);
