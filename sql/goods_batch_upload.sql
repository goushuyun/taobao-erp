create sequence goods_batch_upload_id_seq;

create table goods_batch_upload(

    id text primary key not null default trim(to_char(nextval('goods_batch_upload_id_seq'),'000000000')),--id
    user_id        text            not null,                                       --云店铺id
    success_num     int             not null            default 0,                 --成功数量
    failed_num      int             not null            default 0,                 --失败数量
    origin_file     text            not null            default '',                --商家源文件
    origin_filename text            not null,
    error_file      text            not null            default '',                --处理错误文件
    create_at       timestamptz     not null            default now(),             --创建时间
    update_at       timestamptz     not null            default now()              --更改时间
);
