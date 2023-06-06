--
-- PostgreSQL database dump
--

-- Dumped from database version 15.3
-- Dumped by pg_dump version 15.3

-- Started on 2023-06-04 20:57:49

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 214 (class 1259 OID 16407)
-- Name: urls; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.urls (
    short text NOT NULL PRIMARY KEY,
    original text NOT NULL
);


ALTER TABLE public.urls OWNER TO postgres;

--
-- TOC entry 3316 (class 0 OID 16407)
-- Dependencies: 214
-- Data for Name: urls; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.urls (short, original) FROM stdin;
shwDg8KWly	http://lenta.ru
BfUic0chj2	http://lenta.ru
NR1hrEWc3S	https://joyreactor.cc/new/507498
zT4aSvhkZg	https://joyreactor.cc/new/507498
rx8kzGbsd6	https://joyreactor.cc/new/507498
_7aDHm6UiY	https://joyreactor.cc/new/507498
RCgEPSKoBi	https://joyreactor.cc/new/507498
M2lGptFGkI	https://joyreactor.cc/new/507498
mvyE_kJ0Jq	https://joyreactor.cc/new/507498
zELuYHykBC	https://stackoverflow.com/questions/40096750/how-to-set-http-status-code-on-http-responsewriter
nkk7MpojDR	https://joyreactor.cc/new/507498
OH2uYhZ736	http://google.com
VioTdmfqkN	http://lenta.ru
5QSyN10FOA	https://joyreactor.cc/new/507498
b3_3bNmKI0	https://joyreactor.cc/new/507498
ksYZbuxESU	https://joyreactor.cc/new/507498
\.


--
-- TOC entry 3173 (class 2606 OID 16413)
-- Name: urls urls_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.urls
    ADD CONSTRAINT urls_pkey PRIMARY KEY (short);


-- Completed on 2023-06-04 20:57:49

--
-- PostgreSQL database dump complete
--

