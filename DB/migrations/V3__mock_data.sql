-- ─────────────────────────────────────────────────────────────────────────────
-- Mock seed data – for local development and demo purposes only
-- Credentials:
--   demo@example.com  / password123
--   alice@example.com / password123
-- ─────────────────────────────────────────────────────────────────────────────

-- Users
INSERT INTO auth_users (id, email, password_hash) VALUES
    ('11111111-0000-0000-0000-000000000001', 'demo@example.com',  '$2b$10$pgE5E.8VlgKewcZkhaMmUOBxlAqZTG8FQwAMFkqeb09nyAMXkIETS'),
    ('11111111-0000-0000-0000-000000000002', 'alice@example.com', '$2b$10$Bnwt57OmdqSj65bHRYp.HOEVrQxv.6Y4T67d5XRu7m3WHgbJBmHBG')
ON CONFLICT (email) DO NOTHING;

-- Storages for demo user
INSERT INTO storages (id, user_id, name) VALUES
    ('22222222-0000-0000-0000-000000000001', '11111111-0000-0000-0000-000000000001', 'Photos'),
    ('22222222-0000-0000-0000-000000000002', '11111111-0000-0000-0000-000000000001', 'Documents'),
    ('22222222-0000-0000-0000-000000000003', '11111111-0000-0000-0000-000000000001', 'Projects')
ON CONFLICT DO NOTHING;

-- Storages for alice
INSERT INTO storages (id, user_id, name) VALUES
    ('22222222-0000-0000-0000-000000000004', '11111111-0000-0000-0000-000000000002', 'My Files')
ON CONFLICT DO NOTHING;

-- Items in Photos storage
INSERT INTO items (id, storage_id, name, size_mb, tags) VALUES
    ('33333333-0000-0000-0000-000000000001', '22222222-0000-0000-0000-000000000001', 'vacation-2025.jpg',  3.4,  'photo,vacation,2025'),
    ('33333333-0000-0000-0000-000000000002', '22222222-0000-0000-0000-000000000001', 'birthday-party.png', 5.2,  'photo,birthday'),
    ('33333333-0000-0000-0000-000000000003', '22222222-0000-0000-0000-000000000001', 'sunset.jpg',         2.1,  'photo,nature')
ON CONFLICT DO NOTHING;

-- Items in Documents storage
INSERT INTO items (id, storage_id, name, size_mb, tags) VALUES
    ('33333333-0000-0000-0000-000000000004', '22222222-0000-0000-0000-000000000002', 'resume.pdf',         0.3,  'document,work'),
    ('33333333-0000-0000-0000-000000000005', '22222222-0000-0000-0000-000000000002', 'invoice-march.pdf',  0.1,  'document,finance'),
    ('33333333-0000-0000-0000-000000000006', '22222222-0000-0000-0000-000000000002', 'notes.txt',          0.01, 'text,notes')
ON CONFLICT DO NOTHING;

-- Items in Projects storage
INSERT INTO items (id, storage_id, name, size_mb, tags) VALUES
    ('33333333-0000-0000-0000-000000000007', '22222222-0000-0000-0000-000000000003', 'architecture.png',   8.7,  'diagram,design'),
    ('33333333-0000-0000-0000-000000000008', '22222222-0000-0000-0000-000000000003', 'demo-video.mp4',     120.0,'video,demo')
ON CONFLICT DO NOTHING;

-- Items in alice's storage
INSERT INTO items (id, storage_id, name, size_mb, tags) VALUES
    ('33333333-0000-0000-0000-000000000009', '22222222-0000-0000-0000-000000000004', 'report.xlsx',        1.5,  'spreadsheet,work')
ON CONFLICT DO NOTHING;
