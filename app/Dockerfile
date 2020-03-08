FROM node:latest
COPY vue /app/
WORKDIR /app/
RUN yarn global add @vue/cli && \
  yarn install

CMD ["yarn", "serve"]
  