# sudo docker build -t arkhanna/base:latest .
FROM arkhanna/base:latest

COPY --chmod=755 build/arkhanna-master-server /app/master-server

WORKDIR /app
ENTRYPOINT ["./master-server"]