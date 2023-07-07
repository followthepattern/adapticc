--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2 (Debian 15.2-1.pgdg110+1)
-- Dumped by pg_dump version 15.3

-- Started on 2023-06-17 01:46:17 CEST

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 7 (class 2615 OID 16386)
-- Name: usr; Type: SCHEMA; Schema: -; Owner: dbuser
--

CREATE SCHEMA usr;


ALTER SCHEMA usr OWNER TO dbuser;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 224 (class 1259 OID 16441)
-- Name: role_resource_permissions; Type: TABLE; Schema: usr; Owner: dbuser
--

CREATE TABLE usr.role_resource_permissions (
    role_id character varying NOT NULL,
    resource_id character varying NOT NULL,
    permission integer NOT NULL
);


ALTER TABLE usr.role_resource_permissions OWNER TO dbuser;

--
-- TOC entry 225 (class 1259 OID 16446)
-- Name: roles; Type: TABLE; Schema: usr; Owner: dbuser
--

CREATE TABLE usr.roles (
    id character varying,
    name character varying
);


ALTER TABLE usr.roles OWNER TO dbuser;

--
-- TOC entry 223 (class 1259 OID 16436)
-- Name: user_resource_permissions; Type: TABLE; Schema: usr; Owner: dbuser
--

CREATE TABLE usr.user_resource_permissions (
    user_id character varying NOT NULL,
    resource_id character varying NOT NULL,
    permission integer NOT NULL
);


ALTER TABLE usr.user_resource_permissions OWNER TO dbuser;

--
-- TOC entry 226 (class 1259 OID 16451)
-- Name: user_role; Type: TABLE; Schema: usr; Owner: dbuser
--

CREATE TABLE usr.user_role (
    user_id character varying,
    role_id character varying
);


ALTER TABLE usr.user_role OWNER TO dbuser;

--
-- TOC entry 222 (class 1259 OID 16430)
-- Name: users; Type: TABLE; Schema: usr; Owner: dbuser
--

CREATE TABLE usr.users (
    id character varying NOT NULL,
    email character varying NOT NULL,
    first_name character varying NOT NULL,
    last_name character varying NOT NULL,
    password character varying NOT NULL,
    salt character varying NOT NULL,
    active boolean DEFAULT false NOT NULL,
    registered_at timestamp without time zone NOT NULL
);


ALTER TABLE usr.users OWNER TO dbuser;

--
-- TOC entry 222 (class 1259 OID 16430)
-- Name: products; Type: TABLE; Schema: usr; Owner: dbuser
--

CREATE TABLE usr.products (
    id character varying NOT NULL,
    title character varying NOT NULL,
    description character varying NOT NULL
);


ALTER TABLE usr.products OWNER TO dbuser;

--
-- TOC entry 3398 (class 0 OID 16441)
-- Dependencies: 224
-- Data for Name: role_resource_permissions; Type: TABLE DATA; Schema: usr; Owner: dbuser
--

INSERT INTO usr.role_resource_permissions VALUES ('2b87e6c6-9183-4ee6-ba04-14d3be3f739c', 'PRODUCT', 1);
INSERT INTO usr.role_resource_permissions VALUES ('9606f479-a600-4b71-8042-1eddedb448e', 'PRODUCT', 2);
INSERT INTO usr.role_resource_permissions VALUES ('f5d76752-34a7-41d3-b507-7e0b429574c7', 'PRODUCT', 14);

--
-- TOC entry 3399 (class 0 OID 16446)
-- Dependencies: 225
-- Data for Name: roles; Type: TABLE DATA; Schema: usr; Owner: dbuser
--

INSERT INTO usr.roles VALUES ('f3de1ecb-1e43-4266-ac36-a725bf7b587a', 'Product Creator');
INSERT INTO usr.roles VALUES ('9606f479-a600-4b71-8042-1eddedb448e9', 'Product Reader');
INSERT INTO usr.roles VALUES ('f5d76752-34a7-41d3-b507-7e0b429574c7', 'Product Editor');

--
-- TOC entry 3400 (class 0 OID 16451)
-- Dependencies: 226
-- Data for Name: user_role; Type: TABLE DATA; Schema: usr; Owner: dbuser
--

INSERT INTO usr.user_role VALUES ('2b87e6c6-9183-4ee6-ba04-14d3be3f739c', 'f3de1ecb-1e43-4266-ac36-a725bf7b587a');
INSERT INTO usr.user_role VALUES ('2b87e6c6-9183-4ee6-ba04-14d3be3f739c', 'f5d76752-34a7-41d3-b507-7e0b429574c7');
INSERT INTO usr.user_role VALUES ('9606f479-a600-4b71-8042-1eddedb448e9', '5dde1316-8e1f-4a19-819b-219df533d274');
INSERT INTO usr.user_role VALUES ('9020a1b8-7dbd-4e53-b3b1-9b804645623f', 'f5d76752-34a7-41d3-b507-7e0b429574c7');


