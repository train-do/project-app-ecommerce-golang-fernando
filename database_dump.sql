--
-- PostgreSQL database dump
--

-- Dumped from database version 16rc1
-- Dumped by pg_dump version 16rc1

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
-- Name: Address; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Address" (
    id integer NOT NULL,
    user_id integer NOT NULL,
    name character varying NOT NULL,
    email character varying NOT NULL,
    address character varying NOT NULL,
    is_default boolean DEFAULT false NOT NULL
);


ALTER TABLE public."Address" OWNER TO postgres;

--
-- Name: Address_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Address_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public."Address_id_seq" OWNER TO postgres;

--
-- Name: Address_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Address_id_seq" OWNED BY public."Address".id;


--
-- Name: Banner; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Banner" (
    id integer NOT NULL,
    banner_url character varying NOT NULL,
    title character varying NOT NULL,
    subtitle character varying NOT NULL,
    path_page character varying NOT NULL,
    start_time date NOT NULL,
    end_time date NOT NULL
);


ALTER TABLE public."Banner" OWNER TO postgres;

--
-- Name: Banner_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Banner_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public."Banner_id_seq" OWNER TO postgres;

--
-- Name: Banner_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Banner_id_seq" OWNED BY public."Banner".id;


--
-- Name: Cart; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Cart" (
    id integer NOT NULL,
    user_id integer NOT NULL,
    product_variant_id integer NOT NULL,
    qty integer NOT NULL
);


ALTER TABLE public."Cart" OWNER TO postgres;

--
-- Name: Cart_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Cart_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public."Cart_id_seq" OWNER TO postgres;

--
-- Name: Cart_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Cart_id_seq" OWNED BY public."Cart".id;


--
-- Name: Category; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Category" (
    id integer NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public."Category" OWNER TO postgres;

--
-- Name: Category_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Category_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public."Category_id_seq" OWNER TO postgres;

--
-- Name: Category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Category_id_seq" OWNED BY public."Category".id;


--
-- Name: Gallery; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Gallery" (
    id integer NOT NULL,
    product_id integer NOT NULL,
    image_url character varying NOT NULL
);


ALTER TABLE public."Gallery" OWNER TO postgres;

--
-- Name: Gallery_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Gallery_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public."Gallery_id_seq" OWNER TO postgres;

--
-- Name: Gallery_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Gallery_id_seq" OWNED BY public."Gallery".id;


--
-- Name: Order; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Order" (
    id integer NOT NULL,
    user_id integer NOT NULL,
    address_id integer NOT NULL,
    total_price double precision NOT NULL,
    shipping_price integer DEFAULT 0 NOT NULL,
    status character varying DEFAULT 'pay'::character varying NOT NULL,
    payment_method character varying DEFAULT 'bank'::character varying NOT NULL,
    created_at date NOT NULL
);


ALTER TABLE public."Order" OWNER TO postgres;

--
-- Name: OrderProduct; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."OrderProduct" (
    id integer NOT NULL,
    order_id integer,
    product_variant_id integer NOT NULL,
    qty integer NOT NULL,
    sub_total double precision NOT NULL,
    created_at date NOT NULL
);


ALTER TABLE public."OrderProduct" OWNER TO postgres;

--
-- Name: OrderProduct_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."OrderProduct_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public."OrderProduct_id_seq" OWNER TO postgres;

--
-- Name: OrderProduct_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."OrderProduct_id_seq" OWNED BY public."OrderProduct".id;


--
-- Name: Order_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Order_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public."Order_id_seq" OWNER TO postgres;

--
-- Name: Order_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Order_id_seq" OWNED BY public."Order".id;


--
-- Name: Product; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Product" (
    id integer NOT NULL,
    category_id integer NOT NULL,
    name character varying NOT NULL,
    description character varying NOT NULL,
    price double precision NOT NULL,
    discount integer NOT NULL,
    price_after_discount double precision NOT NULL,
    rating numeric(2,1) NOT NULL,
    qty_sold integer NOT NULL,
    is_best_selling boolean DEFAULT false NOT NULL,
    created_at date NOT NULL
);


ALTER TABLE public."Product" OWNER TO postgres;

--
-- Name: ProductVariant; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."ProductVariant" (
    id integer NOT NULL,
    product_id integer NOT NULL,
    variant_id integer NOT NULL,
    stock integer NOT NULL
);


