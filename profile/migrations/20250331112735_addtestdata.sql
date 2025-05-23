-- +goose Up
-- SQL in this section is executed when the migration is applied.
INSERT INTO profiles (
    id,
    name,
    last_name,
    sex,
    date_of_birth,
    city,
    height,
    chest_size,
    waist_size,
    hip_size,
    weight,
    avatar_url
) VALUES
-- 1-10
('11111111-1111-1111-1111-111111111111', 'Anna', 'Ivanova', 'female', '1990-05-15', 'Moscow', 170, 90, 60, 90, 55, 'https://example.com/avatars/1.jpg'),
('22222222-2222-2222-2222-222222222222', 'Elena', 'Petrova', 'female', '1988-07-22', 'Saint Petersburg', 165, 85, 58, 88, 52, 'https://example.com/avatars/2.jpg'),
('33333333-3333-3333-3333-333333333333', 'Irina', 'Sidorova', 'female', '1995-03-10', 'Novosibirsk', 168, 88, 62, 92, 58, 'https://example.com/avatars/3.jpg'),
('44444444-4444-4444-4444-444444444444', 'Olga', 'Smirnova', 'female', '1992-11-30', 'Yekaterinburg', 172, 92, 64, 94, 60, 'https://example.com/avatars/4.jpg'),
('55555555-5555-5555-5555-555555555555', 'Maria', 'Kuznetsova', 'female', '1985-09-18', 'Kazan', 175, 94, 66, 96, 62, 'https://example.com/avatars/5.jpg'),
('66666666-6666-6666-6666-666666666666', 'Alexey', 'Ivanov', 'male', '1987-04-25', 'Moscow', 185, 105, 80, 95, 80, 'https://example.com/avatars/6.jpg'),
('77777777-7777-7777-7777-777777777777', 'Dmitry', 'Petrov', 'male', '1991-12-05', 'Saint Petersburg', 180, 100, 78, 92, 75, 'https://example.com/avatars/7.jpg'),
('88888888-8888-8888-8888-888888888888', 'Sergey', 'Sidorov', 'male', '1989-08-14', 'Novosibirsk', 178, 98, 76, 90, 72, 'https://example.com/avatars/8.jpg'),
('99999999-9999-9999-9999-999999999999', 'Andrey', 'Smirnov', 'male', '1993-06-20', 'Yekaterinburg', 182, 102, 82, 97, 78, 'https://example.com/avatars/9.jpg'),
('00000000-0000-0000-0000-000000000001', 'Mikhail', 'Kuznetsov', 'male', '1984-02-28', 'Kazan', 188, 108, 85, 100, 85, 'https://example.com/avatars/10.jpg'),

-- 11-20
('00000000-0000-0000-0000-000000000011', 'Svetlana', 'Fedorova', 'female', '1994-01-12', 'Rostov', 167, 86, 59, 89, 54, 'https://example.com/avatars/11.jpg'),
('00000000-0000-0000-0000-000000000012', 'Natalia', 'Morozova', 'female', '1986-10-08', 'Samara', 169, 87, 61, 91, 56, 'https://example.com/avatars/12.jpg'),
('00000000-0000-0000-0000-000000000013', 'Tatiana', 'Volkova', 'female', '1997-07-17', 'Omsk', 171, 89, 63, 93, 57, 'https://example.com/avatars/13.jpg'),
('00000000-0000-0000-0000-000000000014', 'Yulia', 'Alekseeva', 'female', '1991-04-03', 'Chelyabinsk', 173, 91, 65, 95, 59, 'https://example.com/avatars/14.jpg'),
('00000000-0000-0000-0000-000000000015', 'Ekaterina', 'Lebedeva', 'female', '1989-12-19', 'Ufa', 174, 93, 67, 97, 61, 'https://example.com/avatars/15.jpg'),
('00000000-0000-0000-0000-000000000016', 'Vladimir', 'Fedorov', 'male', '1990-03-22', 'Rostov', 183, 103, 83, 98, 79, 'https://example.com/avatars/16.jpg'),
('00000000-0000-0000-0000-000000000017', 'Alexander', 'Morozov', 'male', '1988-11-15', 'Samara', 179, 99, 77, 91, 74, 'https://example.com/avatars/17.jpg'),
('00000000-0000-0000-0000-000000000018', 'Ivan', 'Volkov', 'male', '1996-08-07', 'Omsk', 181, 101, 81, 96, 77, 'https://example.com/avatars/18.jpg'),
('00000000-0000-0000-0000-000000000019', 'Pavel', 'Alekseev', 'male', '1992-05-29', 'Chelyabinsk', 184, 104, 84, 99, 81, 'https://example.com/avatars/19.jpg'),
('00000000-0000-0000-0000-000000000020', 'Nikolay', 'Lebedev', 'male', '1983-09-11', 'Ufa', 187, 107, 86, 101, 84, 'https://example.com/avatars/20.jpg'),

