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

COMMENT ON SCHEMA public IS 'standard public schema';


--
-- TOC entry 6 (class 2615 OID 16385)
-- Name: usr; Type: SCHEMA; Schema: -; Owner: adapticcuser
--

CREATE SCHEMA usr;


ALTER SCHEMA usr OWNER TO adapticcuser;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 220 (class 1259 OID 16412)
-- Name: products; Type: TABLE; Schema: usr; Owner: adapticcuser
--

CREATE TABLE usr.products (
    id character varying NOT NULL,
    title character varying NOT NULL,
    description character varying NOT NULL,
    creation_user_id character varying,
    update_user_id character varying,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp
);


ALTER TABLE usr.products OWNER TO adapticcuser;

--
-- TOC entry 216 (class 1259 OID 16391)
-- Name: roles; Type: TABLE; Schema: usr; Owner: adapticcuser
--

CREATE TABLE usr.roles (
    id character varying,
    code character varying,
    name character varying,
    creation_user_id character varying,
    update_user_id character varying,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp
);


ALTER TABLE usr.roles OWNER TO adapticcuser;

--
-- TOC entry 218 (class 1259 OID 16401)
-- Name: user_role; Type: TABLE; Schema: usr; Owner: adapticcuser
--

CREATE TABLE usr.user_role (
    user_id character varying,
    role_id character varying,
    creation_user_id character varying,
    update_user_id character varying,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp
);


ALTER TABLE usr.user_role OWNER TO adapticcuser;

--
-- TOC entry 219 (class 1259 OID 16406)
-- Name: users; Type: TABLE; Schema: usr; Owner: adapticcuser
--

CREATE TABLE usr.users (
    id character varying NOT NULL,
    email character varying NOT NULL,
    first_name character varying NOT NULL,
    last_name character varying NOT NULL,
    password_hash character varying,
    active boolean DEFAULT false NOT NULL,
    creation_user_id character varying,
    update_user_id character varying,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp
);


ALTER TABLE usr.users OWNER TO adapticcuser;

--
-- TOC entry 3352 (class 0 OID 16412)
-- Dependencies: 220
-- Data for Name: products; Type: TABLE DATA; Schema: usr; Owner: adapticcuser
--

INSERT INTO usr.products VALUES ('b9bee884-0e54-4337-a410-d28865e2789b', 'Test Product', 'Product Description....', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);
INSERT INTO usr.products VALUES ('2c852054-5468-410c-9cbb-5b7a012b58ed', 'Test Product 1', 'Product Description....', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);
INSERT INTO usr.products VALUES ('96822a8f-7416-4a08-b00c-8f67421d9911', 'Test Product 2', 'Product Description....', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);
INSERT INTO usr.products VALUES ('e3dbb5f5-fd9c-4e0f-b3f9-00bf7c758b34', 'Test Product 3', 'Product Description....', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);
INSERT INTO usr.products VALUES ('a8ad7ef1-7664-46cc-883c-5e0ee002067a', 'Test Product 4', 'Product Description....', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);
INSERT INTO usr.products VALUES ('9a73d8f1-1c01-439f-ab07-0c2858ba413a', 'Test Product 6', 'Product Description....', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);
INSERT INTO usr.products VALUES ('7046148d-25b5-45a0-ab13-95206db1b540', 'Test Product 7', 'Product Description....', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);
INSERT INTO usr.products VALUES ('6031f411-46ef-4b81-a744-26ce9ae73ee4', 'Test Product 8', 'Product Description....', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);
INSERT INTO usr.products VALUES ('f5cb161b-0c16-49de-8a5b-8528e3906d0a', 'Test Product 9', 'Product Description....', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);
INSERT INTO usr.products VALUES ('d6606d36-d358-4e7a-a420-0de355b1468b', 'Test Product 10', 'Product Description....', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);
INSERT INTO usr.products VALUES ('b93ab7e4-3551-4a27-a60c-9163ac968208', 'Test Product 12', 'Product Description....', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);
INSERT INTO usr.products VALUES ('e35de786-bd47-404b-b92d-636b5e553596', 'Test Product 13', 'Product Description....', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);
INSERT INTO usr.products VALUES ('91227c6f-f5da-406e-ba7d-e19e0e92bb85', 'Test Product 14', 'Product Description....', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);
INSERT INTO usr.products VALUES ('fcaa12c2-9ee3-47d7-8be0-1e42c0bb8f94', 'Test Product 15', 'Product Description....', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);
INSERT INTO usr.products VALUES ('e1b4da09-ee9a-40ea-828a-8d461c9adf79', 'Test Product 16', 'Product Description....', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);
INSERT INTO usr.products VALUES ('cfdfebf6-5db4-4f1f-84ed-00ffbc6d10fd', 'Test Product 17', 'Product Description....', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);
INSERT INTO usr.products VALUES ('5b6f3413-0af0-4f8f-85b4-af02738e4da6', 'Test Product 19', 'Product Description....', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);
INSERT INTO usr.products VALUES ('0ecdbc3f-ee4e-42f6-831a-217e54b65722', 'Test Product 33', 'Product Description test 1', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);
INSERT INTO usr.products VALUES ('3d0cd537-2de1-4474-9ef2-a0d4581dd407', 'Test Product 36', 'Product Description test 3', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);