ALTER TABLE public."ProductVariant" OWNER TO postgres;

--
-- Name: ProductVariant_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."ProductVariant_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public."ProductVariant_id_seq" OWNER TO postgres;

--
-- Name: ProductVariant_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."ProductVariant_id_seq" OWNED BY public."ProductVariant".id;


--
-- Name: Product_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Product_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public."Product_id_seq" OWNER TO postgres;

--
-- Name: Product_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Product_id_seq" OWNED BY public."Product".id;


--
-- Name: Promo; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Promo" (
    id integer NOT NULL,
    banner_url character varying NOT NULL,
    title character varying NOT NULL,
    subtitle character varying NOT NULL,
    product_id integer NOT NULL,
    start_time date NOT NULL,
    end_time date NOT NULL
);


ALTER TABLE public."Promo" OWNER TO postgres;

--
-- Name: Promo_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Promo_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public."Promo_id_seq" OWNER TO postgres;

--
-- Name: Promo_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Promo_id_seq" OWNED BY public."Promo".id;


--
-- Name: Recommend; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Recommend" (
    id integer NOT NULL,
    banner_url character varying NOT NULL,
    title character varying NOT NULL,
    subtitle character varying NOT NULL,
    product_id integer NOT NULL,
    start_time date NOT NULL,
    end_time date NOT NULL
);


ALTER TABLE public."Recommend" OWNER TO postgres;

--
-- Name: Recommend_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Recommend_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public."Recommend_id_seq" OWNER TO postgres;

--
-- Name: Recommend_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Recommend_id_seq" OWNED BY public."Recommend".id;


--
-- Name: User; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."User" (
    id integer NOT NULL,
    name character varying NOT NULL,
    email character varying NOT NULL,
    phone character varying NOT NULL,
    password character varying NOT NULL,
    token character varying
);


ALTER TABLE public."User" OWNER TO postgres;

--
-- Name: User_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."User_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public."User_id_seq" OWNER TO postgres;

--
-- Name: User_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."User_id_seq" OWNED BY public."User".id;


--
-- Name: Variant; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Variant" (
    id integer NOT NULL,
    description character varying NOT NULL
);


ALTER TABLE public."Variant" OWNER TO postgres;

--
-- Name: Variant_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Variant_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public."Variant_id_seq" OWNER TO postgres;

--
-- Name: Variant_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Variant_id_seq" OWNED BY public."Variant".id;


--
-- Name: Address id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Address" ALTER COLUMN id SET DEFAULT nextval('public."Address_id_seq"'::regclass);


--
-- Name: Banner id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Banner" ALTER COLUMN id SET DEFAULT nextval('public."Banner_id_seq"'::regclass);


--
-- Name: Cart id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Cart" ALTER COLUMN id SET DEFAULT nextval('public."Cart_id_seq"'::regclass);


--
-- Name: Category id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Category" ALTER COLUMN id SET DEFAULT nextval('public."Category_id_seq"'::regclass);


--
-- Name: Gallery id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Gallery" ALTER COLUMN id SET DEFAULT nextval('public."Gallery_id_seq"'::regclass);


--
-- Name: Order id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Order" ALTER COLUMN id SET DEFAULT nextval('public."Order_id_seq"'::regclass);


--
-- Name: OrderProduct id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."OrderProduct" ALTER COLUMN id SET DEFAULT nextval('public."OrderProduct_id_seq"'::regclass);


--
-- Name: Product id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Product" ALTER COLUMN id SET DEFAULT nextval('public."Product_id_seq"'::regclass);


--
-- Name: ProductVariant id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."ProductVariant" ALTER COLUMN id SET DEFAULT nextval('public."ProductVariant_id_seq"'::regclass);


--
-- Name: Promo id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Promo" ALTER COLUMN id SET DEFAULT nextval('public."Promo_id_seq"'::regclass);


--
-- Name: Recommend id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Recommend" ALTER COLUMN id SET DEFAULT nextval('public."Recommend_id_seq"'::regclass);


--
-- Name: User id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."User" ALTER COLUMN id SET DEFAULT nextval('public."User_id_seq"'::regclass);


--
-- Name: Variant id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Variant" ALTER COLUMN id SET DEFAULT nextval('public."Variant_id_seq"'::regclass);