-- 21-30
('00000000-0000-0000-0000-000000000021', 'Anastasia', 'Egorova', 'female', '1993-02-14', 'Krasnodar', 166, 84, 58, 88, 53, 'https://example.com/avatars/21.jpg'),
('00000000-0000-0000-0000-000000000022', 'Polina', 'Sokolova', 'female', '1987-06-27', 'Voronezh', 168, 86, 60, 90, 55, 'https://example.com/avatars/22.jpg'),
('00000000-0000-0000-0000-000000000023', 'Daria', 'Mikhailova', 'female', '1998-01-09', 'Perm', 170, 88, 62, 92, 57, 'https://example.com/avatars/23.jpg'),
('00000000-0000-0000-0000-000000000024', 'Alina', 'Pavlova', 'female', '1990-10-23', 'Volgograd', 172, 90, 64, 94, 59, 'https://example.com/avatars/24.jpg'),
('00000000-0000-0000-0000-000000000025', 'Kristina', 'Romanova', 'female', '1986-07-16', 'Krasnoyarsk', 174, 92, 66, 96, 61, 'https://example.com/avatars/25.jpg'),
('00000000-0000-0000-0000-000000000026', 'Artem', 'Egorov', 'male', '1994-04-05', 'Krasnodar', 182, 102, 82, 97, 78, 'https://example.com/avatars/26.jpg'),
('00000000-0000-0000-0000-000000000027', 'Denis', 'Sokolov', 'male', '1989-12-18', 'Voronezh', 178, 98, 76, 91, 73, 'https://example.com/avatars/27.jpg'),
('00000000-0000-0000-0000-000000000028', 'Maxim', 'Mikhailov', 'male', '1997-09-30', 'Perm', 180, 100, 80, 95, 76, 'https://example.com/avatars/28.jpg'),
('00000000-0000-0000-0000-000000000029', 'Anton', 'Pavlov', 'male', '1991-06-12', 'Volgograd', 183, 103, 83, 98, 80, 'https://example.com/avatars/29.jpg'),
('00000000-0000-0000-0000-000000000030', 'Konstantin', 'Romanov', 'male', '1985-03-24', 'Krasnoyarsk', 186, 106, 85, 100, 83, 'https://example.com/avatars/30.jpg'),

-- 31-40
('00000000-0000-0000-0000-000000000031', 'Victoria', 'Orlova', 'female', '1992-08-07', 'Saratov', 167, 85, 59, 89, 54, 'https://example.com/avatars/31.jpg'),
('00000000-0000-0000-0000-000000000032', 'Valeria', 'Belyaeva', 'female', '1988-05-19', 'Tolyatti', 169, 87, 61, 91, 56, 'https://example.com/avatars/32.jpg'),
('00000000-0000-0000-0000-000000000033', 'Veronika', 'Guseva', 'female', '1999-02-01', 'Barnaul', 171, 89, 63, 93, 58, 'https://example.com/avatars/33.jpg'),
('00000000-0000-0000-0000-000000000034', 'Margarita', 'Titova', 'female', '1991-11-13', 'Ulyanovsk', 173, 91, 65, 95, 60, 'https://example.com/avatars/34.jpg'),
('00000000-0000-0000-0000-000000000035', 'Evgenia', 'Simonova', 'female', '1987-08-26', 'Irkutsk', 175, 93, 67, 97, 62, 'https://example.com/avatars/35.jpg'),
('00000000-0000-0000-0000-000000000036', 'Gleb', 'Orlov', 'male', '1993-05-08', 'Saratov', 181, 101, 81, 96, 77, 'https://example.com/avatars/36.jpg'),
('00000000-0000-0000-0000-000000000037', 'Roman', 'Belyaev', 'male', '1989-12-21', 'Tolyatti', 177, 97, 75, 90, 72, 'https://example.com/avatars/37.jpg'),
('00000000-0000-0000-0000-000000000038', 'Viktor', 'Gusev', 'male', '1998-09-03', 'Barnaul', 179, 99, 79, 94, 75, 'https://example.com/avatars/38.jpg'),
('00000000-0000-0000-0000-000000000039', 'Yuri', 'Titov', 'male', '1992-06-15', 'Ulyanovsk', 182, 102, 82, 97, 79, 'https://example.com/avatars/39.jpg'),
('00000000-0000-0000-0000-000000000040', 'Stanislav', 'Simonov', 'male', '1986-03-28', 'Irkutsk', 185, 105, 84, 99, 82, 'https://example.com/avatars/40.jpg'),