--
-- TOC entry 3348 (class 0 OID 16391)
-- Dependencies: 216
-- Data for Name: roles; Type: TABLE DATA; Schema: usr; Owner: adapticcuser
--

INSERT INTO usr.roles VALUES ('57010e91-b768-49ca-8cac-e681f15ccc0f', 'product:editor', 'Product Editor', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);
INSERT INTO usr.roles VALUES ('f55b268a-9940-4d61-99d9-47a6816e1390', 'product:viewer', 'Product Viewer', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);
INSERT INTO usr.roles VALUES ('3496e6e7-b56e-4595-9a73-94f7c6ad2bc3', 'user:editor', 'User Editor', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);
INSERT INTO usr.roles VALUES ('c5d0f7a9-53fe-450a-bbb9-94d5e12b536e', 'role:editor', 'Role Editor', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);

--
-- TOC entry 3350 (class 0 OID 16401)
-- Dependencies: 218
-- Data for Name: user_role; Type: TABLE DATA; Schema: usr; Owner: adapticcuser
--

INSERT INTO usr.user_role VALUES ('8f9b1e8f-d496-4804-942b-5ea29050370b', 'f55b268a-9940-4d61-99d9-47a6816e1390', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);
INSERT INTO usr.user_role VALUES ('613254df-c779-479c-9d76-b8036e342979', '57010e91-b768-49ca-8cac-e681f15ccc0f', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);
INSERT INTO usr.user_role VALUES ('613254df-c779-479c-9d76-b8036e342979', '3496e6e7-b56e-4595-9a73-94f7c6ad2bc3', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);
INSERT INTO usr.user_role VALUES ('613254df-c779-479c-9d76-b8036e342979', 'c5d0f7a9-53fe-450a-bbb9-94d5e12b536e', '613254df-c779-479c-9d76-b8036e342979', NULL, '2023-07-28 13:23:03.273735', NULL);


--
-- TOC entry 3351 (class 0 OID 16406)
-- Dependencies: 219
-- Data for Name: users; Type: TABLE DATA; Schema: usr; Owner: adapticcuser
--

INSERT INTO usr.users VALUES ('613254df-c779-479c-9d76-b8036e342979', 'admin@admin.com', 'John', 'Jones', '$2a$10$1UZWtyK2f6BvSlqp6SBzkeGiTP5pdkiuRgvlt4Gd4MZIyfpVWCkYq', true, NULL, NULL, '2023-07-28 13:23:03.273735', NULL);
INSERT INTO usr.users VALUES ('8f9b1e8f-d496-4804-942b-5ea29050370b', 'test@test.com', 'Tester', 'Test', '$2a$10$1UZWtyK2f6BvSlqp6SBzkeGiTP5pdkiuRgvlt4Gd4MZIyfpVWCkYq', true, NULL, NULL, '2023-07-28 22:15:07.79185', NULL);


--
-- TOC entry 3204 (class 2606 OID 16424)
-- Name: products products_pkey; Type: CONSTRAINT; Schema: usr; Owner: adapticcuser
--

ALTER TABLE ONLY usr.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- TOC entry 3198 (class 2606 OID 16418)
-- Name: users uq_email; Type: CONSTRAINT; Schema: usr; Owner: adapticcuser
--

ALTER TABLE ONLY usr.users
    ADD CONSTRAINT uq_email UNIQUE (email);


--
-- TOC entry 3200 (class 2606 OID 16420)
-- Name: users users_email_key; Type: CONSTRAINT; Schema: usr; Owner: adapticcuser
--

ALTER TABLE ONLY usr.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- TOC entry 3202 (class 2606 OID 16422)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: usr; Owner: adapticcuser
--

ALTER TABLE ONLY usr.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


-- Completed on 2023-07-29 01:34:55 CEST

--
-- PostgreSQL database dump complete
--