--
-- Data for Name: Address; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Address" (id, user_id, name, email, address, is_default) FROM stdin;
\.


--
-- Data for Name: Banner; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Banner" (id, banner_url, title, subtitle, path_page, start_time, end_time) FROM stdin;
1	https://example.com/banner1.jpg	Laptop Gaming XYZ	Promo potongan harga 20% untuk laptop gaming terbaik	/promo/laptop-gaming-xyz	2024-11-17	2024-11-24
2	https://example.com/banner2.jpg	Smartphone ABC	Belanja smartphone dengan diskon besar-besaran	/promo/smartphone-abc	2024-11-17	2024-11-24
3	https://example.com/banner3.jpg	Smartwatch 123	Dapatkan smartwatch dengan harga spesial hanya di promo ini	/promo/smartwatch-123	2024-11-17	2024-11-24
4	https://example.com/banner4.jpg	Headphone XYZ	Diskon 15% untuk semua produk headphone	/promo/headphone-xyz	2024-11-17	2024-11-24
5	https://example.com/banner5.jpg	Kamera DSLR ABC	Promo spesial akhir tahun untuk kamera DSLR dengan harga diskon	/promo/kamera-dslr-abc	2024-11-17	2024-11-21
\.


--
-- Data for Name: Cart; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Cart" (id, user_id, product_variant_id, qty) FROM stdin;
\.


--
-- Data for Name: Category; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Category" (id, name) FROM stdin;
1	Electronics
2	Books
3	Fashion
4	Home & Kitchen
5	Toys
6	Sports
7	Beauty
\.


