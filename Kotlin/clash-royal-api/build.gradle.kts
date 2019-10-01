import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

group="com.github.shouren"
version="1.0-SNAPSHOT"

buildscript {
    repositories {
        mavenCentral()
    }
    dependencies {
        classpath(kotlin("gradle-plugin", version="1.3.40"))
    }
}

plugins {
    kotlin("jvm") version "1.3.40"
}

repositories {
    mavenCentral()
    jcenter()
}

dependencies {
    implementation(kotlin("stdlib-jdk8"))
    implementation(group = "com.github.kittinunf.fuel", name = "fuel", version = "2.1.0")
}

val compileKotlin: KotlinCompile by tasks
compileKotlin.kotlinOptions {
    jvmTarget = "1.8"
}
val compileTestKotlin: KotlinCompile by tasks
compileTestKotlin.kotlinOptions {
    jvmTarget = "1.8"
}