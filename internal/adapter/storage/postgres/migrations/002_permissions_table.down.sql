ALTER TABLE user_permissions DROP CONSTRAINT user_permissions_permission_id_fkey;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS user_permissions;