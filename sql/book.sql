drop table if exists book;
drop SEQUENCE if exists book_id_seq ;
create SEQUENCE book_id_seq;

--标准图书表
create table book (
    id text primary key not null default 'book_'||to_char(now() AT TIME ZONE 'cct', 'yymmdd') ||trim(to_char(nextval('book_id_seq'), '00000000')), --book id
    isbn            text            not null,       --图书isbn
    book_no         text            default '00',     --如果存在一isbn多书情况，需要分配图书编号
    book_cate       text            default '',     --用于区分一isbn多本书的情况 枚举类型
    title           text            default '',     --标题
    publisher       text            default '',     --出版社
    author          text            default '',     --作者
    edition         text            default '',     --版次
    pubdate         text            default '',     --出版日期
    series_name     text            default '',     --丛书名  eg:北京大学数学教学系列丛书
    image           text            default '',     --图书图片
    price           int             default 0,      --价格
    catalog         text            default '',     --目录
    abstract        text            default '',     --内容简介
    page            text            default '',     --图书页数
    packing         text            default '',     --包装  eg:平装
    format          text            default '',     --版式  eg:32开
    author_intro    text            default '',     --作者介绍
    source_info     text            default '',     --数据源   caiku dangdang bookuu jd
    search_time     int             not null default 0,
    taobao_category text            default '',     --淘宝
    create_at timestamptz not null default now(),
    update_at timestamptz not null default now()
);


CREATE UNIQUE INDEX IF NOT EXISTS book_isbn_no ON  book(isbn,book_no);
CREATE INDEX IF NOT EXISTS book_isbn ON  book(isbn);
CREATE INDEX IF NOT EXISTS book_title ON  book(title);
