FROM gradle:8.10-jdk21 AS build
WORKDIR /app
COPY settings.gradle.kts build.gradle.kts ./
COPY src ./src
RUN gradle --no-daemon installDist

FROM eclipse-temurin:21-jre
WORKDIR /app
COPY --from=build /app/build/install/paperplane-policy /app
ENV CONFIG_FILE=/app/config/application.yaml
ENTRYPOINT ["/app/bin/paperplane-policy"]
