FROM golang:latest as base
WORKDIR /usr/src/app

COPY . .
RUN go mod download && go mod verify
RUN go build -v .
RUN go install

RUN (curl -Ls --tlsv1.2 --proto "=https" --retry 3 https://cli.doppler.com/install.sh || wget -t 3 -qO- https://cli.doppler.com/install.sh) | sh

RUN touch run
RUN echo "#!/bin/bash" >> run
RUN echo "doppler run -- lilypad jobcreator" >> run
RUN chmod +x run

CMD ["/bin/bash", "./run"]
