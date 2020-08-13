---
title: Memory Reservation in Amazon Elastic Container Service
published: true
tags: aws, docker, learning, beginners
cover_image: cover.jpg
---

We use [**task definition**](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definitions.html) to describe how we want a Docker container to be deployed in an ECS cluster. `memoryReservation` is one of the **container definitions** that need to be specified when writing the task definition, see [Task definition parameters by AWS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#container_definitions):

> If a task-level memory value is not specified, you must specify a non-zero integer for one or both of memory or memoryReservation in a container definition.

```json
[
  {
    "name": "worker",
    "image": "alexeiled/stress-ng:latest",
    "command": [
      "--vm-bytes",
      "300m",
      "--vm-keep",
      "--vm",
      "1",
      "-t",
      "1d",
      "-l",
      "0"
    ],
    "memoryReservation": 400
  }
]
```

Essentially, this task definition will launch a container that is constantly consuming `300 MB / 1.049 = 286.102 MiB` of memory (see [`stress-ng`](https://manpages.ubuntu.com/manpages/artful/man1/stress-ng.1.html)). We can think of this as a memory hungry worker. The Terraform configuration of the set up is available in GitHub:

{% github shihanng/ecs-resource-exp %}
