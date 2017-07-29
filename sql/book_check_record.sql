drop table if exists book_check_record;

--标准图书表
create table book_check_record (
    id              uuid primary key default gen_random_uuid(),
    book_id         text            not null,       --图书isbn
    title           text            default '',     --标题
    publisher       text            default '',     --出版社
    author          text            default '',     --作者
    edition         text            default '',     --版次
    image           text            default '',     --图书图片
    price           int             default 0,      --价格
    apply_user_id   text            not null,       --申请人
    check_user_id   text            default '',     --核查人
    apply_reason    text            default '',     --申请原因
    status          int             default 1,      --状态       1:审核中    2:成功   3:失败
    feedback        text            default '',     --申请反馈
    create_at timestamptz not null default now(),
    update_at timestamptz not null default now()
);