--
-- Data for Name: Gallery; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Gallery" (id, product_id, image_url) FROM stdin;
1	1	https://example.com/images/smartphone_a_1.jpg
2	1	https://example.com/images/smartphone_a_2.jpg
3	1	https://example.com/images/smartphone_a_3.jpg
4	2	https://example.com/images/laptop_b_1.jpg
5	2	https://example.com/images/laptop_b_2.jpg
6	2	https://example.com/images/laptop_b_3.jpg
7	3	https://example.com/images/smartwatch_c_1.jpg
8	3	https://example.com/images/smartwatch_c_2.jpg
9	3	https://example.com/images/smartwatch_c_3.jpg
10	4	https://example.com/images/headphones_d_1.jpg
11	4	https://example.com/images/headphones_d_2.jpg
12	4	https://example.com/images/headphones_d_3.jpg
13	5	https://example.com/images/tablet_e_1.jpg
14	5	https://example.com/images/tablet_e_2.jpg
15	5	https://example.com/images/tablet_e_3.jpg
16	6	https://example.com/images/novel_f_1.jpg
17	6	https://example.com/images/novel_f_2.jpg
18	6	https://example.com/images/novel_f_3.jpg
19	7	https://example.com/images/biography_g_1.jpg
20	7	https://example.com/images/biography_g_2.jpg
21	7	https://example.com/images/biography_g_3.jpg
22	8	https://example.com/images/cookbook_h_1.jpg
23	8	https://example.com/images/cookbook_h_2.jpg
24	8	https://example.com/images/cookbook_h_3.jpg
25	9	https://example.com/images/science_fiction_i_1.jpg
26	9	https://example.com/images/science_fiction_i_2.jpg
27	9	https://example.com/images/science_fiction_i_3.jpg
28	10	https://example.com/images/childrens_book_j_1.jpg
29	10	https://example.com/images/childrens_book_j_2.jpg
30	10	https://example.com/images/childrens_book_j_3.jpg
31	11	https://example.com/images/tshirt_k_1.jpg
32	11	https://example.com/images/tshirt_k_2.jpg
33	11	https://example.com/images/tshirt_k_3.jpg
34	12	https://example.com/images/jeans_l_1.jpg
35	12	https://example.com/images/jeans_l_2.jpg
36	12	https://example.com/images/jeans_l_3.jpg
37	13	https://example.com/images/jacket_m_1.jpg
38	13	https://example.com/images/jacket_m_2.jpg
39	13	https://example.com/images/jacket_m_3.jpg
40	14	https://example.com/images/sneakers_n_1.jpg
41	14	https://example.com/images/sneakers_n_2.jpg
42	14	https://example.com/images/sneakers_n_3.jpg
43	15	https://example.com/images/hat_o_1.jpg
44	15	https://example.com/images/hat_o_2.jpg
45	15	https://example.com/images/hat_o_3.jpg
46	16	https://example.com/images/blender_p_1.jpg
47	16	https://example.com/images/blender_p_2.jpg
48	16	https://example.com/images/blender_p_3.jpg
49	17	https://example.com/images/microwave_q_1.jpg
50	17	https://example.com/images/microwave_q_2.jpg
51	17	https://example.com/images/microwave_q_3.jpg
52	18	https://example.com/images/cookware_set_r_1.jpg
53	18	https://example.com/images/cookware_set_r_2.jpg
54	18	https://example.com/images/cookware_set_r_3.jpg
55	19	https://example.com/images/vacuum_cleaner_s_1.jpg
56	19	https://example.com/images/vacuum_cleaner_s_2.jpg
57	19	https://example.com/images/vacuum_cleaner_s_3.jpg
58	20	https://example.com/images/air_purifier_t_1.jpg
59	20	https://example.com/images/air_purifier_t_2.jpg
60	20	https://example.com/images/air_purifier_t_3.jpg
61	21	https://example.com/images/action_figure_u_1.jpg
62	21	https://example.com/images/action_figure_u_2.jpg
63	21	https://example.com/images/action_figure_u_3.jpg
64	22	https://example.com/images/board_game_v_1.jpg
65	22	https://example.com/images/board_game_v_2.jpg
66	22	https://example.com/images/board_game_v_3.jpg
67	23	https://example.com/images/puzzle_w_1.jpg
68	23	https://example.com/images/puzzle_w_2.jpg
69	23	https://example.com/images/puzzle_w_3.jpg
70	24	https://example.com/images/doll_x_1.jpg
71	24	https://example.com/images/doll_x_2.jpg
72	24	https://example.com/images/doll_x_3.jpg
73	25	https://example.com/images/remote_car_y_1.jpg
74	25	https://example.com/images/remote_car_y_2.jpg
75	25	https://example.com/images/remote_car_y_3.jpg
76	26	https://example.com/images/football_z_1.jpg
77	26	https://example.com/images/football_z_2.jpg
78	26	https://example.com/images/football_z_3.jpg
79	27	https://example.com/images/tennis_racket_aa_1.jpg
80	27	https://example.com/images/tennis_racket_aa_2.jpg
81	27	https://example.com/images/tennis_racket_aa_3.jpg
82	28	https://example.com/images/yoga_mat_bb_1.jpg
83	28	https://example.com/images/yoga_mat_bb_2.jpg
84	28	https://example.com/images/yoga_mat_bb_3.jpg
85	29	https://example.com/images/dumbbells_cc_1.jpg
86	29	https://example.com/images/dumbbells_cc_2.jpg
87	29	https://example.com/images/dumbbells_cc_3.jpg
88	30	https://example.com/images/running_shoes_dd_1.jpg
89	30	https://example.com/images/running_shoes_dd_2.jpg
90	30	https://example.com/images/running_shoes_dd_3.jpg
91	31	https://example.com/images/lipstick_ee_1.jpg
92	31	https://example.com/images/lipstick_ee_2.jpg
93	31	https://example.com/images/lipstick_ee_3.jpg
94	32	https://example.com/images/moisturizer_ff_1.jpg
95	32	https://example.com/images/moisturizer_ff_2.jpg
96	32	https://example.com/images/moisturizer_ff_3.jpg
97	33	https://example.com/images/shampoo_gg_1.jpg
98	33	https://example.com/images/shampoo_gg_2.jpg
99	33	https://example.com/images/shampoo_gg_3.jpg
100	34	https://example.com/images/perfume_hh_1.jpg
101	34	https://example.com/images/perfume_hh_2.jpg
102	34	https://example.com/images/perfume_hh_3.jpg
103	35	https://example.com/images/face_mask_ii_1.jpg
104	35	https://example.com/images/face_mask_ii_2.jpg
105	35	https://example.com/images/face_mask_ii_3.jpg
\.


--
-- Data for Name: Order; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Order" (id, user_id, address_id, total_price, shipping_price, status, payment_method, created_at) FROM stdin;
\.


