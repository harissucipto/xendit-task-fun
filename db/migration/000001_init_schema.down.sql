DROP TRIGGER IF EXISTS moveDeleted ON comments;
DROP FUNCTION IF EXISTS moveDeleted();
DROP TABLE IF EXISTS deleted_comments;
DROP TABLE IF EXISTS comments;

