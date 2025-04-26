# social-media-agent
A Social Media Agent written in Go / Playground

# Intro
This is a very basic example of a `smollm` hosted locally in the docker container and used by the backend written in Go to ask questions.

# How to run 

> docker-compose up -d
> air

# How to use

```
curl --location 'http://localhost:8080/api/v1/agent/ask-agent' \
--header 'Content-Type: application/json' \
--data '{
  "query": "Hello AI!"
}'
```

- This API does not persist the conversation context
- This API does not support tool calling (yet?)
