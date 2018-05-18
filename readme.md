# Autoredeploy

Poll Docker hub to see if a new version of the docs image is available.
When that happens, redeploy the docs service.

## How to use this

1. Deploy the docs site
   ```
   docker stack deploy -c docker-compose.yml docs
   ```
2. Add a new cron job with the content of `cron.tab`
   ```
   crontab -e
   ```
3. Monitor the logs
   ```
   tail -f cron.log
   ```
