create table posts (
    id serial not null unique,
    title varchar(64),
    content text,
    primary key(id)
);

insert into posts(title, content)
values
    ('Hello world', 'First hello world post');
    ('Another Post', 'Another blog post');