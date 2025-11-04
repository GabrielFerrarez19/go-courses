#!/bin/bash

# Carregar variáveis de ambiente
export PGPASSWORD=${GOBID_DATABASE_PASSWORD}
export DB_USER=${GOBID_DATABASE_USER:-postgres}
export DB_NAME=${GOBID_DATABASE_NAME:-gobid}
export DB_HOST=${GOBID_DATABASE_HOST:-db}

echo "Waiting for database to be ready..."
until pg_isready -h ${DB_HOST} -U ${DB_USER} -d ${DB_NAME}; do
    sleep 1
done


echo "Creating data if doesn't exits..."
psql -h ${DB_HOST} -U ${DB_USER} -c "CREATE DATABASE ${DB_NAME};" || true


echo "Applying migrations..."
cd /app/internal/store/pgstore/migrations
tern migrate


echo "Starting application..."
cd /app

exec ./main
