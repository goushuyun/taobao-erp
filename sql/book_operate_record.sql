drop table if exists book_operate_record;
--图书操作记录
create table book_operate_record (
    id                      uuid primary key        default gen_random_uuid(),
    book_id                 text                    not null,
    user_id                 text                    not null,
    operate_content         text                    not null,
    create_at               timestamptz             not null default now(),
    update_at               timestamptz             not null default now()
);
