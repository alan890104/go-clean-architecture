#!/bin/bash

# Checking if docker-compose or docker compose command exists
if command -v docker-compose &> /dev/null; then
    DOCKER_COMP="docker-compose"
else
    DOCKER_COMP="docker compose"
fi

# Function to export environment variables based on the OS
function export_env() {
    local mode=$1
    local env_file=".env"

    # Check if mode is provided, then set the env_file
    if [ -n "$mode" ]; then
        env_file=".env.${mode}"
    fi

    unamestr=$(uname)
    if [ "$unamestr" = 'Linux' ]; then
        export $(grep -v '^#' $env_file | xargs -d '\n')
    elif [ "$unamestr" = 'FreeBSD' ] || [ "$unamestr" = 'Darwin' ]; then
        export $(grep -v '^#' $env_file | xargs -0)
    fi
}

mode=$1


case "$mode" in
    install)
        go install github.com/cosmtrek/air@latest
        exit 0
        ;;
        
    dev|prod|test)
        # All actions for these modes are handled in the next switch case block
        ;;

    *)
        echo "Error: Invalid mode. Choose from (install | dev | prod | test)"
        exit 1
        ;;
esac

action=$2

# Switch case to handle the different command options
case "$action" in
    start|stop|teardown)
        export_env $mode
        if [ "$action" = "start" ]; then
            $DOCKER_COMP -f docker-compose.yaml -f docker-compose.${mode}.yaml up -d 
        elif [ "$action" = "stop" ]; then
            $DOCKER_COMP -f docker-compose.yaml -f docker-compose.${mode}.yaml down
        elif [ "$action" = "teardown" ]; then
            $DOCKER_COMP -f docker-compose.yaml -f docker-compose.${mode}.yaml down --remove-orphans -v
        else
            echo "Error: Invalid command. Choose from (start | stop | teardown)"
            exit 1
        fi
        ;;

    generate|migrate|run|serve)
        export_env $mode
        if [ "$action" = "generate" ]; then
            go run ./cmd/gen/gen.go
        elif [ "$action" = "migrate" ]; then
            go run ./cmd/migrate/migrate.go
        elif [ "$action" = "run" ]; then
            go run ./cmd/app/app.go
        elif [ "$action" = "serve" ]; then
            air
        else
            echo "Error: Invalid command. Choose from (generate | migrate | run | serve)"
            exit 1
        fi
        ;;

    *)
        echo "Usage: ./run.sh [dev|prod|test] [init|start|stop|teardown|generate|migrate|run|serve]"
        exit 1
        ;;
esac

exit 0
