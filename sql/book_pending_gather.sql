
--标准图书表
create table book_pending_gather (
    id                      uuid primary key        default gen_random_uuid(),
    book_id                 text                    not null,
    search_time             int                     not null default 0,
    create_at               timestamptz             not null default now(),
    update_at               timestamptz             not null default now()
);

CREATE UNIQUE INDEX IF NOT EXISTS book_pending_gather_book_id ON  book_pending_gather(book_id);
