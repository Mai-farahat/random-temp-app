# 🌡️ Random Temperature App

This is a simple Go application that generates a random temperature (0–40°C)  
and prints a message based on the temperature level.

---

## 🚀 How It Works

- Uses Go’s `math/rand` and `time` packages to generate a random temperature.  
- Prints:
  - 🥶 “It’s cold” if < 10°C  
  - 🍀 “It’s moderate” if between 10°C and 25°C  
  - 😎 “It’s warm” if > 25°C

---

## 🐳 Docker Instructions

### 1. Build the Docker image
```bash
docker build -t random-temp-app .

## 📦 Docker Hub Image
👉 [https://hub.docker.com/r/maifarahat/random-temp-app](https://hub.docker.com/r/maifarahat/random-temp-app)
