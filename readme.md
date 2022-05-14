# Running Atlantis
Atlantis is a tools to run terraform automation.

To run atlantis inside heroku dynos
```shell
atlantis server \
--atlantis-url="$URL" \
--gh-user="$USERNAME" \
--gh-token="$TOKEN" \
--gh-webhook-secret="$SECRET" \
--repo-allowlist="$REPO_ALLOWLIST"
```