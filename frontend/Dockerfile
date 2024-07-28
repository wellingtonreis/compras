FROM node:20.13.1

RUN apt update && apt install -y tzdata \
    && ln -fs /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime \
    && dpkg-reconfigure -f noninteractive tzdata \
    && apt-get clean

WORKDIR /app

RUN npm install -g @vue/cli && \
    npm install -g @vue/cli-init && \
    npm install -g @quasar/cli
    
COPY . .

RUN rm -rf package-lock.json

RUN npm cache clean --force
RUN npm install && npm rebuild node-sass

RUN chown -R node:node /app

EXPOSE 9000

CMD ["quasar", "dev"]
