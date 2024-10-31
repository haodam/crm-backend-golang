-- name: GetUserByEmailSQLC :one
SELECT user_id, user_email FROM pre_go_crm_user_c WHERE user_email = ? LIMIT 1;