SELECT username, hashed_password 
FROM accounts
WHERE username = $1;