#!/bin/bash

PROJECT_NAME="go-clean-arch"

# Checking if docker-compose or docker compose command exists
if command -v docker-compose &> /dev/null; then
    DOCKER_COMP="docker-compose -p ${PROJECT_NAME}"
else
    DOCKER_COMP="docker compose -p ${PROJECT_NAME}"
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

    generate)
        go run ./cmd/gen/gen.go
        exit 0
        ;;
        
    dev|stage|test)
        # All actions for these modes are handled in the next switch case block
        ;;

    *)
        echo "Error: Invalid mode. Choose from (install | dev | stage | test)"
        exit 1
        ;;
esac

action=$2

# Switch case to handle the different command options
case "$action" in
    start|stop|teardown)
        export_env $mode
        if [ "$action" = "start" ]; then
            $DOCKER_COMP -f docker/docker-compose.yaml -f docker/docker-compose.${mode}.yaml up -d ${@:3}
        elif [ "$action" = "stop" ]; then
            $DOCKER_COMP -f docker/docker-compose.yaml -f docker/docker-compose.${mode}.yaml down 
        elif [ "$action" = "teardown" ]; then
            $DOCKER_COMP -f docker/docker-compose.yaml -f docker/docker-compose.${mode}.yaml down --remove-orphans -v
            echo "*** WARNING ***"
            echo "Please run 'sudo rm -rf docker/volumes' by yourself to remove the persistant volumes"
        else
            echo "Error: Invalid command. Choose from (start | stop | teardown)"
            exit 1
        fi
        ;;

    migrate|run|serve)
        export_env $mode
        if [ "$action" = "migrate" ]; then
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
        echo "Usage: ./run.sh [install | generate | dev | stage | test] [start | stop | teardown | migrate | run | serve]"
        exit 1
        ;;
esac

exit 0