--
-- Data for Name: OrderProduct; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."OrderProduct" (id, order_id, product_variant_id, qty, sub_total, created_at) FROM stdin;
\.


--
-- Data for Name: Product; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Product" (id, category_id, name, description, price, discount, price_after_discount, rating, qty_sold, is_best_selling, created_at) FROM stdin;
1	1	Smartphone A	Latest model smartphone	500	10	450	4.5	200	t	2024-01-10
2	1	Laptop B	High-performance laptop	1200	15	1020	4.7	150	t	2024-02-15
3	1	Smartwatch C	Advanced smartwatch	300	5	285	4.2	80	f	2024-03-05
4	1	Headphones D	Noise-canceling headphones	150	20	120	4.3	120	t	2024-03-20
5	1	Tablet E	Lightweight tablet	400	10	360	4.1	90	f	2024-04-10
6	2	Novel F	Bestselling novel	20	5	19	4.8	300	t	2024-02-01
7	2	Biography G	Inspirational biography	25	0	25	4.6	180	t	2024-02-20
8	2	Cookbook H	Popular recipes cookbook	30	10	27	4.4	100	f	2024-03-01
9	2	Science Fiction I	Exciting sci-fi book	15	5	14.25	4.2	150	f	2024-03-25
10	2	Children’s Book J	Illustrated storybook for kids	10	0	10	4.7	250	t	2024-04-05
11	3	T-Shirt K	Comfortable cotton t-shirt	15	5	14.25	4.3	200	t	2024-01-20
12	3	Jeans L	Durable denim jeans	40	10	36	4.5	150	t	2024-02-10
13	3	Jacket M	Warm winter jacket	60	20	48	4.6	100	t	2024-02-25
14	3	Sneakers N	Stylish sneakers	50	15	42.5	4.4	130	f	2024-03-15
15	3	Hat O	Trendy summer hat	20	5	19	4.2	90	f	2024-03-30
16	4	Blender P	High-speed blender	80	10	72	4.3	120	t	2024-01-25
17	4	Microwave Q	Compact microwave oven	150	20	120	4.5	90	t	2024-02-15
18	4	Cookware Set R	Non-stick cookware set	100	15	85	4.6	110	t	2024-03-10
19	4	Vacuum Cleaner S	Powerful vacuum cleaner	200	25	150	4.7	80	f	2024-03-30
20	4	Air Purifier T	Advanced air purifier	250	30	175	4.4	70	f	2024-04-15
21	5	Action Figure U	Popular action figure	25	5	23.75	4.5	300	t	2024-01-05
22	5	Board Game V	Fun family board game	40	10	36	4.6	200	t	2024-01-20
23	5	Puzzle W	Challenging jigsaw puzzle	20	5	19	4.3	180	f	2024-02-05
24	5	Doll X	Beautiful collectible doll	30	0	30	4.2	150	f	2024-02-20
25	5	Remote Car Y	High-speed remote car	50	15	42.5	4.4	100	f	2024-03-05
26	6	Football Z	Durable football	30	10	27	4.7	250	t	2024-01-15
27	6	Tennis Racket AA	Lightweight tennis racket	100	20	80	4.5	120	t	2024-02-01
28	6	Yoga Mat BB	Comfortable yoga mat	40	10	36	4.3	150	t	2024-02-25
29	6	Dumbbells CC	Adjustable dumbbells	80	15	68	4.4	100	f	2024-03-10
30	6	Running Shoes DD	High-performance running shoes	120	25	90	4.6	80	f	2024-03-25
31	7	Lipstick EE	Long-lasting lipstick	20	5	19	4.6	300	t	2024-01-10
32	7	Moisturizer FF	Hydrating moisturizer	40	10	36	4.5	200	t	2024-02-15
33	7	Shampoo GG	Nourishing shampoo	25	5	23.75	4.3	150	f	2024-03-01
34	7	Perfume HH	Luxury perfume	80	15	68	4.7	100	f	2024-03-20
35	7	Face Mask II	Revitalizing face mask	30	0	30	4.2	120	f	2024-04-05
\.


