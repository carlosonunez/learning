FROM ruby:alpine as base_image
MAINTAINER Carlos Nunez <dev@carlosnunez.me>
WORKDIR /app

COPY . /app
RUN \
  apk update && \
  apk add --no-cache ruby-bundler ruby-json ruby-dev make gcc && \
  bundle install
ENTRYPOINT ["sh", "-c"]

FROM ruby:alpine as app_image
WORKDIR /app
COPY --from=base_image . /app
RUN \
  apk update && \
  apk add --no-cache ruby-bundler ruby-json ruby-dev && \
  bundle install --without test && \
  gem build subfactorials_app.gemspec && \
  gem install subfactorials-app.gem
ENTRYPOINT ["subfactorials-app"]
