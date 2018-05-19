# API service

This service receives JSON payloads and adds them to a message queue.
One use case for it is to receive webhook events, and add them to
a message queue so that they can be processed asynchronously.

## Examples

Here are some snippets that are useful for development

```
curl -X POST -H "Content-Type: application/json" -d '{"key":"value"}' http://localhost:8000
```
