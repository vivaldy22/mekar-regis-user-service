create database mekar_regis_db;
use mekar_regis_db;
create table tb_admin (
    admin_id VARCHAR(36) PRIMARY KEY NOT NULL,
    admin_uname VARCHAR(100) NOT NULL,
    admin_pass VARCHAR(255) NOT NULL
);
create table tb_user (
    user_id VARCHAR(36) PRIMARY KEY NOT NULL,
    user_name VARCHAR(100) NOT NULL,
    user_bday DATE NOT NULL,
    user_ktp VARCHAR(16) NOT NULL,
    user_job VARCHAR(36) NOT NULL,
    user_edu VARCHAR(36) NOT NULL,
    user_status INT NOT NULL
);
create table tb_job (
    job_id VARCHAR(36) PRIMARY KEY NOT NULL,
    job_name VARCHAR(100) NOT NULL
);
create table tb_edu (
    edu_id VARCHAR(36) PRIMARY KEY NOT NULL,
    edu_name VARCHAR(100) NOT NULL
);
insert into tb_admin
values (UUID(), 'admin', '$2a$04$2ehoD827XT6Oz0rv9X88S.iJUOpUOx5mU5UQCkLNBXjdWoVvHnDh2');
insert into tb_user
values (UUID(), 'Vivaldy Andhira', '1997-05-22', '1234567891011121', 'fc538cf5-3227-11eb-835d-7085c2a10d57', 'ce374139-3228-11eb-835d-7085c2a10d57', 1);
insert into tb_job
values ('fc5384f7-3227-11eb-835d-7085c2a10d57', 'PNS'),
    (
        'fc538b83-3227-11eb-835d-7085c2a10d57',
        'Wirausaha'
    ),
    (
        'fc538cf5-3227-11eb-835d-7085c2a10d57',
        'Wiraswasta'
    );
insert into tb_edu
values ('ce372226-3228-11eb-835d-7085c2a10d57', 'SD'),
    ('ce373e4a-3228-11eb-835d-7085c2a10d57', 'SMP'),
    ('ce373fbc-3228-11eb-835d-7085c2a10d57', 'SMA'),
    (
        'ce374082-3228-11eb-835d-7085c2a10d57',
        'Diploma'
    ),
    (
        'ce374139-3228-11eb-835d-7085c2a10d57',
        'Sarjana'
    ),
    (
        'ce37420c-3228-11eb-835d-7085c2a10d57',
        'Magister'
    ),
    ('ce37429a-3228-11eb-835d-7085c2a10d57', 'Doktor');