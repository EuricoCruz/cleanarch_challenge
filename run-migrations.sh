#!/bin/sh
set -e

DB_URL="mysql://root:root@tcp(mysql:3306)/orders"

echo "Aguardando MySQL subir..."
sleep 10  # ou use wait-for-it se quiser mais robusto

echo "Rodando migrations..."
migrate -path=/migrations -database "$DB_URL" -verbose down -all
migrate -path=/migrations -database "$DB_URL" -verbose up

echo "Migrations concluídas!"