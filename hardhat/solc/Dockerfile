FROM ethereum/client-go:alltools-v1.12.2 as abigen
FROM ethereum/solc:0.8.21-alpine as solc
FROM alpine

COPY --from=solc /usr/local/bin/solc /usr/local/bin/solc
COPY --from=abigen /usr/local/bin/abigen /usr/local/bin/abigen

ENTRYPOINT ["/bin/sh"]