-- 41-46
('00000000-0000-0000-0000-000000000041', 'Galina', 'Krylova', 'female', '1990-04-11', 'Yaroslavl', 168, 86, 60, 90, 55, 'https://example.com/avatars/41.jpg'),
('00000000-0000-0000-0000-000000000042', 'Larisa', 'Vorobeva', 'female', '1985-01-24', 'Tomsk', 170, 88, 62, 92, 57, 'https://example.com/avatars/42.jpg'),
('00000000-0000-0000-0000-000000000043', 'Zoya', 'Karpova', 'female', '1995-10-07', 'Kemerovo', 172, 90, 64, 94, 59, 'https://example.com/avatars/43.jpg'),
('00000000-0000-0000-0000-000000000044', 'Boris', 'Krylov', 'male', '1991-07-20', 'Yaroslavl', 180, 100, 80, 95, 76, 'https://example.com/avatars/44.jpg'),
('00000000-0000-0000-0000-000000000045', 'Leonid', 'Vorobev', 'male', '1987-04-02', 'Tomsk', 178, 98, 78, 93, 74, 'https://example.com/avatars/45.jpg'),
('00000000-0000-0000-0000-000000000046', 'Arkady', 'Karpov', 'male', '1996-11-15', 'Kemerovo', 181, 101, 81, 96, 77, 'https://example.com/avatars/46.jpg');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DELETE FROM profiles WHERE id IN (
                                  '11111111-1111-1111-1111-111111111111',
                                  '22222222-2222-2222-2222-222222222222',
                                  '33333333-3333-3333-3333-333333333333',
                                  '44444444-4444-4444-4444-444444444444',
                                  '55555555-5555-5555-5555-555555555555',
                                  '66666666-6666-6666-6666-666666666666',
                                  '77777777-7777-7777-7777-777777777777',
                                  '88888888-8888-8888-8888-888888888888',
                                  '99999999-9999-9999-9999-999999999999',
                                  '00000000-0000-0000-0000-000000000001',
                                  '00000000-0000-0000-0000-000000000011',
                                  '00000000-0000-0000-0000-000000000012',
                                  '00000000-0000-0000-0000-000000000013',
                                  '00000000-0000-0000-0000-000000000014',
                                  '00000000-0000-0000-0000-000000000015',
                                  '00000000-0000-0000-0000-000000000016',
                                  '00000000-0000-0000-0000-000000000017',
                                  '00000000-0000-0000-0000-000000000018',
                                  '00000000-0000-0000-0000-000000000019',
                                  '00000000-0000-0000-0000-000000000020',
                                  '00000000-0000-0000-0000-000000000021',
                                  '00000000-0000-0000-0000-000000000022',
                                  '00000000-0000-0000-0000-000000000023',
                                  '00000000-0000-0000-0000-000000000024',
                                  '00000000-0000-0000-0000-000000000025',
                                  '00000000-0000-0000-0000-000000000026',
                                  '00000000-0000-0000-0000-000000000027',
                                  '00000000-0000-0000-0000-000000000028',
                                  '00000000-0000-0000-0000-000000000029',
                                  '00000000-0000-0000-0000-000000000030',
                                  '00000000-0000-0000-0000-000000000031',
                                  '00000000-0000-0000-0000-000000000032',
                                  '00000000-0000-0000-0000-000000000033',
                                  '00000000-0000-0000-0000-000000000034',
                                  '00000000-0000-0000-0000-000000000035',
                                  '00000000-0000-0000-0000-000000000036',
                                  '00000000-0000-0000-0000-000000000037',
                                  '00000000-0000-0000-0000-000000000038',
                                  '00000000-0000-0000-0000-000000000039',
                                  '00000000-0000-0000-0000-000000000040',
                                  '00000000-0000-0000-0000-000000000041',
                                  '00000000-0000-0000-0000-000000000042',
                                  '00000000-0000-0000-0000-000000000043',
                                  '00000000-0000-0000-0000-000000000044',
                                  '00000000-0000-0000-0000-000000000045',
                                  '00000000-0000-0000-0000-000000000046'
    );