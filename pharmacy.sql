create table permission (
    id integer primary key generated always as identity,
    name varchar not null,
    description text,
    created_at timestamptz default now()
);

create table role (
    id integer primary key generated always as identity,
    name varchar not null,
    created_at timestamptz default now()
);

create table account (
    id integer primary key generated always as identity,
    first_name varchar not null,
    last_name varchar not null,
    email varchar,
    password varchar not null,
    created_at timestamptz default now(),
    role_id integer references role (id)
);

create table role_permission (
    role_id integer references role (id) on update cascade on delete cascade,
    permission_id integer references permission (id) on update cascade on delete cascade,
    CONSTRAINT role_permission_id PRIMARY KEY (role_id, permission_id)
);

create table category (
    id integer primary key generated always as identity,
    name varchar not null,
    description text,
    created_at timestamptz default now(),
    user_id integer references account (id)
);

create table supplier (
    id integer primary key generated always as identity,
    name varchar not null,
    address text not null,
    phone_number varchar not null,
    email varchar,
    created_at timestamptz default now()
);

create table product (
    id integer primary key generated always as identity,
    name varchar not null,
    bar_code varchar not null,
    description text,
    category_id integer null references category (id) on update cascade on delete set null,
    purchase_date timestamptz not null,
    production_date timestamptz not null,
    expiry_date timestamptz not null,
    purchase_price numeric not null,
    selling_price numeric not null,
    quantity_available integer,
    reorder_level integer not null,
    sku varchar,
    quantity_sold integer,
    user_id integer null references account (id) on update cascade on delete set null,
    created_at timestamptz default now()
);

create table product_supplier (
    product_id integer references product (id) on update cascade on delete cascade,
    supplier_id integer references supplier (id) on update cascade on delete cascade,
    CONSTRAINT product_supplier_pkey primary key (product_id, supplier_id)
);

create table payment_method (
    id integer primary key generated always as identity,
    name varchar not null,
    is_active bool default true,
    user_id integer references account (id),
    created_at timestamptz default now()
);

create table "order" (
    id integer primary key generated always as identity,
    total_price numeric,
    payment_method_id integer references payment_method (id),
    payment_verified bool,
    cashier_id integer references account (id),
    amount_tendered numeric,
    change numeric,
    created_at timestamptz default now()
);

create table return (
    id integer primary key generated always as identity,
    reason text,
    order_id integer references "order" (id),
    user_id integer references account (id),
    created_at timestamptz default now()
);

create table order_item (
    id integer primary key generated always as identity,
    product_id integer references product (id) on update cascade,
    order_id integer references "order" (id) on update cascade on delete cascade,
    quantity integer not null,
    return_id integer null references return (id) on update cascade on delete set null
);