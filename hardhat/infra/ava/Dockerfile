FROM avaplatform/avalanche-cli:v1.5.2

WORKDIR /app

COPY ./conf/lilypad.json /app/lilypad.json
COPY ./key/ /root/.avalanche-cli/key/

RUN /avalanche subnet import file /app/lilypad.json
RUN /avalanche subnet deploy lilypad -l