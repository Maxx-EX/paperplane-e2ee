plugins {
    kotlin("jvm") version "2.0.21"
    application
}

group = "io.paperplane"
version = "0.1.0"

repositories { mavenCentral() }

dependencies {
    implementation("org.jetbrains.kotlinx:kotlinx-coroutines-core:1.9.0")
    implementation("io.lettuce:lettuce-core:6.5.2.RELEASE")
    implementation("org.yaml:snakeyaml:2.3")
    implementation("org.slf4j:slf4j-simple:2.0.16")
}

application {
    mainClass.set("io.paperplane.policy.MainKt")
}

tasks.withType<org.jetbrains.kotlin.gradle.tasks.KotlinCompile> {
    compilerOptions.freeCompilerArgs.addAll("-Xjsr305=strict")
}

tasks.jar {
    manifest.attributes["Main-Class"] = "io.paperplane.policy.MainKt"
}
