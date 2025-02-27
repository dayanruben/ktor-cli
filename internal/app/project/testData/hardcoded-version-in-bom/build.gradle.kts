plugins {
    kotlin("jvm") version "2.0.20"
}

repositories {
    mavenCentral()
}

dependencies {
    implementation(platform("group:artifact:1.2.3"))
    implementation(platform("io.ktor:ktor-bom:2.3.13"))
    implementation(platform("group:artifact:1.2.3"))
}
