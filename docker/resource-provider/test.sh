
last_line=$(echo "$(bacalhau docker run richbrem/cowsay:latest)" | tail -n 1)
status=$($last_line --json | jq '.["State"]["Executions"][0]["Status"]')
if [ "$status" = "\"Accepted job. execution completed\"" ]; then
    echo "Test Job completed successfully"
else
    echo "Test Job failed with status: $status"
    sleep infinity
fi
