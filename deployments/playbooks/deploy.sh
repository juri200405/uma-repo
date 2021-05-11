#!/bin/sh

ansible-playbook -i hosts/dev.yml create.yml -vvv --skip-tags "docker,aws/ecs_ecr"

