INSERT INTO user_role VALUES (0, 'Администратор');
INSERT INTO user_role VALUES (1, 'Оператор call-центра');
INSERT INTO user_role VALUES (2, 'Работник МФЦ');

INSERT INTO "user" (login, password, role_id) VALUES
('admin', crypt('root', gen_salt('bf')), 0),
('operator_call', crypt('operator2', gen_salt('bf')), 1),
('mfc_empl', crypt('mfc_empl1', gen_salt('bf')), 2);