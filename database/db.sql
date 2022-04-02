
use banking;
create table customers(
    customer_id int(11)not null auto_increment,
    name varchar(100)not null,
    date_of_birth date not null,
    city varchar(100) not null,
    zipcode varchar(10)not null,
    status tinyint(4)not null default "1",    primary key(customer_id)
);

create table accounts(
    account_id int(11) not null auto_increment,
    customer_id int(11) not null,
    openling_date datetime not null default current_timestamp,
    account_type varchar(10) not null,
    pin varchar(10) not null,
    status tinyint(4) not null default 1,
    primary key (account_id),
  
   foreign key (customer_id) references customers(customer_id)
);

create table transactions(
    transaction_id int(11) not null auto_increment,
    account_id int(11) not null,
    amount int(11) not null,
    transaction_type varchar(10) not null,
    transaction_date datetime not null default current_timestamp,
    primary key(transaction_id),
  
    foreign key (account_id) references accounts(account_id)
)

