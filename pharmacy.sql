create table permission
(
    id integer generated always as identity,
    name varchar not null,
    description text,
    created_at timestamp with time zone default now(),
    constraint permission_pkey
        primary key (id)
);

create table role
(
    id integer generated always as identity,
    name varchar not null,
    created_at timestamp with time zone default now(),
    constraint role_pkey
        primary key (id)
);

create table account
(
    id integer generated always as identity,
    first_name varchar not null,
    last_name varchar not null,
    email varchar,
    password varchar not null,
    created_at timestamp with time zone default now(),
    role_id integer,
    phone_number varchar,
    constraint account_pkey
        primary key (id),
    constraint account_role_id_fkey
        foreign key (role_id) references role
);

create unique index account_phone_number_uindex
    on account (phone_number);

create table role_permission
(
    role_id integer not null,
    permission_id integer not null,
    constraint role_permission_id
        primary key (role_id, permission_id),
    constraint role_permission_role_id_fkey
        foreign key (role_id) references role
            on update cascade on delete cascade,
    constraint role_permission_permission_id_fkey
        foreign key (permission_id) references permission
            on update cascade on delete cascade
);

create table category
(
    id integer generated always as identity,
    name varchar not null,
    description text,
    created_at timestamp with time zone default now(),
    user_id integer,
    constraint category_pkey
        primary key (id),
    constraint category_user_id_fkey
        foreign key (user_id) references account
);

create table supplier
(
    id integer generated always as identity,
    name varchar not null,
    address text not null,
    phone_number varchar not null,
    email varchar,
    created_at timestamp with time zone default now(),
    constraint supplier_pkey
        primary key (id)
);

create table product
(
    id integer generated always as identity,
    name varchar not null,
    bar_code varchar not null,
    description text,
    category_id integer,
    purchase_date timestamp with time zone not null,
    production_date timestamp with time zone not null,
    expiry_date timestamp with time zone not null,
    purchase_price numeric not null,
    selling_price numeric not null,
    quantity_available integer,
    reorder_level integer not null,
    sku varchar,
    quantity_sold integer,
    user_id integer,
    created_at timestamp with time zone default now(),
    constraint product_pkey
        primary key (id),
    constraint product_category_id_fkey
        foreign key (category_id) references category
            on update cascade on delete set null,
    constraint product_user_id_fkey
        foreign key (user_id) references account
            on update cascade on delete set null
);

create table product_supplier
(
    product_id integer not null,
    supplier_id integer not null,
    constraint product_supplier_pkey
        primary key (product_id, supplier_id),
    constraint product_supplier_product_id_fkey
        foreign key (product_id) references product
            on update cascade on delete cascade,
    constraint product_supplier_supplier_id_fkey
        foreign key (supplier_id) references supplier
            on update cascade on delete cascade
);

create table payment_method
(
    id integer generated always as identity,
    name varchar not null,
    is_active boolean default true,
    user_id integer,
    created_at timestamp with time zone default now(),
    constraint payment_method_pkey
        primary key (id),
    constraint payment_method_user_id_fkey
        foreign key (user_id) references account
);

create table "order"
(
    id integer generated always as identity,
    total_price numeric,
    payment_method_id integer,
    payment_verified boolean,
    cashier_id integer,
    amount_tendered numeric,
    change numeric,
    created_at timestamp with time zone default now(),
    constraint order_pkey
        primary key (id),
    constraint order_payment_method_id_fkey
        foreign key (payment_method_id) references payment_method,
    constraint order_cashier_id_fkey
        foreign key (cashier_id) references account
);

create table return
(
    id integer generated always as identity,
    reason text,
    order_id integer,
    user_id integer,
    created_at timestamp with time zone default now(),
    constraint return_pkey
        primary key (id),
    constraint return_order_id_fkey
        foreign key (order_id) references "order",
    constraint return_user_id_fkey
        foreign key (user_id) references account
);

create table order_item
(
    id integer generated always as identity,
    product_id integer,
    order_id integer,
    quantity integer not null,
    return_id integer,
    constraint order_item_pkey
        primary key (id),
    constraint order_item_product_id_fkey
        foreign key (product_id) references product
            on update cascade,
    constraint order_item_order_id_fkey
        foreign key (order_id) references "order"
            on update cascade on delete cascade,
    constraint order_item_return_id_fkey
        foreign key (return_id) references return
            on update cascade on delete set null
);