--
-- Data for Name: ProductVariant; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."ProductVariant" (id, product_id, variant_id, stock) FROM stdin;
1	1	2	15
2	1	3	10
3	2	7	8
4	2	8	12
5	3	1	25
6	4	2	30
7	4	3	20
8	5	1	18
9	6	1	10
10	7	1	5
11	8	1	7
12	9	1	20
13	10	1	10
14	11	4	20
15	11	5	15
16	12	4	10
17	12	5	12
18	13	6	8
19	14	1	10
20	15	1	5
21	16	1	20
22	17	1	15
23	18	1	18
24	19	1	12
25	20	1	10
26	21	2	30
27	21	3	20
28	22	1	25
29	23	1	20
30	24	1	15
31	25	2	10
32	25	3	8
33	26	2	20
34	26	3	10
35	27	1	10
36	28	1	15
37	29	1	12
38	30	9	8
39	30	10	10
40	31	2	25
41	31	3	20
42	32	1	18
43	33	1	20
44	34	1	10
45	35	1	12
\.


--
-- Data for Name: Promo; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Promo" (id, banner_url, title, subtitle, product_id, start_time, end_time) FROM stdin;
1	https://example.com/banner1.jpg	Smartphone A	Dapatkan potongan harga untuk Smartphone A	1	2024-11-17	2024-11-24
2	https://example.com/banner2.jpg	T-Shirt K	Promo spesial untuk T-Shirt K	11	2024-11-17	2024-11-24
3	https://example.com/banner3.jpg	Laptop B	Beli Laptop B dan dapatkan diskon!	2	2024-11-17	2024-11-24
4	https://example.com/banner4.jpg	Vacuum Cleaner S	Potongan harga 25% untuk Vacuum Cleaner S	14	2024-11-17	2024-11-24
5	https://example.com/banner5.jpg	Board Game V	Diskon untuk Board Game V, dapatkan segera!	18	2024-11-17	2024-11-21
\.


--
-- Data for Name: Recommend; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Recommend" (id, banner_url, title, subtitle, product_id, start_time, end_time) FROM stdin;
1	https://example.com/banner6.jpg	Jeans L	Dapatkan diskon spesial untuk Jeans L	12	2024-11-17	2024-11-24
2	https://example.com/banner7.jpg	Cookware Set R	Promo besar-besaran untuk Cookware Set R	13	2024-11-17	2024-11-24
3	https://example.com/banner8.jpg	Children’s Book J	Beli Children’s Book J dengan harga spesial!	6	2024-11-17	2024-11-24
4	https://example.com/banner9.jpg	Yoga Mat BB	Promo Yoga Mat BB, diskon 10%!	19	2024-11-17	2024-11-24
5	https://example.com/banner10.jpg	Lipstick EE	Dapatkan Lipstick EE dengan harga diskon!	20	2024-11-17	2024-11-21
\.


--
-- Data for Name: User; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."User" (id, name, email, phone, password, token) FROM stdin;
\.


--
-- Data for Name: Variant; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Variant" (id, description) FROM stdin;
1	
2	Red
3	Blue
4	Small
5	Medium
6	Large
7	10inch
8	14inch
9	32
10	34
\.


--
-- Name: Address_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Address_id_seq"', 1, false);


--
-- Name: Banner_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Banner_id_seq"', 5, true);


--
-- Name: Cart_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Cart_id_seq"', 1, false);


--
-- Name: Category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Category_id_seq"', 7, true);


--
-- Name: Gallery_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Gallery_id_seq"', 105, true);


--
-- Name: OrderProduct_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."OrderProduct_id_seq"', 1, false);


--
-- Name: Order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Order_id_seq"', 1, false);


--
-- Name: ProductVariant_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."ProductVariant_id_seq"', 45, true);


--
-- Name: Product_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Product_id_seq"', 35, true);


--
-- Name: Promo_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Promo_id_seq"', 5, true);


--
-- Name: Recommend_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Recommend_id_seq"', 5, true);


--
-- Name: User_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."User_id_seq"', 1, false);


--
-- Name: Variant_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Variant_id_seq"', 10, true);


--
-- Name: Address Address_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Address"
    ADD CONSTRAINT "Address_pkey" PRIMARY KEY (id);


--
-- Name: Banner Banner_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Banner"
    ADD CONSTRAINT "Banner_pkey" PRIMARY KEY (id);


--
-- Name: Cart Cart_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Cart"
    ADD CONSTRAINT "Cart_pkey" PRIMARY KEY (id);


