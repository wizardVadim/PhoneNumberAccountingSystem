-- Вставка типов номеров
INSERT INTO phone_number_type (type_name) VALUES 
    ('Мобильный'),
    ('Домашний'),
    ('Рабочий')
ON CONFLICT (type_name) DO NOTHING;

-- Вставка 50 физических лиц (простой способ с генерацией)
INSERT INTO physical_person (city, person_address, first_name, last_name, second_name, born_year)
SELECT 
    CASE (random() * 30)::int
        WHEN 0 THEN 'Москва' WHEN 1 THEN 'Санкт-Петербург' WHEN 2 THEN 'Новосибирск'
        WHEN 3 THEN 'Екатеринбург' WHEN 4 THEN 'Казань' WHEN 5 THEN 'Нижний Новгород'
        WHEN 6 THEN 'Челябинск' WHEN 7 THEN 'Омск' WHEN 8 THEN 'Самара'
        WHEN 9 THEN 'Ростов-на-Дону' WHEN 10 THEN 'Уфа' WHEN 11 THEN 'Красноярск'
        WHEN 12 THEN 'Пермь' WHEN 13 THEN 'Воронеж' WHEN 14 THEN 'Волгоград'
        ELSE 'Краснодар'
    END AS city,
    'ул. ' || chr(65 + (random() * 26)::int) || chr(97 + (random() * 26)::int) || 
    ', д. ' || (1 + (random() * 100)::int) AS person_address,
    CASE (random() * 40)::int
        WHEN 0 THEN 'Иван' WHEN 1 THEN 'Александр' WHEN 2 THEN 'Сергей' WHEN 3 THEN 'Дмитрий'
        WHEN 4 THEN 'Андрей' WHEN 5 THEN 'Алексей' WHEN 6 THEN 'Максим' WHEN 7 THEN 'Владимир'
        WHEN 8 THEN 'Николай' WHEN 9 THEN 'Евгений' WHEN 10 THEN 'Михаил' WHEN 11 THEN 'Павел'
        WHEN 12 THEN 'Артем' WHEN 13 THEN 'Роман' WHEN 14 THEN 'Олег' WHEN 15 THEN 'Василий'
        WHEN 16 THEN 'Анна' WHEN 17 THEN 'Елена' WHEN 18 THEN 'Мария' WHEN 19 THEN 'Ольга'
        WHEN 20 THEN 'Татьяна' WHEN 21 THEN 'Наталья' WHEN 22 THEN 'Ирина' WHEN 23 THEN 'Светлана'
        WHEN 24 THEN 'Юлия' WHEN 25 THEN 'Екатерина' WHEN 26 THEN 'Надежда' WHEN 27 THEN 'Людмила'
        ELSE 'Галина'
    END AS first_name,
    CASE (random() * 40)::int
        WHEN 0 THEN 'Иванов' WHEN 1 THEN 'Петров' WHEN 2 THEN 'Сидоров' WHEN 3 THEN 'Кузнецов'
        WHEN 4 THEN 'Смирнов' WHEN 5 THEN 'Попов' WHEN 6 THEN 'Васильев' WHEN 7 THEN 'Соколов'
        WHEN 8 THEN 'Михайлов' WHEN 9 THEN 'Новиков' WHEN 10 THEN 'Федоров' WHEN 11 THEN 'Морозов'
        WHEN 12 THEN 'Волков' WHEN 13 THEN 'Алексеев' WHEN 14 THEN 'Лебедев' WHEN 15 THEN 'Семенов'
        WHEN 16 THEN 'Егоров' WHEN 17 THEN 'Павлов' WHEN 18 THEN 'Козлов' WHEN 19 THEN 'Степанов'
        WHEN 20 THEN 'Николаев' WHEN 21 THEN 'Орлов' WHEN 22 THEN 'Андреев' WHEN 23 THEN 'Макаров'
        ELSE 'Никитин'
    END AS last_name,
    CASE (random() * 30)::int
        WHEN 0 THEN 'Иванович' WHEN 1 THEN 'Петрович' WHEN 2 THEN 'Сидорович' WHEN 3 THEN 'Александрович'
        WHEN 4 THEN 'Сергеевич' WHEN 5 THEN 'Алексеевич' WHEN 6 THEN 'Дмитриевич' WHEN 7 THEN 'Владимирович'
        WHEN 8 THEN 'Николаевич' WHEN 9 THEN 'Евгеньевич' WHEN 10 THEN 'Михайлович' WHEN 11 THEN 'Павлович'
        WHEN 12 THEN 'Андреевич' WHEN 13 THEN 'Ивановна' WHEN 14 THEN 'Петровна' WHEN 15 THEN 'Сидоровна'
        WHEN 16 THEN 'Александровна' WHEN 17 THEN 'Сергеевна' WHEN 18 THEN 'Алексеевна' WHEN 19 THEN 'Дмитриевна'
        ELSE 'Владимировна'
    END AS second_name,
    1950 + (random() * 55)::int AS born_year
FROM generate_series(1, 50);

-- Вставка телефонов (около 150 номеров, 2-4 на человека)
INSERT INTO phone_number (phone_number_value, person_id, phone_number_type_id, comment)
SELECT 
    DISTINCT ON (phone_value) phone_value,
    person_id,
    phone_type_id,
    CASE (random() * 3)::int
        WHEN 0 THEN NULL
        WHEN 1 THEN 'Основной'
        WHEN 2 THEN 'Дополнительный'
        ELSE 'Рабочий'
    END AS comment
FROM (
    SELECT 
        p.id AS person_id,
        CASE (random() * 2)::int
            WHEN 0 THEN '+79' || LPAD((100000000 + (random() * 899999999)::int)::text, 10, '0')
            WHEN 1 THEN '+7' || (495 + (random() * 5)::int) || LPAD((random() * 9999999)::int::text, 7, '0')
            ELSE '+7' || (812 + (random() * 2)::int) || LPAD((random() * 9999999)::int::text, 7, '0')
        END AS phone_value,
        1 + (random() * 3)::int AS phone_type_id
    FROM physical_person p
    CROSS JOIN generate_series(1, 2 + (random() * 3)::int)
    WHERE random() > 0.1  -- Немного рандомизируем количество
) AS phones
WHERE phone_value IS NOT NULL
LIMIT 150;