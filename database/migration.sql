create tabel 'users'
(
    id  bigInteger auto_increment,
    name varchar(255) not null,
    primary key('id')
);
insert into users ('name') values( 'Ram'),('Sam'),('Siva');