--
-- Name: Category Category_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Category"
    ADD CONSTRAINT "Category_pkey" PRIMARY KEY (id);


--
-- Name: Gallery Gallery_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Gallery"
    ADD CONSTRAINT "Gallery_pkey" PRIMARY KEY (id);


--
-- Name: OrderProduct OrderProduct_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."OrderProduct"
    ADD CONSTRAINT "OrderProduct_pkey" PRIMARY KEY (id);


--
-- Name: Order Order_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Order"
    ADD CONSTRAINT "Order_pkey" PRIMARY KEY (id);


--
-- Name: ProductVariant ProductVariant_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."ProductVariant"
    ADD CONSTRAINT "ProductVariant_pkey" PRIMARY KEY (id);


--
-- Name: Product Product_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Product"
    ADD CONSTRAINT "Product_pkey" PRIMARY KEY (id);


--
-- Name: Promo Promo_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Promo"
    ADD CONSTRAINT "Promo_pkey" PRIMARY KEY (id);


--
-- Name: Recommend Recommend_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Recommend"
    ADD CONSTRAINT "Recommend_pkey" PRIMARY KEY (id);


--
-- Name: User User_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."User"
    ADD CONSTRAINT "User_email_key" UNIQUE (email);


--
-- Name: User User_phone_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."User"
    ADD CONSTRAINT "User_phone_key" UNIQUE (phone);


--
-- Name: User User_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."User"
    ADD CONSTRAINT "User_pkey" PRIMARY KEY (id);


--
-- Name: Variant Variant_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Variant"
    ADD CONSTRAINT "Variant_pkey" PRIMARY KEY (id);


--
-- Name: Address Address_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Address"
    ADD CONSTRAINT "Address_user_id_fkey" FOREIGN KEY (user_id) REFERENCES public."User"(id);


--
-- Name: Cart Cart_product_variant_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Cart"
    ADD CONSTRAINT "Cart_product_variant_id_fkey" FOREIGN KEY (product_variant_id) REFERENCES public."ProductVariant"(id);


--
-- Name: Cart Cart_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Cart"
    ADD CONSTRAINT "Cart_user_id_fkey" FOREIGN KEY (user_id) REFERENCES public."User"(id);


--
-- Name: Gallery Gallery_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Gallery"
    ADD CONSTRAINT "Gallery_product_id_fkey" FOREIGN KEY (product_id) REFERENCES public."Product"(id);


--
-- Name: OrderProduct OrderProduct_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."OrderProduct"
    ADD CONSTRAINT "OrderProduct_order_id_fkey" FOREIGN KEY (order_id) REFERENCES public."Order"(id);


--
-- Name: OrderProduct OrderProduct_product_variant_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."OrderProduct"
    ADD CONSTRAINT "OrderProduct_product_variant_id_fkey" FOREIGN KEY (product_variant_id) REFERENCES public."ProductVariant"(id);


--
-- Name: Order Order_address_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Order"
    ADD CONSTRAINT "Order_address_id_fkey" FOREIGN KEY (address_id) REFERENCES public."Address"(id);


--
-- Name: Order Order_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Order"
    ADD CONSTRAINT "Order_user_id_fkey" FOREIGN KEY (user_id) REFERENCES public."User"(id);


--
-- Name: ProductVariant ProductVariant_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."ProductVariant"
    ADD CONSTRAINT "ProductVariant_product_id_fkey" FOREIGN KEY (product_id) REFERENCES public."Product"(id);


--
-- Name: ProductVariant ProductVariant_variant_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."ProductVariant"
    ADD CONSTRAINT "ProductVariant_variant_id_fkey" FOREIGN KEY (variant_id) REFERENCES public."Variant"(id);


--
-- Name: Product Product_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Product"
    ADD CONSTRAINT "Product_category_id_fkey" FOREIGN KEY (category_id) REFERENCES public."Category"(id);


--
-- Name: Promo Promo_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Promo"
    ADD CONSTRAINT "Promo_product_id_fkey" FOREIGN KEY (product_id) REFERENCES public."Product"(id);


--
-- Name: Recommend Recommend_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Recommend"
    ADD CONSTRAINT "Recommend_product_id_fkey" FOREIGN KEY (product_id) REFERENCES public."Product"(id);


--
-- PostgreSQL database dump complete
--

