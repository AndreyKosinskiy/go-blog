package repository

const CreateUserSqlQuery = `INSERT INTO "users" ("id","user_name","email","password","is_deleted","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING *`
const DeleteUserSqlQuery = `UPDATE "users" SET "is_deleted"=$1 WHERE id = $2 AND is_deleted = $3 RETURNING *`
