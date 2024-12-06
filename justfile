set dotenv-load

default:
    @just --list

build-auth:
    go build -C ./src/auth -o build/qme-auth

containerize-auth: 
    docker build ./src/auth -t qmessentials/auth:$AUTH_VERSION

apply-db-migrations-auth:
    kubectl get secret auth-secret -n qmessentials -o jsonpath='{.data.DATABASE_USER}' | base64 --decode > database_user
    kubectl get secret auth-secret -n qmessentials -o jsonpath='{.data.DATABASE_PASSWORD}' | base64 --decode > database_password
    ruby ./migrations/migrate.rb 'auth'
    rm database_user database_password

apply-k8s:
    kubectl apply -f ./kubernetes/namespace.yaml
    kubectl apply -f ./.secrets/auth-secret.yaml
    kubectl apply -f ./kubernetes/statefulsets/auth-db-statefulset.yaml
    kubectl apply -f ./kubernetes/services/auth-db-service.yaml