--
-- TOC entry 3396 (class 0 OID 16430)
-- Dependencies: 222
-- Data for Name: users; Type: TABLE DATA; Schema: usr; Owner: dbuser
--

INSERT INTO usr.users VALUES ('9020a1b8-7dbd-4e53-b3b1-9b804645623f', 'admin@admin.hu', 'Csaba', 'Huszka', '06acf1862e69df3cb9a27c9bfa5f6ae11e572214e79ac7d93c87aefd839d164c', 'JaRp2X+PDS9TvcnRVNaCDvWOmxWnyClU4LvkTLJGK+Q=', false, '2021-05-23 10:00:00');

--
-- TOC entry 3396 (class 0 OID 16430)
-- Dependencies: 222
-- Data for Name: users; Type: TABLE DATA; Schema: usr; Owner: dbuser
--

INSERT INTO usr.products VALUES ('b9bee884-0e54-4337-a410-d28865e2789b', 'Test Product', 'Product Description....');
INSERT INTO usr.products VALUES ('2c852054-5468-410c-9cbb-5b7a012b58ed', 'Test Product 1', 'Product Description....');
INSERT INTO usr.products VALUES ('96822a8f-7416-4a08-b00c-8f67421d9911', 'Test Product 2', 'Product Description....');
INSERT INTO usr.products VALUES ('e3dbb5f5-fd9c-4e0f-b3f9-00bf7c758b34', 'Test Product 3', 'Product Description....');
INSERT INTO usr.products VALUES ('a8ad7ef1-7664-46cc-883c-5e0ee002067a', 'Test Product 4', 'Product Description....');
INSERT INTO usr.products VALUES ('1c9b4c2a-7bb4-43e8-9739-64b0c5188685', 'Test Product 5', 'Product Description....');
INSERT INTO usr.products VALUES ('9a73d8f1-1c01-439f-ab07-0c2858ba413a', 'Test Product 6', 'Product Description....');
INSERT INTO usr.products VALUES ('7046148d-25b5-45a0-ab13-95206db1b540', 'Test Product 7', 'Product Description....');
INSERT INTO usr.products VALUES ('6031f411-46ef-4b81-a744-26ce9ae73ee4', 'Test Product 8', 'Product Description....');
INSERT INTO usr.products VALUES ('f5cb161b-0c16-49de-8a5b-8528e3906d0a', 'Test Product 9', 'Product Description....');
INSERT INTO usr.products VALUES ('d6606d36-d358-4e7a-a420-0de355b1468b', 'Test Product 10', 'Product Description....');
INSERT INTO usr.products VALUES ('0ba62228-e5f3-46ee-83e2-311b4664a9dd', 'Test Product 11', 'Product Description....');
INSERT INTO usr.products VALUES ('b93ab7e4-3551-4a27-a60c-9163ac968208', 'Test Product 12', 'Product Description....');
INSERT INTO usr.products VALUES ('e35de786-bd47-404b-b92d-636b5e553596', 'Test Product 13', 'Product Description....');
INSERT INTO usr.products VALUES ('91227c6f-f5da-406e-ba7d-e19e0e92bb85', 'Test Product 14', 'Product Description....');
INSERT INTO usr.products VALUES ('fcaa12c2-9ee3-47d7-8be0-1e42c0bb8f94', 'Test Product 15', 'Product Description....');
INSERT INTO usr.products VALUES ('e1b4da09-ee9a-40ea-828a-8d461c9adf79', 'Test Product 16', 'Product Description....');
INSERT INTO usr.products VALUES ('cfdfebf6-5db4-4f1f-84ed-00ffbc6d10fd', 'Test Product 17', 'Product Description....');
INSERT INTO usr.products VALUES ('0ecdbc3f-ee4e-42f6-831a-217e54b65722', 'Test Product 18', 'Product Description....');
INSERT INTO usr.products VALUES ('5b6f3413-0af0-4f8f-85b4-af02738e4da6', 'Test Product 19', 'Product Description....');


--
-- TOC entry 3239 (class 2606 OID 16461)
-- Name: users uq_email; Type: CONSTRAINT; Schema: usr; Owner: dbuser
--

ALTER TABLE ONLY usr.users
    ADD CONSTRAINT uq_email UNIQUE (email);

--
-- TOC entry 3241 (class 2606 OID 16467)
-- Name: users users_email_key; Type: CONSTRAINT; Schema: usr; Owner: dbuser
--

ALTER TABLE ONLY usr.users
    ADD CONSTRAINT users_email_key UNIQUE (email);

--
-- TOC entry 3243 (class 2606 OID 16469)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: usr; Owner: dbuser
--

ALTER TABLE ONLY usr.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);

--
-- TOC entry 3243 (class 2606 OID 16469)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: usr; Owner: dbuser
--

ALTER TABLE ONLY usr.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);